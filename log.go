package log

import (
	"io"
	"log"

	"github.com/kyfk/log/level"
)

var defaultLogger = New()

// SetMinLevel sets minumum logging level to the default logger.
func SetMinLevel(lv level.Level) {
	defaultLogger.level = lv
}

// SetFormat sets the format of message output to the default logger.
func SetFormat(fm formatter) {
	defaultLogger.formatter = fm
}

// SetMetadata sets metadata to default logger.
// If you use some querying service for searching specific logs like BigQuery,
// CloudWatch Logs Insight, Elasticsearch and other more, SetMetadata can be used to
// set additional information to be able to search for conveniently.
// For instance, HTTP Request ID, the id of user signed in, EC2 instance-id and other more.
func SetMetadata(md map[string]interface{}) {
	defaultLogger.metadata = md
}

// SetOutput sets io.Writer as destination of logging message to the default logger.
func SetOutput(out io.Writer) {
	defaultLogger.logger = log.New(out, "", 0)
}

// SetStdLogger sets StdLogger that is used output message to the default logger.
func SetStdLogger(lg *log.Logger) {
	defaultLogger.logger = lg
}

// SetFlattenMetadata sets the flag if metadata is going to be flattened.
// If the flag is put on, metadata is going to be flattened in output
func SetFlattenMetadata(b bool) {
	defaultLogger.flattenMetadata = b
}

// Debug logs a message at level Debug on the default logger.
func Debug(v ...interface{}) {
	defaultLogger.Debug(v...)
}

// Info logs a message at level Info on the default logger.
func Info(v ...interface{}) {
	defaultLogger.Info(v...)
}

// Warn logs a message at level Warn on the default logger.
func Warn(v ...interface{}) {
	defaultLogger.Warn(v...)
}

// Error logs a message at level Error on the default logger.
func Error(err error) {
	defaultLogger.Error(err)
}
