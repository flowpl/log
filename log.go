package log

import (
	"reflect"
)

const TIME_FORMAT = "2006-01-02T15:04:05.000000"
const LEVEL_DEBUG = "DEBUG"
const LEVEL_INFO  = "INFO"

type Config struct {
	Level string
	Formatter func(level string, message string, tags map[string]string) string
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
	log.config.Output(log.config.Formatter(LEVEL_INFO, message, mergeTags(log.config.Tags, tags)))
}

func (log Log) Debug(message string, tags interface{}) {
	if log.config.Level == LEVEL_DEBUG {
		log.config.Output(log.config.Formatter(LEVEL_DEBUG, message, mergeTags(log.config.Tags, tags)))
	}
}

func (log Log) ChildLogger(function string, additionalTags map[string]string) *Log {
	childConfig := new(Config)
	childConfig.Level = log.config.Level
	childConfig.Formatter = log.config.Formatter
	childConfig.Output = log.config.Output
	childConfig.DateFormat = log.config.DateFormat
	childConfig.Program = log.config.Program
	childConfig.Tags = log.config.Tags
	childConfig.Tags["function"] = function
	for name, value := range additionalTags {
		childConfig.Tags[name] = value
	}
	return NewLogger(childConfig)

}

func NewLogger(config *Config) *Log {
	config.Tags["program"] = config.Program
	config.Tags["function"] = "main"
	logger := new(Log)
	logger.config = config
	return logger
}

func mergeTags(tags map[string]string, additionalTags interface{}) map[string]string {
	outputTags := map[string]string{}
	for name, value := range tags {
		outputTags[name] = value
	}

	reflectedContext := reflect.TypeOf(&additionalTags).Elem()
	reflectedValue := reflect.ValueOf(&additionalTags).Elem()
	if reflectedValue.Kind() == reflect.Map {
		for _, name := range reflectedValue.MapKeys() {
			outputTags[name] = reflectedValue.MapIndex(name).String()
		}
	} else if reflectedValue.Kind() == reflect.Struct {
		for i := 0; i < reflectedContext.NumField(); i++ {
			currentField := reflectedContext.Field(i)
			outputTags[currentField.Name] = reflectedValue.FieldByName(currentField.Name).String()
		}
	}

	return outputTags
}
