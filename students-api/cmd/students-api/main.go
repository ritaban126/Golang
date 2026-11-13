package main

// fmt a inbuilt package which is use for printing and formating.

import (
	"net/http"

	"github.com/codewithdas/students-api/internal/config"
) 

func main() {
	// load config
	cfg := config.MustLoad()

	// database setup
	// setup router
	http.NewServerMax()


	// setup server
}