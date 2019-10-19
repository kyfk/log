package log

// NopLogger just implements the below simple interface.
// ```
// type Logger interface {
//     Debug(v ...interface{})
//     Info(v ...interface{})
//     Warn(v ...interface{})
//     Error(err error)
// }
// ```
//
// In writting tests, if you use Logger and you don't want any output,
// NopLogger can be used as stabs.
type NopLogger struct{}

// Debug do nothing.
func (NopLogger) Debug(v ...interface{}) {}

// Info do nothing.
func (NopLogger) Info(v ...interface{}) {}

// Warn do nothing.
func (NopLogger) Warn(v ...interface{}) {}

// Error do nothing.
func (NopLogger) Error(err error) {}
