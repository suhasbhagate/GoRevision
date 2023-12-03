package main

import (
	"log"
	"net"

	"github.com/suhasbhagate/GoCode/GoCodeRevision/GRPCDemo/handler"
	"github.com/suhasbhagate/GoCode/GoCodeRevision/GRPCDemo/pb"
	"google.golang.org/grpc"
)

func main() {
	listner, err := net.Listen("tcp", "localhost:2700")
	if err != nil {
		log.Println(err)
	}
	s := grpc.NewServer()
	pb.RegisterServiceServer(s, &handler.Server{})
	err = s.Serve(listner)
	if err != nil {
		log.Println(err)
	}
}
