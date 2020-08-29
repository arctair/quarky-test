package main_test

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"testing"
	"time"

	"github.com/cenkalti/backoff/v4"
)

func TestAcceptance(t *testing.T) {
	baseUrl := os.Getenv("BASE_URL")
	if baseUrl == "" {
		baseUrl = "http://localhost:5000/"

		assertNotError(t, exec.Command("sh", "build").Run())

		command := exec.Command("bin/quarky")
		stderr, err := command.StderrPipe()
		assertNotError(t, err)
		assertNotError(t, command.Start())
		defer dumpPipe("app:", stderr)
		defer command.Process.Kill()

		assertNotError(
			t,
			backoff.Retry(
				func() error {
					_, err := http.Get(baseUrl)
					return err
				},
				NewExponentialBackOff(),
			),
		)
	}

	t.Run("GET /version returns sha1 and version", func(t *testing.T) {
		response, err := http.Get(baseUrl)
		assertNotError(t, err)

		var got map[string]string
		defer response.Body.Close()
		err = json.NewDecoder(response.Body).Decode(&got)
		assertNotError(t, err)

		sha1Pattern := regexp.MustCompile("^[0-9a-f]{40}(-dirty)?$")
		versionPattern := regexp.MustCompile("^\\d+\\.\\d+\\.\\d+$")

		if !sha1Pattern.MatchString(got["sha1"]) {
			t.Errorf("got sha1 %s want 40 hex digits", got["sha1"])
		}
		if !versionPattern.MatchString(got["version"]) && !sha1Pattern.MatchString(got["version"]) {
			t.Errorf("got version %s want semver or 40 hex digits", got["version"])
		}
	})
}

func assertNotError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func dumpPipe(prefix string, p io.ReadCloser) {
	s := bufio.NewScanner(p)
	for s.Scan() {
		log.Printf("%s: %s", prefix, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Printf("Failed to dump pipe: %s", err)
	}
}

func NewExponentialBackOff() *backoff.ExponentialBackOff {
	b := &backoff.ExponentialBackOff{
		InitialInterval:     backoff.DefaultInitialInterval,
		RandomizationFactor: backoff.DefaultRandomizationFactor,
		Multiplier:          backoff.DefaultMultiplier,
		MaxInterval:         backoff.DefaultMaxInterval,
		MaxElapsedTime:      3 * time.Second,
		Stop:                backoff.Stop,
		Clock:               backoff.SystemClock,
	}
	b.Reset()
	return b
}
