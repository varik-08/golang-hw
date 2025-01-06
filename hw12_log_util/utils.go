package main

import (
	"regexp"
	"strings"
)

const (
	ErrorLevel   = "error"
	WarningLevel = "warning"
	InfoLevel    = "info"
	DebugLevel   = "debug"
)

func CheckLogLevel(level string) bool {
	switch level {
	case ErrorLevel, WarningLevel, InfoLevel, DebugLevel:
		return true
	default:
		return false
	}
}

func AnalyseLog(line string, level string) (string, string) {
	line = strings.TrimSpace(line)
	re := regexp.MustCompile(`^(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) - (\w+) - (.+)$`)

	matches := re.FindStringSubmatch(line)
	if matches == nil {
		return "", ""
	}

	date := matches[1]
	lineLevel := matches[2]
	message := matches[3]

	if level != lineLevel {
		return "", ""
	}

	return date, message
}
