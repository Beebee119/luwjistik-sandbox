package handlers

import (
	"net/http"
	"sandbox/service"

	"github.com/gin-gonic/gin"
)

type apiHandler struct {
	container *service.Container
}

func NewApiHandler(container *service.Container) *apiHandler {
	return &apiHandler{
		container: container,
	}
}

func (h *apiHandler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello"})
	return
}
