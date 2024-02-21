package handlers

import (
	"net/http"

	"github.com/fmuharam25/tutorial-golang-gofiber/api/presenter"
	"github.com/fmuharam25/tutorial-golang-gofiber/pkg/department"
	"github.com/fmuharam25/tutorial-golang-gofiber/pkg/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

// AddDepartment is handler/controller which creates Departments
func AddDepartment(service department.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Department
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.DepartmentErrorResponse(err))
		}
		if requestBody.Name == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.DepartmentErrorResponse(errors.New(
				"Please specify Name")))
		}
		result, err := service.InsertDepartment(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.DepartmentErrorResponse(err))
		}
		return c.JSON(presenter.DepartmentSuccessResponse(result))
	}
}

// UpdateDepartment is handler/controller which updates data of Departments
func UpdateDepartment(service department.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Department
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.DepartmentErrorResponse(err))
		}
		result, err := service.UpdateDepartment(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.DepartmentErrorResponse(err))
		}
		return c.JSON(presenter.DepartmentSuccessResponse(result))
	}
}

// RemoveDepartment is handler/controller which removes Departments
func RemoveDepartment(service department.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.DeleteDepartmentRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.DepartmentErrorResponse(err))
		}
		departmentID := requestBody.ID
		err = service.RemoveDepartment(departmentID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.DepartmentErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "deleted successfully",
			"err":    nil,
		})
	}
}

// GetDepartments is handler/controller which lists all Departments from the DepartmentShop
func GetDepartments(service department.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchDepartments()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.DepartmentErrorResponse(err))
		}
		return c.JSON(presenter.DepartmentsSuccessResponse(fetched))
	}
}
