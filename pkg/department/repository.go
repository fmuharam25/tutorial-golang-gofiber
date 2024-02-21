package department

import (
	"context"
	"time"

	"github.com/fmuharam25/tutorial-golang-gofiber/api/presenter"
	"github.com/fmuharam25/tutorial-golang-gofiber/pkg/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	CreateDepartment(department *entities.Department) (*entities.Department, error)
	ReadDepartment() (*[]presenter.Department, error)
	UpdateDepartment(department *entities.Department) (*entities.Department, error)
	DeleteDepartment(ID string) error
}
type repository struct {
	Collection *mongo.Collection
}

// NewRepo is the single instance repo that is being created.
func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

// CreateDepartment is a mongo repository that helps to create departments
func (r *repository) CreateDepartment(department *entities.Department) (*entities.Department, error) {
	department.ID = primitive.NewObjectID()
	department.CreatedAt = time.Now()
	department.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), department)
	if err != nil {
		return nil, err
	}
	return department, nil
}

// ReadDepartment is a mongo repository that helps to fetch departments
func (r *repository) ReadDepartment() (*[]presenter.Department, error) {
	var departments []presenter.Department
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var department presenter.Department
		_ = cursor.Decode(&department)
		departments = append(departments, department)
	}
	return &departments, nil
}

// UpdateDepartment is a mongo repository that helps to update departments
func (r *repository) UpdateDepartment(department *entities.Department) (*entities.Department, error) {
	department.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": department.ID}, bson.M{"$set": department})
	if err != nil {
		return nil, err
	}
	return department, nil
}

// DeleteDepartment is a mongo repository that helps to delete departments
func (r *repository) DeleteDepartment(ID string) error {
	departmentID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": departmentID})
	if err != nil {
		return err
	}
	return nil
}
