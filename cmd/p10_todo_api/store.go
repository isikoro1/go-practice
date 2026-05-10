package main

import "errors"

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
