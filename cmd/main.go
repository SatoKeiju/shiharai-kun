package main

import (
	"net/http"

	goahttp "goa.design/goa/v3/http"

	"github.com/SatoKeiju/shiharai-kun/app/ui"
	"github.com/SatoKeiju/shiharai-kun/app/usecase"
	"github.com/SatoKeiju/shiharai-kun/gen/http/invoices/server"
	goainvoice "github.com/SatoKeiju/shiharai-kun/gen/invoices"
)

func main() {
	// HTTP muxerを作成
	mux := goahttp.NewMuxer()
	// HTTPリクエストデコータをセット
	dec := goahttp.RequestDecoder
	// HTTPレスポンスエンコーダをセット
	enc := goahttp.ResponseEncoder

	// 各種サービスを初期化

	// 請求書
	invoiceUseCase := usecase.NewInvoiceUseCase()
	invoiceService := ui.NewInvoiceService(invoiceUseCase)
	invoiceEndpoints := goainvoice.NewEndpoints(invoiceService)
	invoiceServer := server.New(invoiceEndpoints, mux, dec, enc, nil, nil)

	server.Mount(mux, invoiceServer)

	httpsvr := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
	if err := httpsvr.ListenAndServe(); err != nil {
		panic(err)
	}
}
