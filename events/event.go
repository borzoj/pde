package events

import (
	"time"
)

// Attributes patient attributes
type Attributes map[string]string

// Event patient event
type Event struct {
	Type       string     `json:"type"`
	PatientID  *string    `json:"patient_id,omitempty"`
	Created    *time.Time `json:"created_at,omitempty"`
	Attributes Attributes `json:"attributes"`
}
