package controllers

import (
	"errors"
	"go-ecom/db"
	"go-ecom/dtos"
	"go-ecom/models"
	"go-ecom/validations"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary		Получить все заказы
// @Description	Возвращает массив заказов
// @Tags			order
// @Produce		json
// @Success		200	{object}	ResponseCustom{data=[]models.Orders}
// @Failure		400	{object}	ErrorResponse{error=string} "Orders not found!"
// @Router			/order [get]
func GetOrderAll(c *gin.Context) {
	var orders []models.Orders
	if err := db.DB.Model(&models.Orders{}).Preload("Users").Preload("OrdersItems").Preload("OrdersItems.Items").Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Orders not found!"})
		return
	}

	c.JSON(http.StatusOK, ResponseCustom{Data: orders})
}

// @Summary		Получить однин заказ
// @Description	Возвращает заказ с указанным id
// @Tags			order
// @Produce		json
// @Param			id	path		int	true	"Order ID"
// @Success		200	{object}	ResponseCustom{data=models.Orders}
// @Failure		400	{object}	ErrorResponse{error=string} "Order not found!"
// @Router			/order/{id} [get]
func GetOrderById(c *gin.Context) {
	var order models.Orders
	if err := db.DB.Model(&models.Orders{}).Preload("Users").Preload("OrdersItems").Preload("OrdersItems.Items").First(&order, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Order not found!"})
		return
	}

	c.JSON(http.StatusOK, ResponseCustom{Data: order})
}


// @Summary		Создать заказ
// @Description	Возвращает ID созданного пользователя
// @Tags			order
// @Accept			json
// @Produce		json
// @Param			request	body		dtos.CreateOrdersInput	true	"query params"
// @Success		200	{object}	ResponseCustom{data=ResponseID}
// @Failure		400	{object}	ErrorResponse{error=[]validations.ErrorMsg}
// @Failure		400	{object}	ErrorResponse{error=string} "Can`t create order"
// @Router			/order [post]
func CreateOrder(c *gin.Context) {
	// Validate input
	var input dtos.CreateOrdersInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]validations.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = validations.ErrorMsg{Field: fe.Field(), Message: validations.GetErrorMsg(fe)}
			}
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: out})
			return
		}
	}


	var ordersItems []models.OrdersItems

	for _, item := range input.OrdersItems {
		modelsItem := models.OrdersItems{
			ItemsID: item.ItemsID, Quantity: item.Quantity,
		}
		ordersItems = append(ordersItems, modelsItem)
	}

	// Create order
	order := models.Orders{Status: input.Status, UsersID: input.UsersID, OrdersItems: ordersItems}
	if err := db.DB.Model(&order).Create(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Can`t create order"})
		return
	}

	c.JSON(http.StatusOK, ResponseCustom{Data: ResponseID{ID: order.ID}})
}


// @Summary		Обновить заказ
// @Description	Обновляет заказ по id и возвращает ID этого заказа
// @Tags			order
// @Accept			json
// @Produce			json
// @Param			id		path		int						true	"Order ID"
// @Param			request	body		dtos.UpdateOrdersInput	true	"query params"
// @Success		200	{object}	ResponseCustom{data=ResponseID}
// @Failure		400	{object}	ErrorResponse{error=[]validations.ErrorMsg}
// @Failure		400	{object}	ErrorResponse{error=string} "Order not found!"
// @Failure		400	{object}	ErrorResponse{error=string} "Can`t update order"
// @Router			/order/{id} [patch]
func UpdateOrderById(c *gin.Context) {
	var order models.Orders
	if err := db.DB.First(&order, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest,  ErrorResponse{Error: "Order not found!"})
		return
	}

	// Validate input
	var input dtos.UpdateOrdersInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]validations.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = validations.ErrorMsg{Field: fe.Field(), Message: validations.GetErrorMsg(fe)}
			}
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: out})
		}
		return
	}

	if err := db.DB.Model(&order).Updates(input.Status).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Can`t update order"})
		return
	}

	c.JSON(http.StatusOK, ResponseCustom{Data: ResponseID{ID: order.ID}})
}

// @Summary		Удалить заказ
// @Description	Удаляет заказ по id и возвращает флаг true
// @Tags			order
// @Accept			json
// @Produce			json
// @Param			id		path		int						true	"Order ID"
// @Success		200	{object}	ResponseCustom{data=bool}
// @Failure		400	{object}	ErrorResponse{error=string} "error": "Order not found!"
// @Failure		400	{object}	ErrorResponse{error=string} "error": "error": "Can`t delete order"
// @Router			/order/{id} [delete]
func DeleteOrderById(c *gin.Context) {
	var order models.Orders
	if err := db.DB.First(&order, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Order not found!"})
		return
	}

	if err := db.DB.Delete(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Can`t delete order"})
		return
	}

	c.JSON(http.StatusOK, ResponseCustom{Data: true})
}
