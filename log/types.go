package log

type LogLevel int

const (
	Fatal LogLevel = iota
	Subfatal
	Error
	Information
	Verbose
)
