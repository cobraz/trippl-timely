package tripletex

import (
	"time"

	"github.com/bjerkio/tripletex-go/client/entry"
	"github.com/bjerkio/tripletex-go/models"
)

// GetAllEntries returns all timesheet entries for a given date
func (c TripletexClient) GetAllEntries(employeeId string, d time.Time) ([]*models.TimesheetEntry, error) {
	p := entry.NewTimesheetEntrySearchParams()
	p.DateFrom = d.Format("2006-01-02")
	p.DateTo = d.Add(time.Hour * 24).Format("2006-01-02")
	p.EmployeeID = &employeeId

	res, err := c.client.Entry.TimesheetEntrySearch(p, c.authInfo)
	if err != nil {
		return nil, err
	}

	return res.Payload.Values, nil
}
