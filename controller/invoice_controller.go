package controller

import (
	"korp/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type invoiceController struct {
	//usacase
}

func NewInvoiceController() invoiceController {
	return invoiceController{}
}

func (i *invoiceController) GetInvoice(ctx *gin.Context) {

	invoices := []model.Invoice{
		{
			ID:             1,
			CURRENT_STATUS: "Aberta",
		},
	}

	ctx.JSON(http.StatusOK, invoices)

}
