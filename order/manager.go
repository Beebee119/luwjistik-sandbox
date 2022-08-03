package order

import (
	"sandbox/core"

	"gorm.io/gorm"
)

type OrderManager struct {
	db *gorm.DB
}

func NewManager(db *gorm.DB) *OrderManager {
	return &OrderManager{
		db: db,
	}
}

func (o *OrderManager) UpdateOrder(order *core.Order) error {
	return o.db.Save(order).Error
}

func (o *OrderManager) ListOrders() (orders []core.Order, err error) {
	err = o.db.Model(core.Order{}).Find(&orders).Error
	return orders, err
}

func (o *OrderManager) GetOrder(trackingNumber string) (order core.Order, err error) {
	// err = o.db.Where("").Find(&order).Error
	err = o.db.First(&order, "tracking_number = ?", trackingNumber).Error
	return order, err
}

func (o *OrderManager) CreateOrder(order *core.Order) error {
	return o.db.Create(order).Error
}

func (o *OrderManager) DeleteOrder(order *core.Order) error {
	return o.db.Save(order).Error
}
