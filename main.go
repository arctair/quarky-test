package main

import (
	"log"
	"net/http"
	"sync"
)

// StartHTTPServer ...
func StartHTTPServer(wg *sync.WaitGroup) *http.Server {
	server := &http.Server{
		Addr:    ":5000",
		Handler: &controller{},
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
