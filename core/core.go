package core

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// type Base struct {
// 	ID        uuid.UUID `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }

type User struct {
	ID        uuid.UUID      `gorm:"primaryKey" json:"id"`
	Email     string         `gorm:"unique" json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Order struct {
	ID                  uuid.UUID `gorm:"primaryKey" json:"id"`
	TrackingNumber      string    `gorm:"unique" json:"tracking_number"`
	ConsigneeAddress    string    `json:"consignee_address"`
	ConsigneeCity       string    `json:"consignee_city"`
	ConsigneeProvince   string    `json:"consignee_province"`
	ConsigneePostalCode string    `json:"consignee_postal_code"`
	ConsigneeCountry    string    `json:"consignee_country"`
	Weight              float32   `json:"weight"`
	Height              float32   `json:"height"`
	Width               float32   `json:"width"`
	Length              float32   `json:"length"`
	UserID              uuid.UUID `gorm:"not_null" json:"user_id"`
	User                User
	CreatedAt           time.Time      `json:"created-at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type ServiceProvider struct {
	ID        uuid.UUID      `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID == uuid.Nil {
		o.ID = uuid.New()
	}
	return
}

func (s *ServiceProvider) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return
}
