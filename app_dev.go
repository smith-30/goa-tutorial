// +build app_dev

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"project/goa-tutorial/app"
	"project/goa-tutorial/controllers"
)

func main() {
	// Create service
	service := goa.New("cellar")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "bottle" controller
	c := controllers.NewBottleController(service)
	app.MountBottleController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
