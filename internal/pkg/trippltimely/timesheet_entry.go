package trippltimely

import "time"

type TimesheetEntry struct {
	TotalHours float64
	Note       string
	ProjectID  *int32
	Date       time.Time
	ActivityID *int32
}
