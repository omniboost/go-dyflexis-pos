package timesheet_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestReportFields(t *testing.T) {
	req := client.NewReportFieldsRequest()
	// req.QueryParams().Account = "FLD"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
