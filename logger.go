package tail

import (
	"io/ioutil"
	"log"
	"os"
)

type Logger interface {
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
	Debugf(format string, v ...interface{})
	Debug(v ...interface{})
}

type logger struct {
	*log.Logger
}

// NewDefaultLogger creates new logger when Config.Logger == nil.
func NewDefaultLogger() Logger {
	return &logger{Logger: log.New(os.Stderr, "", log.LstdFlags)}
}

// NewDiscardingLogger creates new disabled logger.
func NewDiscardingLogger() Logger {
	return &logger{Logger: log.New(ioutil.Discard, "", 0)}
}

func (l *logger) Debugf(format string, v ...interface{}) {
	l.Printf(format, v)
}

func (l *logger) Debug(v ...interface{}) {
	l.Print(v)
}
