package timely

import "time"

// Event is used when calling /events in Timely
type Event struct {
	ID   int    `json:"id"`
	UID  string `json:"uid"`
	User struct {
		ID     int    `json:"id"`
		Email  string `json:"email"`
		Name   string `json:"name"`
		Avatar struct {
			LargeRetina  string `json:"large_retina"`
			Large        string `json:"large"`
			MediumRetina string `json:"medium_retina"`
			Medium       string `json:"medium"`
			Timeline     string `json:"timeline"`
		} `json:"avatar"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"user"`
	Project *struct {
		ID          int         `json:"id"`
		Active      bool        `json:"active"`
		AccountID   int         `json:"account_id"`
		Name        string      `json:"name"`
		Color       string      `json:"color"`
		RateType    string      `json:"rate_type"`
		Billable    bool        `json:"billable"`
		UpdatedAt   int         `json:"updated_at"`
		ExternalID  string      `json:"external_id"`
		BudgetScope interface{} `json:"budget_scope"`
		Client      struct {
			ID         int         `json:"id"`
			Name       string      `json:"name"`
			Active     bool        `json:"active"`
			ExternalID interface{} `json:"external_id"`
			UpdatedAt  time.Time   `json:"updated_at"`
		} `json:"client"`
		RequiredNotes    bool          `json:"required_notes"`
		BudgetExpiredOn  interface{}   `json:"budget_expired_on"`
		HasRecurrence    bool          `json:"has_recurrence"`
		Budget           int           `json:"budget"`
		BudgetType       string        `json:"budget_type"`
		HourRate         float64       `json:"hour_rate"`
		HourRateInCents  float64       `json:"hour_rate_in_cents"`
		BudgetProgress   float64       `json:"budget_progress"`
		BudgetPercent    float64       `json:"budget_percent"`
		Labels           []interface{} `json:"labels"`
		LabelIds         []interface{} `json:"label_ids"`
		RequiredLabelIds []interface{} `json:"required_label_ids"`
	} `json:"project"`
	Duration struct {
		Hours        int     `json:"hours"`
		Minutes      int     `json:"minutes"`
		Seconds      int     `json:"seconds"`
		Formatted    string  `json:"formatted"`
		TotalHours   float64 `json:"total_hours"`
		TotalSeconds int     `json:"total_seconds"`
		TotalMinutes int     `json:"total_minutes"`
	} `json:"duration"`
	EstimatedDuration struct {
		Hours        int     `json:"hours"`
		Minutes      int     `json:"minutes"`
		Seconds      int     `json:"seconds"`
		Formatted    string  `json:"formatted"`
		TotalHours   float64 `json:"total_hours"`
		TotalSeconds int     `json:"total_seconds"`
		TotalMinutes int     `json:"total_minutes"`
	} `json:"estimated_duration"`
	Cost struct {
		Fractional int     `json:"fractional"`
		Formatted  string  `json:"formatted"`
		Amount     float64 `json:"amount"`
	} `json:"cost"`
	EstimatedCost struct {
		Fractional int     `json:"fractional"`
		Formatted  string  `json:"formatted"`
		Amount     float64 `json:"amount"`
	} `json:"estimated_cost"`
	Day             string        `json:"day"`
	Note            string        `json:"note"`
	Sequence        int           `json:"sequence"`
	Estimated       bool          `json:"estimated"`
	TimerState      string        `json:"timer_state"`
	TimerStartedOn  int           `json:"timer_started_on"`
	TimerStoppedOn  int           `json:"timer_stopped_on"`
	LabelIds        []interface{} `json:"label_ids"`
	UserIds         []interface{} `json:"user_ids"`
	UpdatedAt       int           `json:"updated_at"`
	CreatedAt       int           `json:"created_at"`
	CreatedFrom     string        `json:"created_from"`
	UpdatedFrom     string        `json:"updated_from"`
	Billed          bool          `json:"billed"`
	To              interface{}   `json:"to"`
	From            interface{}   `json:"from"`
	Deleted         bool          `json:"deleted"`
	HourRate        float64       `json:"hour_rate"`
	HourRateInCents float64       `json:"hour_rate_in_cents"`
	CreatorID       int           `json:"creator_id"`
	UpdaterID       int           `json:"updater_id"`
	ExternalID      interface{}   `json:"external_id"`
	EntryIds        []int         `json:"entry_ids"`
	SuggestionID    interface{}   `json:"suggestion_id"`
	Draft           bool          `json:"draft"`
	Manage          bool          `json:"manage"`
}
