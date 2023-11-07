package cuslog

import (
	"io"
	"os"
)

const (
	fmtEmptySeparate = ""
)

type Level uint8

// const log level
const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

// LevelNameMapping log level string name mapping
var LevelNameMapping = map[Level]string{
	DebugLevel: "DEGUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
	PanicLevel: "PANIC",
	FatalLevel: "FATAL",
}

type options struct {
	output        io.Writer
	level         Level
	stdLevel      Level
	formatter     Formatter
	disableCaller bool
}

type Option func(*options)

func initOptions(opts ...Option) *options {
	o := &options{}

	for _, opt := range opts {
		opt(o)
	}

	if o.output == nil {
		o.output = os.Stderr
	}

	if o.formatter == nil {
		o.formatter = &TextFormtter{}
	}

	return o
}

func WithOutput(output io.Writer) Option {
	return func(o *options) {
		o.output = output
	}
}

func WithLevel(level Level) Option {
	return func(o *options) {
		o.level = level
	}
}

func WithStdLevel(level Level) Option {
	return func(o *options) {
		o.stdLevel = level
	}
}

func WithFormatter(formatter Formatter) Option {
	return func(o *options) {
		o.formatter = formatter
	}
}

func WithDisableCaller(disableCaller bool) Option {
	return func(o *options) {
		o.disableCaller = disableCaller
	}
}
