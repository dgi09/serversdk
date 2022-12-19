package log

type NoneLogger struct {
}

func NewNoneLogger() *NoneLogger {
	return &NoneLogger{}
}

func (l *NoneLogger) Fatal(msg string)       {}
func (l *NoneLogger) Error(msg string)       {}
func (l *NoneLogger) Information(msg string) {}
func (l *NoneLogger) Verbose(msg string)     {}
