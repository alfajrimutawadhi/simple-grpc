package main

import (
	"fmt"
	"log"
	"net"

	"github.com/alfajrimutawadhi/simple-grpc/todo"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	todoGrpc := todo.NewTodoGrpcHandler()
	todo.RegisterTodoServiceServer(grpcServer, todoGrpc)
	fmt.Println("grpc server is running on : 127.0.0.1:5000")
	grpcServer.Serve(listener)
}
