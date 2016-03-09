package game

import "fmt"

// Debug controls visibility of debug messages
var Debug bool

// LogSize defines the log's max number of entries
const LogSize = 32

// Log consists of string entries
type Log struct{
	Entries []string
}

// Write formats input as fmt.Sprintf does, then adds that entry to the
// log, removing the oldest entry to maintain size
func (l *Log) Write(s string, args ...interface{}) {
	l.Entries = append(l.Entries, fmt.Sprintf(s, args...))
	if len(l.Entries) > LogSize {
		l.Entries = l.Entries[len(l.Entries)-LogSize:]
	}
}

// DebugWrite does the same thing as Write, but only appends the entry
// if the Debug flag is true
func (l *Log) DebugWrite(s string, args ...interface{}) {
	if Debug {
		l.Write(s, args...)
	}
}

// Read returns a slice of string entries
func (l *Log) Read(n int) []string {
	if n >= len(l.Entries) {
		return l.Entries
	}
	return l.Entries[len(l.Entries)-n:]
}

// DumpLog uses fmt to print all Log entries
func (l *Log) DumpLog() {
	for _, v := range l.Entries {
		fmt.Println(v)
	}
}
