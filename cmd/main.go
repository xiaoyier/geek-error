package main

import (
	"context"
	"fmt"
	"geek-error/dao"
	"geek-error/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Starting service")

	dao.InitDB()
	mux := http.NewServeMux()
	mux.HandleFunc("/user/signup", service.UserSignUp)
	mux.HandleFunc("/user/login", service.UserLogin)
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	sig := make(chan os.Signal, 0)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		s := <-sig
		fmt.Printf("get os signal: %v", s)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		_ = server.Shutdown(ctx)
	}()

	_ = server.ListenAndServe()
	fmt.Println("Stopping service")
}
