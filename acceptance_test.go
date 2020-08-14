package main_test

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"reflect"
	"sync"
	"testing"

	. "arctair.com/quarky"
)

func assertNotError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func TestAcceptance(t *testing.T) {
	baseUrl := os.Getenv("BASE_URL")
	if baseUrl == "" {
		baseUrl = "http://localhost:5000/"

		serverExit := &sync.WaitGroup{}
		serverExit.Add(1)
		server := StartHTTPServer(serverExit)

		defer func() {
			if err := server.Shutdown(context.TODO()); err != nil {
				panic(err)
			}

			serverExit.Wait()
		}()
	}

	response, err := http.Get(baseUrl)
	assertNotError(t, err)

	var got []string
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&got)
	assertNotError(t, err)

	want := []string{}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
