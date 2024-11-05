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

// Create : 請求書データを保存
func (i invoice) Create(ctx context.Context, m model.Invoice) (model.Invoice, error) {
	q := "INSERT INTO invoices (id, company_id, client_id, issue_date, payment_amount, commission, commission_rate, consumption_tax, consumption_tax_rate, billing_amount, payment_due_date, status)"

	d := dto.DTOFromModel(m)
	if _, err := i.rdb.ExecContext(ctx, q, d.ID, d.CompanyID, d.ClientID, d.IssueDate, d.PaymentAmount, d.Commission, d.CommissionRate, d.ConsumptionTax, d.ConsumptionTaxRate, d.BillingAmount, d.PaymentDueDate, d.Status); err != nil {
		return model.Invoice{}, fmt.Errorf("rdb.ExecContext(model: %+v): %w", m, err)
	}

	return m, nil
}

// FetchListByCompanyID : 指定した企業IDに紐づく期間内の請求書を一覧で取得
func (i invoice) FetchListByCompanyID(ctx context.Context, companyID string, from string, to string) ([]model.Invoice, error) {
	// TODO: "*"をやめて取得するカラムを全て書く
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
