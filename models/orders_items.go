package models

type OrdersItems struct {
	OrdersID uint `json:"orders_id" gorm:"primaryKey"`
	ItemsID  uint  `json:"items_id" gorm:"primaryKey"`
	Items    Items `json:"items"`
	Quantity uint  `json:"quantity" gorm:"not null"`
}
