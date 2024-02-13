package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type todo struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos []todo = make([]todo, 0)

func getTodosHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Convert todos slice to JSON
		todosJSON, err := json.Marshal(todos)
		if err != nil {
			http.Error(rw, "Unable to marshal todos", http.StatusInternalServerError)
			return
		}

		// Set content type and write response
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(todosJSON)
	}
}

func createTodoHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Decode JSON request body
		var newTodo todo
		err := json.NewDecoder(r.Body).Decode(&newTodo)
		if err != nil {
			http.Error(rw, "Unable to decode JSON", http.StatusBadRequest)
			return
		}

		// Append new todo to the todos slice
		todos = append(todos, newTodo)

		// Respond with the updated list of todos
		todosJSON, err := json.Marshal(todos)
		if err != nil {
			http.Error(rw, "Unable to marshal todos", http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Write(todosJSON)
	}
}

func main() {
	// Get all todos: GET /api/todos
	http.HandleFunc("/api/todos", getTodosHandler)

	// Create a new todo: POST /api/todos
	http.HandleFunc("/api/todos/add", createTodoHandler)

	err := http.ListenAndServe("127.0.0.1:5050", nil)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Listening on 127.0.0.1:5050")
}
