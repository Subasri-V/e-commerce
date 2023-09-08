package interfaces

import "github.com/Subasri-V/e-commerce/models"

type ICustomer interface {
	CreateCustomer(user *models.Customer) (*models.DBResponse, error)
}
