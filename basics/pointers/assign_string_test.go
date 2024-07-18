package pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_string_assignment1(t *testing.T) {
	t.Parallel()

	s1 := "abc"

	s2 := &s1

	s3 := s1

	s4 := s2

	assert.Equal(t, s1, "abc")
	assert.Equal(t, *s2, "abc")
	assert.Equal(t, s3, "abc")
	assert.Equal(t, *s4, "abc")

	*s2 = "123"

	assert.Equal(t, s1, "123")
	assert.Equal(t, *s2, "123")
	assert.Equal(t, s3, "abc")
	assert.Equal(t, *s4, "123")

	s3 = "456"

	assert.Equal(t, s1, "123")
	assert.Equal(t, *s2, "123")
	assert.Equal(t, s3, "456")
	assert.Equal(t, *s4, "123")
}

func Test_string_assignment2(t *testing.T) {
	t.Parallel()

	s1 := "abc"

	s2 := &s1

	s3 := s1

	s4 := &s3

	assert.Equal(t, s1, "abc")
	assert.Equal(t, *s2, "abc")
	assert.Equal(t, s3, "abc")
	assert.Equal(t, *s4, "abc")

	*s4 = "123"

	assert.Equal(t, s1, "abc")
	assert.Equal(t, *s2, "abc")
	assert.Equal(t, s3, "123")
	assert.Equal(t, *s4, "123")
}

func Test_string_assignment3(t *testing.T) {
	s1 := "abc"
	s2 := &s1

	s3 := "abc"
	s4 := &s3

	assert.Equal(t, s1, s3)
	// different memory address
	assert.NotSame(t, s2, s4)
	// same value
	assert.Equal(t, *s2, *s4)
}

func Test_string_assignment4(t *testing.T) {
	s1 := "abc"
	s2 := &s1

	s3 := ""
	s4 := (*string)(nil)

	// different value
	assert.NotEqual(t, s1, s3)
	// different memory address
	assert.NotEqual(t, s2, s4)
}
