package domain

// Invoice : 請求書ドメインモデル
// TODO: プライベートフィールドにしてゲッター経由でのみ取得できるようにする
type Invoice struct {
	IssueDate          string
	PaymentAmount      int
	Commission         int
	CommissionRate     float64
	ConsumptionTax     int
	ConsumptionTaxRate float64
	BillingAmount      int
	PaymentDueDate     string
	Status             string
}

// NewInvoiceModel : 請求書の値から請求書ドメインモデルを生成
// TODO: バリデーションや値オブジェクトなどを追記
func NewInvoiceModel(
	issueDate string,
	paymentAmount int,
	commission int,
	commissionRate float64,
	consumptionTax int,
	consumptionTaxRate float64,
	billingAmount int,
	paymentDueDate string,
	status string,
) Invoice {
	return Invoice{
		IssueDate:          issueDate,
		PaymentAmount:      paymentAmount,
		Commission:         commission,
		CommissionRate:     commissionRate,
		ConsumptionTax:     consumptionTax,
		ConsumptionTaxRate: consumptionTaxRate,
		BillingAmount:      billingAmount,
		PaymentDueDate:     paymentDueDate,
		Status:             status,
	}
}
