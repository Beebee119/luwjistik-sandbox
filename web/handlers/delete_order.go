package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *apiHandler) DeleteOrder(c *gin.Context) {
	trackingNumber := c.Param("tracking_number")
	order, err := h.container.OrdersManager.GetOrder(trackingNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.container.OrdersManager.DeleteOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
	return
}
