package usecase

import "korp/model"

type InvoiceUseCase struct {
}

func NewInvoiceUseCase() InvoiceUseCase {
	return InvoiceUseCase{}
}

func (pu *InvoiceUseCase) GetInvoice() ([]model.Invoice, error) {
	return []model.Invoice{}, nil
}
