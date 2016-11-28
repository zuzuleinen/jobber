package sources

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatDate(t *testing.T) {
	var s StackOverflow

	assert.Equal(t, "-2h", s.FormatDate("2 hours ago"))
	assert.Equal(t, "-1h", s.FormatDate("1 hour ago"))
	assert.Equal(t, "-1h", s.FormatDate("< 1 hour ago"))
	assert.Equal(t, "-24h", s.FormatDate("yesterday"))
	assert.Equal(t, "-48h", s.FormatDate("2 days ago"))
	assert.Equal(t, "-168h", s.FormatDate("1 week ago"))
	assert.Equal(t, "-336h", s.FormatDate("2 weeks ago"))
}
