package controller

import (
	"context"
	"fmt"

	"github.com/Subasri-V/e-commerce/interfaces"
	"github.com/Subasri-V/e-commerce/models"

	pro "github.com/Subasri-V/e-commerce/proto"
)

type RPCServer struct {
	pro.UnimplementedCustomerServiceServer
}

var (
	//ctx             gin.Context
	CustomerService interfaces.ICustomer
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *pro.CustomerDetails) (*pro.CustomerResponse, error) {

	// var address models.Address
	// if req != nil {
	// 	address = models.Address{
	// 		Country: req.Address[0].Country,
	// 		Street1: req.Address[0].Street1,
	// 		Street2: req.Address[0].Street2,
	// 		City:    req.Address[0].City,
	// 		State:   req.Address[0].State,
	// 		Zip:     req.Address[0].Zip,
	// 	}
	// }

	// var shippingAddress models.ShippingAddress
	// if req != nil {
	// 	shippingAddress = models.ShippingAddress{
	// 		Street1: req.ShippingAddress[0].Street1,
	// 		Street2: req.ShippingAddress[0].Street2,
	// 		City:    req.ShippingAddress[0].City,
	// 		State:   req.ShippingAddress[0].State,
	// 	}
	// }

	dbCustomer := models.Customer{
		CustomerId: req.CustomerId,
		Firstname:  req.Firstname,
		Lastname:   req.Lastname,
		Password:   req.Password,
		Email:      req.Email,
		// Address:         []models.Address{address},
		// ShippingAddress: []models.ShippingAddress{shippingAddress},
	}
	fmt.Println(dbCustomer.Password)
	result, err := CustomerService.CreateCustomer(&dbCustomer)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &pro.CustomerResponse{
			CustomerId: result.CustomerId,
		}
		return responseCustomer, nil
	}
}
