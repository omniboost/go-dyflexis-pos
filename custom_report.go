package dyflexis

import (
	"net/http"
	"net/url"

	"github.com/cydev/zero"
	"github.com/omniboost/go-dyflexis/omitempty"
	"github.com/omniboost/go-dyflexis/utils"
)

func (c *Client) NewCustomReportRequest() CustomReportRequest {
	r := CustomReportRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomReportRequest struct {
	client      *Client
	queryParams *CustomReportRequestQueryParams
	pathParams  *CustomReportRequestPathParams
	method      string
	headers     http.Header
	requestBody CustomReportRequestBody
}

func (r CustomReportRequest) NewQueryParams() *CustomReportRequestQueryParams {
	return &CustomReportRequestQueryParams{}
}

type CustomReportRequestQueryParams struct {
	ID           string           `schema:"id"`
	StartDate    Date             `schema:"startDate,omitempty"`
	EndDate      Date             `schema:"endDate,omitempty"`
	DateFormat   string           `schema:"dateFormat,omitempty"`
	ReportFormat ReportFormatType `schema:"reportFormat,omitempty"`
}

func (p CustomReportRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomReportRequest) QueryParams() *CustomReportRequestQueryParams {
	return r.queryParams
}

func (r CustomReportRequest) NewPathParams() *CustomReportRequestPathParams {
	return &CustomReportRequestPathParams{}
}

type CustomReportRequestPathParams struct {
}

func (p *CustomReportRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomReportRequest) PathParams() *CustomReportRequestPathParams {
	return r.pathParams
}

func (r *CustomReportRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomReportRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomReportRequest) Method() string {
	return r.method
}

func (r CustomReportRequest) NewRequestBody() CustomReportRequestBody {
	return CustomReportRequestBody{}
}

type CustomReportRequestBody struct{}

func (r CustomReportRequestBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

func (r CustomReportRequestBody) IsEmpty() bool {
	return zero.IsZero(r)
}

func (r *CustomReportRequest) RequestBody() *CustomReportRequestBody {
	return nil
}

func (r *CustomReportRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *CustomReportRequest) SetRequestBody(body CustomReportRequestBody) {
	r.requestBody = body
}

func (r *CustomReportRequest) NewResponseBody() *CustomReportResponseBody {
	return &CustomReportResponseBody{}
}

type CustomReportResponseBody []map[string]interface{}

func (r *CustomReportRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/reports/customreport", r.PathParams())
	return &u, err
}

func (r *CustomReportRequest) Do() (CustomReportResponseBody, error) {
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
