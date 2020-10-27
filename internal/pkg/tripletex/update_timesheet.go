package tripletex

import (
	"fmt"
	"math"
	"time"

	"github.com/bjerkio/tripletex-go/client/entry"
	"github.com/bjerkio/tripletex-go/models"
	"github.com/cobraz/trippl-timely/internal/pkg/trippltimely"
)

func (c *TripletexClient) UpdateTimesheet(d time.Time, entries []trippltimely.TimesheetEntry) error {

	var tEntries []*models.TimesheetEntry
	// yes := true
	emty := " "

	oldEntries, err := c.GetAllEntries(d)

	for _, e := range entries {
		d := e.Date.Format("2006-01-02")
		h := math.Ceil(e.TotalHours*100) / 100
		var p *models.Project

		if e.ProjectID != nil {
			p = &models.Project{
				ID: *e.ProjectID,
			}
		}

		var a models.Activity

		if e.ActivityID != nil {
			a = models.Activity{
				ID: *e.ActivityID,
			}
		} else {
			a = models.Activity{
				ID: c.config.Tripletex.ActivityCode,
			}
		}

		cte := &models.TimesheetEntry{
			Activity: &a,
			Comment:  e.Note,
			Date:     &d,
			Hours:    &h,
			Project:  p,
			Employee: &models.Employee{
				ID:        1772944,
				FirstName: &emty,
				LastName:  &emty,
			},
		}

		var exists bool = false

		for _, te := range oldEntries {
			if te.Project.ID == p.ID && te.Activity.ID == a.ID {
				if *te.Locked == false {
					err := c.UpdateEntry(te.ID, cte)
					if err != nil {
						return err
					}
				}

				exists = true
			}
		}

		if exists == false {
			tEntries = append(tEntries, cte)
		}
	}

	if len(tEntries) > 0 {
		p := entry.NewTimesheetEntryListPostListParams()
		p.Body = tEntries

		_, err = c.client.Entry.TimesheetEntryListPostList(p, c.authInfo)
		if err != nil {
			fmt.Printf("Was not able to insert. Make sure you have a clean slate at Tripletex")
			return nil
		}
	}

	return nil
	// return err
}
