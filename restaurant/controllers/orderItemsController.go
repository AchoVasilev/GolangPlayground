package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrderItems() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func GetOrderItemsByOrderId() (OrderItems []primitive.M, err error) {

}

func CreateOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
