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

// @Summary		Получить весь товар
// @Description	Возвращает массив товара
// @Tags			item
// @Produce		json
// @Success		200	{object}	ResponseCustom{data=[]models.Items}
// @Failure		400	{object}	ErrorResponse{error=string} "Items not found!"
// @Router			/item [get]
func GetItemAll(c *gin.Context) {
	var items []models.Items
	if err := db.DB.Find(&items).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Items not found!"})
		return
	}

	c.JSON(http.StatusOK, ResponseCustom{Data: items})
}

// @Summary		Получить один товар
// @Description	Возвращает товар с указанным id
// @Tags			item
// @Produce		json
// @Param			id	path		int	true	"Item ID"
// @Success		200	{object}	ResponseCustom{data=models.Items}
// @Failure		400	{object}	ErrorResponse{error=string} "Item not found!"
// @Router			/item/{id} [get]
func GetItemById(c *gin.Context) {
	// Get model if exist
	var item models.Items
	if err := db.DB.First(&item, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Item not found!"})
		return
	}

	c.JSON(http.StatusOK, ResponseCustom{Data: item})
}

// @Summary		Создать товар
// @Description	Возвращает ID созданного товара
// @Tags			item
// @Accept			json
// @Produce		json
// @Param			request	body		dtos.CreateItemInput	true	"query params"
// @Success		200	{object}	ResponseCustom{data=ResponseID}
// @Failure		400	{object}	ErrorResponse{error=[]validations.ErrorMsg}
// @Failure		400	{object}	ErrorResponse{error=string} "Can`t create item"
// @Router			/item [post]
func CreateItem(c *gin.Context) {
	// Validate input
	var input dtos.CreateItemInput
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

	// Create item
	item := models.Items{Name: input.Name, Price: input.Price}
	if err := db.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Can`t create item"})
		return
	}

	c.JSON(http.StatusOK,  ResponseCustom{Data: ResponseID{ID: item.ID}})
}


// @Summary		Обновить товар
// @Description	Обновляет товар по ID и возвращает ID этого товара
// @Tags			item
// @Accept			json
// @Produce		json
// @Param			id		path		int						true	"Item ID"
// @Param			request	body		dtos.UpdateItemInput	true	"query params"
// @Success		200	{object}	ResponseCustom{data=ResponseID}
// @Failure		400	{object}	ErrorResponse{error=[]validations.ErrorMsg}
// @Failure		400	{object}	ErrorResponse{error=string} "Item not found!"
// @Failure		400	{object}	ErrorResponse{error=string} "Can`t update iten"
// @Router			/item/{id} [patch]
func UpdateItemById(c *gin.Context) {
	// Get model if exist
	var item models.Items
	if err := db.DB.First(&item, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Item not found!"})
		return
	}

	// Validate input
	var input dtos.UpdateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
        if errors.As(err, &ve) {
            out := make([]validations.ErrorMsg, len(ve))
            for i, fe := range ve {
                out[i] = validations.ErrorMsg{Field: fe.Field(), Message: validations.GetErrorMsg(fe)}
            }
            c.JSON(http.StatusBadRequest, gin.H{"errors": out})
			
        }
		return
	}

	if err := db.DB.Model(&item).Updates(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Item not found!"})
		return
	}

	c.JSON(http.StatusOK, ResponseCustom{Data: ResponseID{ID: item.ID}})
}

// @Summary		Удалить товар
// @Description	Удаляет товар по id и возвращает флаг true
// @Tags			item
// @Accept			json
// @Produce			json
// @Param			id		path		int						true	"Item ID"
// @Success		200	{object}	ResponseCustom{data=bool}
// @Failure		400	{object}	ErrorResponse{error=string} "error": "Item not found!"
// @Failure		400	{object}	ErrorResponse{error=string} "error": "error": "Can`t delete item"
// @Router			/item/{id} [delete]
func DeleteItemById(c *gin.Context) {
	// Get model if exist
	var item models.Items
	if err := db.DB.First(&item, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Item not found!"})
		return
	}

	if err := db.DB.Delete(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Can`t delete item"})
		return
	}

	c.JSON(http.StatusOK, ResponseCustom{Data: true})
}
