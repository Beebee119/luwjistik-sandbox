package handlers

import (
	"net/http"
	"sandbox/web/response"

	"github.com/gin-gonic/gin"
)

func (h *apiHandler) GetOrders(c *gin.Context) {
	orders, err := h.container.OrdersManager.ListOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp []response.OrderResponse
	for _, o := range orders {
		resp = append(resp, response.OrderResponse{
			ID:                  o.ID.String(),
			TrackingNumber:      o.TrackingNumber,
			ConsigneeAddress:    o.ConsigneeAddress,
			ConsigneeCity:       o.ConsigneeCity,
			ConsigneeProvince:   o.ConsigneeProvince,
			ConsigneePostalCode: o.ConsigneePostalCode,
			ConsigneeCountry:    o.ConsigneeCountry,
			Weight:              o.Weight,
			Height:              o.Height,
			Width:               o.Width,
			Length:              o.Length,
			UserID:              o.UserID.String(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
	return
}
