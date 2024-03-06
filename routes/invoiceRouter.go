package routes

import "github.com/gin-gonic/gin"

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controller.GetFoods())
}
