package main

import (
	"log"
	"os"

	"github.com/cobraz/trippl-timely/internal/cmd/config"
	"github.com/cobraz/trippl-timely/internal/cmd/timely"
	"github.com/cobraz/trippl-timely/internal/cmd/trippltimely"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "Trippl Timely",
		Description: "Send Timely events to Tripletex",
		Commands: []*cli.Command{
			&cli.Command{
				Name:   "get-config",
				Action: config.GetConfig,
			},
			&cli.Command{
				Name:   "set-config",
				Action: config.SetConfig,
			},
			&cli.Command{
				Name:   "get-timely-token",
				Action: timely.GetToken,
			},
			&cli.Command{
				Name: "add-timesheet",
				Flags: []cli.Flag{
					&cli.TimestampFlag{
						Name:   "date",
						Layout: "2006-01-02",
					},
				},
				Description: "Lists all events in Timely",
				Action:      trippltimely.AddTimesheet,
			},
			// &cli.Command{
			// 	Name: "list:events",
			// 	Flags: []cli.Flag{
			// 		&cli.StringFlag{
			// 			Name:     "date",
			// 			Required: true,
			// 			Usage:    "Date (format YYYY-MM-DD)",
			// 		},
			// 	},
			// 	Description: "Lists all events in Timely",
			// 	Action:
			// },
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
