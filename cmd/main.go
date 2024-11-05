package main

import (
	"fmt"
	"net/http"
	"time"

	goahttp "goa.design/goa/v3/http"

	"github.com/SatoKeiju/shiharai-kun/app/infra/dao"
	"github.com/SatoKeiju/shiharai-kun/app/infra/db"
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

	// 各種クライアント
	rdb, err := db.NewRDB()
	if err != nil {
		panic(fmt.Errorf("DBクライアントの初期化に失敗しました: %w", err))
	}

	// 各種サービスを初期化

	// ユーザー
	userRepository := dao.NewUser(rdb)
	// 請求書
	invoiceRepository := dao.NewInvoice(rdb)
	invoiceUseCase := usecase.NewInvoiceUseCase(invoiceRepository, userRepository)
	invoiceService := ui.NewInvoiceService(invoiceUseCase)
	invoiceEndpoints := goainvoice.NewEndpoints(invoiceService)
	invoiceServer := server.New(invoiceEndpoints, mux, dec, enc, nil, nil)

	server.Mount(mux, invoiceServer)

	//nolint:exhaustruct
	httpsvr := &http.Server{
		Addr:              ":8081",
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}
	if err := httpsvr.ListenAndServe(); err != nil {
		panic(err)
	}
}
