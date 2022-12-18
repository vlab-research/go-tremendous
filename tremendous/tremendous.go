package tremendous

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
	"strings"
)

type Tremendous struct {
	apiKey string
	sling  *sling.Sling
}

type Client struct {
	tremendous *Tremendous
	Orders     *OrdersService
	Products   *ProductsService
	Funding    *FundingService
}

type ApiError struct {
	Errors struct {
		Message string          `json:"message"`
		Payload json.RawMessage `json:"payload"`
	} `json:"errors"`
}

func (e *ApiError) Error() string {
	return e.Errors.Message
}

func (e *ApiError) Empty() bool {
	return e.Errors.Message == ""
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client, baseUrl string, apiKey string) *Client {

	base := sling.New().Client(httpClient).Base(baseUrl + "/api/v2/")

	t := &Tremendous{
		apiKey: apiKey,
		sling:  base,
	}

	return &Client{
		tremendous: t,
		Orders:     newOrdersService(t),
		Products:   newProductsService(t),
		Funding:    newFundingService(t),
	}
}

func (t *Tremendous) Request(path string, method string, params interface{}, resp interface{}) (*http.Response, error) {

	s := t.sling.New()
	s = s.Set("Accept", "application/json")
	s = s.Set("Authorization", "Bearer "+t.apiKey)

	switch strings.ToUpper(method) {
	case "GET":
		s = s.Get(path).QueryStruct(params)
	case "POST":
		s = s.Post(path).BodyJSON(params)
	}

	fmt.Println(s.New().Request())

	apiError := &ApiError{}
	httpResponse, err := s.New().Receive(resp, apiError)

	if err != nil {
		return httpResponse, err
	}

	if !apiError.Empty() {
		fmt.Println(string(apiError.Errors.Payload))
		return httpResponse, apiError
	}

	return httpResponse, nil
}
