package main

import "github.com/arsu4ka/todo-auth/internal/server"

func main() {
	config := server.DefaultConfig()
	server, err := server.NewServer(config)
	if err != nil {
		panic(err)
	}

	panic(server.Start())
}
