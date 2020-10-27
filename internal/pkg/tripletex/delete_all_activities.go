package tripletex

import (
	"time"

	"github.com/bjerkio/tripletex-go/client/entry"
)

func (c *TripletexClient) deleteActivitiy(id int32) error {
	p := entry.NewTimesheetEntryDeleteParams()
	p.ID = id

	return c.client.Entry.TimesheetEntryDelete(p, c.authInfo)
}

// DeleteAllActivities deletes all activities
func (c *TripletexClient) DeleteAllActivities(d time.Time) error {
	p := entry.NewTimesheetEntrySearchParams()
	p.DateFrom = d.Format("2006-01-02")
	p.DateTo = d.Add(time.Hour * 24).Format("2006-01-02")

	res, err := c.client.Entry.TimesheetEntrySearch(p, c.authInfo)
	if err != nil {
		return err
	}

	for _, d := range res.Payload.Values {
		err := c.deleteActivitiy(d.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
