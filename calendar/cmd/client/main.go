package main

import (
	"calendar/pkg/api"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	x := "{\"id\":\"1\",\"timezone\":\"Europe/Riga\"}"
	conn, err := grpc.Dial(":9090", grpc.WithInsecure())

	if err != nil {
		log.Println(err)
	}

	c := api.NewGHandlersClient(conn)
	res, err := c.UpdateUser(context.Background(), &api.GRequest{X: x})

	if err != nil {
		log.Println(err)
	}

	log.Println(res)
}
