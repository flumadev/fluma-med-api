package main

import (
	"context"
	"log"
	"main/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	serverStatus := handlers.Status(l)
	// gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", serverStatus)
	// sm.Handle("/pacoca", gh)

	server := &http.Server{
		Addr:         ":3000",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			l.Fatal(err)
		}

	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	l.Panicln("Shutting down gracefully", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}
