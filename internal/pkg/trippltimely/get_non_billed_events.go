package trippltimely

import (
	"net/http"

	"github.com/cobraz/trippl-timely/internal/pkg/timely"
)

func getNonBilledEvents(client *http.Client) ([]*timely.Event, error) {
	allEvents, err := timely.GetEvents(client, "")
	if err != nil {
		return nil, err
	}

	var nonBilledEvents []*timely.Event

	for _, event := range allEvents {
		if !event.Billed {
			nonBilledEvents = append(nonBilledEvents, event)
		}
	}

	return nonBilledEvents, nil
}
