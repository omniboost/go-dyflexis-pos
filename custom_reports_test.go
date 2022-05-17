package dyflexis_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCustomReportsGet(t *testing.T) {
	req := client.NewCustomReportsGetRequest()
	// req.QueryParams().Account = "FLD"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
