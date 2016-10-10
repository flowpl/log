package fakes

import (
	"math"
	"io"
	"github.com/flowpl/log"
)

type FakeLogger struct {}
func (fl *FakeLogger) Info(message string, context interface{}) error { return nil }
func (fl *FakeLogger) Debug(message string, context interface{}) error { return nil }
func (fl *FakeLogger) ChildLogger(function string, context interface{}) (log.Logger, error) { return fl, nil }
