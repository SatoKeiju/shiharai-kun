package repository

import (
	"context"

	"github.com/SatoKeiju/shiharai-kun/app/domain/model"
)

// Invoice : 請求書に関するリポジトリ
type Invoice interface {
	FetchListByCompanyID(ctx context.Context, companyID string, from string, to string) ([]model.Invoice, error)
}
