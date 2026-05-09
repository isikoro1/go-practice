package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

// Todo はデータ構造。
// Goでは class ではなく struct を使う。
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// TodoStore は Todo を管理する構造体。
// 本来はDBを使うが、今回はメモリ上の map で管理する。
type TodoStore struct {
	todos  map[int]Todo
	nextID int
}

// NewTodoStore は初期化用の関数。
// Goでは constructor という専用構文はなく、普通の関数で作る。
func NewTodoStore() *TodoStore {
	return &TodoStore{
		todos:  make(map[int]Todo),
		nextID: 1,
	}
}

// Add は TodoStore に紐づくメソッド。
// Goでは func の後ろに receiver を書く。
func (s *TodoStore) Add(title string) Todo {
	todo := Todo{
		ID:    s.nextID,
		Title: title,
		Done:  false,
	}

	s.todos[todo.ID] = todo
	s.nextID++

	return todo
}

// Find は id で Todo を探す。
// 見つからない場合は error を返す。
// Goでは例外 throw ではなく、戻り値で error を返すのが基本。
func (s *TodoStore) Find(id int) (Todo, error) {
	todo, ok := s.todos[id]
	if !ok {
		return Todo{}, errors.New("todo not found")
	}

	return todo, nil
}

// List は Todo一覧を返す。
func (s *TodoStore) List() []Todo {
	result := []Todo{}

	for _, todo := range s.todos {
		result = append(result, todo)
	}

	return result
}

func main() {
	store := NewTodoStore()

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			todos := store.List()

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

			todo := store.Add(requestBody.Title)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(todo)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		idText := r.URL.Path[len("/todos/"):]

		id, err := strconv.Atoi(idText)
		if err != nil {
			http.Error(w, "invalid todo id", http.StatusBadRequest)
			return
		}

		todo, err := store.Find(id)
		if err != nil {
			http.Error(w, "todo not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todo)
	})

	log.Println("server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
