package design

import (
	"goa.design/goa/v3/dsl"
)

var _ = dsl.Service("invoices", func() {
	dsl.Description("invoice service")

	dsl.HTTP(func() {
		dsl.Path("/invoices")
	})

	dsl.Error("bad_request", ErrBadRequest)
	dsl.Error("internal_server_error", ErrInternalServerError)

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
