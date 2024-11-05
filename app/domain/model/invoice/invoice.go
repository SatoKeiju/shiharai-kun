package invoice

import (
	"fmt"
	"math"
	"time"
)

const (
	// commissionRate : 手数料率
	commissionRate = 0.04
	// consumptionTaxRate : 消費税率
	consumptionTaxRate = 0.1
)

// calculateCommission : 手数料を算出
func calculateCommission(price int) int {
	// NOTE_202411: 小数点以下は切り上げ
	return int(math.Ceil(float64(price) * commissionRate))
}

// consumptionTax : 消費税価格を算出
func consumptionTax(price int) int {
	// NOTE_202411: 小数点以下は切り下げ
	return int(float64(price) * consumptionTaxRate)
}

func billingAmount(paymentAmount int) int {
	commission := calculateCommission(paymentAmount)

	return paymentAmount + commission + consumptionTax(commission)
}

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
	if err := validate(issueDate, paymentAmount, paymentDueDate); err != nil {
		return Invoice{}, fmt.Errorf("バリデーションエラー: %w", err)
	}

	commission := calculateCommission(paymentAmount)

	return Invoice{
		CompanyID:          companyID,
		ClientID:           clientID,
		IssueDate:          issueDate,
		PaymentAmount:      paymentAmount,
		Commission:         commission,
		CommissionRate:     commissionRate,
		ConsumptionTax:     consumptionTax(commission),
		ConsumptionTaxRate: consumptionTaxRate,
		BillingAmount:      billingAmount(paymentAmount),
		PaymentDueDate:     paymentDueDate,
		Status:             statusPending.String(),
	}, nil
}

// validate : 請求書ドメインモデルを生成するための値をバリデーション
func validate(issueDate string, paymentAmount int, paymentDueDate string) error {
	// issueDateについて
	if _, err := time.Parse(time.DateOnly, issueDate); err != nil {
		return fmt.Errorf("issueDate(%s)はフォーマット(%s)に変換できません: %w", issueDate, time.DateOnly, err)
	}

	// paymentAmountについて
	if paymentAmount < 1 {
		return fmt.Errorf("paymentAmount(%d)が1未満の整数です", paymentAmount)
	}

	// paymentDueDateについて
	if _, err := time.Parse(time.DateOnly, paymentDueDate); err != nil {
		return fmt.Errorf("paymentDueDate(%s)はフォーマット(%s)に変換できません: %w", paymentDueDate, time.DateOnly, err)
	}

	return nil
}
