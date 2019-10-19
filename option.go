package log

import (
	"io"
	"log"

	"github.com/kyfk/log/level"
)

// Option is a function for initialization in the constructor of Logger.
type Option func(Logger) Logger

// MinLevel returns Option that sets minumum logging level to a new logger.
func MinLevel(lv level.Level) Option {
	return func(l Logger) Logger {
		l.level = lv
		return l
	}
}

// Format returns Option that sets the format of message output to a new logger.
func Format(fm formatter) Option {
	return func(l Logger) Logger {
		l.formatter = fm
		return l
	}
}

// Metadata sets metadata to a new logger.
// If you use some querying service for searching specific logs like BigQuery,
// CloudWatch Logs Insight, Elasticsearch and other more, SetMetadata can be used to
// set additional information to be able to search for conveniently.
// For instance, HTTP Request ID, the id of user signed in, EC2 instance-id and other more.
func Metadata(md map[string]interface{}) Option {
	return func(l Logger) Logger {
		l.metadata = md
		return l
	}
}

// Output returns Option that sets io.Writer as the destination of logging message to a new logger.
func Output(out io.Writer) Option {
	return func(l Logger) Logger {
		l.logger = log.New(out, "", 0)
		return l
	}
}

// StdLogger returns Option that sets StdLogger that is used output message to a new logger.
func StdLogger(lg *log.Logger) Option {
	return func(l Logger) Logger {
		l.logger = lg
		return l
	}
}

// FlattenMetadata returns Option that sets the flag if metadata is going to be flattened.
// If the flag is put on, metadata is going to be flattened in output.
func FlattenMetadata(b bool) Option {
	return func(l Logger) Logger {
		l.flattenMetadata = b
		return l
	}
}
