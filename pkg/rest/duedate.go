package rest

import "time"

type DueDate struct {
	// Date in format YYYY-MM-DD corrected to user's timezone.
	Date string `json:"date"`

	// Datetime is only returned if exact due time set (i.e. it's
	// not a whole-day task), date and time in RFC3339 format in UTC.
	Datetime *time.Time `json:"datetime,omitempty"`

	// Recurring indicates whether the task has a recurring due date.
	Recurring bool `json:"recurring"`

	// String is human defined date in arbitrary format.
	String string `json:"string"`

	// Timezone is only returned if exact due time set, user's timezone
	// definition either in tzdata-compatible format ("Europe/Berlin")
	// or as a string specifying east of UTC offset as "UTC±HH:MM" (i.e.
	// "UTC-01:00").
	Timezone *string `json:"timezone,omitempty"`
}
