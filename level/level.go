package level

// Level is used as minimum logging level.
// If the minimum logging level is less than logging level that wanted to output,
// the message isn't be outputted.
type Level string

const (
	// Debug is debug level.
	Debug Level = "DEBUG"
	// Info is info level.
	Info Level = "INFO"
	// Warn is warn level.
	Warn Level = "WARN"
	// Error is error level.
	Error Level = "ERROR"
)

// LessThan returns true if the level of receiver is less than the level of argument.
func (l Level) LessThan(ll Level) bool {
	return l.Priority() < ll.Priority()
}

// Priority returns number of priorities.
// The greater number is the higher priority.
func (l Level) Priority() int {
	switch l {
	case Debug:
		return 1
	case Info:
		return 2
	case Warn:
		return 3
	case Error:
		return 4
	default:
		return 9999
	}
}
