package handlers

import (
	"log"
	"net/http"
	"sandbox/core"
	"sandbox/web/response"

	"github.com/gin-gonic/gin"
)

func (h *apiHandler) UpdateOrder(c *gin.Context) {
	var temp core.Order
	err := c.BindJSON(&temp)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	order, err := h.container.OrdersManager.GetOrder(temp.TrackingNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	order.ConsigneeAddress = temp.ConsigneeAddress
	order.ConsigneeCity = temp.ConsigneeCity
	order.ConsigneeProvince = temp.ConsigneeProvince
	order.ConsigneePostalCode = temp.ConsigneePostalCode
	order.ConsigneeCountry = temp.ConsigneeCountry
	order.Weight = temp.Weight
	order.Height = temp.Height
	order.Width = temp.Width
	order.Length = temp.Length

	err = h.container.OrdersManager.UpdateOrder(&order)
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

	c.JSON(http.StatusOK, gin.H{"Created data": resp})
	return
}
