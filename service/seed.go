package service

import (
	"math/rand"
	"sandbox/core"

	"github.com/google/uuid"
)

func SeedData(c *Container) {
	userA, _ := c.Auth.RegisterUser("usera@email.com", "password123")
	userB, _ := c.Auth.RegisterUser("userb@email.com", "password123")

	order1 := generateOrder(userA.ID)
	order2 := generateOrder(userA.ID)
	order3 := generateOrder(userB.ID)
	c.OrdersManager.CreateOrder(&order1)
	c.OrdersManager.CreateOrder(&order2)
	c.OrdersManager.CreateOrder(&order3)
}

func generateOrder(userID uuid.UUID) core.Order {
	return core.Order{
		TrackingNumber:      randomString(5),
		ConsigneeAddress:    "Bugis Street",
		ConsigneeCity:       "Singapore",
		ConsigneeProvince:   "Singapore",
		ConsigneePostalCode: "54321",
		ConsigneeCountry:    "SG",
		Weight:              2,
		Height:              3,
		Width:               4,
		Length:              5,
		UserID:              userID,
	}
}

func randomString(n int) string {
	var chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}
