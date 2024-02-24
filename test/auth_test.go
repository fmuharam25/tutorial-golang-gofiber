package test

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/fmuharam25/tutorial-golang-gofiber/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRoute(t *testing.T) {
	// Define a structure for specifying input and output data
	// of a single test case
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
		method       string // method for the test
		body         string //  example"name=foo&surname=bar"
	}{
		// Welcome Route
		{
			description:  "get HTTP status 200",
			route:        "/",
			expectedCode: 200,
			method:       "GET",
		},
		// Login Route
		{
			description:  "post HTTP status 200",
			route:        "/login",
			expectedCode: 200,
			method:       "POST",
			body:         "user=admin&pass=123456",
		},
		// Logout Route
		{
			description:  "post HTTP status 200",
			route:        "/logout",
			expectedCode: 200,
			method:       "POST",
		},
		// Not Found Route
		{
			description:  "get HTTP status 404, when route is not exists",
			route:        "/not-found",
			expectedCode: 404,
			method:       "GET",
		},
		// Invalid Method
		{
			description:  "get HTTP status 405, when method is invalid",
			route:        "/login",
			expectedCode: 405,
			method:       "GET",
		},
	}

	// Define Fiber app.
	app := fiber.New()

	// Create route
	routes.DefaultRoute(app)

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case
		req := httptest.NewRequest(test.method, test.route, bytes.NewBufferString(test.body))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := app.Test(req, 1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
