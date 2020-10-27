package tripletex

import (
	"time"

	"github.com/bjerkio/tripletex-go/client/activity"
	"github.com/bjerkio/tripletex-go/models"
)

// GetActivities returns a list of Tripletex activities
func (c *TripletexClient) GetActivities() ([]*models.Activity, error) {

	yes := true

	req := activity.ActivitySearchParams{
		IsProjectActivity: &yes,
		IsGeneral:         &yes,
	}

	res, err := c.client.Activity.ActivitySearch(req.WithTimeout(10*time.Second), c.authInfo)
	if err != nil {
		return nil, err
	}

	return res.Payload.Values, nil
}
