package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"

	_ "net/http/pprof"
)
import "net/http"

func main() {
	beforeCount := runtime.NumGoroutine()

	var mux = http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("pong")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	var srv = http.Server{
		Addr:    ":6060",
		Handler: mux,
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		// Access profiling data at http://localhost:6060/debug/pprof/goroutine
		log.Println("server is starting on http://localhost:6060")
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		stopSignal := make(chan os.Signal, 1)
		signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM)
		<-stopSignal

		log.Println("server is shutting down")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Fatal(err)
		}

		log.Println("server is shut down")
	}()

	afterCount := runtime.NumGoroutine()
	if afterCount > beforeCount {
		fmt.Printf("Possible leak: goroutine count increased from %d to %d\n",
			beforeCount, afterCount)
	}
	wg.Wait()
}
