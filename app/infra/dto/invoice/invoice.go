package invoice

import (
	"github.com/google/uuid"

	"github.com/SatoKeiju/shiharai-kun/app/domain/model/invoice"
)

// DTO : 請求書情報DTO
type DTO struct {
	ID                 string  `db:"id"`
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
func ModelFromDTO(d DTO) invoice.Invoice {
	i, err := invoice.New(d.CompanyID, d.ClientID, d.IssueDate, d.PaymentAmount, d.PaymentDueDate)
	if err != nil {
		//nolint:exhaustruct
		return invoice.Invoice{}
	}

	return i
}

func DTOFromModel(m invoice.Invoice) DTO {
	// TODO: IDの生成ロジックを置き場所も含めて検討
	newUUID, _ := uuid.NewRandom()

	return DTO{
		ID:                 newUUID.String(),
		CompanyID:          m.CompanyID,
		ClientID:           m.ClientID,
		IssueDate:          m.IssueDate,
		PaymentAmount:      m.PaymentAmount,
		Commission:         m.Commission,
		CommissionRate:     m.CommissionRate,
		ConsumptionTax:     m.ConsumptionTax,
		ConsumptionTaxRate: m.ConsumptionTaxRate,
		BillingAmount:      m.BillingAmount,
		PaymentDueDate:     m.PaymentDueDate,
		Status:             m.Status,
	}
}
