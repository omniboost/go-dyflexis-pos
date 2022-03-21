package timesheet_test

import (
	"context"
	"os"
	"testing"

	netsuite "github.com/omniboost/go-timesheetportal"
	"golang.org/x/oauth2"
)

var (
	client *netsuite.Client
)

func TestMain(m *testing.M) {
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	refreshToken := os.Getenv("REFRESH_TOKEN")
	debug := os.Getenv("DEBUG")

	oauthConfig := netsuite.NewOauth2Config()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	token := &oauth2.Token{
		RefreshToken: refreshToken,
	}

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(context.Background(), token)

	client = netsuite.NewClient(httpClient)
	if debug != "" {
		client.SetDebug(true)
	}

	client.SetDisallowUnknownFields(true)
	m.Run()
}
