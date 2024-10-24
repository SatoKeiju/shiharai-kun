package usecase

import (
	"context"
	"github.com/SatoKeiju/shiharai-kun/app/domain"
)

// InvoiceUseCase : 請求書に関するユースケースが満たすべきユースケース
type InvoiceUseCase interface {
	FetchList(ctx context.Context, userID string, from string, to string) ([]domain.Invoice, error)
}

type invoiceUseCase struct{}

// NewInvoiceUseCase : 請求書に関するユースケースを作成
func NewInvoiceUseCase() InvoiceUseCase {
	return invoiceUseCase{}
}

func (u invoiceUseCase) FetchList(ctx context.Context, userID string, from string, to string) ([]domain.Invoice, error) {
	// TODO: 後続で実装
	return nil, nil
}
