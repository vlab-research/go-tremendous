package tremendous

import (
	"time"
)

type FundingService struct {
	base *Tremendous
}

type Meta struct {
	AvailableCents      int       `json:"available_cents"`
	PendingCents        int       `json:"pending_cents"`
	AccountholderName   string    `json:"accountholder_name"`
	AccountType         string    `json:"account_type"`
	BankName            string    `json:"bank_name"`
	AccountNumberMask   string    `json:"account_number_mask"`
	AccountRoutingMask  string    `json:"account_routing_mask"`
	Refundable          bool      `json:"refundable"`
	Network             string    `json:"network"`
	Last4               string    `json:"last4"`
	Expired             bool      `json:"expired"`
	LastPaymentFailedAt time.Time `json:"last_payment_failed_at"`
}

type FundingSource struct {
	ID     string `json:"id"`
	Method string `json:"method"`
	Type   string `json:"type"`
	Meta   Meta   `json:"meta"`
}

type ListFundingResponse struct {
	FundingSources []FundingSource `json:"funding_sources"`
}

func newFundingService(t *Tremendous) *FundingService {
	return &FundingService{
		base: t,
	}
}

func (p *FundingService) List() (*ListFundingResponse, error) {
	resp := &ListFundingResponse{}

	_, err := p.base.Request("funding_sources", "GET", nil, resp)
	return resp, err
}
