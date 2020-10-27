package trippltimely

import (
	"fmt"
	"strconv"
	"time"

	cnf "github.com/cobraz/trippl-timely/internal/pkg/config"
	"github.com/cobraz/trippl-timely/internal/pkg/timely"
	"github.com/cobraz/trippl-timely/internal/pkg/tripletex"
	ttly "github.com/cobraz/trippl-timely/internal/pkg/trippltimely"
	"github.com/urfave/cli/v2"
)

func AddTimesheet(c *cli.Context) error {
	config, err := cnf.GetConfig()
	if err != nil {
		return err
	}

	client, err := timely.Client(config)
	if err != nil {
		return err
	}

	tx, err := tripletex.New(config)
	if err != nil {
		return err
	}

	d := c.Timestamp("date")

	// List events from Timely with projects
	// events, err := ttly.GetEventsWithProjectExternalID(client, *d)
	events, err := timely.GetEventsByDate(client, *d)
	if err != nil {
		return err
	}

	if len(events) < 1 {
		fmt.Println("There are no events for this date.")
		return nil
	}

	// TODO: Do we really need to delete stuff?
	err = tx.DeleteAllActivities(*d)
	if err != nil {
		return err
	}

	var timeentries []ttly.TimesheetEntry

	for _, e := range events {
		date, err := time.Parse("2006-01-02", e.Day)
		if err != nil {
			return err
		}

		if e.Project != nil {
			if e.Project.ExternalID != "" {
				extID, err := strconv.ParseInt(e.Project.ExternalID, 10, 32)
				if err != nil {
					return err
				}
				pID := int32(extID)
				timeentries = append(timeentries, ttly.TimesheetEntry{
					TotalHours: e.Duration.TotalHours,
					Note:       e.Note,
					ProjectID:  &pID,
					Date:       date,
				})
			} else {
				// TODO: Move this to config? Mapping?
				// if e.Project.ID == 3344426 {
				// 	// var aID int32 = 728489
				// 	// timeentries = append(timeentries, ttly.TimesheetEntry{
				// 	// 	TotalHours: e.Duration.TotalHours,
				// 	// 	Note:       e.Note,
				// 	// 	Date:       date,
				// 	// 	ActivityID: &aID,
				// 	// })
				// } else if e.Project.ID == 3344427 {
				// 	// var aID int32 = 728489
				// 	// timeentries = append(timeentries, ttly.TimesheetEntry{
				// 	// 	TotalHours: e.Duration.TotalHours,
				// 	// 	Note:       e.Note,
				// 	// 	Date:       date,
				// 	// 	ActivityID: &aID,
				// 	// })
				// } else {
				fmt.Printf("%s is missing external ID!\n", e.Project.Name)
				// }
			}

		}
	}

	// Step three: Set them in Tripletex
	err = tx.UpdateTimesheet(timeentries)
	if err != nil {
		return err
	}
	// Step four: Update billing status

	return nil
}
