package controllers

import (
	"errors"
	"go-ecom/db"
	"go-ecom/dtos"
	"go-ecom/models"
	"go-ecom/validations"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary		Получить всех пользователей
// @Description	Возвращает массив пользователей
// @Tags			user
// @Produce		json
// @Success		200	{object}	ResponseCustom{data=[]models.Users}
// @Failure		400	{object}	ErrorResponse{error=string} "Users not found!"
// @Router			/user [get]
func GetUserAll(c *gin.Context) {
	var users []models.Users
	if err := db.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Users not found!"})
		return
	}
	c.JSON(http.StatusOK, ResponseCustom{Data: users})
}

// @Summary		Получить одного пользователя
// @Description	Возвращает пользователя с указанным id
// @Tags			user
// @Produce		json
// @Param			id	path		int	true	"User ID"
// @Success		200	{object}	ResponseCustom{data=models.Users}
// @Failure		400	{object}	ErrorResponse{error=string} "User not found!"
// @Router			/user/{id} [get]
func GetUserById(c *gin.Context) {
	var user models.Users
	if err := db.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "User not found!"})
		return
	}

	c.JSON(http.StatusOK, ResponseCustom{Data: user})
}

// @Summary		Создать пользователя
// @Description	Возвращает ID созданного пользователя
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			request	body		dtos.CreateUserInput	true	"query params"
// @Success		200	{object}	ResponseCustom{data=ResponseID}
// @Failure		400	{object}	ErrorResponse{error=[]validations.ErrorMsg}
// @Failure		400	{object}	ErrorResponse{error=string} "Can`t create user"
// @Router			/user [post]
func CreateUser(c *gin.Context) {
	// Validate input
	var input dtos.CreateUserInput
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

	// Create order
	user := models.Users{Name: strings.ToLower(input.Name), Email: input.Email, Phone: input.Phone, Password: input.Password}
	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Can`t create user"})
		return
	}

	c.JSON(http.StatusOK, ResponseCustom{Data: ResponseID{ID: user.ID}})
}

// @Summary		Обновить пользователя
// @Description	Обновляет пользователя по ID и возвращает ID этого пользователя
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			id		path		int						true	"User ID"
// @Param			request	body		dtos.UpdateUserInput	true	"query params"
// @Success		200	{object}	ResponseCustom{data=ResponseID}
// @Failure		400	{object}	ErrorResponse{error=[]validations.ErrorMsg}
// @Failure		400	{object}	ErrorResponse{error=string} "User not found!"
// @Failure		400	{object}	ErrorResponse{error=string} "Can`t update user"
// @Router			/user/{id} [patch]
func UpdateUserById(c *gin.Context) {
	// Get model if exist
	var user models.Users
	if err := db.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "User not found!"})
		return
	}

	// Validate input
	var input dtos.UpdateUserInput
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


	if err := db.DB.Model(&user).Updates(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Can`t update user"})
		return
	}


	c.JSON(http.StatusOK, ResponseCustom{Data: ResponseID{ID: user.ID}})
}


// @Summary		Удалить пользователя
// @Description	Удаляет пользователя по id и возвращает флаг true
// @Tags			user
// @Accept			json
// @Produce			json
// @Param			id		path		int						true	"User ID"
// @Success		200	{object}	ResponseCustom{data=bool}
// @Failure		400	{object}	ErrorResponse{error=string} "error": "Record not found!"
// @Failure		400	{object}	ErrorResponse{error=string} "error": "error": "Can`t delete user"
// @Router			/user/{id} [delete]
func DeleteUserById(c *gin.Context) {
	// Get model if exist
	var user models.Users
	if err := db.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "User not found!"})
		return
	}

	if err := db.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Can`t delete user"})
		return
	}

	c.JSON(http.StatusOK,  ResponseCustom{Data: true})
}
