package log

import (
	"log"
	"os"
	"reflect"
	"runtime"
	"testing"

	"github.com/kyfk/log/format"
	"github.com/kyfk/log/level"
	"github.com/stretchr/testify/assert"
)

func TestMinLevel(t *testing.T) {
	assert := assert.New(t)
	lg1 := MinLevel(level.Debug)(Logger{})
	assert.Equal(level.Debug, lg1.level)
	lg2 := MinLevel(level.Error)(Logger{})
	assert.Equal(level.Error, lg2.level)
}

func TestFormat(t *testing.T) {
	assert := assert.New(t)
	lg1 := Format(format.JSON)(Logger{})
	f1 := runtime.FuncForPC(reflect.ValueOf(format.JSON).Pointer()).Name()
	f2 := runtime.FuncForPC(reflect.ValueOf(lg1.formatter).Pointer()).Name()
	assert.Equal(f1, f2)

	lg2 := Format(format.JSONPretty)(Logger{})
	f3 := runtime.FuncForPC(reflect.ValueOf(format.JSONPretty).Pointer()).Name()
	f4 := runtime.FuncForPC(reflect.ValueOf(lg2.formatter).Pointer()).Name()
	assert.Equal(f3, f4)
}

func TestMetadata(t *testing.T) {
	assert := assert.New(t)

	meta1 := map[string]interface{}{}
	lg1 := Metadata(meta1)(Logger{})
	assert.Equal(meta1, lg1.metadata)

	meta2 := map[string]interface{}{}
	lg2 := Metadata(meta2)(Logger{})
	assert.Equal(meta2, lg2.metadata)
}

func TestFlattenMetadata(t *testing.T) {
	assert := assert.New(t)

	lg1 := FlattenMetadata(true)(Logger{})
	assert.Equal(true, lg1.flattenMetadata)

	lg2 := FlattenMetadata(false)(Logger{})
	assert.Equal(false, lg2.flattenMetadata)
}

func TestStdLogger(t *testing.T) {
	assert := assert.New(t)

	slg1 := log.New(os.Stdout, "", 0)
	lg1 := StdLogger(slg1)(Logger{})
	assert.Equal(slg1, lg1.logger)

	slg2 := log.New(os.Stdout, "", 0)
	lg2 := StdLogger(slg2)(Logger{})
	assert.Equal(slg2, lg2.logger)
}
