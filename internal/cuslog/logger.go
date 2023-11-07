package cuslog

import (
	"fmt"
	"os"
	"sync"
)

// 默认的logger实例
var std = New()

type logger struct {
	opt       *options
	mu        sync.Mutex
	entryPool *sync.Pool
}

func New(opts ...Option) *logger {
	lg := &logger{opt: initOptions(opts...)}
	lg.entryPool = &sync.Pool{New: func() interface{} { return entry(lg) }}
	return lg
}

func SetOptions(opts ...Option) {
	std.SetOptions(opts...)
}

func (l *logger) SetOptions(opts ...Option) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, opt := range opts {
		opt(l.opt)
	}
}

func (l *logger) entry() *Entry {
	return l.entryPool.Get().(*Entry)
}

func (l *logger) Debug(args ...interface{}) {
	l.entry().write(DebugLevel, fmtEmptySeparate, args...)
}

func (l *logger) Info(args ...interface{}) {
	l.entry().write(InfoLevel, fmtEmptySeparate, args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.entry().write(WarnLevel, fmtEmptySeparate, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.entry().write(ErrorLevel, fmtEmptySeparate, args...)
}

func (l *logger) Panic(args ...interface{}) {
	l.entry().write(PanicLevel, fmtEmptySeparate, args...)
	panic(fmt.Sprint(args...))
}

func (l *logger) Fatal(args ...interface{}) {
	l.entry().write(FatalLevel, fmtEmptySeparate, args...)
	os.Exit(1)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.entry().write(DebugLevel, format, args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.entry().write(InfoLevel, format, args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.entry().write(WarnLevel, format, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.entry().write(ErrorLevel, format, args...)
}

func (l *logger) Panicf(format string, args ...interface{}) {
	l.entry().write(PanicLevel, format, args...)
	panic(fmt.Sprintf(format, args...))
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.entry().write(FatalLevel, format, args...)
	os.Exit(1)
}

// std logger

func Debug(args ...interface{}) {
	std.Debug(args...)
}

func Info(args ...interface{}) {
	std.Info(args...)
}

func Warn(args ...interface{}) {
	std.Warn(args...)
}

func Error(args ...interface{}) {
	std.Error(args...)
}

func Panic(args ...interface{}) {
	std.Panic(args...)
}

func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}
