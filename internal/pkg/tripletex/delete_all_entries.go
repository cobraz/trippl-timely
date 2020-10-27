package tripletex

import (
	"time"

	"github.com/bjerkio/tripletex-go/client/entry"
)

func (c *TripletexClient) deleteEntries(id int32) error {
	p := entry.NewTimesheetEntryDeleteParams()
	p.ID = id

	return c.client.Entry.TimesheetEntryDelete(p, c.authInfo)
}

// DeleteAllEntries deletes all activities
func (c *TripletexClient) DeleteAllEntries(d time.Time) error {
	res, err := c.GetAllEntries(d)
	if err != nil {
		return err
	}

	for _, d := range res {
		err := c.deleteEntries(d.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
