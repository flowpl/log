package log

import (
	"reflect"
	"fmt"
)

const TIME_FORMAT = "2006-01-02T15:04:05.000000"
const LEVEL_DEBUG = "DEBUG"
const LEVEL_INFO  = "INFO"

type Config struct {
	Level string
	Formatter func(level string, message string, tags map[string]string, dateFormat string) string
	Output func(formattedMessage string)
	Program string
	DateFormat string
	Tags map[string]string
}

type Log struct {
	config *Config
}
type LogFormattingFailed string
func (err LogFormattingFailed) String() string {
	return "LogFormattingFailed"
}

func (log Log) Info(message string, tags interface{}) {
	log.config.Output(log.config.Formatter(LEVEL_INFO, message, mergeTags(log.config.Tags, tags), log.config.DateFormat))
}

func (log Log) Debug(message string, tags interface{}) {
	if log.config.Level == LEVEL_DEBUG {
		log.config.Output(log.config.Formatter(LEVEL_DEBUG, message, mergeTags(log.config.Tags, tags), log.config.DateFormat))
	}
}

func (log Log) ChildLogger(function string, context interface{}) *Log {
	childConfig := new(Config)
	childConfig.Level = log.config.Level
	childConfig.Formatter = log.config.Formatter
	childConfig.Output = log.config.Output
	childConfig.DateFormat = log.config.DateFormat
	childConfig.Program = log.config.Program
	childConfig.Tags = log.config.Tags
	childConfig.Tags["function"] = function
	childConfig.Tags = mergeTags(log.config.Tags, context)
	return NewLogger(childConfig)

}

func NewLogger(config *Config) *Log {
	if config.Tags == nil {
		config.Tags = map[string]string{}
	}
	config.Tags["program"] = config.Program
	config.Tags["function"] = "main"
	logger := new(Log)
	logger.config = config
	return logger
}

func mergeTags(tags map[string]string, context interface{}) map[string]string {
	outputTags := map[string]string{}
	for name, value := range tags {
		outputTags[name] = value
	}

	if context == nil {
		return outputTags
	}

	switch aTags := context.(type) {
	case error:
		outputTags["error"] = aTags.Error()
		return outputTags
	case fmt.Stringer:
		outputTags["context"] = aTags.String()
		return outputTags
	}

	reflectedValue := reflect.ValueOf(context)
	for {
		if reflectedValue.Kind() == reflect.Ptr {
			reflectedValue = reflectedValue.Elem()
		} else {
			break
		}
	}

	if reflectedValue.Kind() == reflect.Map {
		for _, name := range reflectedValue.MapKeys() {
			outputTags[name.String()] = reflectedValue.MapIndex(name).String()
		}
	} else if reflectedValue.Kind() == reflect.Struct {
		reflectedContext := reflect.TypeOf(context)
		for {
			if reflectedContext.Kind() == reflect.Ptr {
				reflectedContext = reflectedContext.Elem()
			} else {
				break
			}
		}

		for i := 0; i < reflectedContext.NumField(); i++ {
			currentField := reflectedContext.Field(i)
			if currentField.PkgPath == "" {  // merge only exported fields
				outputTags[currentField.Name] = reflectedValue.FieldByName(currentField.Name).String()
			}
		}
	} else {
		panic("invalid type for context. Must be map, struct, ptr(map) or ptr(struct)")
	}

	return outputTags
}
