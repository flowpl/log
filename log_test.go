package log

import (
	"testing"
)

const FORMATTED_MESSAGE = "message returned from dummyFormatter"

type DummyFormatOutput struct {
	level string
	message string
	tags map[string]string
	dateFormat string
	outputMessage string
}
func (dfo *DummyFormatOutput) createDummyFormatter() func(string, string, map[string]string, string) string{
	return func(level string, message string, tags map[string]string, dateFormat string) string {
		dfo.level = level
		dfo.message = message
		dfo.tags = tags
		dfo.dateFormat = dateFormat
		return FORMATTED_MESSAGE
	}
}
func (dfo *DummyFormatOutput) createDummyOutput() func(string) {
	return func (message string) {
		dfo.outputMessage = message
	}
}

func TestNewLoggerCreatesANewLogThatUsesGivenConfig(t *testing.T) {
	dfo := new(DummyFormatOutput)
	config := &Config{
		LEVEL_DEBUG,
		dfo.createDummyFormatter(),
		dfo.createDummyOutput(),
		"log_test",
		"2006",
		map[string]string{"tag1":"value1"},
	}
	result := NewLogger(config)

	result.Debug("message", map[string]string{})
	if dfo.level != LEVEL_DEBUG {
		t.Errorf("expected level to be \"%s\", actual: \"%s\"", LEVEL_DEBUG, dfo.level)
	}

	if dfo.message != "message" {
		t.Errorf("expected message to be \"%s\", actual: \"%s\"", "message", dfo.message)
	}

	if dfo.tags["tag1"] != "value1" {
		t.Errorf("expected tag \"tag1\" to be \"value1\", actual: \"%s\"", dfo.tags["tag1"])
	}

	if dfo.tags["function"] != "main" {
		t.Errorf("expected tag \"function\" to be \"main\", actual: \"%s\"", dfo.tags["function"])
	}

	if dfo.tags["program"] != "log_test" {
		t.Errorf("expected tag \"program\" to be \"log_test\", actual: \"%s\"", dfo.tags["program"])
	}

	if len(dfo.tags) != 3 {
		t.Errorf("expected exactly %d tags, actual: %d", 3, len(dfo.tags))
	}

	if dfo.dateFormat != "2006" {
		t.Errorf("expected dateFormat to be \"%s\", actual: \"%s\"", "2006", dfo.dateFormat)
	}

	if dfo.outputMessage != FORMATTED_MESSAGE {
		t.Errorf("expected outputMessage to be \"%s\", actual: \"%s\"", FORMATTED_MESSAGE, dfo.outputMessage)
	}
}

func TestNewLoggerShouldAcceptNilAsTagsConfig(t *testing.T) {
	dfo := new(DummyFormatOutput)
	config := &Config{
		LEVEL_DEBUG,
		dfo.createDummyFormatter(),
		dfo.createDummyOutput(),
		"log_test",
		"2006",
		nil,
	}
	result := NewLogger(config)
	result.Debug("message", map[string]string{})

	if len(dfo.tags) != 2 {
		t.Errorf("expected exactly 2 tags, actual: %d", len(dfo.tags))
	}
}

func TestLog_InfoShouldOutputMessagesIfLevelIsDebug(t *testing.T) {
	dfo := new(DummyFormatOutput)
	config := &Config{
		LEVEL_DEBUG,
		dfo.createDummyFormatter(),
		dfo.createDummyOutput(),
		"log_test",
		"2006",
		nil,
	}
	result := NewLogger(config)
	result.Info("message", map[string]string{})

	if dfo.message != "message" {
		t.Errorf("expected message to be \"%s\", actual: \"%s\"", "message", dfo.message)
	}

	if dfo.outputMessage != FORMATTED_MESSAGE {
		t.Errorf("expected outputMessage to be \"%s\", actual \"%s\"", FORMATTED_MESSAGE, dfo.outputMessage)
	}
}

func TestLog_InfoShouldOutputMessagesIfLevelIsInfo(t *testing.T) {
	dfo := new(DummyFormatOutput)
	config := &Config{
		LEVEL_INFO,
		dfo.createDummyFormatter(),
		dfo.createDummyOutput(),
		"log_test",
		"2006",
		nil,
	}
	result := NewLogger(config)
	result.Info("message", map[string]string{})

	if dfo.message != "message" {
		t.Errorf("expected message to be \"%s\", actual: \"%s\"", "message", dfo.message)
	}

	if dfo.outputMessage != FORMATTED_MESSAGE {
		t.Errorf("expected outputMessage to be \"%s\", actual \"%s\"", FORMATTED_MESSAGE, dfo.outputMessage)
	}
}

func TestLog_InfoShouldReturnInvalidContext(t *testing.T) {
	dfo := new(DummyFormatOutput)
	config := &Config{
		LEVEL_INFO,
		dfo.createDummyFormatter(),
		dfo.createDummyOutput(),
		"log_test",
		"2006",
		nil,
	}
	result := NewLogger(config)
	err := result.Info("message", "")

	if _, ok := err.(*InvalidContext); !ok {
		t.Errorf("expected InvalidContext error")
	}
}

func TestLog_DebugShouldOutputMessagesIfLevelIsDebug(t *testing.T) {
	dfo := new(DummyFormatOutput)
	config := &Config{
		LEVEL_DEBUG,
		dfo.createDummyFormatter(),
		dfo.createDummyOutput(),
		"log_test",
		"2006",
		nil,
	}
	result := NewLogger(config)
	result.Debug("message", map[string]string{})

	if dfo.message != "message" {
		t.Errorf("expected message to be \"%s\", actual: \"%s\"", "message", dfo.message)
	}

	if dfo.outputMessage != FORMATTED_MESSAGE {
		t.Errorf("expected outputMessage to be \"%s\", actual \"%s\"", FORMATTED_MESSAGE, dfo.outputMessage)
	}
}

func TestLog_DebugShouldNotOutputMessagesIfLevelIsInfo(t *testing.T) {
	dfo := new(DummyFormatOutput)
	config := &Config{
		LEVEL_INFO,
		dfo.createDummyFormatter(),
		dfo.createDummyOutput(),
		"log_test",
		"2006",
		nil,
	}
	result := NewLogger(config)
	result.Debug("message", map[string]string{})

	if dfo.message != "" {
		t.Errorf("expected message to be \"%s\", actual: \"%s\"", "message", dfo.message)
	}

	if dfo.outputMessage != "" {
		t.Errorf("expected outputMessage to be [empty string], actual \"%s\"", dfo.outputMessage)
	}
}

func TestLog_DebugShouldReturnInvalidContext(t *testing.T) {
	dfo := new(DummyFormatOutput)
	config := &Config{
		LEVEL_DEBUG,
		dfo.createDummyFormatter(),
		dfo.createDummyOutput(),
		"log_test",
		"2006",
		nil,
	}
	result := NewLogger(config)
	err := result.Debug("message", "")

	if _, ok := err.(*InvalidContext); !ok {
		t.Errorf("expected InvalidContext error")
	}
}

func TestLog_ChildLoggerShouldCreateANewLoggerInstance(t *testing.T) {
	dfo := new(DummyFormatOutput)
	config := &Config{
		LEVEL_INFO,
		dfo.createDummyFormatter(),
		dfo.createDummyOutput(),
		"log_test",
		"2006",
		nil,
	}
	parentLogger := NewLogger(config)
	result, _ := parentLogger.ChildLogger("child_test", map[string]string{})

	result.Info("message", map[string]string{})
	if parentLogger == result {
		t.Errorf("expected logger and childLogger to be different instances, but they are the same")
	}
}

func TestLog_ChildLoggerShouldMergeItsContextWithParentLoggersTags(t *testing.T) {
	dfo := new(DummyFormatOutput)
	config := &Config{
		LEVEL_INFO,
		dfo.createDummyFormatter(),
		dfo.createDummyOutput(),
		"log_test",
		"2006",
		nil,
	}
	parentLogger := NewLogger(config)
	result, _ := parentLogger.ChildLogger("child_test", map[string]string{"child_tag":"value"})

	result.Info("message", map[string]string{})
	if len(dfo.tags) != 3 {
		t.Errorf("expected exactly 3 tags, actual: \"%d\"", len(dfo.tags))
	}

	if dfo.tags["child_tag"] != "value" {
		t.Errorf("expected tag child_tag to be \"%s\", actual: \"%s\"", "message", dfo.tags["child_tag"])
	}
}

func TestLog_ChildLoggerShouldAllowNilAsContext(t *testing.T) {
	dfo := new(DummyFormatOutput)
	config := &Config{
		LEVEL_INFO,
		dfo.createDummyFormatter(),
		dfo.createDummyOutput(),
		"log_test",
		"2006",
		nil,
	}
	parentLogger := NewLogger(config)
	result, _ := parentLogger.ChildLogger("child_test", nil)

	result.Info("message", map[string]string{})

	if len(dfo.tags) != 2 {
		t.Errorf("expected exactls 2 tags, actual: \"%d\"", len(dfo.tags))
	}
}

type arbitraryStruct struct {
	ExportedStructTag string
	structTag string
}
func TestLog_ChildLoggerShouldAcceptAnArbitraryStructAsContextAndMergeExportedFieldsIntoTags(t *testing.T) {
	dfo := new(DummyFormatOutput)
	config := &Config{
		LEVEL_INFO,
		dfo.createDummyFormatter(),
		dfo.createDummyOutput(),
		"log_test",
		"2006",
		nil,
	}
	parentLogger := NewLogger(config)
	result, _ := parentLogger.ChildLogger("child_test", arbitraryStruct{"exportedValue", "value"})

	result.Info("message", map[string]string{})

	if len(dfo.tags) != 3 {
		t.Errorf("expected exactly 3 tags, actual: \"%d\"", len(dfo.tags))
	}
	if dfo.tags["ExportedStructTag"] != "exportedValue" {
		t.Errorf("expected tag ExportedStructTag to be \"exportedValue\", actual: \"%s\"", dfo.tags["ExportedStructTag"])
	}
}

func TestLog_ChildLoggerShouldAcceptAPtrToArbitraryStructAsContextAndMergeExportedFieldsIntoTags(t *testing.T) {
	dfo := new(DummyFormatOutput)
	config := &Config{
		LEVEL_INFO,
		dfo.createDummyFormatter(),
		dfo.createDummyOutput(),
		"log_test",
		"2006",
		nil,
	}
	parentLogger := NewLogger(config)
	result, _ := parentLogger.ChildLogger("child_test", &arbitraryStruct{"exportedValue", "value"})

	result.Info("message", map[string]string{})

	if len(dfo.tags) != 3 {
		t.Errorf("expected exactly 3 tags, actual: \"%d\"", len(dfo.tags))
	}
	if dfo.tags["ExportedStructTag"] != "exportedValue" {
		t.Errorf("expected tag ExportedStructTag to be \"exportedValue\", actual: \"%s\"", dfo.tags["ExportedStructTag"])
	}
}


type arbitraryStringer struct {
	ExportedStructTag string
	structTag string
}
func (as arbitraryStringer) String() string {
	return "result from stringer"
}
func TestLog_ChildLoggerShouldUseTheStringerInterfaceOfContextIfPresent(t *testing.T) {
	dfo := new(DummyFormatOutput)
	config := &Config{
		LEVEL_INFO,
		dfo.createDummyFormatter(),
		dfo.createDummyOutput(),
		"log_test",
		"2006",
		nil,
	}
	parentLogger := NewLogger(config)
	result, _ := parentLogger.ChildLogger("child_test", arbitraryStringer{"exportedValue", "value"})

	result.Info("message", map[string]string{})
	if len(dfo.tags) != 3 {
		t.Errorf("expected exactly 3 tags, actual: \"%s\"", len(dfo.tags))
	}

	if dfo.tags["context"] != "result from stringer" {
		t.Errorf("expected tag context to be \"%s\", actual: \"%s\"", "result from stringer", dfo.tags["context"])
	}
}


type arbitraryError struct {
	ExportedStructTag string
	structTag string
}
func (as arbitraryError) Error() string {
	return "result from error"
}
func TestLog_ChildLoggerShouldUseTheErrorInterfaceOfContextIfPresent(t *testing.T) {
	dfo := new(DummyFormatOutput)
	config := &Config{
		LEVEL_INFO,
		dfo.createDummyFormatter(),
		dfo.createDummyOutput(),
		"log_test",
		"2006",
		nil,
	}
	parentLogger := NewLogger(config)
	result, _ := parentLogger.ChildLogger("child_test", arbitraryError{"exportedValue", "value"})

	result.Info("message", map[string]string{})
	if len(dfo.tags) != 3 {
		t.Errorf("expected exactly 3 tags, actual: \"%s\"", len(dfo.tags))
	}

	if dfo.tags["error"] != "result from error" {
		t.Errorf("expected tag context to be \"%s\", actual: \"%s\"", "result from error", dfo.tags["error"])
	}
}


type arbitraryErrorStringer struct {
	ExportedStructTag string
	structTag string
}
func (as arbitraryErrorStringer) Error() string {
	return "result from error"
}
func (as arbitraryErrorStringer) String() string {
	return "result from stringer"
}
func TestLog_ChildLoggerShouldPreferErrorOverStringer(t *testing.T) {
	dfo := new(DummyFormatOutput)
	config := &Config{
		LEVEL_INFO,
		dfo.createDummyFormatter(),
		dfo.createDummyOutput(),
		"log_test",
		"2006",
		nil,
	}
	parentLogger := NewLogger(config)
	result, _ := parentLogger.ChildLogger("child_test", arbitraryErrorStringer{"exportedValue", "value"})

	result.Info("message", map[string]string{})
	if len(dfo.tags) != 3 {
		t.Errorf("expected exactly 3 tags, actual: \"%s\"", len(dfo.tags))
	}

	if dfo.tags["error"] != "result from error" {
		t.Errorf("expected tag context to be \"%s\", actual: \"%s\"", "result from error", dfo.tags["error"])
	}
}


func TestLog_ChildLoggerShouldReturnInvalidContextOnInvalidContext(t *testing.T) {
	dfo := new(DummyFormatOutput)
	config := &Config{
		LEVEL_INFO,
		dfo.createDummyFormatter(),
		dfo.createDummyOutput(),
		"log_test",
		"2006",
		nil,
	}
	parentLogger := NewLogger(config)
	_, err := parentLogger.ChildLogger("child_test", "")

	if _, ok := err.(*InvalidContext); !ok {
		t.Errorf("expected InvalidContext error")
	}
}
