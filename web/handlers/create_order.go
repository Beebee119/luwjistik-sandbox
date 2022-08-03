package handlers

import (
	"log"
	"net/http"
	"sandbox/core"
	"sandbox/web/response"

	"github.com/gin-gonic/gin"
)

func (h *apiHandler) CreateOrder(c *gin.Context) {
	var temp core.Order
	err := c.BindJSON(&temp)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.container.OrdersManager.CreateOrder(&temp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := response.CreateOrderResponse{
		ID:             temp.ID.String(),
		TrackingNumber: temp.TrackingNumber,
	}

	c.JSON(http.StatusCreated, gin.H{"Created data": resp})
	return
}
