package design

import "goa.design/goa/v3/dsl"

// invoice : 請求書
var invoice = dsl.Type("invoice", func() {
	dsl.Attribute("issue_date", dsl.String, "発行日", func() {
		dsl.Example("2024-10-01")
	})
	dsl.Attribute("payment_amount", dsl.Int, "支払金額", func() {
		dsl.Minimum(1)
		dsl.Example(10000)
	})
	dsl.Attribute("commission", dsl.Int, "手数料", func() {
		dsl.Minimum(0)
		dsl.Example(400)
	})
	dsl.Attribute("commission_rate", dsl.Float64, "手数料率", func() {
		dsl.Minimum(0)
		dsl.Example(0.04)
	})
	dsl.Attribute("consumption_tax", dsl.Int, "消費税", func() {
		dsl.Minimum(0)
		dsl.Example(40)
	})
	dsl.Attribute("consumption_tax_rate", dsl.Float64, "消費税率", func() {
		dsl.Minimum(0)
		dsl.Example(0.1)
	})
	dsl.Attribute("billing_amount", dsl.Int, "請求金額", func() {
		dsl.Minimum(1)
		dsl.Example(10440)
	})
	dsl.Attribute("payment_due_date", dsl.String, "支払期日", func() {
		dsl.Example("2024-10-27")
	})
	dsl.Attribute("status", dsl.String, "ステータス", func() {
		dsl.Enum("未処理", "処理中", "支払い済み", "エラー")
		dsl.Example("未処理")
	})

	dsl.Required("issue_date", "payment_amount", "commission", "commission_rate", "consumption_tax", "consumption_tax_rate", "billing_amount", "payment_due_date", "status")
})
