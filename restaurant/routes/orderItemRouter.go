package routes

import (
	"github.com/gin-gonic/gin"
	"restaurant/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders/items", controllers.GetOrderItems())
	incomingRoutes.GET("/orders/items/:id", controllers.GetOrderItem())
	incomingRoutes.GET("/orders/items/order/:id", controllers.GetOrderItemsByOrderId())
	incomingRoutes.POST("/orders/items", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orders/items/:id", controllers.UpdateOrderItem())
}
