package cuslog

import (
	"bytes"
	"runtime"
	"strings"
	"time"
)

type Entry struct {
	logger *logger
	Buffer *bytes.Buffer
	Map    map[string]interface{}
	Level  Level
	Time   time.Time
	File   string
	Line   int
	Func   string
	Format string
	Args   []interface{}
}

func entry(logger *logger) *Entry {
	return &Entry{logger: logger, Buffer: new(bytes.Buffer), Map: make(map[string]interface{}, 5)}
}

func (e *Entry) write(level Level, format string, args ...interface{}) {
	if e.logger.opt.level > level {
		return
	}
	e.Time = time.Now()
	e.Level = level
	e.Format = format
	e.Args = args

	if !e.logger.opt.disableCaller {
		if pc, file, line, ok := runtime.Caller(2); !ok {
			e.File = "???"
			e.Func = "???"
		} else {
			e.File, e.Line, e.Func = file, line, runtime.FuncForPC(pc).Name()
			e.Func = e.Func[strings.LastIndex(e.Func, "/")+1:]
		}
	}

	// 根据logger.opt.formatter类型进行格式化，并将结果暂时存储到e.Buffer
	e.format()
	// 将e.Buffer中暂存的数据，写到logger.opt.output中
	e.writeOutput()
	// 重置entry
	e.release()
}

func (e *Entry) format() {
	_ = e.logger.opt.formatter.Format(e)
}

func (e *Entry) writeOutput() {
	e.logger.mu.Lock()
	defer e.logger.mu.Unlock()

	_, _ = e.logger.opt.output.Write(e.Buffer.Bytes())
}

func (e *Entry) release() {
	e.Args, e.Line, e.File, e.Format, e.Func = nil, 0, "", "", ""
	fields := []string{"level", "time", "file", "func", "message"}
	for _, field := range fields {
		if _, ok := e.Map[field]; ok {
			delete(e.Map, field)
		}
	}
	e.Buffer.Reset()
	// entry重置之后，put回entryPool，下次使用可以直接从临时对象池子取
	e.logger.entryPool.Put(e)
}
