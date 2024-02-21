package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Employee Constructs your employee model under entities.
type Employee struct {
	ID           primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name"`
	DepartmentID primitive.ObjectID `json:"department_id" bson:"department_id"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// DeleteRequest struct is used to parse Delete Requests for Employees
type DeleteEmployeeRequest struct {
	ID string `json:"id"`
}
