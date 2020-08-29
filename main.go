package main

import (
	"log"
	"net/http"
	"sync"
)

var (
	sha1    string
	version string
)

type BuildVars struct{}

func (b *BuildVars) getSha1() string {
	return sha1
}

func (b *BuildVars) getVersion() string {
	return version
}

// StartHTTPServer ...
func StartHTTPServer(wg *sync.WaitGroup) *http.Server {
	server := &http.Server{
		Addr: ":5000",
		Handler: &controller{
			&BuildVars{},
		},
	}

	go func() {
		defer wg.Done()

		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	return server
}

func main() {
	serverExit := &sync.WaitGroup{}
	serverExit.Add(1)
	StartHTTPServer(serverExit)
	serverExit.Wait()
}
