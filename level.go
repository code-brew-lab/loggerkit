package loggerkit

import "errors"

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarning
	LevelError
)

type LogLevel int

func (ll LogLevel) String() string {
	switch ll {
	case 0:
		return "DEBUG "
	case 1:
		return "INFO "
	case 2:
		return "WARNING "
	case 3:
		return "ERROR "
	default:
		return ""
	}
}

func newLogLevel(level string) (LogLevel, error) {
	switch level {
	case "DEBUG":
		return LevelDebug, nil
	case "INFO":
		return LevelInfo, nil
	case "WARNING":
		return LevelWarning, nil
	case "ERROR":
		return LevelError, nil
	default:
		return -1, errors.New("invalid log level")
	}
}
