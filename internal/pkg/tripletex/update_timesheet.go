package tripletex

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/bjerkio/tripletex-go/client/entry"
	"github.com/bjerkio/tripletex-go/models"
	"github.com/cobraz/trippl-timely/internal/pkg/trippltimely"
)

func (c *TripletexClient) UpdateTimesheet(entries []trippltimely.TimesheetEntry) error {

	var tEntries []*models.TimesheetEntry
	// yes := true
	emty := " "

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

		tEntries = append(tEntries, &models.TimesheetEntry{
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
		})
	}

	body, err := json.Marshal(tEntries)
	fmt.Println(string(body))

	p := entry.NewTimesheetEntryListPostListParams()
	p.Body = tEntries

	_, err = c.client.Entry.TimesheetEntryListPostList(p, c.authInfo)
	if err != nil {
		// TODO: Add a way where we can update 👇
		// p := entry.NewTimesheetEntryListPutListParams()
		// p.Body = tEntries
		// _, err = c.client.Entry.TimesheetEntryListPutList(p, c.authInfo)
		// if err != nil {
		// 	return err
		// }
		fmt.Printf("Was not able to insert. Make sure you have a clean slate at Tripletex")
		return nil
	}

	return nil
	// return err
}
