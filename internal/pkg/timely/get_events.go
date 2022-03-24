package timely

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetEvents returns all Timely events
func GetEvents(client *http.Client, arguments string) ([]*Event, error) {
	url := "https://api.timelyapp.com/1.1/930822/users/1960450/events"
	if arguments != "" {
		url = fmt.Sprintf("%s%s", url, arguments)
	}

	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	var events []*Event

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, err
	}

	return events, nil
}
