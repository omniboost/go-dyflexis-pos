package timesheet

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-timesheetportal/utils"
)

func (c *Client) NewCustomReportsGetRequest() CustomReportsGetRequest {
	r := CustomReportsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomReportsGetRequest struct {
	client      *Client
	queryParams *CustomReportsGetRequestQueryParams
	pathParams  *CustomReportsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody CustomReportsGetRequestBody
}

func (r CustomReportsGetRequest) NewQueryParams() *CustomReportsGetRequestQueryParams {
	return &CustomReportsGetRequestQueryParams{}
}

type CustomReportsGetRequestQueryParams struct {
	Account string `schema:"Account"`
	// PageSize       int    `schema:"PageSize"`
}

func (p CustomReportsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomReportsGetRequest) QueryParams() *CustomReportsGetRequestQueryParams {
	return r.queryParams
}

func (r CustomReportsGetRequest) NewPathParams() *CustomReportsGetRequestPathParams {
	return &CustomReportsGetRequestPathParams{}
}

type CustomReportsGetRequestPathParams struct {
}

func (p *CustomReportsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomReportsGetRequest) PathParams() *CustomReportsGetRequestPathParams {
	return r.pathParams
}

func (r *CustomReportsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomReportsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomReportsGetRequest) Method() string {
	return r.method
}

func (r CustomReportsGetRequest) NewRequestBody() CustomReportsGetRequestBody {
	return CustomReportsGetRequestBody{}
}

type CustomReportsGetRequestBody struct {
}

func (r *CustomReportsGetRequest) RequestBody() *CustomReportsGetRequestBody {
	return nil
}

func (r *CustomReportsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *CustomReportsGetRequest) SetRequestBody(body CustomReportsGetRequestBody) {
	r.requestBody = body
}

func (r *CustomReportsGetRequest) NewResponseBody() *CustomReportsGetResponseBody {
	return &CustomReportsGetResponseBody{}
}

type CustomReportsGetResponseBody struct {
	Success      bool        `json:"success"`
	Message      interface{} `json:"message"`
	ReturnObject []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"returnObject"`
	RecordsUpdated      int           `json:"recordsUpdated"`
	RecordsCreated      int           `json:"recordsCreated"`
	RecordsDeleted      int           `json:"recordsDeleted"`
	RecordsCreatedCodes []interface{} `json:"recordsCreatedCodes"`
	ErrorCode           int           `json:"errorCode"`
}

func (r *CustomReportsGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/reports/customreports", r.PathParams())
	return &u, err
}

func (r *CustomReportsGetRequest) Do() (CustomReportsGetResponseBody, error) {
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
