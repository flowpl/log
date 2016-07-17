package log

import (
	"fmt"
	"time"
	"strings"
)

func JsonFormatter(level string, message string, tags map[string]string, dateFormat string) string {
	outputMessage := fmt.Sprintf(
		"{\"time\":\"%s\",\"level\":\"%s\",\"message\":\"%s\"",
		time.Now().UTC().Format(dateFormat),
		level,
		message,
	)
	for name, value := range tags {
		outputMessage += fmt.Sprintf(",\"%s\":\"%s\"", name, value)
	}
	return outputMessage + "}"
}

func TextFormatter(level string, message string, tags map[string]string, dateFormat string) string {
	outputMessage := fmt.Sprintf(
		"%s\t%s\t%s\t%s",
		time.Now().UTC().Format(dateFormat),
		level,
		tags["function"],
		message,
	)

	for name, value := range tags {
		if name != "function" {
			outputMessage += fmt.Sprintf("%s:%s,", name, value)
		}
	}
	return strings.Trim(outputMessage, ",")
}
