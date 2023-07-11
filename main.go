package main

import (
	"fmt"
	"log"
	"net"

	"github.com/adityarifqyfauzan/user-services/config"
	"github.com/adityarifqyfauzan/user-services/routes"
	"google.golang.org/grpc"
)

func main() {

	db := config.InitPostgresDB()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.Env("APP_PORT")))
	if err != nil {
		log.Fatalf("unable to listen server: %v", err)
	}

	server := grpc.NewServer()

	routes.InitServices(server, db)

	log.Printf("Server running at %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("unable to start the server: %v", err)
	}

}
