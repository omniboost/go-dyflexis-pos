package dyflexis

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-dyflexis/utils"
)

func (c *Client) NewUsersGetRequest() UsersGetRequest {
	r := UsersGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type UsersGetRequest struct {
	client      *Client
	queryParams *UsersGetRequestQueryParams
	pathParams  *UsersGetRequestPathParams
	method      string
	headers     http.Header
	requestBody UsersGetRequestBody
}

func (r UsersGetRequest) NewQueryParams() *UsersGetRequestQueryParams {
	return &UsersGetRequestQueryParams{}
}

type UsersGetRequestQueryParams struct {
	ID int `schema:"id"`
}

func (p UsersGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *UsersGetRequest) QueryParams() *UsersGetRequestQueryParams {
	return r.queryParams
}

func (r UsersGetRequest) NewPathParams() *UsersGetRequestPathParams {
	return &UsersGetRequestPathParams{}
}

type UsersGetRequestPathParams struct {
}

func (p *UsersGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *UsersGetRequest) PathParams() *UsersGetRequestPathParams {
	return r.pathParams
}

func (r *UsersGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *UsersGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *UsersGetRequest) Method() string {
	return r.method
}

func (r UsersGetRequest) NewRequestBody() UsersGetRequestBody {
	return UsersGetRequestBody{}
}

type UsersGetRequestBody struct {
}

func (r *UsersGetRequest) RequestBody() *UsersGetRequestBody {
	return nil
}

func (r *UsersGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *UsersGetRequest) SetRequestBody(body UsersGetRequestBody) {
	r.requestBody = body
}

func (r *UsersGetRequest) NewResponseBody() *UsersGetResponseBody {
	return &UsersGetResponseBody{}
}

type UsersGetResponseBody Timesheet

func (r *UsersGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/users", r.PathParams())
	return &u, err
}

func (r *UsersGetRequest) Do() (UsersGetResponseBody, error) {
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
