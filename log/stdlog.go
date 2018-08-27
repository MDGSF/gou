package log

import (
	"fmt"
	"io"
	"os"
)

var std = New(os.Stdout, "", "", LLevel|LstdFlags|Lshortfile, InfoLevel, IsTerminal)

func DefaultLog() *Logger {
	return std
}

func Clone() *Logger {
	return std.Clone()
}

// IncrOneCallDepth call depth add one
func IncrOneCallDepth() {
	std.mu.Lock()
	defer std.mu.Unlock()
	std.callDepth++
}

// SetCallDepth set whether log output is terminal
func SetCallDepth(callDepth int) {
	std.mu.Lock()
	defer std.mu.Unlock()
	std.callDepth = callDepth
}

// SetIsTerminal set whether log output is terminal
func SetIsTerminal(isTerminal int) {
	std.mu.Lock()
	defer std.mu.Unlock()
	std.isTerminal = isTerminal
}

// SetLevel sets the log level.
func SetLevel(level Level) {
	std.mu.Lock()
	defer std.mu.Unlock()
	std.level = level
}

// Level returns the log level.
func GetLevel() Level {
	return std.Level()
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

// ContentPrefix returns the output content prefix for the standard logger.
func ContentPrefix() string {
	return std.ContentPrefix()
}

// SetContentPrefix sets the output content prefix for the standard logger.
func SetContentPrefix(prefix string) {
	std.SetContentPrefix(prefix)
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

// These functions write to the standard logger.

// Panic is equivalent to Print() followed by a call to panic().
func Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	std.Output(std.callDepth, s, PanicLevel)
	panic(s)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	std.Output(std.callDepth, s, PanicLevel)
	panic(s)
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	std.Output(std.callDepth, s, PanicLevel)
	panic(s)
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	std.Output(std.callDepth, fmt.Sprint(v...), FatalLevel)
	os.Exit(1)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	std.Output(std.callDepth, fmt.Sprintf(format, v...), FatalLevel)
	os.Exit(1)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...interface{}) {
	std.Output(std.callDepth, fmt.Sprintln(v...), FatalLevel)
	os.Exit(1)
}

// Error is the same as Errorf
func Error(format string, v ...interface{}) {
	if std.level >= ErrorLevel {
		std.Output(std.callDepth, fmt.Sprintf(format, v...), ErrorLevel)
	}
}

// Errorf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, v ...interface{}) {
	if std.level >= ErrorLevel {
		std.Output(std.callDepth, fmt.Sprintf(format, v...), ErrorLevel)
	}
}

// Errorln debug level log
func Errorln(v ...interface{}) {
	if std.level >= ErrorLevel {
		std.Output(std.callDepth, fmt.Sprintln(v...), ErrorLevel)
	}
}

// Warn is the same as Warnf
func Warn(format string, v ...interface{}) {
	if std.level >= WarnLevel {
		std.Output(std.callDepth, fmt.Sprintf(format, v...), WarnLevel)
	}
}

// Warnf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Warnf(format string, v ...interface{}) {
	if std.level >= WarnLevel {
		std.Output(std.callDepth, fmt.Sprintf(format, v...), WarnLevel)
	}
}

// Warnln debug level log
func Warnln(v ...interface{}) {
	if std.level >= WarnLevel {
		std.Output(std.callDepth, fmt.Sprintln(v...), WarnLevel)
	}
}

// Info is the same as Infof
func Info(format string, v ...interface{}) {
	if std.level >= InfoLevel {
		std.Output(std.callDepth, fmt.Sprintf(format, v...), InfoLevel)
	}
}

// Infof calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Infof(format string, v ...interface{}) {
	if std.level >= InfoLevel {
		std.Output(std.callDepth, fmt.Sprintf(format, v...), InfoLevel)
	}
}

// Infoln debug level log
func Infoln(v ...interface{}) {
	if std.level >= InfoLevel {
		std.Output(std.callDepth, fmt.Sprintln(v...), InfoLevel)
	}
}

// Debug is the same as Debugf
func Debug(format string, v ...interface{}) {
	if std.level >= DebugLevel {
		std.Output(std.callDepth, fmt.Sprintf(format, v...), DebugLevel)
	}
}

// Debugf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, v ...interface{}) {
	if std.level >= DebugLevel {
		std.Output(std.callDepth, fmt.Sprintf(format, v...), DebugLevel)
	}
}

// Debugln debug level log
func Debugln(v ...interface{}) {
	if std.level >= DebugLevel {
		std.Output(std.callDepth, fmt.Sprintln(v...), DebugLevel)
	}
}

// Verbose is the same as Verbosef
func Verbose(format string, v ...interface{}) {
	if std.level >= VerboseLevel {
		std.Output(std.callDepth, fmt.Sprintf(format, v...), VerboseLevel)
	}
}

// Verbosef calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Verbosef(format string, v ...interface{}) {
	if std.level >= VerboseLevel {
		std.Output(std.callDepth, fmt.Sprintf(format, v...), VerboseLevel)
	}
}

// Verboseln verbose level log
func Verboseln(v ...interface{}) {
	if std.level >= VerboseLevel {
		std.Output(std.callDepth, fmt.Sprintln(v...), VerboseLevel)
	}
}

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	std.Output(std.callDepth, fmt.Sprint(v...), std.level)
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	std.Output(std.callDepth, fmt.Sprintf(format, v...), std.level)
}

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	std.Output(std.callDepth, fmt.Sprintln(v...), std.level)
}
