package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type Order struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

var (
	orders    = []Order{}
	nextID    = 1
	orderLock sync.Mutex
)

func main() {
	r := gin.Default()

	// GET /orders - Return all orders
	r.GET("/orders", func(c *gin.Context) {
		c.JSON(http.StatusOK, orders)
	})

	// GET /orders/:id - Return order by ID
	r.GET("/orders/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		for _, order := range orders {
			if order.ID == id {
				c.JSON(http.StatusOK, order)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
	})

	// POST /orders - Create new order with PENDING status
	r.POST("/orders", func(c *gin.Context) {
		orderLock.Lock()
		defer orderLock.Unlock()

		order := Order{
			ID:     nextID,
			Status: "PENDING",
		}
		nextID++
		orders = append(orders, order)

		c.JSON(http.StatusCreated, order)
	})

	r.Run(":8082")
}
