package design

import "goa.design/goa/v3/dsl"

// ErrBadRequest : 400 Bad Request
var ErrBadRequest = dsl.ResultType("application/vnd.err.bad_request", func() {
	dsl.Description("400 Bad Request")
	dsl.Attributes(func() {
		dsl.Attribute("message", dsl.String, func() {
			dsl.Example("不正なリクエスト")
		})

		dsl.Required("message")
	})
})

// ErrInternalServerError : 500 Internal Server Error
var ErrInternalServerError = dsl.ResultType("application/vnd.err.internal_server_error", func() {
	dsl.Description("500 Internal Server Error")
	dsl.Attributes(func() {
		dsl.Attribute("message", dsl.String, func() {
			dsl.Example("サーバーエラー")
		})

		dsl.Required("message")
	})
})
