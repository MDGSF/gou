package log

// TextFormatter format logs into text.
type TextFormatter struct {
}

// Format returns the bytes log.
func (f *TextFormatter) Format(prefix string,
	suffix string,
	time string,
	file string,
	line int,
	contentPrefix string,
	content string) []byte {
	return nil
}
