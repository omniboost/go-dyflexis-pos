package timesheet_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestTimesheetsGet(t *testing.T) {
	req := client.NewTimesheetsGetRequest()
	req.QueryParams().TimesheetID = 20785
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
