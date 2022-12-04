package main

import (
	"UacademyGo/grpc_example/proto-gen/dice"
	boxDice "UacademyGo/grpc_example/services/dice"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


func main() {
	println("gRPC server tutorial in GO")

	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	dice.RegisterTutorialServer(s, &boxDice.TutorialService{})

	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}