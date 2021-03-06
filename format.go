package log

import (
	"fmt"
	"time"
	"strings"
	"sort"
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

	outputMessage = outputMessage + "\t"
	keys := make([]string, len(tags))
	for key, _ := range tags {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		if key != "function" && len(key) > 0 {
			outputMessage += fmt.Sprintf("%s:%s,", key, tags[key])
		}
	}
	return strings.Trim(outputMessage, ",")
}
