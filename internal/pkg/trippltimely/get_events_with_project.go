package trippltimely

import (
	"net/http"
	"time"

	"github.com/cobraz/trippl-timely/internal/pkg/timely"
)

func GetEventsWithProjectExternalID(client *http.Client, date time.Time) ([]*timely.Event, error) {
	allEvents, err := timely.GetEventsByDate(client, date)
	if err != nil {
		return nil, err
	}

	var hasProjectEvents []*timely.Event

	for _, event := range allEvents {
		if event.Project != nil {
			if event.Project.ExternalID != "" {
				hasProjectEvents = append(hasProjectEvents, event)
			}
		}
	}

	return hasProjectEvents, nil
}
