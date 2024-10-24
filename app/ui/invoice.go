package ui

import (
	"context"
	"fmt"

	"github.com/SatoKeiju/shiharai-kun/app/usecase"
	gen "github.com/SatoKeiju/shiharai-kun/gen/invoices"
)

type invoicesService struct {
	u usecase.InvoiceUseCase
}

// NewInvoiceService : 請求書サービスを作成
func NewInvoiceService(u usecase.InvoiceUseCase) gen.Service {
	return &invoicesService{u: u}
}

// FetchList : 指定期間内に支払いが発生する請求書データの一覧を取得
func (s *invoicesService) FetchList(ctx context.Context, p *gen.FetchListPayload) ([]*gen.Invoice, error) {
	list, err := s.u.FetchList(ctx, p.UserID, p.FromDate, p.ToDate)
	if err != nil {
		return nil, fmt.Errorf("InvoiceUseCase.FetchList(userID: %s, from: %s, to: %s): %w", p.UserID, p.FromDate, p.ToDate, err)
	}

	res := make([]*gen.Invoice, len(list))
	for i, invoice := range list {
		res[i] = &gen.Invoice{
			IssueDate:          invoice.IssueDate,
			PaymentAmount:      invoice.PaymentAmount,
			Commission:         invoice.Commission,
			CommissionRate:     invoice.CommissionRate,
			ConsumptionTax:     invoice.ConsumptionTax,
			ConsumptionTaxRate: invoice.ConsumptionTaxRate,
			BillingAmount:      invoice.BillingAmount,
			PaymentDueDate:     invoice.PaymentDueDate,
			Status:             invoice.Status,
		}
	}
	return res, nil
}
