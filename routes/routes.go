package routes

import (
	"github.com/gin-gonic/gin"
	"go-ecom/controllers"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

func GroupRoutes(router *gin.Engine) {
	docs := router.Group("/docs")
	{
		docs.GET("*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	}

	order_routes := router.Group("/order")
	{
		order_routes.GET("", controllers.GetOrderAll)
		order_routes.GET(":id", controllers.GetOrderById)
		order_routes.POST("", controllers.CreateOrder)
		order_routes.PATCH(":id", controllers.UpdateOrderById)
		order_routes.DELETE(":id", controllers.DeleteOrderById)
	}

	user_routes := router.Group("/user")
	{
		user_routes.GET("", controllers.GetUserAll)
		user_routes.GET(":id", controllers.GetUserById)
		user_routes.POST("", controllers.CreateUser)
		user_routes.PATCH(":id", controllers.UpdateUserById)
		user_routes.DELETE(":id", controllers.DeleteUserById)
	}


	item_routes := router.Group("/item")
	{
		item_routes.GET("", controllers.GetItemAll)
		item_routes.GET(":id", controllers.GetItemById)
		item_routes.POST("", controllers.CreateItem)
		item_routes.PATCH(":id", controllers.UpdateItemById)
		item_routes.DELETE(":id", controllers.DeleteItemById)
	}
}