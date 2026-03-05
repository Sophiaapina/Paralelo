package main

import (
	"log"
	"os"
	"time"
)

type AppLogger struct {
	file   *os.File
	logger *log.Logger
}

func InitLogger(path string) (*AppLogger, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	l := log.New(f, "", 0)
	return &AppLogger{file: f, logger: l}, nil
}

func (a *AppLogger) Close() error {
	if a == nil || a.file == nil {
		return nil
	}
	return a.file.Close()
}

func (a *AppLogger) Infof(format string, args ...any) {
	a.write("INFO", format, args...)
}

func (a *AppLogger) Errorf(format string, args ...any) {
	a.write("ERROR", format, args...)
}

func (a *AppLogger) write(level, format string, args ...any) {
	if a == nil || a.logger == nil {
		return
	}
	ts := time.Now().Format(time.RFC3339)
	a.logger.Printf(ts+" ["+level+"] "+format, args...)
}
