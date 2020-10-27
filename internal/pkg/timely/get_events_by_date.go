package timely

import (
	"fmt"
	"net/http"
	"time"
)

func GetEventsByDate(client *http.Client, date time.Time) ([]*Event, error) {
	d := date.Format("2006-01-02")
	filter := fmt.Sprintf("?day=%s", d)

	e, err := GetEvents(client, filter)
	if err != nil {
		return nil, err
	}

	return e, nil
}
