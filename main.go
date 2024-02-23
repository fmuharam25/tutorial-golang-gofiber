package main

import (
	"context"
	"fmt"
	"log"
	"time"

	handlers "github.com/fmuharam25/tutorial-golang-gofiber/api/handlres"
	"github.com/fmuharam25/tutorial-golang-gofiber/api/routes"
	"github.com/fmuharam25/tutorial-golang-gofiber/pkg/department"
	"github.com/fmuharam25/tutorial-golang-gofiber/pkg/employee"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
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

	//Routes Basic
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome API V1")
	})
	app.Post("/login", handlers.Login)
	app.Post("/logout", handlers.Logout)

	//Route with basic auth
	apiv1 := app.Group("/api/v1")
	routes.DepartmentRouter(apiv1, departmentService,
		basicauth.New(basicauth.Config{
			Users: map[string]string{
				"admin": "123456",
			},
		}))

	//Route with JWT auth
	apiv2 := app.Group("/api/v2")
	routes.EmployeeRouter(apiv2, employeeService,
		jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: []byte("secret")},
		}))

	//Defer error
	defer cancel()
	log.Fatal(app.Listen(":8080"))

}

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017").SetServerSelectionTimeout(5*time.Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database("company")
	return db, cancel, nil
}
