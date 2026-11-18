package main

// fmt a inbuilt package which is use for printing and formating.

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/codewithdas/students-api/internal/config"
) 

func main() {
	// load config
	cfg := config.MustLoad()

	// database setup
	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("welcome to students api"))
	})
	// setup server
	server := http.Server {
		Addr: cfg.Addr,
		Handler: router,
	}
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

	
}