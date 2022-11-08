package paymentintent

import (
	"github.com/li31727/oxpay-go"
	"net/http"
)

// Client is used to invoke /payment_links APIs.
type Client struct {
	B          oxpay.Backend
	McpTID     string
	ApiBackend oxpay.SupportedBackend
}

// New creates a new payment link.
func New(params *oxpay.PaymentIntentParams) (*oxpay.PaymentIntent, error) {
	return getC().New(params)
}

func getC() Client {
	return Client{oxpay.GetBackend(oxpay.APIBackend), oxpay.McpTID, oxpay.APIBackend}
}
func (c Client) getPath(relativePath string) string {
	return "/" + string(c.ApiBackend) + relativePath
}

// New creates a new payment just for card quick pay.
func (c Client) New(params *oxpay.PaymentIntentParams) (*oxpay.PaymentIntent, error) {
	paymentlink := &oxpay.PaymentIntent{}
	params.Data = params
	err := c.B.Call(
		http.MethodPost,
		c.getPath("/v5/sale"),
		c.McpTID,
		&params.Params,
		paymentlink,
	)
	return paymentlink, err
}
func Get(id string, params *oxpay.PaymentIntentParams) (*oxpay.PaymentIntent, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *oxpay.PaymentIntentParams) (*oxpay.PaymentIntent, error) {
	paymentintent := &oxpay.PaymentIntent{}
	params.TransactionId = id
	params.Data = params

	err := c.B.Call(
		http.MethodPost,
		c.getPath("/v5/query/detail"),
		c.McpTID,
		&params.Params,
		paymentintent,
	)
	return paymentintent, err
}

func Cancel(id string, params *oxpay.PaymentIntentParams) (*oxpay.PaymentIntent, error) {
	return getC().Cancel(id, params)
}
func (c Client) Cancel(id string, params *oxpay.PaymentIntentParams) (*oxpay.PaymentIntent, error) {
	paymentintent := &oxpay.PaymentIntent{}
	params.TransactionId = id
	params.Data = params

	err := c.B.Call(
		http.MethodPost,
		c.getPath("/v5/ewallet/revokeOrder"),
		c.McpTID,
		&params.Params,
		paymentintent,
	)
	return paymentintent, err
}
