package handlers

import (
	"net/http"

	"github.com/fmuharam25/tutorial-golang-gofiber/api/presenter"
	"github.com/fmuharam25/tutorial-golang-gofiber/pkg/employee"
	"github.com/fmuharam25/tutorial-golang-gofiber/pkg/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

// AddEmployee is handler/controller which creates Employees
func AddEmployee(service employee.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Employee
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}
		if requestBody.Name == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.EmployeeErrorResponse(errors.New(
				"Please specify Name")))
		}
		result, err := service.InsertEmployee(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}
		return c.JSON(presenter.EmployeeSuccessResponse(result))
	}
}

// UpdateEmployee is handler/controller which updates data of Employees
func UpdateEmployee(service employee.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Employee
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}
		result, err := service.UpdateEmployee(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}
		return c.JSON(presenter.EmployeeSuccessResponse(result))
	}
}

// RemoveEmployee is handler/controller which removes Employees
func RemoveEmployee(service employee.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.DeleteEmployeeRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}
		employeeID := requestBody.ID
		err = service.RemoveEmployee(employeeID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "deleted successfully",
			"err":    nil,
		})
	}
}

// GetEmployees is handler/controller which lists all Employees from the EmployeeShop
func GetEmployees(service employee.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchEmployees()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}
		return c.JSON(presenter.EmployeesSuccessResponse(fetched))
	}
}
