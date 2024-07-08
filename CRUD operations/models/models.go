package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID       uint      `json:"id" gorm:"primary_key"`
	Name     string    `json:"name"  binding:"required" gorm:"not null"`
	Dob      time.Time `json:"dob" gorm:"type:date;not null"`
	Email    string    `json:"email" gorm:"unique;not null"`
	Password string    `json:"password" gorm:"not null"`
}

// Ticket represents a ticket in the system
type Ticket struct {
	TicketID    uint      `json:"ticket_id" gorm:"primary_key"`
	UserID      uint      `json:"user_id" gorm:"foreignkey:UserID;not null"`
	DateCreated time.Time `json:"date_created" gorm:"not null"`
	DatePaid    *time.Time `json:"date_paid"`
}



// Order represents an order in the system
type Order struct {
	OrderID       uint      `json:"order_id" gorm:"primary_key"`
	TicketID      uint      `json:"ticket_id" gorm:"foreignkey:TicketID;not null"`
	CreatedAtTime time.Time `json:"created_at_time" gorm:"not null"`
	MenuItem      string    `json:"menu_item" gorm:"not null"`
	Quantity      int       `json:"quantity" gorm:"not null"`
	Price         float64   `json:"price" gorm:"not null"`
}

// Payment represents a payment in the system
type Payment struct {
	PaymentID     uint      `json:"payment_id" gorm:"primary_key;not null"`
	TicketID      uint      `json:"ticket_id" gorm:"foreignkey:TicketID;not null"`
	CreatedAtTime time.Time `json:"created_at" gorm:"not null"`
	Amount        float64   `json:"amount" gorm:"not null"`
	Method        string    `json:"method" gorm:"not null"`
}

func (t Ticket) GetUserID() uint{
	return t.UserID
}

func (t Ticket) GetTicketID() uint{
	return t.TicketID
}
