package dyflexis

import "time"

type ReportFormatType string

// type CustomReports

type Timesheet struct {
	StartDate           string      `json:"startDate"`
	CreateDate          string      `json:"createDate"`
	SubmitDate          time.Time   `json:"submitDate"`
	ApprovedDate        time.Time   `json:"approvedDate"`
	RejectedDate        interface{} `json:"rejectedDate"`
	TimesheetGUID       string      `json:"timesheetGuid"`
	InternalTimesheetID int         `json:"internalTimesheetId"`
	SubmitterCode       interface{} `json:"submitterCode"`
	ApproverCode        interface{} `json:"approverCode"`
	APITimesheetCode    string      `json:"apiTimesheetCode"`
	EmployeeID          int         `json:"employeeId"`
	EmployeeName        string      `json:"employeeName"`
	AllowUpdates        bool        `json:"allowUpdates"`
	Status              int         `json:"status"`
	TimesheetFrequency  int         `json:"timesheetFrequency"`
	EntryGroups         []struct {
		ChargeCodeReference string `json:"chargeCodeReference"`
		RateCode            string `json:"rateCode"`
		Entries             []struct {
			EntryQuantity    float64       `json:"entryQuantity"`
			BreakTime        float64       `json:"breakTime"`
			BreakStartTime   string        `json:"breakStartTime"`
			BreakEndTime     string        `json:"breakEndTime"`
			Date             string        `json:"date"`
			StartTime        string        `json:"startTime"`
			EndTime          string        `json:"endTime"`
			Valid            bool          `json:"valid"`
			ValidationErrors []interface{} `json:"validationErrors"`
		} `json:"entries"`
		Valid            bool          `json:"valid"`
		ValidationErrors []interface{} `json:"validationErrors"`
	} `json:"entryGroups"`
	EndDate          string        `json:"endDate"`
	Valid            bool          `json:"valid"`
	ValidationErrors []interface{} `json:"validationErrors"`
}

type Projects []Project

type Project struct {
	ChargeCodeGroupCode            interface{} `json:"chargeCodeGroupCode"`
	ChargeCodeGroupDescription     interface{} `json:"chargeCodeGroupDescription"`
	ProjectName                    string      `json:"projectName"`
	ProjectCode                    string      `json:"projectCode"`
	ClientCode                     string      `json:"clientCode"`
	AccountingCode                 string      `json:"accountingCode"`
	CategoryCode                   interface{} `json:"categoryCode"`
	StartDate                      time.Time   `json:"startDate"`
	StartDateNull                  bool        `json:"startDateNull"`
	EndDate                        time.Time   `json:"endDate"`
	EndDateNull                    bool        `json:"endDateNull"`
	CurrencyCode                   string      `json:"currencyCode"`
	PayBudget                      float64     `json:"payBudget"`
	ChargeBudget                   float64     `json:"chargeBudget"`
	ChargeBudgetCurrencyCode       string      `json:"chargeBudgetCurrencyCode"`
	PayBudgetcurrencyCode          string      `json:"payBudgetcurrencyCode"`
	ProjectManagerCode             string      `json:"projectManagerCode"`
	ProjectManagerName             string      `json:"projectManagerName"`
	AssignedUsers                  interface{} `json:"assignedUsers"`
	AutoAssignToEveryone           bool        `json:"autoAssignToEveryone"`
	CostCentreCode                 string      `json:"costCentreCode"`
	CostCentreDescription          string      `json:"costCentreDescription"`
	TaxCode                        interface{} `json:"taxCode"`
	PurchaseOrder                  interface{} `json:"purchaseOrder"`
	CalculatePayBudgetFromTasks    bool        `json:"calculatePayBudgetFromTasks"`
	CalculateChargeBudgetFromTasks bool        `json:"calculateChargeBudgetFromTasks"`
	CalculateTimeBudgetFromTasks   bool        `json:"calculateTimeBudgetFromTasks"`
	CustomFields                   struct {
		CustomFieldString0  string      `json:"customFieldString0"`
		CustomFieldString1  interface{} `json:"customFieldString1"`
		CustomFieldString2  interface{} `json:"customFieldString2"`
		CustomFieldString3  interface{} `json:"customFieldString3"`
		CustomFieldString4  interface{} `json:"customFieldString4"`
		CustomFieldString5  interface{} `json:"customFieldString5"`
		CustomFieldString6  interface{} `json:"customFieldString6"`
		CustomFieldString7  interface{} `json:"customFieldString7"`
		CustomFieldString8  interface{} `json:"customFieldString8"`
		CustomFieldString9  interface{} `json:"customFieldString9"`
		CustomFieldList0    interface{} `json:"customFieldList0"`
		CustomFieldList1    interface{} `json:"customFieldList1"`
		CustomFieldList2    interface{} `json:"customFieldList2"`
		CustomFieldList3    interface{} `json:"customFieldList3"`
		CustomFieldList4    interface{} `json:"customFieldList4"`
		CustomFieldList5    interface{} `json:"customFieldList5"`
		CustomFieldList6    interface{} `json:"customFieldList6"`
		CustomFieldList7    interface{} `json:"customFieldList7"`
		CustomFieldList8    interface{} `json:"customFieldList8"`
		CustomFieldList9    interface{} `json:"customFieldList9"`
		CustomFieldInt0     interface{} `json:"customFieldInt0"`
		CustomFieldInt1     interface{} `json:"customFieldInt1"`
		CustomFieldInt2     interface{} `json:"customFieldInt2"`
		CustomFieldInt3     interface{} `json:"customFieldInt3"`
		CustomFieldInt4     interface{} `json:"customFieldInt4"`
		CustomFieldInt5     interface{} `json:"customFieldInt5"`
		CustomFieldInt6     interface{} `json:"customFieldInt6"`
		CustomFieldInt7     interface{} `json:"customFieldInt7"`
		CustomFieldInt8     interface{} `json:"customFieldInt8"`
		CustomFieldInt9     interface{} `json:"customFieldInt9"`
		CustomFieldDecimal0 interface{} `json:"customFieldDecimal0"`
		CustomFieldDecimal1 interface{} `json:"customFieldDecimal1"`
		CustomFieldDecimal2 interface{} `json:"customFieldDecimal2"`
		CustomFieldDecimal3 interface{} `json:"customFieldDecimal3"`
		CustomFieldDecimal4 interface{} `json:"customFieldDecimal4"`
		CustomFieldDecimal5 interface{} `json:"customFieldDecimal5"`
		CustomFieldDecimal6 interface{} `json:"customFieldDecimal6"`
		CustomFieldDecimal7 interface{} `json:"customFieldDecimal7"`
		CustomFieldDecimal8 interface{} `json:"customFieldDecimal8"`
		CustomFieldDecimal9 interface{} `json:"customFieldDecimal9"`
		CustomFieldBool0    interface{} `json:"customFieldBool0"`
		CustomFieldBool1    interface{} `json:"customFieldBool1"`
		CustomFieldBool2    interface{} `json:"customFieldBool2"`
		CustomFieldBool3    interface{} `json:"customFieldBool3"`
		CustomFieldBool4    interface{} `json:"customFieldBool4"`
		CustomFieldBool5    interface{} `json:"customFieldBool5"`
		CustomFieldBool6    interface{} `json:"customFieldBool6"`
		CustomFieldBool7    interface{} `json:"customFieldBool7"`
		CustomFieldBool8    interface{} `json:"customFieldBool8"`
		CustomFieldBool9    interface{} `json:"customFieldBool9"`
	} `json:"customFields"`
	CustomFieldValues []struct {
		FieldName   string `json:"fieldName"`
		FieldTarget int    `json:"fieldTarget"`
		FieldIndex  int    `json:"fieldIndex"`
		DataType    int    `json:"dataType"`
		Value       string `json:"value"`
	} `json:"customFieldValues"`
	BillingContactUserCode             interface{}   `json:"billingContactUserCode"`
	BillingContactEmail                interface{}   `json:"billingContactEmail"`
	BillingContactName                 interface{}   `json:"billingContactName"`
	LastModified                       time.Time     `json:"lastModified"`
	Address                            interface{}   `json:"address"`
	InvoiceDescription                 interface{}   `json:"invoiceDescription"`
	AssignedApproverCodes              []interface{} `json:"assignedApproverCodes"`
	ShouldSetAssignedApproversOnUpdate bool          `json:"shouldSetAssignedApproversOnUpdate"`
	Valid                              bool          `json:"valid"`
	ValidationErrors                   []interface{} `json:"validationErrors"`
}
