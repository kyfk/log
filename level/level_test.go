package level

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevel(t *testing.T) {
	assert := assert.New(t)
	assert.True(Debug.LessThan(Info))
	assert.True(Info.LessThan(Warn))
	assert.True(Warn.LessThan(Error))
}
