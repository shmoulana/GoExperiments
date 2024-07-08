package main

import (
	"go-gin-postgres/auth"
	"go-gin-postgres/database"
	"go-gin-postgres/handlers"
	"go-gin-postgres/middleware"
	"go-gin-postgres/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)


func main() {

	// Initialize the logger
	logger := logrus.New()	
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.SetLevel(logrus.DebugLevel) // Set log level to debug for capturing SQL queries
	
	// Initialize the database
	db, err := database.Initialize(logger)
	if err != nil {
		// Handle error if database initialization fails
		logger.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Enable Gorm's debug mode to log SQL queries
	db.LogMode(true)

	router := gin.Default()

	// Use logging middleware
	router.Use(middleware.LoggingMiddleware())

	router.POST("/login", handlers.Login())

	// Group routes that require authentication
	authorized := router.Group("/")
	authorized.Use(auth.Authenticate)

	// User routes
	authorized.GET("/users", handlers.GetAll[models.User]())
	authorized.POST("/users", handlers.Create[models.User]())
	authorized.GET("/users/:id", handlers.GetByID[models.User]())
	authorized.PUT("/users/:id", handlers.UpdateByID[models.User]())
	authorized.DELETE("/users/:id", handlers.DeleteByID[models.User]())
	authorized.GET("/users/range/:start_id/:end_id", handlers.GetUsersByRange[models.User]())
	authorized.GET("/users/byname/:name", handlers.GetUserByName[models.User]())

	// Ticket routes
	authorized.GET("/tickets/date/:start_date/:end_date", handlers.GetTicketsByDate[models.Ticket]())
	authorized.GET("/tickets/date/time/:start_date/:end_date", handlers.GetTicketsByDateTime[models.Ticket]())
	authorized.GET("/tickets/:user_id", handlers.GetTicketsByUserId[models.Ticket]())
	authorized.GET("/tickets/payment/:status", handlers.GetTicketsByPaymentStatus[models.Ticket]())
	authorized.GET("/records/date/:date_created", handlers.GetRecordsByTicketDateCreated[models.Ticket, models.User, models.Order, models.Payment]())
	authorized.GET("/records/:date/:start_time/:end_time", handlers.GetRecordsByDateTimeRange[models.Ticket, models.User, models.Order, models.Payment]())



	// Order routes
	authorized.GET("/orders/date/:start_date/:end_date", handlers.GetOrdersByDate[models.Order]())

	// Payment routes
	authorized.GET("/payments/date/:start_date/:end_date", handlers.GetPaymentsByDate[models.Payment]())

	router.Run(":8080")
}
