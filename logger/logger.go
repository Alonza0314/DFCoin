package logger

import (
	"sync"
	"time"

	loggergo "github.com/Alonza0314/logger-go"
)

func getLogFile() string {
	return LOG_DB_PATH + time.Now().Format("2006-01-02") + ".log"
}

var (
	Log *Logger
)

type Logger struct {
	logger *loggergo.FileLogger
	mtx    sync.Mutex
}

func NewLogger() *Logger {
	return &Logger{
		logger: loggergo.NewFileLogger(getLogFile()),
		mtx:    sync.Mutex{},
	}
}

func (l *Logger) Error(tag, msg string) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	l.logger.Error(tag, msg)
	loggergo.Error(tag, msg)
}

func (l *Logger) Info(tag, msg string) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	l.logger.Info(tag, msg)
	loggergo.Info(tag, msg)
}

func (l *Logger) Debug(tag, msg string) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	l.logger.Debug(tag, msg)
	loggergo.Debug(tag, msg)
}

func (l *Logger) Warn(tag, msg string) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	l.logger.Warn(tag, msg)
	loggergo.Warn(tag, msg)
}
