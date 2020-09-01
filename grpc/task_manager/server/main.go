package main

import (
	"context"
	"fmt"
	"golang_start/grpc/task_manager/models"
	proto "golang_start/grpc/task_manager/proto/task"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	ts models.TodoManager
}

func main() {
	listener, err := net.Listen("tcp", ":3030")
	if err != nil {
		panic(err)
	}

	ts, err := models.NewTodos()
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterTodoManagerServer(srv, &server{ts})
	reflection.Register(srv)
	fmt.Println("Server is running on port 3030 ...")
	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *server) Add(ctx context.Context, t *proto.Todo) (*proto.Todo, error) {
	err := s.ts.Add(*t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (s *server) UpdateTitle(ctx context.Context, t *proto.Todo) (*proto.Status, error) {
	err := s.ts.UpdateTitle(t.GetId(), t.GetTitle())
	if err != nil {
		return nil, err
	}

	return &proto.Status{Status: "success"}, nil
}

func (s *server) UpdateStatus(ctx context.Context, id *proto.TodoId) (*proto.Status, error) {
	err := s.ts.UpdateStatus(id.GetId())
	if err != nil {
		return nil, err
	}

	return &proto.Status{Status: "success"}, nil
}

func (s *server) Delete(ctx context.Context, id *proto.TodoId) (*proto.Status, error) {
	err := s.ts.Delete(id.GetId())
	if err != nil {
		return nil, err
	}

	return &proto.Status{Status: "success"}, nil
}

func (s *server) ShowAllTodos(ctx context.Context, e *empty.Empty) (*proto.Todos, error) {
	todos, err := s.ts.ShowAllTodos()
	if err != nil {
		return nil, err
	}

	return &proto.Todos{Todos: todos}, nil
}

func (s *server) ShowDoneTodos(ctx context.Context, e *empty.Empty) (*proto.Todos, error) {
	todos, err := s.ts.ShowDoneTodos()
	if err != nil {
		return nil, err
	}

	return &proto.Todos{Todos: todos}, nil
}
