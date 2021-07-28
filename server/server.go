package server

import (
	"github.com/fellippemendonca/gaia/services"
	// "github.com/fellippemendonca/gaia/services/databases/postgresql"
	"github.com/gin-gonic/gin"
)

// Run Server
func Run() {

	// Connect to required services
	// postgreConn := postgresql.Connect()
	// defer postgreConn.Close()

	// Assign services connection to an injectable Services struct
	svc := services.Services{}
	// svc.Postgresql = postgreConn

	// Create Router and middleware
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Init routes and inject started services access
	InitRoutes(router, &svc)

	// Start service on assigned port
	router.Run(":3000")
}
