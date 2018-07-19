package mlog

import "github.com/MDGSF/utils/log"

var std = New()

func Clone() *MLogger {
	return std.Clone()
}

func AddOneLogger(logger *log.Logger) {
	logger.IncrOneCallDepth()
	std.loggers = append(std.loggers, logger)
}

func SetLevel(level log.Level) {
	std.SetLevel(level)
}

// Panic is equivalent to Print() followed by a call to panic().
func Panic(v ...interface{}) {
	for _, log := range std.loggers {
		log.Panic(v...)
	}
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	for _, log := range std.loggers {
		log.Panicf(format, v...)
	}
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...interface{}) {
	for _, log := range std.loggers {
		log.Panicln(v...)
	}
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	for _, log := range std.loggers {
		log.Fatal(v...)
	}
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	for _, log := range std.loggers {
		log.Fatalf(format, v...)
	}
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...interface{}) {
	for _, log := range std.loggers {
		log.Fatalln(v...)
	}
}

// Error is the same as Errorf
func Error(format string, v ...interface{}) {
	for _, log := range std.loggers {
		log.Error(format, v...)
	}
}

// Errorf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, v ...interface{}) {
	for _, log := range std.loggers {
		log.Errorf(format, v...)
	}
}

// Errorln debug level log
func Errorln(v ...interface{}) {
	for _, log := range std.loggers {
		log.Errorln(v...)
	}
}

// Warn is the same as Warnf
func Warn(format string, v ...interface{}) {
	for _, log := range std.loggers {
		log.Warn(format, v...)
	}
}

// Warnf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Warnf(format string, v ...interface{}) {
	for _, log := range std.loggers {
		log.Warnf(format, v...)
	}
}

// Warnln debug level log
func Warnln(v ...interface{}) {
	for _, log := range std.loggers {
		log.Warnln(v...)
	}
}

// Info is the same as Infof
func Info(format string, v ...interface{}) {
	for _, log := range std.loggers {
		log.Info(format, v...)
	}
}

// Infof calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Infof(format string, v ...interface{}) {
	for _, log := range std.loggers {
		log.Infof(format, v...)
	}
}

// Infoln debug level log
func Infoln(v ...interface{}) {
	for _, log := range std.loggers {
		log.Infoln(v...)
	}
}

// Debug is the same as Debugf
func Debug(format string, v ...interface{}) {
	for _, log := range std.loggers {
		log.Debug(format, v...)
	}
}

// Debugf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, v ...interface{}) {
	for _, log := range std.loggers {
		log.Debugf(format, v...)
	}
}

// Debugln debug level log
func Debugln(v ...interface{}) {
	for _, log := range std.loggers {
		log.Debugln(v...)
	}
}

// Verbose is the same as Verbosef
func Verbose(format string, v ...interface{}) {
	for _, log := range std.loggers {
		log.Verbose(format, v...)
	}
}

// Verbosef calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Verbosef(format string, v ...interface{}) {
	for _, log := range std.loggers {
		log.Verbosef(format, v...)
	}
}

// Verboseln verbose level log
func Verboseln(v ...interface{}) {
	for _, log := range std.loggers {
		log.Verboseln(v...)
	}
}

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	for _, log := range std.loggers {
		log.Print(v...)
	}
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	for _, log := range std.loggers {
		log.Printf(format, v...)
	}
}

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	for _, log := range std.loggers {
		log.Println(v...)
	}
}
