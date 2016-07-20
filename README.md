# log

[![travis build](https://travis-ci.org/flowpl/log.svg)](https://travis-ci.org/flowpl/log)
[![Coverage Status](https://coveralls.io/repos/github/flowpl/log/badge.svg?branch=master)](https://coveralls.io/github/flowpl/log?branch=master)

A minimal logging library for the go language, inspired by this blog post:
http://dave.cheney.net/2015/11/05/lets-talk-about-logging

## Philosophy
- there are informational messages that you want to see while the program is running in production
- there are debug messages that you want to see while debugging

Consequently there are just two log levels: INFO and DEBUG, which are used for output control only. 
When thinking: "I need additional WARNING, ERROR and CRITICAL levels", consider this:
Errors are supposed to be handled... or not. 
When handling the error, the developer is expected to bring the program back into a working state without loosing any data.
When not handling the error, the program should not do anything anymore, especially not have side effects, and exit. 
In either case the developer has to decide if she wants to see that the error happened while the program is running in production (INFO)
or just during development (DEBUG).
For a more detailed explanation, read Dave Cheney's post linked above. 

## Features
- simplified log levels that are easy to reason about
- pluggable output formatters (plain text and JSON are currently supported)
  - the plain text formatter outputs lines that can easily be processed with standard command line tools
  - use the error and fmt.Stringer interfaces to serialize context tag objects
- pluggable output handlers (stdout and stderr are currently supported)
- cascading context handling using child loggers and tags

## Installation

```bash
go get github.com/flowpl/log
```

or using glide

```bash
glide get github.com/flowpl/log
```

## Examples

### create a logger instance

```go
func setupLogger() {
    logConfig := new(log.Config)
    logConfig.Level = log.LEVEL_INFO
    logConfig.Formatter = log.TextFormatter
    logConfig.Output = log.StdOutOutput
    logConfig.Program = "log_test"
    logConfig.DateFormat = "2006-01-02T15:04:05.000000"
    logConfig.Tags = map[string]string{"context1":"value1"}
    return log.NewLogger(logConfig)
}
```

### write log messages

```go
package main

import github.com/flowpl/log

func main() {
    logger := setupLogger()
    logger.Info("my first log message", map[string]string{"additional_tag":"value"})
}
```

Output to stdout:

```
2016-07-14T13:09:51.678678 INFO    main    my first log message    additional_tag:value,program:log_test,function:main
```

### manage contexts

```go
package main

import github.com/flowpl/log

func main() {
    logger := setupLogger()
    anotherFunction("some parameter", logger)
}

func anotherFunction(parameter1 string, parentLogger log.Log) {
    logger := parentLogger.ChildLogger("anotherFunction", map[string]string{"functionContext":"functionValue"})
    logger.Debug("message", map[string]string{"logContext":"logValue"})
}

```

Output to stdout:

```
2016-07-14T13:09:51.678678 INFO    main    message    functionContext:functionValue,logContext:logValue,program:log_test,function:main
```

## Contributing

Create github issues for feature requests and bug reports.
If you want to contribute code, fork the repo and create a pull request on github.

## License
MIT, see LICENSE
