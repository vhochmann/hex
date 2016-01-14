package game

import "fmt"

const LogSize = 32

type Log struct{
	Entries []string
}

func (l *Log) Write(s string, args ...interface{}) {
	l.Entries = append(l.Entries, fmt.Sprintf(s, args...))
	if len(l.Entries) > LogSize {
		l.Entries = l.Entries[len(l.Entries)-LogSize:]
	} 
}

func (l *Log) Read(n int) []string {
	if n >= len(l.Entries) {
		return l.Entries
	}
	return l.Entries[len(l.Entries)-n:]
}
