package routes

import (
	"github.com/gin-gonic/gin"
)

func InvoiceRouters(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("")
}
