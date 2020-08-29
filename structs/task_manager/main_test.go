package main

import (
	"testing"
	"time"
)

var ts = NewTodos()

func TestTodosAdd(t *testing.T) {
	v := &Todo{Title: "New todo", ID: 0, CreatedAt: time.Now(), ExpiresAt: time.Now().AddDate(0, 0, 2), Done: false}
	added := ts.Add(v)
	if added.Title != v.Title {
		t.Error("Error while adding todo")
	}
}

func TestTodosChangeTitle(t *testing.T) {
	title := "New todo updated"
	upd := ts.ChangeTitle(0, title)
	if upd.Title != title {
		t.Error("Error while updating title")
	}
}

func TestTodosChangeDone(t *testing.T) {
	id := 0
	ts.ChangeDone(id)
	if true != ts.todos[id].Done {
		t.Error("Error while changing status")
	}
}

func TestTodosShowDoneTodos(t *testing.T) {
	done := ts.ShowDoneTodos(true)
	if len(done.todos) != 1 {
		t.Error("Error while showing done todos")
	}
}

func TestTodosShowExpiredTodos(t *testing.T) {
	v := &Todo{Title: "New todo", ID: 0, CreatedAt: time.Now().AddDate(0, 0, -2), ExpiresAt: time.Now().AddDate(0, 0, -1), Done: false}
	ts.Add(v)
	expired := ts.ShowExpiredTodos()
	if len(expired.todos) != 1 {
		t.Error("Error while showing expired todos")
	}
}

func TestTodosShowAllTodos(t *testing.T) {
	todos := ts.ShowAllTodos()
	if len(ts.todos) != len(todos.todos) {
		t.Error("Error while showing all todos")
	}
}

func TestTodosDelete(t *testing.T) {
	id := 0
	ts.Delete(id)
	if len(ts.todos) != 1 {
		t.Error("Error while deleting the todo")
	}
}
