package ini

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

// Line one line
type Line struct {
	comment string
	key     string
	value   string
}

// Section section
type Section struct {
	comment string
	name    string
	lines   []*Line
}

// NewSection create an section
func NewSection(name string) *Section {
	return &Section{
		name:  name,
		lines: make([]*Line, 0),
	}
}

func (sect *Section) replaceOrCreate(key, value, comment string) {
	line := sect.findLineByKey(key)
	if line == nil {
		line = &Line{
			key:   key,
			value: value,
		}
		sect.lines = append(sect.lines, line)
	} else {
		line.value = value
	}
}

func (sect *Section) findLineByKey(key string) *Line {
	for k := range sect.lines {
		line := sect.lines[k]
		if line.key == key {
			return line
		}
	}
	return nil
}

// Config ini config
type Config struct {
	sections       []*Section
	defaultsection *Section
}

// NewConfig create ini config from file
func NewConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return NewConfigFromData(data)
}

// NewConfigFromData create ini config from data.
// bytes.Buffer()
// bytes.NewReader()
// bufio.NewReader()
func NewConfigFromData(data []byte) (*Config, error) {
	conf := &Config{
		sections:       make([]*Section, 0),
		defaultsection: NewSection(""),
	}
	br := bytes.NewReader(data)
	r := bufio.NewReader(br)

	curSection := conf.defaultsection
	var comment bytes.Buffer

	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		line = bytes.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		if line[0] == '#' || line[0] == ';' || line[0] == '"' {
			if comment.Len() > 0 {
				comment.WriteByte('\n')
			}
			comment.Write(line)
			continue
		}

		if len(line) > 2 && line[0] == '[' && line[len(line)-1] == ']' {

			sectionName := string(bytes.TrimSpace(line[1 : len(line)-1]))
			section := NewSection(sectionName)
			conf.addOneSection(section)
			curSection = section

			if comment.Len() > 0 {
				curSection.comment = comment.String()
				comment.Reset()
			}

			continue
		}

		lineparts := bytes.SplitN(line, []byte{'='}, 2)
		if len(lineparts) != 2 {
			return nil, errors.New("read the content error: \"" + string(line) + "\", should key = val")
		}

		linekey := string(bytes.TrimSpace(lineparts[0]))
		linevalue := string(bytes.TrimSpace(lineparts[1]))
		var linecomment string
		if comment.Len() > 0 {
			linecomment = comment.String()
			comment.Reset()
		}
		curSection.replaceOrCreate(linekey, linevalue, linecomment)
	}
	return conf, nil
}

// SaveConfigFile save config to file.
func (conf *Config) SaveConfigFile(filename string) error {
	data := conf.FormatBeauty()
	return ioutil.WriteFile(filename, data, 0644)
}

/*
Set set key = value
key can be "section::key"
*/
func (conf *Config) Set(key, value string) error {
	var section *Section
	var linekey string
	if strings.Contains(key, "::") {
		keyparts := strings.SplitN(key, "::", 2)
		sectionName := keyparts[0]
		section = conf.findOrCreate(sectionName)
		linekey = keyparts[1]
	} else {
		section = conf.defaultsection
		linekey = key
	}

	section.replaceOrCreate(linekey, value, "")
	return nil
}

// String get string by key
func (conf *Config) String(key string) string {
	section, linekey := conf.getSectionAndKey(key)
	if section == nil {
		return ""
	}

	line := section.findLineByKey(linekey)
	if line == nil {
		return ""
	}

	return line.value
}

// DefaultString returns the string value for a given key.
// if err != nil return defaultval
func (conf *Config) DefaultString(key string, defaultval string) string {
	value := conf.String(key)
	if len(value) == 0 {
		return defaultval
	}
	return value
}

// Strings get string slice by key
func (conf *Config) Strings(key string) []string {
	value := conf.String(key)
	if len(value) == 0 {
		return nil
	}
	return strings.Split(value, ";")
}

// DefaultStrings returns the []string value for a given key.
// if err != nil return defaultval
func (conf *Config) DefaultStrings(key string, defaultval []string) []string {
	value := conf.Strings(key)
	if value == nil || len(value) == 0 {
		return defaultval
	}
	return value
}

// Int get int by key
func (conf *Config) Int(key string) (int, error) {
	value := conf.String(key)
	return strconv.Atoi(value)
}

// DefaultInt returns the integer value for a given key.
// if err != nil return defaultval
func (conf *Config) DefaultInt(key string, defaultval int) int {
	value, err := conf.Int(key)
	if err != nil {
		return defaultval
	}
	return value
}

// Int64 get int64 by key
func (conf *Config) Int64(key string) (int64, error) {
	value := conf.String(key)
	return strconv.ParseInt(value, 10, 64)
}

// DefaultInt64 returns the int64 value for a given key.
// if err != nil return defaultval
func (conf *Config) DefaultInt64(key string, defaultval int64) int64 {
	value, err := conf.Int64(key)
	if err != nil {
		return defaultval
	}
	return value
}

// Bool get bool by key
func (conf *Config) Bool(key string) (bool, error) {
	value := conf.String(key)
	switch value {
	case "1", "t", "T", "true", "TRUE", "True", "YES", "yes", "Yes", "Y", "y", "ON", "on", "On":
		return true, nil
	case "0", "f", "F", "false", "FALSE", "False", "NO", "no", "No", "N", "n", "OFF", "off", "Off":
		return false, nil
	}
	return false, fmt.Errorf("invalid bool value")
}

// DefaultBool returns the boolean value for a given key.
// if err != nil return defaultval
func (conf *Config) DefaultBool(key string, defaultval bool) bool {
	value, err := conf.Bool(key)
	if err != nil {
		return defaultval
	}
	return value
}

// Float returns float64 value for a given key.
func (conf *Config) Float(key string) (float64, error) {
	value := conf.String(key)
	return strconv.ParseFloat(value, 64)
}

// DefaultFloat returns float64 value for a given key.
// if err != nil return defaultval
func (conf *Config) DefaultFloat(key string, defaultval float64) float64 {
	value, err := conf.Float(key)
	if err != nil {
		return defaultval
	}
	return value
}

// DIY return raw value for a given key.
func (conf *Config) DIY(key string) (interface{}, error) {
	return nil, nil
}

// GetSection returns map for given section
func (conf *Config) GetSection(sectionName string) (map[string]string, error) {
	section := conf.findSection(sectionName)
	if section == nil {
		return nil, errors.New("Not exist section")
	}
	result := make(map[string]string)
	for k := range section.lines {
		oneline := section.lines[k]
		result[oneline.key] = oneline.value
	}
	return result, nil
}

func (conf *Config) getSectionAndKey(key string) (*Section, string) {
	if !strings.Contains(key, "::") {
		return conf.defaultsection, key
	}

	keyparts := strings.SplitN(key, "::", 2)
	sectionName := keyparts[0]
	linekey := keyparts[1]
	section := conf.findSection(sectionName)
	return section, linekey
}

// findOrCreate 查找 sectionName 对应的 section，找不到的话就创建一个
func (conf *Config) findOrCreate(sectionName string) *Section {
	section := conf.findSection(sectionName)
	if section == nil {
		section = NewSection(sectionName)
		conf.addOneSection(section)
	}
	return section
}

// findSection find section by sectionName
func (conf *Config) findSection(sectionName string) *Section {
	for k := range conf.sections {
		section := conf.sections[k]
		if section.name == sectionName {
			return section
		}
	}
	return nil
}

// addOneSection add one section
func (conf *Config) addOneSection(section *Section) {
	conf.sections = append(conf.sections, section)
}

// FormatBeauty format beauty style
func (conf *Config) FormatBeauty() []byte {
	buf := bytes.NewBufferString("")

	if len(conf.defaultsection.comment) > 0 {
		buf.WriteString(fmt.Sprintf("%v\n", conf.defaultsection.comment))
	}
	for k := range conf.defaultsection.lines {
		oneline := conf.defaultsection.lines[k]
		if len(oneline.comment) > 0 {
			buf.WriteString(fmt.Sprintf("%v\n", oneline.comment))
		}
		buf.WriteString(fmt.Sprintf("%v=%v\n", oneline.key, oneline.value))
	}
	if len(conf.sections) > 0 {
		buf.WriteByte('\n')
	}

	for k := range conf.sections {
		section := conf.sections[k]
		if len(section.comment) > 0 {
			buf.WriteString(fmt.Sprintf("%v\n", section.comment))
		}
		buf.WriteString(fmt.Sprintf("[%v]\n", section.name))
		for i := range section.lines {
			oneline := section.lines[i]
			if len(oneline.comment) > 0 {
				buf.WriteString(fmt.Sprintf("%v\n", oneline.comment))
			}
			buf.WriteString(fmt.Sprintf("%v=%v\n", oneline.key, oneline.value))
		}
		buf.WriteByte('\n')
	}
	return buf.Bytes()
}
