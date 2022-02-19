package router

import (
	"cashflow/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	payments := router.Group("/payments")
	{
		payments.GET("/", controllers.ListPayments)
		payments.POST("/", controllers.CreatePayment)
		payments.PUT("/:id", controllers.UpdatePayment)
		payments.DELETE("/:id", controllers.DeletePayments)
	}
	return router
}
