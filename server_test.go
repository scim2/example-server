package server

import (
	"context"
	test "github.com/scim2/test-suite"
	"log"
	"net/http"
	"sync"
	"testing"
)

func startHTTPServer(wg *sync.WaitGroup) *http.Server {
	server := &http.Server{Addr: ":3000"}
	http.Handle("/", Server)

	go func() {
		defer wg.Done()

		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	return server
}

func TestServer(t *testing.T) {
	httpServerExitDone := &sync.WaitGroup{}
	httpServerExitDone.Add(1)
	server := startHTTPServer(httpServerExitDone)

	testSuite := new(test.SCIMTestSuite)
	testSuite.BaseURL("http://localhost:3000")
	// suite.Run(t, testSuite)

	if err := server.Shutdown(context.TODO()); err != nil {
		panic(err)
	}
	httpServerExitDone.Wait()
}
