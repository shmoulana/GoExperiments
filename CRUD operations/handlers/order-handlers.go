package handlers

import (
	"go-gin-postgres/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetOrdersByDate[T Model]() gin.HandlerFunc {
	return func(c *gin.Context) {
		startDate, _ := time.Parse("2006-01-02", c.Param("start_date"))
		endDate, _ := time.Parse("2006-01-02", c.Param("end_date"))
		var records []T
		db := database.GetDB()
		db = db.Debug()
		db.Where("created_at_time >= ? AND created_at_time <= ?", startDate, endDate).Find(&records)
		c.JSON(http.StatusOK, records)
	}
}