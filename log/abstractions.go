package log

type ILogger interface {
	Fatal(msg string)
	Error(msg string)
	Information(msg string)
	Verbose(msg string)
}
