package main

import (
	"fmt"

	// "os/user"

	"github.com/codewithdas/podcast/auth"
	"github.com/codewithdas/podcast/user"
)

// to install any package the use (go get then package name(github.com/faith/colour)) like jaisa npm in jsor ts then for go use "go get"
// go mod tidy -> command use hoti hai jab koi package ko use karti hai without isntall in main.gpo file thne this command is help for install

// imp command for create a packages like this go mod init github.com/codewithdas/podcast

func main() {
	auth.LoginCredintials("codersgyan","secrect")
	session := auth.GetSession()
	fmt.Println("session",session)

	user := user.User{
		Email: "user@gmail.com",
		// Name: "Jhon Doe",
	}
	fmt.Println(user.Email, user.Name)
	// fmt.Println("\033[31m%s\033[0m\n",user.Email)

	
}