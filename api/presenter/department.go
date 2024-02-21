package presenter

import (
	"github.com/fmuharam25/tutorial-golang-gofiber/pkg/entities"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Department is the presenter object which will be passed in the response by Handler
type Department struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name"`
}

// DepartmentSuccessResponse is the singular SuccessResponse that will be passed in the response by
// Handler
func DepartmentSuccessResponse(data *entities.Department) *fiber.Map {
	book := Department{
		ID:   data.ID,
		Name: data.Name,
	}
	return &fiber.Map{
		"status": true,
		"data":   book,
		"error":  nil,
	}
}

// DepartmentsSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func DepartmentsSuccessResponse(data *[]Department) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// DepartmentErrorResponse is the ErrorResponse that will be passed in the response by Handler
func DepartmentErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
