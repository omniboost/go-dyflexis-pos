package timesheet

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-timesheetportal/utils"
)

func (c *Client) NewProjectsRequest() ProjectsRequest {
	r := ProjectsRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ProjectsRequest struct {
	client      *Client
	queryParams *ProjectsRequestQueryParams
	pathParams  *ProjectsRequestPathParams
	method      string
	headers     http.Header
	requestBody ProjectsRequestBody
}

func (r ProjectsRequest) NewQueryParams() *ProjectsRequestQueryParams {
	return &ProjectsRequestQueryParams{}
}

type ProjectsRequestQueryParams struct {
	Account string `schema:"Account"`
	// PageSize       int    `schema:"PageSize"`
}

func (p ProjectsRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ProjectsRequest) QueryParams() *ProjectsRequestQueryParams {
	return r.queryParams
}

func (r ProjectsRequest) NewPathParams() *ProjectsRequestPathParams {
	return &ProjectsRequestPathParams{}
}

type ProjectsRequestPathParams struct {
}

func (p *ProjectsRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ProjectsRequest) PathParams() *ProjectsRequestPathParams {
	return r.pathParams
}

func (r *ProjectsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ProjectsRequest) SetMethod(method string) {
	r.method = method
}

func (r *ProjectsRequest) Method() string {
	return r.method
}

func (r ProjectsRequest) NewRequestBody() ProjectsRequestBody {
	return ProjectsRequestBody{}
}

type ProjectsRequestBody struct {
}

func (r *ProjectsRequest) RequestBody() *ProjectsRequestBody {
	return nil
}

func (r *ProjectsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *ProjectsRequest) SetRequestBody(body ProjectsRequestBody) {
	r.requestBody = body
}

func (r *ProjectsRequest) NewResponseBody() *ProjectsResponseBody {
	return &ProjectsResponseBody{}
}

type ProjectsResponseBody Projects

func (r *ProjectsRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/projects", r.PathParams())
	return &u, err
}

func (r *ProjectsRequest) Do() (ProjectsResponseBody, error) {
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
