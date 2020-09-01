package main

import (
	"context"
	"fmt"
	proto "golang_start/grpc/task_manager/proto/task"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var client proto.TodoManagerClient

func TestAdd(t *testing.T) {
	conn, err := grpc.Dial("localhost:3030", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client = proto.NewTodoManagerClient(conn)

	task := &proto.Todo{
		Title:     "Finish grpc",
		Done:      false,
		CreatedAt: int64(time.Now().Day()),
		ExpiresAt: int64(time.Now().AddDate(0, 0, 1).Day()),
	}

	res, err := client.Add(context.Background(), task)
	if err != nil {
		t.Error("Error while adding task: ", err)
	}
	fmt.Println("Response:", res)
}

func TestUpdateTitle(t *testing.T) {
	conn, err := grpc.Dial("localhost:3030", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client = proto.NewTodoManagerClient(conn)

	task := &proto.Todo{
		Id:    1,
		Title: "Finish grpc updated",
	}

	res, err := client.UpdateTitle(context.Background(), task)
	if err != nil {
		t.Error("Error while updating title: ", err)
	}
	fmt.Println("Response:", res)
}

func TestUpdateDone(t *testing.T) {
	conn, err := grpc.Dial("localhost:3030", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client = proto.NewTodoManagerClient(conn)
	id := &proto.TodoId{
		Id: 1,
	}
	res, err := client.UpdateStatus(context.Background(), id)
	if err != nil {
		t.Error("Error while updating title: ", err)
	}
	fmt.Println("Response:", res)
}

func TestShowAllTodos(t *testing.T) {
	conn, err := grpc.Dial("localhost:3030", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client = proto.NewTodoManagerClient(conn)
	res, err := client.ShowAllTodos(context.Background(), &empty.Empty{})
	if err != nil {
		t.Error("Error while updating title: ", err)
	}
	fmt.Println("Response:", res)
}

func TestShowDoneTodos(t *testing.T) {
	conn, err := grpc.Dial("localhost:3030", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client = proto.NewTodoManagerClient(conn)
	res, err := client.ShowDoneTodos(context.Background(), &empty.Empty{})
	if err != nil {
		t.Error("Error while updating title: ", err)
	}
	fmt.Println("Response:", res)
}
