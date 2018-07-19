package mlog

import (
	"github.com/MDGSF/utils/log"
)

type MLogger struct {
	loggers []*log.Logger
}

func New() *MLogger {
	mlog := &MLogger{}
	mlog.loggers = make([]*log.Logger, 0)
	return mlog
}

func (mlog *MLogger) Clone() *MLogger {
	newMultiLog := New()
	for k, _ := range mlog.loggers {
		curLog := mlog.loggers[k]
		newLog := curLog.Clone()
		newMultiLog.loggers = append(newMultiLog.loggers, newLog)
	}
	return newMultiLog
}

func (mlog *MLogger) SetLevel(level log.Level) {
	for _, log := range mlog.loggers {
		log.SetLevel(level)
	}
}

func (mlog *MLogger) AddOneLogger(logger *log.Logger) {
	logger.IncrOneCallDepth()
	mlog.loggers = append(mlog.loggers, logger)
}

// Panic is equivalent to Print() followed by a call to panic().
func (mlog *MLogger) Panic(v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Panic(v...)
	}
}

// Panicf is equivalent to Printf() followed by a call to panic().
func (mlog *MLogger) Panicf(format string, v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Panicf(format, v...)
	}
}

// Panicln is equivalent to Println() followed by a call to panic().
func (mlog *MLogger) Panicln(v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Panicln(v...)
	}
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func (mlog *MLogger) Fatal(v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Fatal(v...)
	}
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func (mlog *MLogger) Fatalf(format string, v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Fatalf(format, v...)
	}
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func (mlog *MLogger) Fatalln(v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Fatalln(v...)
	}
}

// Error is the same as Errorf
func (mlog *MLogger) Error(format string, v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Error(format, v...)
	}
}

// Errorf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (mlog *MLogger) Errorf(format string, v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Errorf(format, v...)
	}
}

// Errorln debug level log
func (mlog *MLogger) Errorln(v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Errorln(v...)
	}
}

// Warn is the same as Warnf
func (mlog *MLogger) Warn(format string, v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Warn(format, v...)
	}
}

// Warnf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (mlog *MLogger) Warnf(format string, v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Warnf(format, v...)
	}
}

// Warnln debug level log
func (mlog *MLogger) Warnln(v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Warnln(v...)
	}
}

// Info is the same as Infof
func (mlog *MLogger) Info(format string, v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Info(format, v...)
	}
}

// Infof calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (mlog *MLogger) Infof(format string, v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Infof(format, v...)
	}
}

// Infoln debug level log
func (mlog *MLogger) Infoln(v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Infoln(v...)
	}
}

// Debug is the same as Debugf
func (mlog *MLogger) Debug(format string, v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Debug(format, v...)
	}
}

// Debugf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (mlog *MLogger) Debugf(format string, v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Debugf(format, v...)
	}
}

// Debugln debug level log
func (mlog *MLogger) Debugln(v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Debugln(v...)
	}
}

// Verbose is the same as Verbosef
func (mlog *MLogger) Verbose(format string, v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Verbose(format, v...)
	}
}

// Verbosef calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (mlog *MLogger) Verbosef(format string, v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Verbosef(format, v...)
	}
}

// Verboseln verbose level log
func (mlog *MLogger) Verboseln(v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Verboseln(v...)
	}
}

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func (mlog *MLogger) Print(v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Print(v...)
	}
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (mlog *MLogger) Printf(format string, v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Printf(format, v...)
	}
}

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func (mlog *MLogger) Println(v ...interface{}) {
	for _, log := range mlog.loggers {
		log.Println(v...)
	}
}
