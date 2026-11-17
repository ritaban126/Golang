package main

// fmt a inbuilt package which is use for printing and formating.

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/codewithdas/students-api/internal/config"
	"github.com/codewithdas/students-api/internal/http/handeler/student"
) 

func main() {
	// load config
	cfg := config.MustLoad()

	// database setup
	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	// setup server
	server := http.Server {
		Addr: cfg.Addr,
		Handler: router,
	}

	slog.Info("server started", slog.String("address", cfg.Addr))
	// fmt.Println("server started")
	fmt.Printf("server started %s", cfg.HTTPServer.Addr)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start server")
	}

	} ()

	<- done

	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	if err := server.Shutdown(ctx); err != nil{
	   slog.Error("failed to shutdown server", slog.String("error",err.Error()))
	}
slog.Info("server shutdown succesfully")

	
}