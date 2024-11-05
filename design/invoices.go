package design

import (
	"goa.design/goa/v3/dsl"
)

var _ = dsl.Service("invoices", func() {
	dsl.Description("invoice.go service")

	dsl.HTTP(func() {
		dsl.Path("/invoices")
	})

	dsl.Error("bad_request", ErrBadRequest)
	dsl.Error("internal_server_error", ErrInternalServerError)

	dsl.Method("create", func() {
		dsl.Meta("swagger:summary", "請求書データを作成")
		dsl.Payload(func() {
			dsl.Attribute("user_id", dsl.String, "ユーザID")
			dsl.Attribute("client_id", dsl.String, "取引先ID")
			dsl.Attribute("issue_date", dsl.String, "発行日", func() {
				dsl.Example("2024-10-01")
			})
			dsl.Attribute("payment_amount", dsl.Int, "支払金額", func() {
				dsl.Minimum(1)
				dsl.Example(10000)
			})
			dsl.Attribute("payment_due_date", dsl.String, "支払期日", func() {
				dsl.Example("2024-10-27")
			})

			dsl.Required("user_id", "client_id", "issue_date", "payment_amount", "payment_due_date")
		})

		dsl.Result(invoice)

		dsl.HTTP(func() {
			dsl.POST("/")
			dsl.Param("user_id")
			dsl.Response(dsl.StatusCreated)
			dsl.Response("bad_request", dsl.StatusBadRequest)
			dsl.Response("internal_server_error", dsl.StatusInternalServerError)
		})
	})

	dsl.Method("fetch list", func() {
		dsl.Meta("swagger:summary", "指定期間内に支払いが発生する請求書データの一覧を取得")
		dsl.Payload(func() {
			dsl.Attribute("user_id", dsl.String, "ユーザID")
			dsl.Attribute("from_date", dsl.String, func() {
				dsl.Description("指定する期間の開始日")
				dsl.Example("2024-10-01")
			})
			dsl.Attribute("to_date", dsl.String, func() {
				dsl.Description("指定する期間の最終日")
				dsl.Example("2024-10-25")
			})

			dsl.Required("user_id", "from_date", "to_date")
		})

		dsl.Result(dsl.ArrayOf(invoice))

		dsl.HTTP(func() {
			dsl.GET("/")
			dsl.Param("user_id")
			dsl.Param("from_date")
			dsl.Param("to_date")
			dsl.Response(dsl.StatusOK)
			dsl.Response("bad_request", dsl.StatusBadRequest)
			dsl.Response("internal_server_error", dsl.StatusInternalServerError)
		})
	})
})
