package logrus

import (
	"bytes"
	"fmt"
	"path"
	"runtime"
	"sort"
	"strings"
	"sync"
)

// HJFormatter formats logs into text
type HJFormatter struct {
	// QuoteEmptyFields will wrap empty fields in quotes if true
	QuoteEmptyFields bool

	// Whether the logger's out is to a terminal
	isTerminal bool

	sync.Once
}

func (f *HJFormatter) init(entry *Entry) {
	if entry.Logger != nil {
		f.isTerminal = checkIfTerminal(entry.Logger.Out)
	}
}

// Format renders a single log entry
func (f *HJFormatter) Format(entry *Entry) ([]byte, error) {
	var b *bytes.Buffer
	keys := make([]string, 0, len(entry.Data))
	for k := range entry.Data {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	prefixFieldClashes(entry.Data)

	f.Do(func() { f.init(entry) })

	timestampFormat := "2006-01-02 15:04:05.999"

	file := ""
	line := 0
	funcName := ""
	fileName := ""
	pc, file, line, ok := runtime.Caller(4)
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
		fileName = path.Base(file)
	}

	if f.isTerminal {
		f.printColored(b, entry, keys, timestampFormat, fileName, funcName, line)
	} else {
		f.printNoColored(b, entry, keys, timestampFormat, fileName, funcName, line)
	}

	b.WriteByte('\n')

	return b.Bytes(), nil
}

func (f *HJFormatter) printColored(b *bytes.Buffer, entry *Entry, keys []string,
	timestampFormat string, file string, funcName string, line int) {

	var levelColor int
	switch entry.Level {
	case DebugLevel:
		levelColor = gray
	case WarnLevel:
		levelColor = yellow
	case ErrorLevel, FatalLevel, PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	levelText := strings.ToUpper(entry.Level.String())[0:4]

	fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m[%s:%s:%s:%d] %-44s ", levelColor, levelText, entry.Time.Format(timestampFormat),
		file, funcName, line, entry.Message)
	for _, k := range keys {
		v := entry.Data[k]
		fmt.Fprintf(b, " \x1b[%dm%s\x1b[0m=", levelColor, k)
		f.appendValue(b, v)
	}
}

func (f *HJFormatter) printNoColored(b *bytes.Buffer, entry *Entry, keys []string,
	timestampFormat string, file string, funcName string, line int) {

	levelText := strings.ToUpper(entry.Level.String())[0:4]

	fmt.Fprintf(b, "%s[%s:%s:%s:%d] %-44s ", levelText, entry.Time.Format(timestampFormat),
		file, funcName, line, entry.Message)
	for _, k := range keys {
		v := entry.Data[k]
		fmt.Fprintf(b, " %s=", k)
		f.appendValue(b, v)
	}
}

func (f *HJFormatter) needsQuoting(text string) bool {
	if f.QuoteEmptyFields && len(text) == 0 {
		return true
	}
	for _, ch := range text {
		if !((ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') ||
			ch == '-' || ch == '.' || ch == '_' || ch == '/' || ch == '@' || ch == '^' || ch == '+') {
			return true
		}
	}
	return false
}

func (f *HJFormatter) appendValue(b *bytes.Buffer, value interface{}) {
	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}

	if !f.needsQuoting(stringVal) {
		b.WriteString(stringVal)
	} else {
		b.WriteString(fmt.Sprintf("%q", stringVal))
	}
}
