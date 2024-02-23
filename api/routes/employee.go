package routes

import (
	handlers "github.com/fmuharam25/tutorial-golang-gofiber/api/handlres"
	"github.com/fmuharam25/tutorial-golang-gofiber/pkg/employee"
	"github.com/gofiber/fiber/v2"
)

// EmployeeRouter is the Router for GoFiber App
func EmployeeRouter(app fiber.Router, service employee.Service, middleware fiber.Handler) {
	app.Use(middleware)
	app.Get("/employees", handlers.GetEmployees(service))
	app.Post("/employees", handlers.AddEmployee(service))
	app.Put("/employees", handlers.UpdateEmployee(service))
	app.Delete("/employees", handlers.RemoveEmployee(service))
}
