package pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_return_array(t *testing.T) {
	t.Parallel()

	a := getArray()
	assert.Equal(t, a, []string{"111", "222", "333"})

	a[0] = "aaa"
	a[1] = "bbb"

	b := getArray()
	assert.Equal(t, b, []string{"111", "222", "333"})

	c := &b
	(*c)[0] = "aaa"
	(*c)[1] = "bbb"
	assert.Equal(t, c, &[]string{"aaa", "bbb", "333"})

	d := getArray()
	assert.Equal(t, d, []string{"111", "222", "333"})
}

func getArray() []string {
	return []string{
		"111",
		"222",
		"333",
	}
}
