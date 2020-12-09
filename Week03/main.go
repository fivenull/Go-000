package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	done := make(chan error, 1)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
	// 用于控制server shutdown
	stop := make(chan struct{})
	g := new(errgroup.Group)
	g.Go(func() error {
		return serverApp(stop)
	})
	go func() {
		done <- g.Wait()
	}()

	select {
	case sig := <-sc:
		log.Printf("receive signal %s", sig.String())
		close(stop)
	case err := <-done:
		log.Printf("serverApp error: %v", err)
	}

	log.Println("server exiting.....")
}

func serverApp(stop <-chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(rw, "Hello Golang")
	})
	s := http.Server{
		Addr:    "0.0.0.0:5454",
		Handler: mux,
	}
	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}
