package routes

import (
	handlers "github.com/fmuharam25/tutorial-golang-gofiber/api/handlres"
	"github.com/fmuharam25/tutorial-golang-gofiber/pkg/department"
	"github.com/gofiber/fiber/v2"
)

// DepartmentRouter is the Router for GoFiber App
func DepartmentRouter(app fiber.Router, service department.Service, middleware fiber.Handler) {
	app.Use(middleware)
	app.Get("/departments", handlers.GetDepartments(service))
	app.Post("/departments", handlers.AddDepartment(service))
	app.Put("/departments", handlers.UpdateDepartment(service))
	app.Delete("/departments", handlers.RemoveDepartment(service))
}
