package tripletex

import (
	"time"

	"github.com/bjerkio/tripletex-go/client/employee"
	"github.com/bjerkio/tripletex-go/models"
)

// GetEmployees returns a list of Tripletex employees
func (c *TripletexClient) GetEmployees() ([]*models.Employee, error) {

	req := employee.EmployeeSearchParams{}

	res, err := c.client.Employee.EmployeeSearch(req.WithTimeout(10*time.Second), c.authInfo)
	if err != nil {
		return nil, err
	}

	return res.Payload.Values, nil
}
