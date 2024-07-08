package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"go-gin-postgres/database"
	"go-gin-postgres/models"
)

// Seeder is responsible for seeding data into the database
type Seeder struct {
	DB *gorm.DB
}

// NewSeeder creates a new instance of Seeder with the provided database connection
func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{
		DB: db,
	}
}

// SeedUsers seeds user data into the database
func (s *Seeder) SeedUsers(numUsers int) error {
	for i := 0; i < numUsers; i++ {
		user := models.User{
			Name:     gofakeit.Name(),
			Dob:      stripTime(gofakeit.DateRange(time.Now().AddDate(-50, 0, 0), time.Now().AddDate(-18, 0, 0))),
			Email:    gofakeit.Email(),
			Password: gofakeit.Password(true, true, true, false, false, 12),
		}
		if err := s.DB.Create(&user).Error; err != nil {
			log.Printf("Failed to create user: %v", err)
			return err
		}
	}
	return nil
}

// SeedTickets seeds ticket data into the database
func (s *Seeder) SeedTickets(numTickets int, numUsers int) error {
	startDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 5, 30, 23, 59, 59, 0, time.UTC)

	for i := 0; i < numTickets; i++ {
		var datePaid *time.Time
		if rand.Float64() < 0.2 {
			datePaid = nil
		} else {
			datePaidValue := stripTime(gofakeit.DateRange(startDate, endDate)).Add(randomTime())
			datePaid = &datePaidValue
		}
		dateCreated := stripTime(gofakeit.DateRange(startDate, endDate)).Add(randomTime())

		ticket := models.	Ticket{
			UserID:      uint(rand.Intn(numUsers) + 1),
			DateCreated: dateCreated,
			DatePaid:    datePaid,
		}
		if err := s.DB.Create(&ticket).Error; err != nil {
			log.Printf("Failed to create ticket: %v", err)
			return err
		}
	}
	return nil
}


// SeedOrders seeds order data into the database
func (s *Seeder) SeedOrders(numOrders int, numTickets int) error {
	startDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 5, 30, 0, 0, 0, 0, time.UTC)
	for i := 0; i < numOrders; i++ {
		order := models.Order{
			TicketID:      uint(rand.Intn(numTickets) + 1),
			CreatedAtTime: stripTime(gofakeit.DateRange(startDate, endDate)).Add(randomTime()),
			MenuItem:      gofakeit.BeerName(),
			Quantity:      gofakeit.Number(1, 5),
			Price:         gofakeit.Price(1, 100),
		}
		if err := s.DB.Create(&order).Error; err != nil {
			log.Printf("Failed to create order: %v", err)
			return err
		}
	}
	return nil
}

// SeedPayments seeds payment data into the database
func (s *Seeder) SeedPayments(numPayments int, numTickets int) error {
	startDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 5, 30, 0, 0, 0, 0, time.UTC)
	for i := 0; i < numPayments; i++ {
		payment := models.Payment{
			TicketID:      uint(rand.Intn(numTickets) + 1),
			CreatedAtTime: stripTime(gofakeit.DateRange(startDate, endDate)).Add(randomTime()),
			Amount:        gofakeit.Price(1, 100),
			Method:        gofakeit.RandomString([]string{"Credit Card", "PayPal", "Bitcoin", "Bank Transfer"}),
		}
		if err := s.DB.Create(&payment).Error; err != nil {
			log.Printf("Failed to create payment: %v", err)
			return err
		}
	}
	return nil
}

// Helper function to strip time from date
func stripTime(t time.Time) time.Time {
	return t.Truncate(24 * time.Hour)
}

// Helper function to generate a random time duration within a day
func randomTime() time.Duration {
	return time.Duration(rand.Intn(24))*time.Hour + time.Duration(rand.Intn(60))*time.Minute + time.Duration(rand.Intn(60))*time.Second
}

func main() {
	// Initialize the logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.SetLevel(logrus.DebugLevel) // Set log level to debug for capturing SQL queries

	
	// Initialize the database
	db, err := database.Initialize(logger)
	if err != nil {
		// Handle error if database initialization fails
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Create a new instance of Seeder
	seeder := NewSeeder(db)

	// Seed initial data into the database
	numUsers := 1000000
	numTickets := 2000000
	numOrders := 5000000
	numPayments := 4000000
	if err := seeder.SeedUsers(numUsers); err != nil {
		log.Fatalf("Failed to seed users: %v", err)
	}
	if err := seeder.SeedTickets(numTickets, numUsers); err != nil {
		log.Fatalf("Failed to seed tickets: %v", err)
	}
	if err := seeder.SeedOrders(numOrders, numTickets); err != nil {
		log.Fatalf("Failed to seed orders: %v", err)
	}
	if err := seeder.SeedPayments(numPayments, numTickets); err != nil {
		log.Fatalf("Failed to seed payments: %v", err)
	}

	log.Println("Data seeding completed successfully.")
}
