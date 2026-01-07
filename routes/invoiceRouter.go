package routes

import (
	"github.com/gin-gonic/gin"
)

func InvoiceRouters(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controller.GetInvoice())
	incomingRoutes.POST("/invoices", controller.GetInvoice())
	incomingRoutes.PATCH("invoices/:invoice_id", controller.UpdateInvoice())
}
