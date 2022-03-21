package timesheet

import (
	"net/http"
	"net/url"

	"github.com/cydev/zero"
	"github.com/omniboost/go-timesheetportal/omitempty"
	"github.com/omniboost/go-timesheetportal/utils"
)

func (c *Client) NewBuiltinReportRequest() BuiltinReportRequest {
	r := BuiltinReportRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type BuiltinReportRequest struct {
	client      *Client
	queryParams *BuiltinReportRequestQueryParams
	pathParams  *BuiltinReportRequestPathParams
	method      string
	headers     http.Header
	requestBody BuiltinReportRequestBody
}

func (r BuiltinReportRequest) NewQueryParams() *BuiltinReportRequestQueryParams {
	return &BuiltinReportRequestQueryParams{}
}

type BuiltinReportRequestQueryParams struct {
	Account string `schema:"Account"`
	// PageSize       int    `schema:"PageSize"`
}

func (p BuiltinReportRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *BuiltinReportRequest) QueryParams() *BuiltinReportRequestQueryParams {
	return r.queryParams
}

func (r BuiltinReportRequest) NewPathParams() *BuiltinReportRequestPathParams {
	return &BuiltinReportRequestPathParams{}
}

type BuiltinReportRequestPathParams struct {
}

func (p *BuiltinReportRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *BuiltinReportRequest) PathParams() *BuiltinReportRequestPathParams {
	return r.pathParams
}

func (r *BuiltinReportRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *BuiltinReportRequest) SetMethod(method string) {
	r.method = method
}

func (r *BuiltinReportRequest) Method() string {
	return r.method
}

func (r BuiltinReportRequest) NewRequestBody() BuiltinReportRequestBody {
	return BuiltinReportRequestBody{
		ReportFields: []string{},
	}
}

type BuiltinReportRequestBody struct {
	UseIsoDateFormat                     bool     `json:"UseIsoDateFormat,omitempty"`
	StartDate                            Time     `json:"StartDate,omitempty"`
	EndDate                              Time     `json:"EndDate,omitempty"`
	ReportTimeGrouping                   string   `json:"ReportTimeGrouping,omitempty"`
	ReportFields                         []string `json:"ReportFields,omitempty"`
	ClientFilter                         string   `json:"ClientFilter,omitempty"`
	CostCentreFilter                     string   `json:"CostCentreFilter,omitempty"`
	EmployeeGroupFilter                  string   `json:"EmployeeGroupFilter,omitempty"`
	ChargeCodeGroupFilter                string   `json:"ChargeCodeGroupFilter,omitempty"`
	ChargeCodeGroupFilterMultiple        []string `json:"ChargeCodeGroupFilterMultiple,omitempty"`
	ChargeCodeGroupCategoryFilter        string   `json:"ChargeCodeGroupCategoryFilter,omitempty"`
	TaskCodeFilter                       string   `json:"TaskCodeFilter,omitempty"`
	ModifiedAfterDate                    Time     `json:"ModifiedAfterDate,omitempty"`
	ModifiedAfterTimesheetSequenceNumber int      `json:"ModifiedAfterTimesheetSequenceNumber,omitempty"`
	ModifiedAfterExpenseSequenceNumber   int      `json:"ModifiedAfterExpenseSequenceNumber,omitempty"`
	TimesheetStatusFilters               []string `json:"TimesheetStatusFilters,omitempty"`
	ExpenseStatusFilters                 []string `json:"ExpenseStatusFilters,omitempty"`
	ReportFieldGroupings                 []string `json:"ReportFieldGroupings,omitempty"`
	ApprovedDateFilterStart              Time     `json:"ApprovedDateFilterStart,omitempty"`
	ApprovedDateFilterEnd                Time     `json:"ApprovedDateFilterEnd,omitempty"`
	InvoicedStatusFilter                 string   `json:"InvoicedStatusFilter,omitempty"`
	BillableStatusFilter                 string   `json:"BillableStatusFilter,omitempty"`
	ProductiveStatusFilter               string   `json:"ProductiveStatusFilter,omitempty"`
	BuiltInReportType                    string   `json:"BuiltInReportType,omitempty"`
	TimesheetStatusFrequency             string   `json:"TimesheetStatusFrequency,omitempty"`
	ExportedDateFilterStart              Time     `json:"ExportedDateFilterStart,omitempty"`
	ExportedDateFilterEnd                Time     `json:"ExportedDateFilterEnd,omitempty"`
	PageIndex                            int      `json:"PageIndex,omitempty"`
	ReportFormat                         string   `json:"ReportFormat,omitempty"`
}

func (r BuiltinReportRequestBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

func (r BuiltinReportRequestBody) IsEmpty() bool {
	return zero.IsZero(r)
}

func (r *BuiltinReportRequest) RequestBody() *BuiltinReportRequestBody {
	return &r.requestBody
}

func (r *BuiltinReportRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *BuiltinReportRequest) SetRequestBody(body BuiltinReportRequestBody) {
	r.requestBody = body
}

func (r *BuiltinReportRequest) NewResponseBody() *BuiltinReportResponseBody {
	return &BuiltinReportResponseBody{}
}

type BuiltinReportResponseBody []map[string]interface{}

func (r *BuiltinReportRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/reports/builtinreport", r.PathParams())
	return &u, err
}

func (r *BuiltinReportRequest) Do() (BuiltinReportResponseBody, error) {
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
