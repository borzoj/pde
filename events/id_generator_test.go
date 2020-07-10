package events

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultIdGeneratorSame(t *testing.T) {
	e1 := Event{
		Type:      "event.test",
		PatientID: nil,
		Created:   nil,
		Attributes: Attributes{
			"three": "four",
			"one":   "two",
			"five":  "six",
		},
	}
	e2 := Event{
		Type:      "event.test",
		PatientID: nil,
		Created:   nil,
		Attributes: Attributes{
			"one":   "two",
			"five":  "six",
			"three": "four",
		},
	}
	id1, err1 := DefaultIDGenerator(e1)
	id2, err2 := DefaultIDGenerator(e2)
	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.Equal(t, id1, id2)
}

func TestDefaultIdGeneratorDifferent(t *testing.T) {
	e1 := Event{
		Type:      "event.test",
		PatientID: nil,
		Created:   nil,
		Attributes: Attributes{
			"three": "four",
			"one":   "two",
			"five":  "six",
		},
	}
	e2 := Event{
		Type:      "event.test",
		PatientID: nil,
		Created:   nil,
		Attributes: Attributes{
			"one":   "two",
			"five":  "sdsdsd",
			"three": "four",
		},
	}
	id1, err1 := DefaultIDGenerator(e1)
	id2, err2 := DefaultIDGenerator(e2)
	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.NotEqual(t, id1, id2)
}

func TestDefaultIdGeneratorReturnsHexString(t *testing.T) {
	e1 := Event{
		Type:      "event.test",
		PatientID: nil,
		Created:   nil,
		Attributes: Attributes{
			"three": "four",
			"one":   "two",
			"five":  "six",
		},
	}
	id1, err1 := DefaultIDGenerator(e1)
	assert.Nil(t, err1)
	match, _ := regexp.MatchString(`^[0-9a-f]+$`, id1)
	assert.True(t, match)
}
