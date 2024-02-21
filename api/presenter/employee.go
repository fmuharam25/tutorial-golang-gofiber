package presenter

import (
	"github.com/fmuharam25/tutorial-golang-gofiber/pkg/entities"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Employee is the presenter object which will be passed in the response by Handler
type Employee struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `json:"name"`
	DepartmentID primitive.ObjectID `json:"department_id" bson:"department_id"`
}

// EmployeeSuccessResponse is the singular SuccessResponse that will be passed in the response by
// Handler
func EmployeeSuccessResponse(data *entities.Employee) *fiber.Map {
	book := Employee{
		ID:           data.ID,
		Name:         data.Name,
		DepartmentID: data.DepartmentID,
	}
	return &fiber.Map{
		"status": true,
		"data":   book,
		"error":  nil,
	}
}

// EmployeesSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func EmployeesSuccessResponse(data *[]Employee) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// EmployeeErrorResponse is the ErrorResponse that will be passed in the response by Handler
func EmployeeErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
