package cuslog

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"strings"
	"time"
)

type JsonFormatter struct {
	IgnoreBasicFields bool
}

func (f *JsonFormatter) Format(e *Entry) error {
	if !f.IgnoreBasicFields {
		e.Map["time"] = e.Time.Format(time.RFC3339)
		e.Map["level"] = LevelNameMapping[e.Level]
		if e.File != "" {
			short := e.File[strings.LastIndex(e.File, "/")+1:]
			e.Map["file"] = fmt.Sprintf("%s:%d", short, e.Line)
			e.Map["func"] = e.Func
		}

		switch e.Format {
		case fmtEmptySeparate:
			e.Map["message"] = fmt.Sprint(e.Args...)
		default:
			e.Map["message"] = fmt.Sprintf(e.Format, e.Args...)
		}
		return jsoniter.NewEncoder(e.Buffer).Encode(e.Map)
	}

	switch e.Format {
	case fmtEmptySeparate:
		e.Map["message"] = fmt.Sprint(e.Args...)
	default:
		e.Map["message"] = fmt.Sprintf(e.Format, e.Args...)
	}
	return jsoniter.NewEncoder(e.Buffer).Encode(e.Map)
}
