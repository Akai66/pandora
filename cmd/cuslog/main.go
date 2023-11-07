package main

import (
	"os"

	"github.com/Akai66/pandora/internal/cuslog"
)

func main() {
	// 使用std logger
	cuslog.Debug("use std debug log")
	cuslog.SetOptions(cuslog.WithLevel(cuslog.InfoLevel))
	cuslog.Debug("some debug log")
	cuslog.Info("change level to info")
	cuslog.Error("some error log")
	//cuslog.Panic("some panic log", " test panic")
	//cuslog.Fatalf("server: failed to listen :%d", 8080)

	// 自定义logger，输出到文件
	fd, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		cuslog.Fatalf("failed to open test.log: %v", err)
	}
	defer fd.Close()
	lg := cuslog.New(cuslog.WithLevel(cuslog.WarnLevel),
		cuslog.WithOutput(fd),
		cuslog.WithFormatter(&cuslog.JsonFormatter{}),
		cuslog.WithDisableCaller(false),
	)
	lg.Info("to test.log info msg")
	lg.Warnf("to test.log some %s log", "warn")
	lg.Error("to test.log some error log")
}
