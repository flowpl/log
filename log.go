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
	ProgramName string
	FunctionName string
	DateFormat string
	Tags map[string]string
}

type LogFormattingFailed string
func (err LogFormattingFailed) String() string {
	return "LogFormattingFailed"
}

type InvalidContext string
func (err InvalidContext) Error() string {
	return "invalid type for context. Must be map, struct, ptr(map) or ptr(struct)"
}

type Logger interface {
	Info(string, interface{}) error
	Debug(string, interface{}) error
	ChildLogger(string, interface{}) (Logger, error)
}

type Log struct {
	config *Config
}
func (log Log) Info(message string, tags interface{}) error {
	mergedTags, err := mergeTags(log.config.Tags, tags)
	if err != nil {
		return err
	}
	log.config.Output(log.config.Formatter(LEVEL_INFO, message, mergedTags, log.config.DateFormat))
	return nil
}

func (log Log) Debug(message string, tags interface{}) error {
	if log.config.Level == LEVEL_DEBUG {
		mergedTags, err := mergeTags(log.config.Tags, tags)
		if err != nil {
			return err
		}
		log.config.Output(log.config.Formatter(LEVEL_DEBUG, message, mergedTags, log.config.DateFormat))
	}
	return nil
}

func (log Log) ChildLogger(functionName string, context interface{}) (Logger, error) {
	childConfig := new(Config)
	childConfig.Level = log.config.Level
	childConfig.Formatter = log.config.Formatter
	childConfig.Output = log.config.Output
	childConfig.DateFormat = log.config.DateFormat
	childConfig.ProgramName = log.config.ProgramName
	childConfig.Tags = make(map[string]string, len(log.config.Tags))
	for k, v := range log.config.Tags {
		childConfig.Tags[k] = v
	}
	childConfig.FunctionName = functionName
	mergedTags, err := mergeTags(childConfig.Tags, context)
	if err != nil {
		return nil, err
	}
	childConfig.Tags = mergedTags
	return NewLogger(childConfig), nil
}

func NewLogger(config *Config) Logger {
	if config.Tags == nil {
		config.Tags = map[string]string{}
	}
	config.Tags["program"] = config.ProgramName
	config.Tags["function"] = config.FunctionName
	logger := new(Log)
	logger.config = config
	return logger
}

func mergeTags(tags map[string]string, context interface{}) (map[string]string, error) {
	outputTags := map[string]string{}
	for name, value := range tags {
		outputTags[name] = value
	}

	if context == nil {
		return outputTags, nil
	}

	switch aTags := context.(type) {
	case error:
		outputTags["error"] = aTags.Error()
		return outputTags, nil
	case fmt.Stringer:
		outputTags["context"] = aTags.String()
		return outputTags, nil
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
		return nil, new(InvalidContext)
	}

	return outputTags, nil
}
