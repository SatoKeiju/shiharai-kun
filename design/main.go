package design

import (
	"goa.design/goa/v3/dsl"

	"github.com/SatoKeiju/shiharai-kun/app"
)

var _ = dsl.API("pointclub", func() {
	dsl.Title("DMM PointClub Service")
	dsl.Description("DMM PointClub API System")

	dsl.Server("pointclub", func() {
		if app.IsStg() {
			dsl.Host("stg", func() { dsl.URI("https://stg.api.pointclub.dmm.com/") })
		} else {
			dsl.Host("local", func() { dsl.URI("http://localhost:8088") })
		}
	})
	dsl.HTTP(func() {
		dsl.Path("/v1")
	})
})
