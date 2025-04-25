package json

import (
	"testing"
	"time"

	gj "github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

type Date time.Time

type UnmarshalError struct {
	Date string
}

func (e UnmarshalError) Error() string {
	return "invalid date format: " + e.Date
}

func (d Date) IsZero() bool {
	return time.Time(d).IsZero()
}

func (d *Date) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*d = Date(time.Time{})
		return nil
	}

	tDate, err := time.Parse(`"2006-01-02"`, string(data))
	if err != nil {
		return UnmarshalError{
			Date: string(data),
		}
	}

	*d = Date(tDate)

	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	// return null rather than "0001-01-01", which is not friendly to other languages
	if d.IsZero() {
		return []byte("null"), nil
	}

	return []byte(time.Time(d).Format(`"2006-01-02"`)), nil
}

func Test_Date_null(t *testing.T) {
	t.Parallel()

	type Event struct {
		Date Date `json:"date"`
	}

	event := Event{
		Date: Date{},
	}

	marshaledEvent, err := gj.Marshal(event)
	if err != nil {
		t.Fatal(err)
	}

	var unmarshaledEvent Event
	err = gj.Unmarshal(marshaledEvent, &unmarshaledEvent)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, `{"date":null}`, string(marshaledEvent))
	assert.Equal(t, event, unmarshaledEvent)
}
