package main

import (
	"log"
	"net"
	pb "wc_grpc/service"
	"wc_grpc/wcservice"

	"google.golang.org/grpc"
)

const port = ":50051"

//main function\
func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen0: ", err)
	}

	//Initialize new Server
	s := grpc.NewServer()

	//Regester the server as a new grpc service
	pb.RegisterWordCountServieceServer(s, &wcservice.WcServer{})
	log.Println("server listening at ", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}

}
