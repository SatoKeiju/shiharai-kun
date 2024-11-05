package invoice

import "github.com/SatoKeiju/shiharai-kun/app/domain/model/invoice"

// DTO : 請求書情報DTO
type DTO struct {
	CompanyID          string  `db:"company_id"`
	ClientID           string  `db:"client_id"`
	IssueDate          string  `db:"issue_date"`
	PaymentAmount      int     `db:"payment_amount"`
	Commission         int     `db:"commission"`
	CommissionRate     float64 `db:"commission_rate"`
	ConsumptionTax     int     `db:"consumption_tax"`
	ConsumptionTaxRate float64 `db:"consumption_tax_rate"`
	BillingAmount      int     `db:"billing_amount"`
	PaymentDueDate     string  `db:"payment_due_date"`
	Status             string  `db:"status"`
}

// ModelFromDTO : 請求書情報DTOから請求書ドメインモデルを生成
// TODO: バリデーションなども含めたモデル生成関数を払い出してそれ経由で生成
func ModelFromDTO(d DTO) invoice.Invoice {
	return invoice.Invoice{
		CompanyID:          d.CompanyID,
		ClientID:           d.ClientID,
		IssueDate:          d.IssueDate,
		PaymentAmount:      d.PaymentAmount,
		Commission:         d.Commission,
		CommissionRate:     d.CommissionRate,
		ConsumptionTax:     d.ConsumptionTax,
		ConsumptionTaxRate: d.ConsumptionTaxRate,
		BillingAmount:      d.BillingAmount,
		PaymentDueDate:     d.PaymentDueDate,
		Status:             d.Status,
	}
}
