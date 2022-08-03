package web

import (
	"sandbox/service"
	"sandbox/web/handlers"
)

func RegisterAPIRoutes(c *service.Container) {
	api := handlers.NewApiHandler(c)

	c.Web.GET("/hello", api.Hello)
	c.Web.GET("/orders", api.GetOrders)
	c.Web.GET("/orders/:tracking_number", api.GetOrder)
	c.Web.DELETE("/orders/:tracking_number", api.DeleteOrder)
	c.Web.POST("/orders", api.CreateOrder)
	c.Web.PUT("/orders", api.UpdateOrder)
}
