package server

import (
	"github.com/fellippemendonca/gaia/db/postgresql"
	"github.com/fellippemendonca/gaia/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"os"
)

// Run Server
func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to required services
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASS")
	db := os.Getenv("POSTGRES_DB")
	postgresUrl := "postgres://" + user + ":" + pass + "@" + host + "/" + db + "?sslmode=disable"
	postgreConn := postgresql.Connect(postgresUrl)
	defer postgreConn.Close()

	// Assign services connection to an injectable Services struct
	svc := services.Services{}
	svc.Postgresql = postgreConn

	// Create Router and middleware
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Init routes and inject started services access
	InitRoutes(router, &svc)

	// Start service on assigned port
	router.Run(":3000")
}
