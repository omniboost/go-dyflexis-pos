package dyflexis_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestCustomReport(t *testing.T) {
	req := client.NewCustomReportRequest()
	req.QueryParams().ID = "5b54cb46-ffcd-435d-8ae4-a1678cf9cbd6"
	req.QueryParams().StartDate.Time = time.Date(2022, 3, 1, 0, 0, 0, 0, time.UTC)
	req.QueryParams().EndDate.Time = time.Date(2022, 3, 2, 0, 0, 0, 0, time.UTC)
	// req.QueryParams().DateFormat = "%Y%m%d"
	req.QueryParams().ReportFormat = "KeyValuePairsArray"
	resp, err := req.Do()

	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
