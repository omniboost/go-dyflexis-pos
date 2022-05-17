package dyflexis_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestProjects(t *testing.T) {
	req := client.NewProjectsRequest()
	// req.QueryParams().Account = "FLD"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
