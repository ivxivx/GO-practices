package times

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_UTC(t *testing.T) {
	t.Parallel()

	today := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	_, todayOffset := today.Zone()

	yesterday := today.AddDate(0, 0, -1)
	_, yesterdayOffset := yesterday.Zone()

	assert.Equal(t, 2024, today.Year())
	assert.Equal(t, 1, int(today.Month()))
	assert.Equal(t, 1, today.Day())
	assert.Equal(t, 0, today.Second())
	assert.Equal(t, 0, todayOffset)

	assert.Equal(t, 2023, yesterday.Year())
	assert.Equal(t, 12, int(yesterday.Month()))
	assert.Equal(t, 31, yesterday.Day())
	assert.Equal(t, 0, yesterday.Second())
	assert.Equal(t, 0, yesterdayOffset)
}

func Test_NonUTC(t *testing.T) {
	t.Parallel()

	today, err := time.Parse("2006-01-02T15:04:05Z07:00", "2024-01-01T00:00:00+07:00")
	assert.NoError(t, err)

	_, todayOffset := today.Zone()

	yesterday := today.AddDate(0, 0, -1)

	_, yesterdayOffset := yesterday.Zone()

	assert.Equal(t, 2024, today.Year())
	assert.Equal(t, 1, int(today.Month()))
	assert.Equal(t, 1, today.Day())
	assert.Equal(t, 0, today.Second())
	assert.Equal(t, int64(7*time.Hour/time.Second), int64(todayOffset))

	assert.Equal(t, 2023, yesterday.Year())
	assert.Equal(t, 12, int(yesterday.Month()))
	assert.Equal(t, 31, yesterday.Day())
	assert.Equal(t, 0, yesterday.Second())
	assert.Equal(t, int64(7*time.Hour/time.Second), int64(yesterdayOffset))
}

func Test_convertToUTC(t *testing.T) {
	t.Parallel()

	from, err := time.Parse("2006-01-02T15:04:05Z07:00", "2024-01-01T00:00:00+08:00")
	assert.NoError(t, err)

	_, fromOffset := from.Zone()
	assert.Equal(t, int64(8*time.Hour/time.Second), int64(fromOffset))

	to := from.In(time.UTC)

	_, toOffset := to.Zone()

	assert.Equal(t, 2023, to.Year())
	assert.Equal(t, 12, int(to.Month()))
	assert.Equal(t, 31, to.Day())
	assert.Equal(t, 16, to.Hour())
	assert.Equal(t, 0, to.Minute())
	assert.Equal(t, 0, to.Second())
	assert.Equal(t, 0, toOffset)
}