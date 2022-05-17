package dyflexis

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-dyflexis/utils"
)

func (c *Client) NewTimesheetsGetRequest() TimesheetsGetRequest {
	r := TimesheetsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type TimesheetsGetRequest struct {
	client      *Client
	queryParams *TimesheetsGetRequestQueryParams
	pathParams  *TimesheetsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody TimesheetsGetRequestBody
}

func (r TimesheetsGetRequest) NewQueryParams() *TimesheetsGetRequestQueryParams {
	return &TimesheetsGetRequestQueryParams{}
}

type TimesheetsGetRequestQueryParams struct {
	TimesheetID int `schema:"TimeSheetId,omitempty"`
}

func (p TimesheetsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *TimesheetsGetRequest) QueryParams() *TimesheetsGetRequestQueryParams {
	return r.queryParams
}

func (r TimesheetsGetRequest) NewPathParams() *TimesheetsGetRequestPathParams {
	return &TimesheetsGetRequestPathParams{}
}

type TimesheetsGetRequestPathParams struct {
}

func (p *TimesheetsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *TimesheetsGetRequest) PathParams() *TimesheetsGetRequestPathParams {
	return r.pathParams
}

func (r *TimesheetsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *TimesheetsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *TimesheetsGetRequest) Method() string {
	return r.method
}

func (r TimesheetsGetRequest) NewRequestBody() TimesheetsGetRequestBody {
	return TimesheetsGetRequestBody{}
}

type TimesheetsGetRequestBody struct {
}

func (r *TimesheetsGetRequest) RequestBody() *TimesheetsGetRequestBody {
	return nil
}

func (r *TimesheetsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *TimesheetsGetRequest) SetRequestBody(body TimesheetsGetRequestBody) {
	r.requestBody = body
}

func (r *TimesheetsGetRequest) NewResponseBody() *TimesheetsGetResponseBody {
	return &TimesheetsGetResponseBody{}
}

type TimesheetsGetResponseBody Timesheet

func (r *TimesheetsGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/v2/timesheets", r.PathParams())
	return &u, err
}

func (r *TimesheetsGetRequest) Do() (TimesheetsGetResponseBody, error) {
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
