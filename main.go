package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fmuharam25/tutorial-golang-gofiber/api/routes"
	"github.com/fmuharam25/tutorial-golang-gofiber/pkg/department"
	"github.com/fmuharam25/tutorial-golang-gofiber/pkg/employee"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	db, cancel, err := databaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!")
	departmentCollection := db.Collection("departments")
	departmentRepo := department.NewRepo(departmentCollection)
	departmentService := department.NewService(departmentRepo)

	employeeCollection := db.Collection("employees")
	employeeRepo := employee.NewRepo(employeeCollection)
	employeeService := employee.NewService(employeeRepo)

	//Welcome Routes
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the API V1 Department and Employee!"))
	})

	//API Route
	api := app.Group("/api")
	routes.DepartmentRouter(api, departmentService)
	routes.EmployeeRouter(api, employeeService)
	defer cancel()
	log.Fatal(app.Listen(":8080"))
}

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://username:password@localhost:27017/fiber").SetServerSelectionTimeout(5*time.Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database("departments")
	return db, cancel, nil
}
