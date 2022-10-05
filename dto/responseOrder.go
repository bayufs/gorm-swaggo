package dto

type ResponseOrder[T any] struct {
	Message string `example:"Success"`
	Status  int    `example:"200"`
	Data    []T
}

type ResponseCreateOrder struct {
	Message string `example:"Success"`
	Status  int    `example:"201"`
	Data    Order
}
