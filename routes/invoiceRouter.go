package routes

import (
	controller "golang-restourant-management/controlles"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controller.GetInvoice())
	incomingRoutes.POST("/invoices", controller.GetInvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id", controller.UpdateInvoice())
}
