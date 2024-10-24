// Code generated by goa v3.19.1, DO NOT EDIT.
//
// invoices client HTTP transport
//
// Command:
// $ goa gen github.com/SatoKeiju/shiharai-kun/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the invoices service endpoint HTTP clients.
type Client struct {
	// FetchList Doer is the HTTP client used to make requests to the fetch list
	// endpoint.
	FetchListDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the invoices service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		FetchListDoer:       doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// FetchList returns an endpoint that makes HTTP requests to the invoices
// service fetch list server.
func (c *Client) FetchList() goa.Endpoint {
	var (
		encodeRequest  = EncodeFetchListRequest(c.encoder)
		decodeResponse = DecodeFetchListResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildFetchListRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.FetchListDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("invoices", "fetch list", err)
		}
		return decodeResponse(resp)
	}
}
