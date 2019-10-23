package log

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/kyfk/log/format"
	"github.com/kyfk/log/level"
	"github.com/pkg/errors"
)

type formatter func(map[string]interface{}) (string, error)

// Logger has fields that Option or setter SetXXXX set.
type Logger struct {
	level           level.Level
	logger          *log.Logger
	formatter       formatter
	metadata        map[string]interface{}
	flattenMetadata bool
	isMergeFailed   bool
	isFormatFailed  bool

	// these fields are only for testing
	nowFunc      func() time.Time
	withoutTrace bool
}

// New initialize new Logger with options.
func New(ops ...Option) *Logger {
	lg := Logger{
		level:     level.Debug,
		logger:    log.New(os.Stdout, "", 0),
		formatter: format.JSONPretty,
		metadata:  map[string]interface{}{},
		nowFunc:   time.Now,
	}

	for _, o := range ops {
		lg = o(lg)
	}
	return &lg
}

// SetMetadata sets a metadata to a logger.
func (l *Logger) SetMetadata(meta map[string]interface{}) {
	l.metadata = meta
}

// Debug logs a message at level Debug.
func (l *Logger) Debug(v ...interface{}) {
	if level.Debug.LessThan(l.level) || len(v) == 0 {
		return
	}
	l.println(map[string]interface{}{
		"level":   level.Debug,
		"message": fmt.Sprint(v...),
		"time":    l.nowFunc(),
	})
}

// Info logs a message at level Info.
func (l *Logger) Info(v ...interface{}) {
	if level.Info.LessThan(l.level) || len(v) == 0 {
		return
	}
	l.println(map[string]interface{}{
		"level":   level.Info,
		"message": fmt.Sprint(v...),
		"time":    l.nowFunc(),
	})
}

// Warn logs a message at level Warn.
func (l *Logger) Warn(v ...interface{}) {
	if level.Warn.LessThan(l.level) || len(v) == 0 {
		return
	}

	data := map[string]interface{}{
		"level": level.Warn,
		"time":  l.nowFunc(),
	}

	switch v0 := v[0].(type) {
	case interface{ StackTrace() errors.StackTrace }:
		if !l.withoutTrace {
			data["trace"] = v0.StackTrace()
			data["message"] = fmt.Sprint(v[0:]...)
		} else {
			data["message"] = fmt.Sprint(v...)
		}
	default:
		if !l.withoutTrace {
			data["trace"] = callers().framesString()
		}
		data["message"] = fmt.Sprint(v...)
	}

	l.println(data)
}

// Error logs a message at level Error.
func (l *Logger) Error(err error) {
	if level.Error.LessThan(l.level) || err == nil {
		return
	}

	data := map[string]interface{}{
		"level": level.Error,
		"time":  l.nowFunc(),
	}

	if !l.withoutTrace {
		switch v := err.(type) {
		case interface{ StackTrace() errors.StackTrace }:
			data["trace"] = v.StackTrace()
		default:
			data["trace"] = callers().framesString()
		}
	}

	err = errors.Cause(err)
	data["error"] = fmt.Sprintf("%s: %s", reflect.TypeOf(err), err.Error())

	l.println(data)
}

func (l *Logger) println(v map[string]interface{}) {
	var data map[string]interface{}
	if l.flattenMetadata && !l.isMergeFailed {
		var err error
		data, err = merge(v, l.metadata)
		if err != nil {
			l.isMergeFailed = true
			l.Error(err)
			return
		}
	} else {
		data = v
		if l.metadata != nil {
			v["meta"] = l.metadata
		}
	}

	s, err := l.formatter(data)
	if err != nil {
		if l.isFormatFailed {
			return
		}
		l.isFormatFailed = true
		l.Error(err)
		return
	}
	l.logger.Println(s)
}

func merge(a, b map[string]interface{}) (map[string]interface{}, error) {
	for k, v := range b {
		_, ok := a[k]
		if ok {
			return map[string]interface{}{}, errors.Errorf("the key of metadata conflicted: key=%s", k)
		}
		a[k] = v
	}
	return a, nil
}
