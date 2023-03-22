package main

import (
	"log"
	"net"
	"net/http"
	trippb "uber/proto/gen/go"
	trip "uber/tripservice"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var port string = "localhost:3333"

func main() {
	go startGRPCGateway()

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()
	trippb.RegisterTripServiceServer(server, trip.Service{})
	log.Fatal(server.Serve(listen))
}

func startGRPCGateway() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseEnumNumbers: true,
				UseProtoNames:  false, // lower case with _ when true, by default is false
			}}),
	)
	err := trippb.RegisterTripServiceHandlerFromEndpoint(ctx, mux, port, []grpc.DialOption{grpc.WithInsecure()})

	if err != nil {
		log.Fatalln("Failed to start grpc gateway %v", err)
	}

	err = http.ListenAndServe("localhost:3000", mux)
	if err != nil {
		log.Fatalln("Failed to listen 3000 server")
	}
}
