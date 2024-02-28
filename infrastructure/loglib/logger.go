package loglib

import (
	"parmod/go-htmx-sample/infrastructure/appenv"

	"github.com/sirupsen/logrus"
)

var baseLogger = logrus.WithField("app", appenv.GetWithDefault("APPNAME", "APP")).WithField("env", appenv.GetWithDefault("APPENV", "local"))
var defaultConciseLogger = NewConciseLogger(baseLogger)

type Formatter string

const (
	FormatterText      Formatter = "text"
	FormatterTextColor Formatter = "text_color"
	FormatterJSON      Formatter = "json"
)

func init() {
	// Set default log level
	SetLevel(appenv.GetWithDefault("LOG_LEVEL", "INFO"))
	// Structured logging
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

// SetLevel sets log level
func SetLevel(levelString string) {
	if level, err := logrus.ParseLevel(levelString); err == nil {
		logrus.SetLevel(level)
	}
}

// DefaultConciseLogger returns the default concise logger
func DefaultConciseLogger() Logger {
	return defaultConciseLogger
}

type Logger interface {
	WithField(key string, value interface{}) Logger
	WithFields(map[string]interface{}) Logger
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Fatalf(string, ...interface{})
	Panicf(string, ...interface{})
}
