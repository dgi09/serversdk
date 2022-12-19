package log

type LogLevel int

const (
	Fatal LogLevel = iota
	Error
	Information
	Verbose
)
