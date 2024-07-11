package main

import (
	"log"
	"net"
	"os"

	"github.com/bobadojo/go/pkg/menus/v1/menuspb"
	"google.golang.org/grpc"
)

type menusServer struct {
	menuspb.UnimplementedMenusServer
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listening on port %s", port)
	grpcServer := grpc.NewServer()
	menuspb.RegisterMenusServer(grpcServer, &menusServer{})
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
