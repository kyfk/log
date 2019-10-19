package log

import (
	"errors"
	"testing"
)

func TestNopLogger(t *testing.T) {
	lg := NopLogger{}
	lg.Debug()
	lg.Info()
	lg.Warn()
	lg.Error(errors.New("error"))
}
