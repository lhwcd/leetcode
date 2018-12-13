package validparentheses

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
)

func TestIsValid(t *testing.T) {
	//
	s := ""
	output := isValid(s)
	assert.Equal(t, output, true)

	//
	s = "()"
	output = isValid(s)
	assert.Equal(t, output, true)
	//
	s = "{[]}"
	output = isValid(s)
	assert.Equal(t, output, true)
	//
	s = "{[)]}"
	output = isValid(s)
	assert.Equal(t, output, false)
}
