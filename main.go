package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
)

func main() {
	muxRouter := http.NewServeMux()
	muxRouter.Handle("/", http.FileServer(http.Dir("./public")))

	chiRouter := chi.NewRouter()
	chiRouter.Handle("/", http.FileServer(http.Dir("./public")))

	sig := make(chan os.Signal, 1)
	go func() {
		http.ListenAndServe(":3001", muxRouter)
	}()

	go func() {
		http.ListenAndServe(":3002", chiRouter)
	}()

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
