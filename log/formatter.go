package log

// Formatter is used to format the output.
type Formatter interface {
	Format(prefix string,
		suffix string,
		time string,
		file string,
		line int,
		contentPrefix string,
		content string) []byte
}
