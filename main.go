package main

import (
	"database/sql"
	"fmt"
	"golang_project/controller"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "gamex_db"
)

func main() {
	// Setup database connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a connection to the database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}
	defer db.Close()

	// Verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	} else {
		fmt.Println("Successfully connected to the database!")
	}

	// Initialize Gin
	r := gin.Default()

	// Serve static files (CSS, JS, images)
	r.Static("/assets", "./gamex-master/assets")

	// Serve index.html page
	r.GET("/index", func(c *gin.Context) {
		c.File("./gamex-master/index.html")
	})

	// Serve index2 page (view)
	r.GET("/index2", func(c *gin.Context) {
		c.File("./gamex-master/index2.html")
	})

	// API endpoint to get player rankings (controller)
	r.GET("/api/players", controller.GetPlayersController(db))

	// Start the HTTP server
	fmt.Println("Starting server on :8080...")
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
