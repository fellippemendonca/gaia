package server

import (
	"github.com/fellippemendonca/gaia/db/postgresql"
	"github.com/fellippemendonca/gaia/db/postgresql/orm/orm_forecast"
	"github.com/fellippemendonca/gaia/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"os"
)

// Run Server
func Run() {

	// Loading .env file variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to required services
	postgresUrl := os.Getenv("POSTGRES_URL")
	postgreConn := postgresql.Connect(postgresUrl)
	defer postgreConn.Close()

	// Assign services connection to an injectable Services struct
	svc := services.Services{}
	svc.Postgresql = postgreConn

	// Create new tables if it not yet exist
	result, err := orm_forecast.CreateForecastsTable(&svc)
	if err != nil {
		log.Fatal("Error creating tables")
	}
	log.Println(result)

	// Create Router and middleware
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Init routes and inject started services access
	InitRoutes(router, &svc)

	// Start service on assigned port
	router.Run(":3000")
}
