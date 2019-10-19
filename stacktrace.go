package log

import (
	"fmt"
	"runtime"
)

type frame uintptr

// pc returns the program counter for this frame;
// multiple frames may have the same PC value.
func (f frame) pc() uintptr { return uintptr(f) - 1 }

// file returns the full path to the file that contains the
// function for this frame's pc.
func (f frame) file() string {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown"
	}
	file, _ := fn.FileLine(f.pc())
	return file
}

// line returns the line number of source code of the
// function for this frame's pc.
func (f frame) line() int {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return 0
	}
	_, line := fn.FileLine(f.pc())
	return line
}

// name returns the name of this function if known.
func (f frame) name() string {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown"
	}
	return fn.Name()
}

func (f frame) marshalText() string {
	name := f.name()
	if name == "unknown" {
		return name
	}
	return fmt.Sprintf("%s %s:%d", name, f.file(), f.line())
}

type stackTrace []frame

func (s stackTrace) framesString() []string {
	arr := make([]string, len(s))
	for i, ss := range s {
		arr[i] = ss.marshalText()
	}
	return arr
}

func callers() stackTrace {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	st := pcs[0:n]

	f := make([]frame, len(st))
	for i := 0; i < len(f); i++ {
		f[i] = frame((st)[i])
	}
	return f
}
