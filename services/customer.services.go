package services

import (
	"context"

	"github.com/Subasri-V/e-commerce/config"
	"github.com/Subasri-V/e-commerce/interfaces"
	"github.com/Subasri-V/e-commerce/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type CustomerService struct {
	ProfileCollection *mongo.Collection
	tokenCollection   *mongo.Collection
	ctx               context.Context
}

func InitCustomerService(collection, tokenCollection *mongo.Collection, ctx context.Context) interfaces.ICustomer {
	return &CustomerService{collection, tokenCollection, ctx}
}

func (p *CustomerService) CreateCustomer(user *models.Customer) (*models.DBResponse, error) {

	res, err := p.ProfileCollection.InsertOne(p.ctx, &user)
	if err != nil {
		return nil, err
	}
	mongoclient, _ := config.ConnectDataBase()
	collection := mongoclient.Database("Ecommerce").Collection("CustomerProfile")
	query := bson.M{"customerid": user.CustomerId}
	var customer models.Customer
	err2 := collection.FindOne(p.ctx, query).Decode(&customer)
	if err != nil {
		return nil, err2
	}

	var newUser models.DBResponse
	query2 := bson.M{"_id": res.InsertedID}
	err = p.ProfileCollection.FindOne(p.ctx, query2).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}
