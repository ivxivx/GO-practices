package errors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SampleError1 struct {
	RequestID string
}

func (e SampleError1) Error() string {
	return fmt.Sprintf("request %s", e.RequestID)
}

type SampleError2 struct {
	RequestID string
}

func (e *SampleError2) Error() string {
	return fmt.Sprintf("request %s", e.RequestID)
}

func Test_as(t *testing.T) {
	t.Parallel()

	err11 := SampleError1{RequestID: "123"}

	assert.True(t, errors.As(err11, &SampleError1{}))
	assert.True(t, errors.As(err11, ptr(SampleError1{})))

	err12 := &SampleError1{RequestID: "123"}

	assert.False(t, errors.As(err12, &SampleError1{}))
	assert.True(t, errors.As(*err12, &SampleError1{}))

	err21 := SampleError2{RequestID: "123"}

	assert.True(t, errors.As(&err21, ptr(&SampleError2{})))
	assert.True(t, errors.As(&err21, ptr(ptr(SampleError2{}))))

	// runtime: second argument to errors.As must be a non-nil pointer to either a type that implements error, or to any interface type
	// assert.True(t, errors.As(&err21, &SampleError2{}))

	// err22 := &SampleError2{RequestID: "123"}

	// runtime: second argument to errors.As must be a non-nil pointer to either a type that implements error, or to any interface type
	// assert.False(t, errors.As(err22, SampleError2{}))

	var err211 *SampleError2
	assert.True(t, errors.As(&err21, &err211))
}

func ptr[T any](s T) *T {
	return &s
}
