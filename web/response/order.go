package response

type OrderResponse struct {
	ID                  string
	TrackingNumber      string
	ConsigneeAddress    string
	ConsigneeCity       string
	ConsigneeProvince   string
	ConsigneePostalCode string
	ConsigneeCountry    string
	Weight              float32
	Height              float32
	Width               float32
	Length              float32
	UserID              string
}

type CreateOrderResponse struct {
	ID             string
	TrackingNumber string
}
