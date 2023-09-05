package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	ID uint `json:"id" gorm:"primaryKey"`

	TotalAmount float64 `json:"total_amount" gorm:"default=0"`
	Status      string  `json:"status" gorm:"size:255 not null"`

	UsersID     uint          `json:"users_id" gorm:"not null"`
	Users       Users         `json:"users" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OrdersItems []OrdersItems `json:"orders_items" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (o *Orders) AfterCreate(db *gorm.DB) (err error) {
	var totalAmount float64
	for _, orderItem := range o.OrdersItems {
		var item Items

		// Предполагаем, что orderItem.ItemsID - это ID товара для текущего элемента OrdersItems
		result := db.First(&item, orderItem.ItemsID)
		if result.Error != nil {
			// Обработка ошибки, если не удалось найти товар по ID			
			return errors.New("откат, item не найден")
		}

		// Вычисляем сумму для текущего заказа и текущего элемента OrdersItems
		subtotal := float64(orderItem.Quantity) * item.Price
		totalAmount += subtotal
	}
	
	if err := db.Model(o).Update("TotalAmount", totalAmount).Error; err != nil {
		return errors.New("откат транзакции, обновить orders.total_amount не удалось")
	}

	return
}
