package invoice

// Invoice : 請求書ドメインモデル
// TODO: プライベートフィールドにしてゲッター経由でのみ取得できるようにする
type Invoice struct {
	CompanyID          string
	ClientID           string
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

// New : 請求書の値から請求書ドメインモデルを生成
func New(companyID string, clientID string, issueDate string, paymentAmount int, paymentDueDate string) (Invoice, error) {
	// TODO: バリデーションや値オブジェクトなどを追記
	return Invoice{
		CompanyID:          companyID,
		ClientID:           clientID,
		IssueDate:          issueDate,
		PaymentAmount:      paymentAmount,
		Commission:         0,
		CommissionRate:     0,
		ConsumptionTax:     0,
		ConsumptionTaxRate: 0,
		BillingAmount:      0,
		PaymentDueDate:     paymentDueDate,
		Status:             "",
	}, nil
}
