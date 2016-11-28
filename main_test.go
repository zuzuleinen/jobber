package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseTime(t *testing.T) {
	assert.Equal(t, "-2h", ParseTime("2 hours ago"))
	assert.Equal(t, "-1h", ParseTime("1 hour ago"))
	assert.Equal(t, "-1h", ParseTime("< 1 hour ago"))
}
