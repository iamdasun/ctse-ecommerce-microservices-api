package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Order struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Status string `json:"status"`
}

var db *gorm.DB

func initDatabase() {
	var err error
	db, err = gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Order{})

	var count int64
	db.Model(&Order{}).Count(&count)
	if count == 0 {
		db.Create(&Order{Status: "PENDING"})
		db.Create(&Order{Status: "PENDING"})
		db.Create(&Order{Status: "PENDING"})
	}
}

func main() {
	initDatabase()

	r := gin.Default()

	// GET /orders - Return all orders
	r.GET("/orders", func(c *gin.Context) {
		var orders []Order
		db.Find(&orders)
		c.JSON(http.StatusOK, orders)
	})

	// GET /orders/:id - Return order by ID
	r.GET("/orders/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var order Order
		result := db.First(&order, id)

		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}

		c.JSON(http.StatusOK, order)
	})

	// POST /orders - Create new order with PENDING status
	r.POST("/orders", func(c *gin.Context) {
		order := Order{Status: "PENDING"}
		db.Create(&order)

		c.JSON(http.StatusCreated, order)
	})

	r.Run(":8082")
}
