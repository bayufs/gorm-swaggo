package dto

type InputOrder struct {
	CustomerName string      `json:"customerName,omitempty"`
	Items        []InputItem `json:"items,omitempty"`
}

type InputItem struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type Order struct {
	ID           string `json:"id,omitempty"`
	CustomerName string `json:"customerName,omitempty"`
	Items        []Item `json:"items,omitempty" gorm:"foreignKey:OrderId"`
}

type Item struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     string `json:"orderId,omitempty"`
}
