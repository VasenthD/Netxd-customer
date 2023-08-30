package interfaces

import models "Netxd_project/Nexd_Customer_Dal/Netxd_Customer_models"

type ICustomer interface{
	CreateCustomer (customer *models.Customer)(*models.CustomerResponse,error)
}
