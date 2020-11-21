package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

const addr = ":9000"

func main() {
	go func() {
		fmt.Printf("Listening on %s\n", addr)
		log.Fatal(http.ListenAndServe(addr, http.HandlerFunc(healthHandler)))
	}()

	waitForShutdown()
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

// waitForShutdown blocks until the service exists.
//
// Listens on os.Interrupt. If the signal is received for a second
// time, the process is killed with status code 1.
func waitForShutdown() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	sig := <-sigint
	fmt.Printf("Received signal %v\n", sig)

	go func() {
		sig = <-sigint
		fmt.Printf("Received signal %v for the second time. Killing process.\n", sig)
		os.Exit(1)
	}()
}
