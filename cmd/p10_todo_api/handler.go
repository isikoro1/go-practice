package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type TodoHandler struct {
	store *TodoStore
}

func NewTodoHandler(store *TodoStore) *TodoHandler {
	return &TodoHandler{store: store}
}

func (h *TodoHandler) HandleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		todos := h.store.List()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)

	case http.MethodPost:
		var requestBody struct {
			Title string `json:"title"`
		}

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil || requestBody.Title == "" {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		todo := h.store.Add(requestBody.Title)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(todo)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)

	}
}

func (h *TodoHandler) HandleTodoByID(w http.ResponseWriter, r *http.Request) {
	idText := r.URL.Path[len("/todos/"):]

	id, err := strconv.Atoi(idText)
	if err != nil {
		http.Error(w, "invalid todo id", http.StatusBadRequest)
		return
	}

	todo, err := h.store.Find(id)
	if err != nil {
		http.Error(w, "todo not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
