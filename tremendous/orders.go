package tremendous

import (
	"time"
)

type OrdersService struct {
	base *Tremendous
}

type Refund struct {
	Total float64 `json:"total"`
}
type Payment struct {
	FundingSourceID string  `json:"funding_source_id,omitempty"`
	Subtotal        float64 `json:"subtotal,omitempty"`
	Total           float64 `json:"total,omitempty"`
	Fees            float64 `json:"fees,omitempty"`
	Refund          Refund  `json:"refund,omitempty"`
	Channel         string  `json:"channel,omitempty"`
}
type Value struct {
	Denomination float64 `json:"denomination,omitempty"`
	CurrencyCode string  `json:"currency_code,omitempty"`
}
type Delivery struct {
	Method string `json:"method,omitempty"`
	Status string `json:"status,omitempty"`
	Link   string `json:"link,omitempty"`
}
type Recipient struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}
type CustomField struct {
	ID    string `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
}
type Reward struct {
	CampaignID   string        `json:"campaign_id,omitempty"`
	Products     []string      `json:"products,omitempty"`
	ID           string        `json:"id,omitempty"`
	OrderID      string        `json:"order_id,omitempty"`
	CreatedAt    time.Time     `json:"created_at,omitempty"`
	Value        Value         `json:"value,omitempty"`
	Delivery     Delivery      `json:"delivery,omitempty"`
	Recipient    Recipient     `json:"recipient,omitempty"`
	CustomFields []CustomField `json:"custom_fields,omitempty"`
}

type Order struct {
	ID         string    `json:"id,omitempty"`
	ExternalID string    `json:"external_id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	Status     string    `json:"status,omitempty"`
	Payment    Payment   `json:"payment,omitempty"`
	Rewards    []Reward  `json:"rewards,omitempty"`
	InvoiceID  string    `json:"invoice_id,omitempty"`
}

type ListOrdersResponse struct {
	Orders     []Order `json:"orders"`
	TotalCount int     `json:"total_count"`
}

type CreateOrderResponse struct {
	Order Order `json:"order"`
}

type CreateOrderRequest struct {
	ExternalID string   `json:"external_id"`
	Payment    Payment  `json:"payment"`
	Rewards    []Reward `json:"rewards"`
}

func newOrdersService(t *Tremendous) *OrdersService {
	return &OrdersService{
		base: t,
	}
}

func (o *OrdersService) List() (*ListOrdersResponse, error) {
	resp := &ListOrdersResponse{}
	_, err := o.base.Request("orders", "GET", nil, resp)
	return resp, err
}

func (o *OrdersService) Create(
	funding string,
	products []string,
	denomination float64,
	currency string,
	delivery string,
	recipient Recipient,
) (*CreateOrderResponse, error) {

	params := &CreateOrderRequest{
		Payment: Payment{
			FundingSourceID: funding,
		},
		Rewards: []Reward{
			{
				Value: Value{
					Denomination: denomination,
					CurrencyCode: currency,
				},
				Delivery: Delivery{
					Method: delivery,
				},
				Recipient: recipient,
				Products:  products,
			},
		},
	}

	resp := &CreateOrderResponse{}
	_, err := o.base.Request("orders", "POST", params, resp)

	return resp, err
}
