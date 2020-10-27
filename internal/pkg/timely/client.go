package timely

import (
	"context"
	"errors"
	"net/http"

	"github.com/cobraz/trippl-timely/internal/pkg/config"
	"golang.org/x/oauth2"
)

func Client(config config.Config) (*http.Client, error) {

	if config.Timely.Token == nil {
		return nil, errors.New("No token found. Please run trippl-timley get-timley-token")
	}

	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     config.Timely.ClientID,
		ClientSecret: config.Timely.ClientSecret,
		RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.timelyapp.com/1.1/oauth/authorize",
			TokenURL: "https://api.timelyapp.com/1.1/oauth/token",
		},
	}

	return conf.Client(ctx, config.Timely.Token), nil
}
