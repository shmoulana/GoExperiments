package handlers

import (
	"net/http"
	"time"

	"go-gin-postgres/database"
	"go-gin-postgres/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TicketModel interface {
	models.Ticket
    GetUserID() uint
    GetTicketID() uint
}

type OrderModel interface {
	models.Order
}

type PaymentModel interface {
	models.Payment
}


func GetTicketsByDate[T Model]() gin.HandlerFunc {
	return func(c *gin.Context) {
		startDate, _ := time.Parse("2006-01-02", c.Param("start_date"))
		endDate, _ := time.Parse("2006-01-02", c.Param("end_date"))
		var records []T
		db := database.GetDB()
		db = db.Debug()
		db.Where("date_created >= ? AND date_created <= ?", startDate, endDate).Find(&records)
		c.JSON(http.StatusOK, records)
	}
}

func GetTicketsByDateTime[T Model]() gin.HandlerFunc {
	return func(c *gin.Context) {
		startDate, err1 := time.Parse("2006-01-02 15:04:05", c.Param("start_date"))
		endDate, err2 := time.Parse("2006-01-02 15:04:05", c.Param("end_date"))
		
		if err1 != nil || err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD HH:MM:SS"})
			return
		}
		
		var records []T
		db := database.GetDB().Debug()
		db.Where("date_created >= ? AND date_created <= ?", startDate, endDate).Find(&records)
		c.JSON(http.StatusOK, records)
	}
}



func GetTicketsByUserId[T Model]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var records []T
		db := database.GetDB()
		db.Where("user_id = ?", c.Param("user_id")).Find(&records)
		c.JSON(http.StatusOK, records)
	}
}

func GetTicketsByPaymentStatus[T Model]() gin.HandlerFunc {
	return func(c *gin.Context) {
		status := c.Param("status")
		var records []T
		db := database.GetDB()
		if status == "paid" {
			db.Where("date_paid IS NOT NULL").Find(&records)
		} else if status == "unpaid" {
			db.Where("date_paid IS NULL").Find(&records)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
			return
		}
		c.JSON(http.StatusOK, records)
	}
}




func GetRecordsByTicketDateCreated[T TicketModel, U Model, O OrderModel, P PaymentModel]() gin.HandlerFunc {
    return func(c *gin.Context) {

		start := time.Now()

        // Parse date from URL parameter
        dateCreated, err := time.Parse("2006-01-02", c.Param("date_created"))

        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
            return
        }

        var tickets []T
        var userIDs []uint
        var ticketIDs []uint

        db := database.GetDB().Debug()

        // Find tickets with the specific date_created
        if err := db.Where("DATE(date_created) = ?", dateCreated).Find(&tickets).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        for _, ticket := range tickets {
            userIDs = append(userIDs, ticket.GetUserID())
            ticketIDs = append(ticketIDs, ticket.GetTicketID())
        }

        var users []U
        var orders []O
        var payments []P

        // Find users related to the tickets
        if err := db.Where("id IN (?)", userIDs).Find(&users).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // Find orders related to the tickets
        if err := db.Where("ticket_id IN (?)", ticketIDs).Find(&orders).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // Find payments related to the tickets
        if err := db.Where("ticket_id IN (?)", ticketIDs).Find(&payments).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // Create a custom response struct
        type Response struct {
            Users    []U    `json:"users"`
            Tickets  []T `json:"tickets"`
            Orders   []O `json:"orders"`
            Payments []P `json:"payments"`
        }

        // Populate the response struct
        response := Response{
            Users:    users,
            Tickets:  tickets,
            Orders:   orders,
            Payments: payments,
        }
		// Log request details and execution time
        logrus.Infof("Handler: GetRecordsByTicketDateCreated | Execution time: %v", time.Since(start))

        c.JSON(http.StatusOK, response)
    }
}


func GetRecordsByDateTimeRange[T TicketModel, U Model, O OrderModel, P PaymentModel]() gin.HandlerFunc {
    return func(c *gin.Context) {
		start := time.Now()


        // Parse date parameter from URL
        date, err := time.Parse("2006-01-02", c.Param("date"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
            return
        }

        // Parse start and end times from URL
        startTime, err := time.Parse("15:04:05", c.Param("start_time"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start time format. Use HH:MM:SS"})
            return
        }

        endTime, err := time.Parse("15:04:05", c.Param("end_time"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end time format. Use HH:MM:SS"})
            return
        }

        // Combine date with start and end times
        startDateTime := time.Date(date.Year(), date.Month(), date.Day(), startTime.Hour(), startTime.Minute(), startTime.Second(), 0, time.UTC)
        endDateTime := time.Date(date.Year(), date.Month(), date.Day(), endTime.Hour(), endTime.Minute(), endTime.Second(), 0, time.UTC)

        var tickets []T
        var userIDs []uint
        var ticketIDs []uint

        db := database.GetDB().Debug()

        // Find tickets within the specified date and time range
        if err := db.Where("date_created BETWEEN ? AND ?", startDateTime, endDateTime).Find(&tickets).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        for _, ticket := range tickets {
            userIDs = append(userIDs, ticket.GetUserID())
            ticketIDs = append(ticketIDs, ticket.GetTicketID())
        }

        var users []U
        var orders []O
        var payments []P

        // Find users related to the tickets
        if err := db.Where("id IN (?)", userIDs).Find(&users).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // Find orders related to the tickets
        if err := db.Where("ticket_id IN (?)", ticketIDs).Find(&orders).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // Find payments related to the tickets
        if err := db.Where("ticket_id IN (?)", ticketIDs).Find(&payments).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // Create a custom response struct
        type Response struct {
            Users    []U `json:"users"`
            Tickets  []T  `json:"tickets"`
            Orders   []O  `json:"orders"`
            Payments []P `json:"payments"`
        }

        // Populate the response struct
        response := Response{
            Users:    users,
            Tickets:  tickets,
            Orders:   orders,
            Payments: payments,
        }

		// Log request details and execution time
        logrus.Infof("Handler: GetRecordsByDateTimeRange | Execution time: %v", time.Since(start))

        c.JSON(http.StatusOK, response)
    }
}

