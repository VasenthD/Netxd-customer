package Services

import (
	interfaces "Netxd_project/Nexd_Customer_Dal/Netxd_Customer_Interfaces"
	models "Netxd_project/Nexd_Customer_Dal/Netxd_Customer_models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}

func InitCustomer(collection *mongo.Collection, ctx context.Context) interfaces.ICustomer {
	return &CustomerService{collection, ctx}
}

func (c *CustomerService) CreateCustomer(user *models.Customer) (*models.CustomerResponse, error) {

	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.IsActive = true

	_, err := c.CustomerCollection.InsertOne(c.ctx, &user)
	if err != nil {
		return nil, err
	}
	response := &models.CustomerResponse{
		CustomerId: user.Customer_Id,
		CreatedAt:  user.CreatedAt.UTC(),
	}

	return response, nil
}
