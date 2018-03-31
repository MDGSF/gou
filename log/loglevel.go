package log

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
		return gray
	}
	return nocolor

}
