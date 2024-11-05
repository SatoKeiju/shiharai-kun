package dao

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/SatoKeiju/shiharai-kun/app/domain/model"
	"github.com/SatoKeiju/shiharai-kun/app/domain/repository"
)

type invoice struct {
	rdb *sqlx.DB
}

// NewInvoice : Invoiceリポジトリを生成
func NewInvoice(rdb *sqlx.DB) repository.Invoice {
	return invoice{rdb: rdb}
}

// FetchListByCompanyID : 指定した企業IDに紐づく期間内の請求書を一覧で取得
func (i invoice) FetchListByCompanyID(ctx context.Context, companyID string, from string, to string) ([]model.Invoice, error) {
	// TODO implement me
	return nil, nil
}
