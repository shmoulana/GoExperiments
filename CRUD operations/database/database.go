package database

import (
	"go-gin-postgres/models"
	"log"
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Initialize initializes the database connection using the singleton pattern
func Initialize(logger *logrus.Logger) (*gorm.DB, error) {
	var err error

	once.Do(func() {
		// Connect to PostgreSQL database
		db, err = gorm.Open("postgres", "host=localhost user=postgres dbname=myapi sslmode=disable password=12345678")
		if err != nil {
			logger.Fatalf("Failed to connect to database: %v", err)
		}

		// Set logger for GORM
		db.SetLogger(logger)

		// Auto-migrate models
		db.AutoMigrate(&models.User{}, &models.Ticket{}, &models.Order{}, &models.Payment{})
		db.LogMode(true)
		db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	})

	return db, err
}

// GetDB returns the singleton database instance
func GetDB() *gorm.DB {
	return db
}
