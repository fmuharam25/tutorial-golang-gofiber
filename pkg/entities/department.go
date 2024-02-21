package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Department Constructs your department model under entities.
type Department struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// DeleteRequest struct is used to parse Delete Requests for Departments
type DeleteDepartmentRequest struct {
	ID string `json:"id"`
}
