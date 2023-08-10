package log

import (
	golog "log"
)

var lvlDescrMap []string = []string{"[FTL]", "[ERR]", "[INF]", "[VRB]"}

type DefaultLogger struct {
	minLvl LogLevel
}

func NewDefaultLogger(minLvl LogLevel) *DefaultLogger {
	return &DefaultLogger{
		minLvl: minLvl,
	}
}

func (l *DefaultLogger) execute(reqLvl LogLevel, msg string) {
	if reqLvl <= l.minLvl {
		golog.Printf("%s %s", lvlDescrMap[reqLvl], msg)
	}
}

func (l *DefaultLogger) SetMinLevel(lvl LogLevel) {
	l.minLvl = lvl
}

func (l *DefaultLogger) Fatal(msg string) {
	l.execute(Fatal, msg)
}

func (l *DefaultLogger) Error(msg string) {
	l.execute(Error, msg)
}

func (l *DefaultLogger) InformationTest(msg string) {
	l.execute(Information, msg)
}

func (l *DefaultLogger) Verbose(msg string) {
	l.execute(Verbose, msg)
}
