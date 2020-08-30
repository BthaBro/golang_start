package main

import (
	"fmt"
	"testing"
	"time"
)

var ts, err = NewTodos()

func TestTodosAdd(t *testing.T) {
	v := &Todo{Title: "Expired", Done: false, CreatedAt: time.Now(), ExpiresAt: time.Now().AddDate(0, 0, 2)}
	added, err := ts.Add(v)

	if err != nil || added.Title != v.Title {
		t.Error("Error while adding task")
	}
}

func TestTodosUpdateStatus(t *testing.T) {
	id := 2
	_, err := ts.UpdateStatus(id)

	if err != nil {
		t.Error("Error while updating status")
	}
}

func TestTodosUpdateTitle(t *testing.T) {
	id := 2
	title := "New todo here"
	upd, err := ts.UpdateTitle(title, id)
	if err != nil || upd.Title != title {
		t.Error("Error while updating title")
	}
}

func TestTodosDelete(t *testing.T) {
	id := 2
	err := ts.Delete(id)
	if err != nil {
		t.Error("Error while deleting task")
	}
}

func TestTodosGetALlTodos(t *testing.T) {
	todos, err := ts.GetAllTodos()
	fmt.Println(todos)
	if err != nil {
		t.Error("Error while getting all tasks")
	}
}

func TestTodosGetDoneTodos(t *testing.T) {
	done := true
	_, err := ts.GetDoneTodos(done)
	if err != nil {
		t.Error("Error while getting done tasks")
	}
}

func TestTodosGetExpiredTodos(t *testing.T) {
	exp, err := ts.GetExpiredTodos()
	if err != nil || len(exp) != 1 {
		t.Error("Error while getting expired todos")
	}
}
