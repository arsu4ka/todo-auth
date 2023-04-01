package main

import "github.com/arsu4ka/todo-auth/internal/controller"

func main() {
	config := controller.DefaultConfig()
	server, err := controller.NewController(config)
	if err != nil {
		panic(err)
	}

	panic(server.Start())
}
