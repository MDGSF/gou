package log

import "strings"

// Level type
type Level uint32

func (level Level) String() string {
	switch level {
	case PanicLevel:
		return "panic"
	case FatalLevel:
		return "fatal"
	case ErrorLevel:
		return "ERRO"
	case WarnLevel:
		return "WARN"
	case InfoLevel:
		return "INFO"
	case DebugLevel:
		return "DEBU"
	case VerboseLevel:
		return "VEBO"
	}
	return "unknown"
}

// Color get level color
func (level Level) Color() int {
	switch level {
	case PanicLevel:
		return red
	case FatalLevel:
		return red
	case ErrorLevel:
		return red
	case WarnLevel:
		return yellow
	case InfoLevel:
		return blue
	case DebugLevel:
		return green
	case VerboseLevel:
		return gray
	}
	return nocolor
}

// NameToLevel get log level from level name.
func NameToLevel(name string) Level {
	lname := strings.ToLower(name)
	if strings.Contains(lname, "panic") {
		return PanicLevel
	} else if strings.Contains(lname, "fatal") {
		return FatalLevel
	} else if strings.Contains(lname, "error") {
		return ErrorLevel
	} else if strings.Contains(lname, "warn") {
		return WarnLevel
	} else if strings.Contains(lname, "info") {
		return InfoLevel
	} else if strings.Contains(lname, "debug") {
		return DebugLevel
	} else if strings.Contains(lname, "verbose") {
		return VerboseLevel
	}
	return VerboseLevel
}
