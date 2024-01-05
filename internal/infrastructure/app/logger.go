package app

import (
	"errors"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	ErrInvalidLogFormat = errors.New("unknown log format, valid values are 'text' or 'json'")
	ErrInvalidLogLevel  = errors.New("unknown log level, valid values are 'trace', 'debug', 'info', 'warn', 'error', 'fatal' or 'panic'")
)

func ConfigureLogger(cfg LogConfig) error {
	if err := setLogLevel(cfg.Level); err != nil {
		return fmt.Errorf("failed to set log-level: %w", err)
	}

	if err := setLogFormat(cfg.Formatter); err != nil {
		return fmt.Errorf("failed to set log format: %w", err)
	}

	return nil
}

func setLogFormat(format string) error {
	switch strings.ToLower(format) {
	case "text":
		log.SetFormatter(&log.TextFormatter{})
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		return ErrInvalidLogFormat
	}

	return nil
}

func setLogLevel(levelName string) error {
	var logLevel log.Level

	switch strings.ToLower(levelName) {
	case "trace":
		logLevel = log.TraceLevel
	case "debug":
		logLevel = log.DebugLevel
	case "info":
		logLevel = log.InfoLevel
	case "warn":
		logLevel = log.WarnLevel
	case "error":
		logLevel = log.ErrorLevel
	case "fatal":
		logLevel = log.FatalLevel
	case "panic":
		logLevel = log.PanicLevel
	default:
		return ErrInvalidLogLevel
	}

	log.SetLevel(logLevel)
	return nil
}
