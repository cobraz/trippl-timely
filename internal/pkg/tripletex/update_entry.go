package tripletex

import (
	"github.com/bjerkio/tripletex-go/client/entry"
	"github.com/bjerkio/tripletex-go/models"
)

// GetAllEntries returns all timesheet entries for a given date
func (c TripletexClient) UpdateEntry(ID int32, e *models.TimesheetEntry) error {
	p := entry.NewTimesheetEntryPutParams()
	p.Body = e
	p.ID = ID

	_, err := c.client.Entry.TimesheetEntryPut(p, c.authInfo)
	if err != nil {
		return err
	}

	return nil
}
