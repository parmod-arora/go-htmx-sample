package loglib

import "github.com/sirupsen/logrus"

var (
	// Make sure ConciseLogger implements Logger
	_ Logger = (*ConciseLogger)(nil)
)

// NewConciseLogger returns an instance of ConciseLogger based on the provided logrus entry
func NewConciseLogger(logger *logrus.Entry) Logger {
	return &ConciseLogger{LogEntry: logger}
}

// ConciseLogger is an implementation of Logger interface
type ConciseLogger struct {
	LogEntry *logrus.Entry
}

// WithField adds a single field to the logger's data
func (v *ConciseLogger) WithField(key string, value interface{}) Logger {
	return &ConciseLogger{LogEntry: v.LogEntry.WithField(key, value)}
}

// WithFields adds a map of fields to the logger's data
func (v *ConciseLogger) WithFields(fields map[string]interface{}) Logger {
	return &ConciseLogger{LogEntry: v.LogEntry.WithFields(fields)}
}

// Debugf logs the message using debug level
func (v *ConciseLogger) Debugf(message string, args ...interface{}) {
	v.LogEntry.Debugf(message, args...)
}

// Infof logs the message using info level
func (v *ConciseLogger) Infof(message string, args ...interface{}) {
	v.LogEntry.Infof(message, args...)
}

// Warnf logs the message using warn level
func (v *ConciseLogger) Warnf(message string, args ...interface{}) {
	v.LogEntry.Warnf(message, args...)
}

// Errorf logs the message using error level
func (v *ConciseLogger) Errorf(message string, args ...interface{}) {
	v.LogEntry.Errorf(message, args...)
}

// Fatalf logs the message using fatal level
// Exits using os.Exit(1) after logging
func (v *ConciseLogger) Fatalf(message string, args ...interface{}) {
	v.LogEntry.Fatalf(message, args...)
}

// Panicf logs the message using panic level
// Triggers a panic after logging
func (v *ConciseLogger) Panicf(message string, args ...interface{}) {
	v.LogEntry.Panicf(message, args...)
}
