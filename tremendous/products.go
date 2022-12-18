package tremendous

type ProductsService struct {
	base *Tremendous
}

type Countries struct {
	Abbr string `json:"abbr"`
}
type Products struct {
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	Category      string      `json:"category"`
	CurrencyCodes []string    `json:"currency_codes"`
	Countries     []Countries `json:"countries"`
}

type ListProductsResponse struct {
	Products []Products `json:"products"`
}

type ListProductsRequest struct {
	Country  string `url:"country,omitempty"`
	Currency string `url:"currency,omitempty"`
}

func newProductsService(t *Tremendous) *ProductsService {
	return &ProductsService{
		base: t,
	}
}

func (p *ProductsService) List(country, currency string) (*ListProductsResponse, error) {

	resp := &ListProductsResponse{}

	params := &ListProductsRequest{
		Country:  country,
		Currency: currency,
	}

	_, err := p.base.Request("products", "GET", params, resp)

	return resp, err
}
