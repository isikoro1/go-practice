package main

// Todo はデータ構造。
// Goでは class ではなく struct を使う。
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
