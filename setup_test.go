package dyflexis_test

import (
	"context"
	"os"
	"testing"

	timesheetportal "github.com/omniboost/go-dyflexis"
	"golang.org/x/oauth2"
)

var (
	client *timesheetportal.Client
)

func TestMain(m *testing.M) {
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	refreshToken := os.Getenv("REFRESH_TOKEN")
	debug := os.Getenv("DEBUG")

	oauthConfig := timesheetportal.NewOauth2Config()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	token := &oauth2.Token{
		RefreshToken: refreshToken,
	}

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(context.Background(), token)

	client = timesheetportal.NewClient(httpClient)
	if debug != "" {
		client.SetDebug(true)
	}

	client.SetDisallowUnknownFields(true)
	m.Run()
}
