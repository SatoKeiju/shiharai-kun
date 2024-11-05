package invoice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestModel_Invoice_New(t *testing.T) {
	t.Parallel()

	t.Run("正常系", func(t *testing.T) {
		t.Run("請求書ドメインモデルを生成できる", func(t *testing.T) {
			companyID := "testCompanyID"
			clientID := "testClientID"
			issueDate := "2024-10-01"
			paymentAmount := 10000
			paymentDueDate := "2024-10-25"

			want := Invoice{
				CompanyID:          "testCompanyID",
				ClientID:           "testClientID",
				IssueDate:          "2024-10-01",
				PaymentAmount:      10000,
				Commission:         400,
				CommissionRate:     0.04,
				ConsumptionTax:     40,
				ConsumptionTaxRate: 0.1,
				BillingAmount:      10440,
				PaymentDueDate:     "2024-10-25",
				Status:             statusPending.String(),
			}

			got, err := New(companyID, clientID, issueDate, paymentAmount, paymentDueDate)
			require.NoError(t, err)
			require.Equal(t, want, got)
		})
	})
}
