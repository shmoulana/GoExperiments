package handlers

import (
	"net/http"

	"go-gin-postgres/database"
	"go-gin-postgres/models"

	"github.com/gin-gonic/gin"
)

type Model interface{
	models.User | models.Ticket | models.Order | models.Payment
}



func GetAll[T Model]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var records []T
		db := database.GetDB()
		db = db.Debug()
		db.Find(&records)
		c.JSON(http.StatusOK, records)
	}
}

func Create[T Model]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var record T
		if err := c.ShouldBindJSON(&record); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db := database.GetDB()
		db.Create(&record)
		c.JSON(http.StatusOK, record)
	}
}



func GetByID[T Model]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var record T
		id := c.Param("id")
		db := database.GetDB()
		if err := db.First(&record, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusOK, record)
	}
}

func UpdateByID[T Model]() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var record T
		db := database.GetDB()
		if err := db.First(&record, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		if err := c.ShouldBindJSON(&record); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Save(&record)
		c.JSON(http.StatusOK, record)
	}
}

func DeleteByID[T Model]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var record T
		id := c.Param("id")
		db := database.GetDB()
		if err := db.First(&record, id).Error; err != nil {
			c.JSON(404, gin.H{"error": "user not found"})
			return
		}
		db.Delete(&record)
		c.JSON(200, gin.H{"message": "user deleted"})
	}
}

func GetUsersByRange[T Model]() gin.HandlerFunc {
	return func(c *gin.Context) {
		startID := c.Param("start_id")
		endID := c.Param("end_id")
		var record []T
		db := database.GetDB()
		db.Where("id >= ? AND id <= ?", startID, endID).Find(&record)
		c.JSON(200, record)
	}
}

func GetUserByName[T Model]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var record T
		name := c.Param("name")
		db := database.GetDB()
		db.Where("name = ?", name).Find(&record)
		c.JSON(http.StatusOK, record)
	}
}

