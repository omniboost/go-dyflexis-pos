package dyflexis

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/omniboost/go-dyflexis-pos/utils"
)

func (c *Client) NewDyflexisPOSPostRequest() DyflexisPOSPostRequest {
	r := DyflexisPOSPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type DyflexisPOSPostRequest struct {
	client      *Client
	queryParams *DyflexisPOSPostRequestQueryParams
	pathParams  *DyflexisPOSPostRequestPathParams
	method      string
	headers     http.Header
	requestBody DyflexisPOSPostRequestBody
}

func (r DyflexisPOSPostRequest) NewQueryParams() *DyflexisPOSPostRequestQueryParams {
	return &DyflexisPOSPostRequestQueryParams{}
}

type DyflexisPOSPostRequestQueryParams struct {
}

func (p DyflexisPOSPostRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(Time{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *DyflexisPOSPostRequest) QueryParams() *DyflexisPOSPostRequestQueryParams {
	return r.queryParams
}

func (r DyflexisPOSPostRequest) NewPathParams() *DyflexisPOSPostRequestPathParams {
	return &DyflexisPOSPostRequestPathParams{}
}

type DyflexisPOSPostRequestPathParams struct{}

func (p *DyflexisPOSPostRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DyflexisPOSPostRequest) PathParams() *DyflexisPOSPostRequestPathParams {
	return r.pathParams
}

func (r *DyflexisPOSPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *DyflexisPOSPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *DyflexisPOSPostRequest) Method() string {
	return r.method
}

func (r DyflexisPOSPostRequest) NewRequestBody() DyflexisPOSPostRequestBody {
	return DyflexisPOSPostRequestBody{}
}

type DyflexisPOSPostRequestBody struct {
	BusinessDate Date `json:"BusinessDate"`
	Reports      struct {
		Turnover struct {
			Turnover json.Number `json:"turnover"`
			Tax      json.Number `json:"tax"`
		} `json:"Turnover"`
		Hourly []struct {
			Time     Date        `json:"time"`
			Turnover json.Number `json:"turnover"`
			Tax      json.Number `json:"tax"`
		} `json:"Hourly"`
		Departments []struct {
			DepartmentID int         `json:"departmentID"`
			Turnover     json.Number `json:"turnover"`
			Tax          json.Number `json:"tax"`
		} `json:"Departments"`
	} `json:"Reports"`
}

func (r *DyflexisPOSPostRequest) RequestBody() *DyflexisPOSPostRequestBody {
	return &r.requestBody
}

func (r *DyflexisPOSPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *DyflexisPOSPostRequest) SetRequestBody(body DyflexisPOSPostRequestBody) {
	r.requestBody = body
}

func (r *DyflexisPOSPostRequest) NewResponseBody() *InformationStreamPutResponseBody {
	return &InformationStreamPutResponseBody{}
}

type InformationStreamPutResponseBody struct {
	Date  string `json:"date"`
	Value int    `json:"value"`
}

func (r *DyflexisPOSPostRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/dyflexispos", r.PathParams())
	return &u, err
}

func (r *DyflexisPOSPostRequest) Do() (InformationStreamPutResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
