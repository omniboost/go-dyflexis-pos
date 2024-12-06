package dyflexis_test

import (
	"os"
	"testing"

	dyflexis "github.com/omniboost/go-dyflexis-pos"
)

var (
	client *dyflexis.Client
)

func TestMain(m *testing.M) {
	debug := os.Getenv("DEBUG")
	posClientID := os.Getenv("POS_CLIENT_ID")
	posToken := os.Getenv("POS_TOKEN")

	client = dyflexis.NewClient(nil)
	if debug != "" {
		client.SetDebug(true)
	}

	client.SetPOSClientID(posClientID)
	client.SetPOSToken(posToken)

	client.SetDisallowUnknownFields(true)
	m.Run()
}
