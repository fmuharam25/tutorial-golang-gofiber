package employee

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
	CreateEmployee(employee *entities.Employee) (*entities.Employee, error)
	ReadEmployee() (*[]presenter.Employee, error)
	UpdateEmployee(employee *entities.Employee) (*entities.Employee, error)
	DeleteEmployee(ID string) error
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

// CreateEmployee is a mongo repository that helps to create employees
func (r *repository) CreateEmployee(employee *entities.Employee) (*entities.Employee, error) {
	employee.ID = primitive.NewObjectID()
	employee.DepartmentID = primitive.NewObjectID()
	employee.CreatedAt = time.Now()
	employee.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), employee)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

// ReadEmployee is a mongo repository that helps to fetch employees
func (r *repository) ReadEmployee() (*[]presenter.Employee, error) {
	var employees []presenter.Employee
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var employee presenter.Employee
		_ = cursor.Decode(&employee)
		employees = append(employees, employee)
	}
	return &employees, nil
}

// UpdateEmployee is a mongo repository that helps to update employees
func (r *repository) UpdateEmployee(employee *entities.Employee) (*entities.Employee, error) {
	employee.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": employee.ID}, bson.M{"$set": employee})
	if err != nil {
		return nil, err
	}
	return employee, nil
}

// DeleteEmployee is a mongo repository that helps to delete employees
func (r *repository) DeleteEmployee(ID string) error {
	employeeID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": employeeID})
	if err != nil {
		return err
	}
	return nil
}
