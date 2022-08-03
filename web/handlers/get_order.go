package handlers

import (
	"net/http"
	"sandbox/web/response"

	"github.com/gin-gonic/gin"
)

func (h *apiHandler) GetOrder(c *gin.Context) {
	trackingNumber := c.Param("tracking_number")
	order, err := h.container.OrdersManager.GetOrder(trackingNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := response.OrderResponse{
		ID:                  order.ID.String(),
		TrackingNumber:      order.TrackingNumber,
		ConsigneeAddress:    order.ConsigneeAddress,
		ConsigneeCity:       order.ConsigneeCity,
		ConsigneeProvince:   order.ConsigneeProvince,
		ConsigneePostalCode: order.ConsigneePostalCode,
		ConsigneeCountry:    order.ConsigneeCountry,
		Weight:              order.Weight,
		Height:              order.Height,
		Width:               order.Width,
		Length:              order.Length,
		UserID:              order.UserID.String(),
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
	return
}
