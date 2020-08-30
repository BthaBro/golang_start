package main

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Todo structure
type Todo struct {
	ID                   int
	Title                string
	CreatedAt, ExpiresAt time.Time
	Done                 bool
}

// Todos - List of all Todos
type Todos struct {
	db *sqlx.DB
}

// TodoManager - methods on Todos structure
type TodoManager interface {
	Add()
	ChangeTitle()
	ChangeDone()
	Delete()
	ShowAllTodos()
	ShowDoneTodos()
	ShowExpiredTodos()
}

// NewTodos - constructor for Todos
func NewTodos() (*Todos, error) {
	ts := Todos{}
	var err error
	ts.db, err = sqlx.Connect("postgres", "user=postgres dbname=mydb password=123456 host=127.0.0.1")
	if err != nil {
		return nil, err
	}

	return &ts, err
}

// Add - adds new todo
func (ts *Todos) Add(t *Todo) (*Todo, error) {
	query := `INSERT INTO tasks (title, expiresat, createdat, done) values ($1, $2, $3, $4)`

	ts.db.MustExec(query, t.Title, t.ExpiresAt, t.CreatedAt, t.Done)

	rows, err := ts.db.Query(`SELECT * FROM tasks ORDER BY id DESC LIMIT 1`)

	if err != nil {
		return nil, err
	}

	var id int
	var title string
	var e, c time.Time
	var d bool

	for rows.Next() {
		err = rows.Scan(&id, &title, &d, &c, &e)
		if err != nil {
			return nil, err
		}
		t.ID = id
		t.Title = title
		t.ExpiresAt = e
		t.CreatedAt = c
		t.Done = d
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (ts *Todos) UpdateStatus(id int) (*Todo, error) {
	rows, err := ts.db.Query(`SELECT * FROM tasks WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	var title string
	var e, c time.Time
	var d bool

	for rows.Next() {
		err = rows.Scan(&id, &title, &d, &c, &e)
		if err != nil {
			return nil, err
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	query := `UPDATE tasks SET done = $1 WHERE id = $2`

	ts.db.MustExec(query, !d, id)

	t := &Todo{}
	t.Title = title
	t.Done = !d
	t.ExpiresAt = e
	t.CreatedAt = c
	t.ID = id
	return t, nil
}

func (ts *Todos) UpdateTitle(title string, id int) (*Todo, error) {
	rows, err := ts.db.Query(`SELECT * FROM tasks WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	var tit string
	var e, c time.Time
	var d bool

	for rows.Next() {
		err = rows.Scan(&id, &tit, &d, &c, &e)
		if err != nil {
			return nil, err
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	query := "UPDATE tasks SET title=$1 WHERE id=$2"

	ts.db.MustExec(query, title, id)

	t := &Todo{}
	t.Title = title
	t.Done = d
	t.ExpiresAt = e
	t.CreatedAt = c
	t.ID = id

	return t, nil
}

func (ts *Todos) Delete(id int) error {
	query := `DELETE FROM tasks WHERE id = $1`

	_, err := ts.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (ts *Todos) GetAllTodos() ([]Todo, error) {
	query := `SELECT * FROM tasks`

	rows, err := ts.db.Query(query)

	if err != nil {
		return nil, err
	}

	var id int
	var tit string
	var e, c time.Time
	var d bool
	tasks := []Todo{}
	t := &Todo{}

	for rows.Next() {
		err = rows.Scan(&id, &tit, &d, &c, &e)
		if err != nil {
			return nil, err
		}

		t.ID = id
		t.Title = tit
		t.Done = d
		t.ExpiresAt = e
		t.CreatedAt = c
		tasks = append(tasks, *t)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (ts *Todos) GetDoneTodos(done bool) ([]Todo, error) {
	query := `SELECT * FROM tasks WHERE done = $1`

	rows, err := ts.db.Query(query, done)

	if err != nil {
		return nil, err
	}

	var id int
	var tit string
	var e, c time.Time
	var d bool
	tasks := []Todo{}
	t := &Todo{}

	for rows.Next() {
		err = rows.Scan(&id, &tit, &d, &c, &e)
		if err != nil {
			return nil, err
		}

		t.ID = id
		t.Title = tit
		t.Done = d
		t.ExpiresAt = e
		t.CreatedAt = c
		tasks = append(tasks, *t)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (ts *Todos) GetExpiredTodos() ([]Todo, error) {
	todos, err := ts.GetAllTodos()
	if err != nil {
		return nil, err
	}
	tasks := []Todo{}
	timeNow := time.Now()

	for _, v := range todos {
		if v.ExpiresAt.Day() < timeNow.Day() {
			tasks = append(tasks, v)
		}
	}

	// query := `SELECT * FROM tasks WHERE expiresat = $1`

	// rows, err := ts.db.Query(query, timeNow)

	// if err != nil {
	// 	return nil, err
	// }

	// var id int
	// var tit string
	// var e, c time.Time
	// var d bool

	// for rows.Next() {
	// 	err = rows.Scan(&id, &tit, &d, &c, &e)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	t.ID = id
	// 	t.Title = tit
	// 	t.Done = d
	// 	t.ExpiresAt = e
	// 	t.CreatedAt = c
	// 	tasks = append(tasks, *t)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	return nil, err
	// }

	return tasks, nil
}

func main() {

}
