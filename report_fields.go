package dyflexis

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-dyflexis/utils"
)

func (c *Client) NewReportFieldsRequest() ReportFieldsRequest {
	r := ReportFieldsRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ReportFieldsRequest struct {
	client      *Client
	queryParams *ReportFieldsRequestQueryParams
	pathParams  *ReportFieldsRequestPathParams
	method      string
	headers     http.Header
	requestBody ReportFieldsRequestBody
}

func (r ReportFieldsRequest) NewQueryParams() *ReportFieldsRequestQueryParams {
	return &ReportFieldsRequestQueryParams{}
}

type ReportFieldsRequestQueryParams struct {
	Account string `schema:"Account"`
	// PageSize       int    `schema:"PageSize"`
}

func (p ReportFieldsRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ReportFieldsRequest) QueryParams() *ReportFieldsRequestQueryParams {
	return r.queryParams
}

func (r ReportFieldsRequest) NewPathParams() *ReportFieldsRequestPathParams {
	return &ReportFieldsRequestPathParams{}
}

type ReportFieldsRequestPathParams struct {
}

func (p *ReportFieldsRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ReportFieldsRequest) PathParams() *ReportFieldsRequestPathParams {
	return r.pathParams
}

func (r *ReportFieldsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ReportFieldsRequest) SetMethod(method string) {
	r.method = method
}

func (r *ReportFieldsRequest) Method() string {
	return r.method
}

func (r ReportFieldsRequest) NewRequestBody() ReportFieldsRequestBody {
	return ReportFieldsRequestBody{}
}

type ReportFieldsRequestBody struct {
}

func (r *ReportFieldsRequest) RequestBody() *ReportFieldsRequestBody {
	return nil
}

func (r *ReportFieldsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *ReportFieldsRequest) SetRequestBody(body ReportFieldsRequestBody) {
	r.requestBody = body
}

func (r *ReportFieldsRequest) NewResponseBody() *ReportFieldsResponseBody {
	return &ReportFieldsResponseBody{}
}

type ReportFieldsResponseBody []struct {
	ReportFieldID   int    `json:"reportFieldId"`
	ReportFieldName string `json:"reportFieldName"`
	Description     string `json:"description"`
	DataType        int    `json:"dataType"`
}

func (r *ReportFieldsRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/reports/reportfields", r.PathParams())
	return &u, err
}

func (r *ReportFieldsRequest) Do() (ReportFieldsResponseBody, error) {
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
