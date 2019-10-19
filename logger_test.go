package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/kyfk/log/format"
	"github.com/kyfk/log/level"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	logger := New(
		FlattenMetadata(true),
		Format(format.JSON),
	)

	assert := assert.New(t)
	assert.Equal(level.Debug, logger.level)
	f1 := runtime.FuncForPC(reflect.ValueOf(format.JSON).Pointer()).Name()
	f2 := runtime.FuncForPC(reflect.ValueOf(logger.formatter).Pointer()).Name()
	assert.Equal(f1, f2)
	assert.Equal(map[string]interface{}{}, logger.metadata)
}

func TestDebug(t *testing.T) {
	assert := assert.New(t)

	t.Run("if minimum level is Debug", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		logger := New(
			MinLevel(level.Debug),
			Format(format.JSONPretty),
			Output(buf),
		)
		logger.nowFunc = func() time.Time { return time.Time{} }

		t.Run("output messages correctly", func(t *testing.T) {
			logger.Debug("debug")
			assert.Equal(`{
  "level": "DEBUG",
  "message": "debug",
  "meta": {},
  "time": "0001-01-01T00:00:00Z"
}
`, buf.String())
		})

		buf.Reset()

		t.Run("if no arguments is passed, output nothing", func(t *testing.T) {
			logger.Debug()
			assert.Empty(buf.String())
		})

	})

	t.Run("if minimum level is lower than Info, output nothing", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		logger := New(
			MinLevel(level.Info),
			Format(format.JSON),
			Output(buf),
		)
		logger.nowFunc = func() time.Time { return time.Time{} }
		logger.Debug("debug")
		assert.Empty(buf.String())
	})
}

func TestInfo(t *testing.T) {
	assert := assert.New(t)

	t.Run("if minimum level is Info", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		logger := New(
			MinLevel(level.Info),
			Format(format.JSONPretty),
			Output(buf),
		)
		logger.nowFunc = func() time.Time { return time.Time{} }

		t.Run("output messages correctly", func(t *testing.T) {
			logger.Info("info")
			assert.Equal(`{
  "level": "INFO",
  "message": "info",
  "meta": {},
  "time": "0001-01-01T00:00:00Z"
}
`, buf.String())
		})

		buf.Reset()

		t.Run("if no arguments is passed, output nothing", func(t *testing.T) {
			logger.Info()
			assert.Empty(buf.String())
		})

	})

	t.Run("if minimum level is lower than Info, output nothing", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		logger := New(
			MinLevel(level.Warn),
			Format(format.JSON),
			Output(buf),
		)
		logger.nowFunc = func() time.Time { return time.Time{} }
		logger.Info("info")
		assert.Empty(buf.String())
	})
}

func TestWarn(t *testing.T) {
	assert := assert.New(t)

	t.Run("if minimum level is Warn", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		logger := New(
			MinLevel(level.Warn),
			Format(format.JSONPretty),
			Output(buf),
		)
		logger.nowFunc = func() time.Time { return time.Time{} }

		t.Run("output messages correctly", func(t *testing.T) {
			logger.Warn(errors.New("warn"))
			var mp map[string]interface{}
			assert.NoError(json.Unmarshal(buf.Bytes(), &mp))
			assert.Equal("WARN", mp["level"])
			assert.Equal("warn", mp["message"])
			assert.Equal("0001-01-01T00:00:00Z", mp["time"])
			assert.NotEmpty(mp["trace"])
		})

		buf.Reset()

		t.Run("if no arguments is passed, output nothing", func(t *testing.T) {
			logger.Warn()
			assert.Empty(buf.String())
		})

	})

	t.Run("if minimum level is lower than Warn, output nothing", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		logger := New(
			MinLevel(level.Error),
			Format(format.JSON),
			Output(buf),
		)
		logger.nowFunc = func() time.Time { return time.Time{} }
		logger.Warn("warn")
		assert.Empty(buf.String())
	})
}

type MyError error

func TestError(t *testing.T) {
	assert := assert.New(t)

	t.Run("if minimum level is Error", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		logger := New(
			MinLevel(level.Error),
			Format(format.JSONPretty),
			Output(buf),
		)
		logger.nowFunc = func() time.Time { return time.Time{} }

		t.Run("output messages correctly", func(t *testing.T) {
			logger.Error(MyError(fmt.Errorf("error")))
			var mp map[string]interface{}
			assert.NoError(json.Unmarshal(buf.Bytes(), &mp))
			assert.Equal("ERROR", mp["level"])
			assert.Equal("*errors.errorString: error", mp["error"])
			assert.Equal("0001-01-01T00:00:00Z", mp["time"])
			assert.NotEmpty(mp["trace"])
		})

		buf.Reset()

		t.Run("if nil is passed, output nothing", func(t *testing.T) {
			logger.Error(nil)
			assert.Empty(buf.String())
		})

	})

	t.Run("if minimum level is lower than Error, output nothing", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		logger := New(
			MinLevel(level.Level("EXTREAM")),
			Format(format.JSON),
			Output(buf),
		)
		logger.nowFunc = func() time.Time { return time.Time{} }
		logger.Error(MyError(fmt.Errorf("error")))
		assert.Empty(buf.String())
	})
}

func TestErrorMetadataConflicted(t *testing.T) {
	assert := assert.New(t)
	buf := bytes.NewBuffer(nil)
	meta := map[string]interface{}{
		"error": "error1",
	}
	logger := New(
		Output(buf),
		Metadata(meta),
		FlattenMetadata(true),
	)

	logger.nowFunc = func() time.Time { return time.Time{} }

	logger.Error(fmt.Errorf("error"))

	var mp map[string]interface{}
	assert.NoError(json.Unmarshal(buf.Bytes(), &mp))
	assert.Equal("ERROR", mp["level"])
	assert.Equal("*errors.fundamental: the key of metadata conflicted: key=error", mp["error"])
	assert.Equal(meta, mp["meta"])
	assert.Equal("0001-01-01T00:00:00Z", mp["time"])
	assert.NotEmpty(mp["trace"])
}
