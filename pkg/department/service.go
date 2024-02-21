package department

import (
	"github.com/fmuharam25/tutorial-golang-gofiber/api/presenter"
	"github.com/fmuharam25/tutorial-golang-gofiber/pkg/entities"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertDepartment(department *entities.Department) (*entities.Department, error)
	FetchDepartments() (*[]presenter.Department, error)
	UpdateDepartment(department *entities.Department) (*entities.Department, error)
	RemoveDepartment(ID string) error
}

type service struct {
	repository Repository
}

// NewService is used to create a single instance of the service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// InsertDepartment is a service layer that helps insert department in DepartmentShop
func (s *service) InsertDepartment(department *entities.Department) (*entities.Department, error) {
	return s.repository.CreateDepartment(department)
}

// FetchDepartments is a service layer that helps fetch all departments in DepartmentShop
func (s *service) FetchDepartments() (*[]presenter.Department, error) {
	return s.repository.ReadDepartment()
}

// UpdateDepartment is a service layer that helps update departments in DepartmentShop
func (s *service) UpdateDepartment(department *entities.Department) (*entities.Department, error) {
	return s.repository.UpdateDepartment(department)
}

// RemoveDepartment is a service layer that helps remove departments from DepartmentShop
func (s *service) RemoveDepartment(ID string) error {
	return s.repository.DeleteDepartment(ID)
}
