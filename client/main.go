package main

import (
	"context"
	"fmt"
	"log"
	trippb "uber/proto/gen/go"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:3333", grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		log.Fatal("Cannot connect to server %v", err)
	}

	tsClient := trippb.NewTripServiceClient(conn)
	response, err := tsClient.GetTrip(context.Background(), &trippb.GetTripRequest{Id: "trip456"})

	if err != nil {
		log.Fatal("Cannot call GetTrip %v", err)
	}

	fmt.Println(response)
}
