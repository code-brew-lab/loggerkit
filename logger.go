package loggerkit

import (
	"io"
	"log"
	"os"
)

type (
	Logger interface {
		Info(msg string, v ...interface{})
		Debug(msg string, v ...interface{})
		Warning(msg string, v ...interface{})
		Error(msg string, v ...interface{})
		FatalError(msg string, v ...interface{})
		PanicError(msg string, v ...interface{})
	}

	OptFunc func(*opt) (*opt, error)

	logger struct {
		*opt
	}

	opt struct {
		*log.Logger
		level LogLevel
	}
)

func defaultOpts() *opt {
	return &opt{log.New(os.Stdout, "", log.Default().Flags()), LevelDebug}
}

func WithLevel(level string) OptFunc {
	return func(o *opt) (*opt, error) {
		l, err := newLogLevel(level)
		if err != nil {
			return nil, err
		}
		o.level = l
		return o, nil
	}
}

func WithWriter(out io.Writer) OptFunc {
	return func(o *opt) (*opt, error) {
		o.SetOutput(out)
		return o, nil
	}
}

func New(opts ...OptFunc) (Logger, error) {
	var (
		opt *opt
		err error
	)

	opt = defaultOpts()
	for _, optFunc := range opts {
		opt, err = optFunc(opt)
		if err != nil {
			return nil, err
		}
	}

	return &logger{opt}, nil
}

func (l *logger) Debug(msg string, v ...interface{}) {
	if l.level > LevelDebug {
		return
	}

	l.SetPrefix(LevelDebug.String())
	l.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	l.Printf(msg, v...)
}

func (l *logger) Info(msg string, v ...interface{}) {
	l.SetPrefix(LevelInfo.String())
	l.SetFlags(log.Ltime)

	l.Printf(msg, v...)
}

func (l *logger) Warning(msg string, v ...interface{}) {
	if l.level > LevelWarning {
		return
	}

	l.SetPrefix(LevelWarning.String())
	l.SetFlags(log.Ltime | log.Llongfile)

	l.Printf(msg, v...)
}

func (l *logger) Error(msg string, v ...interface{}) {
	if l.level > LevelError {
		return
	}

	l.SetPrefix(LevelError.String())
	l.SetFlags(log.Ltime | log.Llongfile)

	l.Printf(msg, v...)
}

func (l *logger) FatalError(msg string, v ...interface{}) {
	l.SetPrefix(LevelError.String())
	l.SetFlags(log.Ltime | log.Llongfile)

	l.Fatalf(msg, v...)
}

func (l *logger) PanicError(msg string, v ...interface{}) {
	l.SetPrefix(LevelError.String())
	l.SetFlags(log.Ltime | log.Llongfile)

	l.Panicf(msg, v...)
}
