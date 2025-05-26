package utils

import (
    "fmt"
    "log"
    "os"
    "time"
)

// LogLevel defines supported log levels.
type LogLevel string

const (
    LogInfo    LogLevel = "INFO"
    LogWarn    LogLevel = "WARN"
    LogError   LogLevel = "ERROR"
    LogDebug   LogLevel = "DEBUG"
    LogCrit    LogLevel = "CRITICAL"
)

// Logger provides namespaced logging for systemd-compatible environments.
type Logger struct {
    namespace string
    subject   string
    logger    *log.Logger
}

// NewLogger returns a new logger for a given namespace and subject.
// These appear as LOG_SUBJECT and LOG_NAMESPACE in journald.
func NewLogger(namespace, subject string) *Logger {
    prefix := fmt.Sprintf("LOG_NAMESPACE=%s LOG_SUBJECT=%s", namespace, subject)
    return &Logger{
        namespace: namespace,
        subject:   subject,
        logger:    log.New(os.Stdout, prefix+" ", 0),
    }
}

// logf outputs a log message with timestamp and level, systemd-readable.
func (l *Logger) logf(level LogLevel, format string, args ...interface{}) {
    timestamp := time.Now().Format(time.RFC3339)
    msg := fmt.Sprintf(format, args...)
    logLine := fmt.Sprintf("LEVEL=%s TIMESTAMP=%s MESSAGE=%q", level, timestamp, msg)
    l.logger.Println(logLine)
}

func (l *Logger) Infof(format string, args ...interface{}) {
    l.logf(LogInfo, format, args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
    l.logf(LogDebug, format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
    l.logf(LogWarn, format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
    l.logf(LogError, format, args...)
}

func (l *Logger) Criticalf(format string, args ...interface{}) {
    l.logf(LogCrit, format, args...)
}
