package main

import (
	grpc2 "calendar/internal/grpc"
	"calendar/internal/repository"
	storage "calendar/internal/repository/postgre"
	"log"
	"net"

	"calendar/pkg/api"
	"google.golang.org/grpc"
)

func main() {

	pgConn := storage.New()
	s := &storage.Repository{Pool: pgConn}

	str := *repository.NewStorage(s)

	server := grpc.NewServer()
	srv := &grpc2.SHandler{
		Storage: str,
	}
	api.RegisterGHandlersServer(server, srv)

	l, err := net.Listen("tcp", ":9090")

	if err != nil {
		log.Println(err)
	}

	err = server.Serve(l)
	if err != nil {
		log.Println(err)
	}
}
