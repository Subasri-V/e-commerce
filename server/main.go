package main

import (
	"context"
	"fmt"
	"net"

	"github.com/Subasri-V/e-commerce/config"
	"github.com/Subasri-V/e-commerce/constants"
	"github.com/Subasri-V/e-commerce/services"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	controllers "github.com/Subasri-V/e-commerce/controller"
	pro "github.com/Subasri-V/e-commerce/proto"
)

func initDatabase(client *mongo.Client) {
	profileCollection := config.GetCollection(client, "Ecommerce", "CustomerProfile")
	tokenCollection := config.GetCollection(client, "Ecommerce", "Tokens")
	controllers.CustomerService = services.InitCustomerService(profileCollection, tokenCollection, context.Background())

}

func main() {
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)

	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, healthServer)
	pro.RegisterCustomerServiceServer(s, &controllers.RPCServer{})
	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
