package dyflexis_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestBuiltinReport(t *testing.T) {
	req := client.NewBuiltinReportRequest()
	// req.RequestBody().ReportTimeGrouping = "Daily"
	req.RequestBody().BuiltInReportType = "TimesheetStatus"
	req.RequestBody().ReportFields = []string{
		"EmployeeName",
		"InternalTimesheetId",
		"TotalTimesheetCount",
		"TimesheetStatusId",
		"TimesheetLog",
		"TimesheetId",
		"EmployeeId",
		"DateOfBirth",
		"EmployeeJoinDate",
		"TimesheetStatusId",
		"EmployeeEmailAddress",
	}
	req.RequestBody().ReportFormat = "KeyValuePairsArray"
	req.RequestBody().StartDate.Time = time.Date(2022, 3, 14, 0, 0, 0, 0, time.UTC)
	req.RequestBody().EndDate.Time = time.Date(2022, 3, 20, 0, 0, 0, 0, time.UTC)
	req.RequestBody().TimesheetStatusFilters = []string{
		"Approved",
	}
	req.RequestBody().PageIndex = 0
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
