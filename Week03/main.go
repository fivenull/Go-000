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
	g, ctx := errgroup.WithContext(context.Background())

	// 用于控制server shutdown
	stop := make(chan struct{})
	g.Go(func() error {
		return serverApp(stop)
	})

	g.Go(func() error {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case sig := <-sc:
				log.Printf("receive signal %s", sig.String())
				close(stop)
				return nil
			}
		}
	})
	if err := g.Wait(); err != nil {
		log.Printf("errgroup wait error:%v", err)
	}

	log.Println("server exiting.....")
}

func serverApp(stop <-chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(rw, "Hello Golang")
	})
	s := http.Server{
		Addr:    "0.0.0.0:5455",
		Handler: mux,
	}
	go func() {
		<-stop
		s.Shutdown(context.TODO())
	}()
	return s.ListenAndServe()
}
