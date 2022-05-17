package dyflexis_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestUsersGet(t *testing.T) {
	req := client.NewUsersGetRequest()
	req.QueryParams().ID = 448450
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
