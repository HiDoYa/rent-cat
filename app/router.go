package app

import (
	"github.com/HiDoYa/rent-cat/app/api"
	// Required to initialize swagger docs
	_ "github.com/HiDoYa/rent-cat/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Launch ...
func Launch() {
	router := gin.Default()

	// Serve swagger doc & ui
	swagHandler := ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("/swagger/doc.json"),
		ginSwagger.InstanceName("swagger"))
	router.GET("/swagger/*any", swagHandler)

	// Add API routes
	api.Routes(router)

	address := ":8080"
	err := router.Run(address)

	if err != nil {
	}

}