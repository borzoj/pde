package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetPatientEventsEmpty(t *testing.T) {
	store := &ArrayStore{}
	events, err := store.GetPatientEvents("id 1")
	assert.Equal(t, 0, len(events))
	assert.Nil(t, err)
	events, err = store.GetPatientEvents("id 2")
	assert.Equal(t, 0, len(events))
	assert.Nil(t, err)
}

func TestAppend(t *testing.T) {
	id1 := "patient id 1"
	t1 := time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC)
	e1 := Event{
		Type:      "event.test",
		PatientID: &id1,
		Created:   &t1,
	}
	t2 := time.Date(2017, time.February, 16, 0, 0, 0, 0, time.UTC)
	e2 := Event{
		Type:      "event.test",
		PatientID: &id1,
		Created:   &t2,
	}
	store := NewArrayStore()

	err := store.Append(e1)
	assert.Nil(t, err)
	p1Events, err := store.GetPatientEvents("patient id 1")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(p1Events))
	assert.Equal(t, e1, p1Events[0])
	p2Events, err := store.GetPatientEvents("patient id 2")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(p2Events))

	err = store.Append(e2)
	assert.Nil(t, err)
	p1Events, err = store.GetPatientEvents("patient id 1")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(p1Events))
	assert.Equal(t, e1, p1Events[0])
	assert.Equal(t, e2, p1Events[1])
	p2Events, err = store.GetPatientEvents("patient id 2")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(p2Events))

}

func TestAppendEmptyId(t *testing.T) {
	id1 := ""
	t1 := time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC)
	e1 := Event{
		Type:      "event.test",
		PatientID: &id1,
		Created:   &t1,
	}
	store := NewArrayStore()

	err := store.Append(e1)
	assert.NotNil(t, err)
}

func TestAppendNilId(t *testing.T) {
	t1 := time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC)
	e1 := Event{
		Type:      "event.test",
		PatientID: nil,
		Created:   &t1,
	}
	store := NewArrayStore()

	err := store.Append(e1)
	assert.NotNil(t, err)
}
