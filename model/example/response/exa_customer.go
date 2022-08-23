package response

import "github.com/opisnoeasy/course-service/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
