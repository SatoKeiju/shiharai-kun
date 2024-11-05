package repository

import (
	"context"

	"github.com/SatoKeiju/shiharai-kun/app/domain/model/invoice"
)

// Invoice : 請求書に関するリポジトリ
type Invoice interface {
	Create(ctx context.Context, model invoice.Invoice) (invoice.Invoice, error)
	FetchListByCompanyID(ctx context.Context, companyID string, from string, to string) ([]invoice.Invoice, error)
}
