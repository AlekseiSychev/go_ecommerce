package dtos

type CreateOrdersInput struct {
	UsersID uint `json:"user_id" binding:"required,gte=0"`
	Status string `json:"status" binding:"required,min=3,max=20"`

	OrdersItems []*CreateOrdersItemsInput `json:"orders_items" binding:"required,min=1"`
}

type UpdateOrdersInput struct {
	Status string `json:"status" binding:"required,min=3,max=20"`
}

type CreateOrdersItemsInput struct {
	ItemsID  uint `json:"items_id" binding:"required,gte=0"`
	Quantity uint `json:"quantity" binding:"required,gte=1"`
}
