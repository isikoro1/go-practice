package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewTodoStore()
	handler := NewTodoHandler(store)

	http.HandleFunc("/todos", handler.HandleTodos)
	http.HandleFunc("/todos/", handler.HandleTodoByID)

	log.Println("server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
