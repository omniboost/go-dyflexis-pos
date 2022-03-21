package timesheet_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

func TestCustomReport(t *testing.T) {
	req := client.NewCustomReportRequest()
	req.QueryParams().ID = "b6736ab3-1290-4111-806e-407535f5466d"
	req.QueryParams().StartDate.Time = time.Date(2022, 3, 1, 0, 0, 0, 0, time.UTC)
	req.QueryParams().EndDate.Time = time.Date(2022, 3, 2, 0, 0, 0, 0, time.UTC)
	// req.QueryParams().DateFormat = "%Y%m%d"
	// req.QueryParams().ReportFormat = "json"
	resp, err := req.Do()

	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
