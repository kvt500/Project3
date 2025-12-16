package main

import (
	"fmt"
	"restapi/http"
	"restapi/todo"
)

func main() {
	todoList := todo.NewList()
    //coment
    httpHandlers := http.NewHTTPHandlers(todoList)
	httpServer := http.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server:", err)
	}
}
