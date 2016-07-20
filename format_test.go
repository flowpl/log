package log

import (
	"testing"
	"encoding/json"
	"time"
	"strings"
)

type JsonLogMessage struct {
	Time string
	Level string
	Message string
	Tag1 string
	Something string
}

func TestJsonFormatterFormatsLevelAndMessageIntoValidJSON(t *testing.T) {
	resultString := JsonFormatter("LOG_LEVEL", "some message", map[string]string{}, "2006")
	resultJson := new(JsonLogMessage)
	err := json.Unmarshal([]byte(resultString), resultJson)

	if err != nil {
		t.Errorf("json.Unmarshal failed with error: %s", err.Error())
	}

	if resultJson.Message != "some message" {
		t.Errorf("expected message: \"%s\", actual: \"%s\"", "some message", resultJson.Message)
	}

	if resultJson.Level != "LOG_LEVEL" {
		t.Errorf("expected level: \"%s\", actual: \"%s\"", "LOG_LEVEL", resultJson.Level)
	}
}

func TestJsonFormatterFormatsDateInUTC(t *testing.T) {
	const DATE_FORMAT_STRING = "2006-01-02T15"
	resultString := JsonFormatter("LOG_LEVEL", "some message", map[string]string{}, DATE_FORMAT_STRING)
	resultJson := new(JsonLogMessage)
	err := json.Unmarshal([]byte(resultString), resultJson)

	if err != nil {
		t.Errorf("json.Unmarshal failed with error: %s", err.Error())
	}

	currentDateUTC := time.Now().UTC()
	if resultJson.Time != currentDateUTC.Format(DATE_FORMAT_STRING) {
		t.Errorf("expected time: \"%s\", actual: \"%s\"", currentDateUTC.Format(DATE_FORMAT_STRING), resultJson.Time)
	}
}

func TestJsonFormatterAppendsAllTagsToMessage(t *testing.T) {
	resultString := JsonFormatter("LOG_LEVEL", "some message", map[string]string{"tag1":"value1", "something":"anything"}, "2006")
	resultJson := new(JsonLogMessage)
	err := json.Unmarshal([]byte(resultString), resultJson)

	if err != nil {
		t.Errorf("json.Unmarshal failed with error: %s", err.Error())
	}

	if resultJson.Tag1 != "value1" || resultJson.Something != "anything" {
		t.Errorf("expected tags: \"%s\", actual: \"tag1:%s,something:%s\"", "tag1:value1,something:anything", resultJson.Tag1, resultJson.Something)
	}
}


func TestTextFormatterCreatesTheCorrectFormat(t *testing.T) {
	resultString := TextFormatter("LOG_LEVEL", "some message", map[string]string{"function":"function name", "something":"anything"}, "2006")
	resultParts := strings.Split(resultString, "\t")

	if len(resultParts) != 5 {
		t.Errorf("expected parts count: %d, actual: %d", 5, len(resultParts))
	}

	if resultParts[0] != time.Now().UTC().Format("2006") {
		t.Errorf("expected first message part to be timestamp, actual: \"%s\"", resultParts[0])
	}

	if resultParts[1] != "LOG_LEVEL" {
		t.Errorf("expected second message part to be timestamp, actual: \"%s\"", resultParts[1])
	}

	if resultParts[2] != "function name" {
		t.Errorf("expected third messaeg part to be the function name, actual: \"%s\"", resultParts[2])
	}

	if resultParts[3] != "some message" {
		t.Errorf("expected fourth message part to be the message, actual: \"%s\"", resultParts[3])
	}

	if resultParts[4] != "something:anything" {
		t.Errorf("expected fifth message part to be all tags but function, actual: \"%s\"", resultParts[4])
	}
}

func TestTextFormatterOutputsTimestampInUTC(t *testing.T) {
	const DATE_FORMAT_STRING = "2006-01-02T15"
	resultString := TextFormatter("LOG_LEVEL", "some message", map[string]string{"function":"function name", "something":"anything"}, DATE_FORMAT_STRING)
	resultParts := strings.Split(resultString, "\t")

	if resultParts[0] != time.Now().UTC().Format(DATE_FORMAT_STRING) {
		t.Errorf("expected time: \"%s\", actual: \"%s\"", time.Now().UTC().Format(DATE_FORMAT_STRING), resultParts[0])
	}
}

func TestTextFormatterSortsTagsAlphabetically(t *testing.T) {
	resultString := TextFormatter("L", "s", map[string]string{"function":"function name", "something":"anything", "before":"tag1"}, "2006")
	resultParts := strings.Split(resultString, "\t")

	if resultParts[4] != "before:tag1,something:anything" {
		t.Errorf("expected tags to be: \"%s\", actual: \"%s\"", "before:tag1,something:anything", resultParts[5])
	}
}
