package timely

import (
	"context"
	"fmt"

	cnf "github.com/cobraz/trippl-timely/internal/pkg/config"
	"github.com/urfave/cli/v2"
	"golang.org/x/oauth2"
)

func GetToken(c *cli.Context) error {
	config, err := cnf.GetConfig()
	if err != nil {
		return err
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

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return err
	}
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		return err
	}

	config.Timely.Token = tok
	err = cnf.SetConfig(config)
	if err != nil {
		return err
	}

	cnf.Print()

	return nil
}
