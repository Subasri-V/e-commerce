package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	pb "github.com/Subasri-V/e-commerce/proto"
	"github.com/Subasri-V/e-commerce/services"
)

// var (
// 	mongoclient *mongo.Client
// 	ctx         context.Context
// 	server      *gin.Engine
// )

func main() {
	r := gin.Default()
	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)
	r.POST("/signup", func(c *gin.Context) {
		var request pb.CustomerDetails

		// Parse incoming JSON
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		pass, err2 := services.HashPassword(request.Password)
		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
			return
		}
		request.Password = pass

		// Call the gRPC service
		response, err := client.CreateCustomer(c.Request.Context(), &request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"value": response})
	})
	r.Run(":8080")
}
