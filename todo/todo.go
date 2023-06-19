package todo

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Todo struct {
	UnimplementedTodoServiceServer
	Id    int
	Title string
}

func NewTodoGrpcHandler() TodoServiceServer {
	return &Todo{}
}

func (td *Todo) GetTodo(ctx context.Context, req *emptypb.Empty) (*TodoContent, error) {
	return &TodoContent{
		Id:    1,
		Title: "simple gRPC",
	}, nil
}
