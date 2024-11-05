package dao

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	model "github.com/SatoKeiju/shiharai-kun/app/domain/model/invoice"
	"github.com/SatoKeiju/shiharai-kun/app/domain/repository"
	dto "github.com/SatoKeiju/shiharai-kun/app/infra/dto/invoice"
)

type invoice struct {
	rdb *sqlx.DB
}

// NewInvoice : Invoiceリポジトリを生成
func NewInvoice(rdb *sqlx.DB) repository.Invoice {
	return invoice{rdb: rdb}
}

func (i invoice) Create(_ context.Context, model model.Invoice) (model.Invoice, error) {
	// TODO: 実装
	return model, nil
}

// FetchListByCompanyID : 指定した企業IDに紐づく期間内の請求書を一覧で取得
func (i invoice) FetchListByCompanyID(ctx context.Context, companyID string, from string, to string) ([]model.Invoice, error) {
	q := "SELECT * FROM invoices WHERE company_id = ? AND payment_due_date BETWEEN ? AND ?"
	var dList []dto.DTO
	if err := i.rdb.SelectContext(ctx, dList, q, companyID, from, to); err != nil {
		return nil, fmt.Errorf("rdb.SelectContext(companyID: %s, from: %s, to: %s): %w", companyID, from, to, err)
	}

	list := make([]model.Invoice, len(dList))
	for j, invoiceDTO := range dList {
		list[j] = dto.ModelFromDTO(invoiceDTO)
	}

	return list, nil
}
