package log

import (
	"fmt"
	"io"
	"os"
)

var std = New(os.Stderr, "", "", LstdFlags|Lshortfile, DebugLevel, IsTerminal)

// SetLevel sets the log level.
func SetLevel(level Level) {
	std.mu.Lock()
	defer std.mu.Unlock()
	std.level = level
}

// SetOutput sets the output destination for the standard logger.
func SetOutput(w io.Writer) {
	std.mu.Lock()
	defer std.mu.Unlock()
	std.out = w
}

// Flags returns the output flags for the standard logger.
func Flags() int {
	return std.Flags()
}

// SetFlags sets the output flags for the standard logger.
func SetFlags(flag int) {
	std.SetFlags(flag)
}

// Prefix returns the output prefix for the standard logger.
func Prefix() string {
	return std.Prefix()
}

// SetPrefix sets the output prefix for the standard logger.
func SetPrefix(prefix string) {
	std.SetPrefix(prefix)
}

// Suffix returns the output suffix for the standard logger.
func Suffix() string {
	return std.Suffix()
}

// SetSuffix sets the output suffix for the standard logger.
func SetSuffix(suffix string) {
	std.SetSuffix(suffix)
}

// These functions write to the standard logger.

// Error is the same as Errorf
func Error(format string, v ...interface{}) {
	if std.level >= ErrorLevel {
		std.Output(2, fmt.Sprintf(format, v...), ErrorLevel)
	}
}

// Errorf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, v ...interface{}) {
	if std.level >= ErrorLevel {
		std.Output(2, fmt.Sprintf(format, v...), ErrorLevel)
	}
}

// Errorln debug level log
func Errorln(v ...interface{}) {
	if std.level >= ErrorLevel {
		std.Output(2, fmt.Sprintln(v...), ErrorLevel)
	}
}

// Warn is the same as Warnf
func Warn(format string, v ...interface{}) {
	if std.level >= WarnLevel {
		std.Output(2, fmt.Sprintf(format, v...), WarnLevel)
	}
}

// Warnf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Warnf(format string, v ...interface{}) {
	if std.level >= WarnLevel {
		std.Output(2, fmt.Sprintf(format, v...), WarnLevel)
	}
}

// Warnln debug level log
func Warnln(v ...interface{}) {
	if std.level >= WarnLevel {
		std.Output(2, fmt.Sprintln(v...), WarnLevel)
	}
}

// Info is the same as Infof
func Info(format string, v ...interface{}) {
	if std.level >= InfoLevel {
		std.Output(2, fmt.Sprintf(format, v...), InfoLevel)
	}
}

// Infof calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Infof(format string, v ...interface{}) {
	if std.level >= InfoLevel {
		std.Output(2, fmt.Sprintf(format, v...), InfoLevel)
	}
}

// Infoln debug level log
func Infoln(v ...interface{}) {
	if std.level >= InfoLevel {
		std.Output(2, fmt.Sprintln(v...), InfoLevel)
	}
}

// Verbose is the same as Debug
func Verbose(format string, v ...interface{}) {
	if std.level >= DebugLevel {
		std.Output(2, fmt.Sprintf(format, v...), DebugLevel)
	}
}

// Debug is the same as Debugf
func Debug(format string, v ...interface{}) {
	if std.level >= DebugLevel {
		std.Output(2, fmt.Sprintf(format, v...), DebugLevel)
	}
}

// Debugf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, v ...interface{}) {
	if std.level >= DebugLevel {
		std.Output(2, fmt.Sprintf(format, v...), DebugLevel)
	}
}

// Debugln debug level log
func Debugln(v ...interface{}) {
	if std.level >= DebugLevel {
		std.Output(2, fmt.Sprintln(v...), DebugLevel)
	}
}

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	std.Output(2, fmt.Sprint(v...), std.level)
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf(format, v...), std.level)
}

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	std.Output(2, fmt.Sprintln(v...), std.level)
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	std.Output(2, fmt.Sprint(v...), FatalLevel)
	os.Exit(1)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf(format, v...), FatalLevel)
	os.Exit(1)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...interface{}) {
	std.Output(2, fmt.Sprintln(v...), FatalLevel)
	os.Exit(1)
}

// Panic is equivalent to Print() followed by a call to panic().
func Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	std.Output(2, s, PanicLevel)
	panic(s)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	std.Output(2, s, PanicLevel)
	panic(s)
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	std.Output(2, s, PanicLevel)
	panic(s)
}

// Output writes the output for a logging event. The string s contains
// the text to print after the prefix specified by the flags of the
// Logger. A newline is appended if the last character of s is not
// already a newline. Calldepth is the count of the number of
// frames to skip when computing the file name and line number
// if Llongfile or Lshortfile is set; a value of 1 will print the details
// for the caller of Output.
func Output(calldepth int, s string) error {
	return std.Output(calldepth+1, s, std.level) // +1 for this frame.
}
