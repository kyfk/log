package log

import (
	"errors"
	"log"
	"os"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/kyfk/log/format"
	"github.com/kyfk/log/level"
	"github.com/stretchr/testify/assert"
)

func TestSetMinLevel(t *testing.T) {
	SetMinLevel(level.Debug)
	assert.Equal(t, level.Debug, defaultLogger.level)
	SetMinLevel(level.Warn)
	assert.Equal(t, level.Warn, defaultLogger.level)
}

func TestSetFormat(t *testing.T) {
	SetFormat(format.JSON)
	f1 := runtime.FuncForPC(reflect.ValueOf(format.JSON).Pointer()).Name()
	f2 := runtime.FuncForPC(reflect.ValueOf(defaultLogger.formatter).Pointer()).Name()
	assert.Equal(t, f1, f2)

	SetFormat(format.JSONPretty)
	f3 := runtime.FuncForPC(reflect.ValueOf(format.JSONPretty).Pointer()).Name()
	f4 := runtime.FuncForPC(reflect.ValueOf(defaultLogger.formatter).Pointer()).Name()
	assert.Equal(t, f3, f4)
}

func TestSetMetadata(t *testing.T) {
	var meta1 = map[string]interface{}{"meta1": "meta1"}
	SetMetadata(meta1)
	assert.Equal(t, meta1, defaultLogger.metadata)
	var meta2 = map[string]interface{}{"meta2": "meta2"}
	SetMetadata(meta2)
	assert.Equal(t, meta2, defaultLogger.metadata)
}

func TestSetStdLogger(t *testing.T) {
	lg := log.New(os.Stdout, "", 0)
	SetStdLogger(lg)
	assert.Equal(t, lg, defaultLogger.logger)

	lg1 := log.New(os.Stdout, "", 0)
	SetStdLogger(lg1)
	assert.Equal(t, lg1, defaultLogger.logger)
}

func TestSetFlattenMetadata(t *testing.T) {
	SetFlattenMetadata(true)
	assert.Equal(t, true, defaultLogger.flattenMetadata)

	SetFlattenMetadata(false)
	assert.Equal(t, false, defaultLogger.flattenMetadata)
}

func ExampleOutput() {
	SetMinLevel(level.Debug)
	SetFormat(format.JSONPretty)
	SetMetadata(map[string]interface{}{
		"uesr_id":    "86f32b8b-ec0d-479f-aed1-1070aa54cecf",
		"request_id": "943ad105-7543-11e6-a9ac-65e093327849",
	})
	SetFlattenMetadata(true)
	SetOutput(os.Stdout)

	defaultLogger.nowFunc = func() time.Time { return time.Time{} }
	defaultLogger.withoutTrace = true

	Debug("debug")
	Info("info")
	Warn("warn")
	Error(errors.New("error"))

	// Output:
	//
	// {
	//   "level": "DEBUG",
	//   "message": "debug",
	//   "request_id": "943ad105-7543-11e6-a9ac-65e093327849",
	//   "time": "0001-01-01T00:00:00Z",
	//   "uesr_id": "86f32b8b-ec0d-479f-aed1-1070aa54cecf"
	// }
	// {
	//   "level": "INFO",
	//   "message": "info",
	//   "request_id": "943ad105-7543-11e6-a9ac-65e093327849",
	//   "time": "0001-01-01T00:00:00Z",
	//   "uesr_id": "86f32b8b-ec0d-479f-aed1-1070aa54cecf"
	// }
	// {
	//   "level": "WARN",
	//   "message": "warn",
	//   "request_id": "943ad105-7543-11e6-a9ac-65e093327849",
	//   "time": "0001-01-01T00:00:00Z",
	//   "uesr_id": "86f32b8b-ec0d-479f-aed1-1070aa54cecf"
	// }
	// {
	//   "error": "*errors.errorString: error",
	//   "level": "ERROR",
	//   "request_id": "943ad105-7543-11e6-a9ac-65e093327849",
	//   "time": "0001-01-01T00:00:00Z",
	//   "uesr_id": "86f32b8b-ec0d-479f-aed1-1070aa54cecf"
	// }
	//
}
