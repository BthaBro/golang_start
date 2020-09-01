package models

import (
	pr "golang_start/grpc/task_manager/proto/task"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Todos - List of all Todos
type Todos struct {
	db *sqlx.DB
}

// TodoManager - methods on Todos structure
type TodoManager interface {
	Add(t pr.Todo) error
	UpdateTitle(id int64, title string) error
	UpdateStatus(id int64) error
	Delete(id int64) error
	ShowAllTodos() ([]*pr.Todo, error)
	ShowDoneTodos() ([]*pr.Todo, error)
}

// NewTodos - constructor for Todos
func NewTodos() (*Todos, error) {
	ts := Todos{}
	var err error
	ts.db, err = sqlx.Connect("postgres", "user=postgres dbname=tasklist password=123456 host=127.0.0.1")
	if err != nil {
		return nil, err
	}

	return &ts, err
}

// Add - adds new todo
func (ts *Todos) Add(t pr.Todo) error {
	query := `INSERT INTO tasks (title, expiresat, createdat, done) values ($1, $2, $3, $4)`

	ts.db.MustExec(query, t.Title, t.ExpiresAt, t.CreatedAt, t.Done)

	rows, err := ts.db.Query(`SELECT * FROM tasks ORDER BY id DESC LIMIT 1`)

	if err != nil {
		return err
	}

	var id, e, c int
	var title string
	var d bool

	for rows.Next() {
		err = rows.Scan(&id, &title, &d, &c, &e)
		if err != nil {
			return err
		}
		t.Id = int64(id)
		t.Title = title
		t.ExpiresAt = int64(e)
		t.CreatedAt = int64(c)
		t.Done = d
	}

	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

func (ts *Todos) UpdateStatus(id int64) error {
	rows, err := ts.db.Query(`SELECT * FROM tasks WHERE id = $1`, id)
	if err != nil {
		return err
	}

	var title, e, c string
	var d bool

	for rows.Next() {
		err = rows.Scan(&id, &title, &d, &c, &e)
		if err != nil {
			return err
		}
	}
	err = rows.Err()
	if err != nil {
		return err
	}

	query := `UPDATE tasks SET done = $1 WHERE id = $2`

	ts.db.MustExec(query, !d, id)

	return nil
}

func (ts *Todos) UpdateTitle(id int64, title string) error {
	query := `UPDATE tasks SET title = $1 WHERE id = $2`
	ts.db.MustExec(query, title, id)

	return nil
}

func (ts *Todos) Delete(id int64) error {
	query := `DELETE FROM tasks WHERE id = $1`

	_, err := ts.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (ts *Todos) ShowAllTodos() ([]*pr.Todo, error) {
	query := `SELECT * FROM tasks`

	var tasks []*pr.Todo

	err := ts.db.Select(&tasks, query)

	return tasks, err
}

func (ts *Todos) ShowDoneTodos() ([]*pr.Todo, error) {
	query := `SELECT * FROM tasks WHERE done = $1`

	var tasks []*pr.Todo

	err := ts.db.Select(&tasks, query, true)

	return tasks, err
}
