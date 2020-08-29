package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
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
	todos []Todo
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

// NewTodo - constructor for Todo
func NewTodo() Todo {
	t := Todo{}
	t.Done = false
	return t
}

// NewTodos - constructor for Todos
func NewTodos() Todos {
	ts := Todos{}
	ts.todos = []Todo{}
	return ts
}

// Add - adds new todo
func (ts *Todos) Add(t *Todo) *Todo {
	t.ID = len(ts.todos)
	ts.todos = append(ts.todos, *t)
	fmt.Println("Task has been added successfully.")
	return &ts.todos[t.ID]
}

// ChangeTitle - changes title
func (ts *Todos) ChangeTitle(id int, title string) *Todo {
	t := &ts.todos[id]
	t.Title = title
	fmt.Println("Title has been successfully changed")
	return t
}

func EnterTitle(t *Todo) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter text to change title: ")
	title, _ := reader.ReadString('\n')

	t.Title = title
}

// ChangeDone - changes done status
func (ts *Todos) ChangeDone(id int) {
	t := &ts.todos[id]
	t.Done = !t.Done

	fmt.Println("Status has been successfully changed")
}

// Delete - deletes todo by id
func (ts *Todos) Delete(id int) {
	for _, v := range ts.todos {
		if v.ID == id {
			ts.todos = append(ts.todos[:id], ts.todos[id+1:]...)
			fmt.Println("Task has been successfully deleted")
			return
		}
	}
}

// ShowDoneTodos - shows list of done/undone todos
func (ts *Todos) ShowDoneTodos(d bool) *Todos {
	todos := Todos{}
	for _, v := range ts.todos {
		if v.Done == d {
			fmt.Printf("ID: %v\n", v.ID)
			fmt.Printf("Task: %v", v.Title)
			if d == true {
				fmt.Println("Status: Done")
			} else {
				fmt.Println("Status: Not finished")
			}
			fmt.Printf("Created at: %v\n", v.CreatedAt.Day())
			fmt.Printf("Deadline: %v\n", v.ExpiresAt.Day())
			fmt.Println("---------------------------")
			todos.todos = append(todos.todos, v)
		}
	}
	return &todos
}

// ShowExpiredTodos - shows all todos that are expired
func (ts *Todos) ShowExpiredTodos() *Todos {
	todos := Todos{}
	for _, v := range ts.todos {
		if v.ExpiresAt.Before(time.Now()) {
			fmt.Printf("ID: %v\n", v.ID)
			fmt.Printf("Task: %v\n", v.Title)
			if v.Done == true {
				fmt.Println("Status: Done")
			} else {
				fmt.Println("Status: Not finished")
			}
			fmt.Printf("Created at: %v\n", v.CreatedAt.Day())
			fmt.Printf("Deadline: %v\n", v.ExpiresAt.Day())
			fmt.Println("---------------------------")
			todos.todos = append(todos.todos, v)
		}
	}
	return &todos
}

// GetAllTodos - show all todos
func (ts *Todos) ShowAllTodos() *Todos {
	for _, v := range ts.todos {
		fmt.Printf("ID: %v\n", v.ID)
		fmt.Printf("Task: %v", v.Title)
		if v.Done == true {
			fmt.Println("Status: Done")
		} else {
			fmt.Println("Status: Not finished")
		}
		fmt.Printf("Created at: %v\n", v.CreatedAt.Day())
		fmt.Printf("Deadline: %v\n", v.ExpiresAt.Day())
		fmt.Println("---------------------------")
	}
	return ts
}

func EnterDetails(t *Todo) {
	var day int
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please, enter following details")
	fmt.Println("Title: ")
	title, _ := reader.ReadString('\n')
	fmt.Println("Deadline is after how many days?: ")
	fmt.Scanln(&day)
	t.Title = title
	t.ExpiresAt = time.Now().AddDate(0, 0, day)
	t.CreatedAt = time.Now()
}

func EnterId(ts *Todos) int {
	var id int

	// Get id
	fmt.Println("Please enter the id of the task you would like to change title: ")
	fmt.Scanln(&id)
	if id < len(ts.todos) {
		return id
	}
	fmt.Println("ID doesn't exist in task list")
	return -1
}

// Menu - outputs menu
func Menu() {
	fmt.Println("********************")
	fmt.Println("Task Manager")
	fmt.Println("********************")
	fmt.Println("1. Add new task")
	fmt.Println("2. Change title of task")
	fmt.Println("3. Change status of task")
	fmt.Println("4. Delete task")
	fmt.Println("5. Show all tasks")
	fmt.Println("6. Show not finished tasks")
	fmt.Println("7. Show finished tasks")
	fmt.Println("8. Show expired tasks")
	fmt.Println("0. Exit")
	fmt.Println("\nEnter your choice: ")
}

func main() {
	// 	var ts Todos
	// 	var t Todo
	// 	var i int

	// F:
	// 	for {
	// 		Menu()
	// 		fmt.Scanln(&i)

	// 		switch i {
	// 		case 1:
	// 			ts.Add(&t)
	// 		// case 2:
	// 		// 	ts.ChangeTitle(&t]]\)
	// 		case 3:
	// 			ts.ChangeDone(&t)
	// 		case 4:
	// 			ts.Delete()
	// 		case 5:
	// 			ts.GetAllTodos()
	// 		case 6:
	// 			ts.ShowDoneTodos(false)
	// 		case 7:
	// 			ts.ShowDoneTodos(true)
	// 		case 8:
	// 			ts.ShowExpiredTodos()
	// 		case 0:
	// 			break F
	// 		}
	// 	}

}
