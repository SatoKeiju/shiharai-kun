package design

import (
	"goa.design/goa/v3/dsl"
)

var _ = dsl.API("shiharai-kun", func() {
	dsl.Title("shiharai kun service")
	dsl.Description("shiharai kun API service")

	dsl.Server("shiharai-kun", func() {
		dsl.Host("local", func() { dsl.URI("http://localhost:8088") })
	})
	dsl.HTTP(func() {
		dsl.Path("/api")
	})
})
