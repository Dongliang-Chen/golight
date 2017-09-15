// Copyright 2017 The Golight Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Logger provides a "structured" logging util and simple interface with JSON output.
// It has built-in context where Key/Value pairs may be added so that they don't have 
// to be part of the final logging call to improve logging performance.
// It also supports hierarchical loggers structure using the Sublogger function.
// It's built on top of zerolog, for its "Blazing fast" performance and "Zero Allocation"
package logger

import (
	"io"
	"time"
	zlog "github.com/rs/zerolog"
)

// LogObj or the composite of it that can be passed to Log.Print(), 
// and e.Interface() functions.
type LogObj map[string]interface{}

type LogLevel uint8

const (
	LogDebug = iota				// LogDebug defines debug log level.
	LogInfo						// LogInfo defines info log level.
	LogWarn						// LogWarn defines warn log level.
	LogError						// LogError defines error log level.
	LogFatal						// LogFatal defines fatal log level.
	LogPanic						// LogPanic defines panic log level.
	LogDisabled					// LogDisabled disables the logger.
)

type Logger struct {
	zlog.Logger
}

type Context struct {
	zlog.Context
}

// New creates a new Logger using the io.Writer argument.
// If withTimestamp is set to true, the logger will append the timestamp to the log message.
// On creation, the key, value store is empty.
func New(w io.Writer, withTimestamp bool) Logger {
	//zlog.TimeFieldFormat
	zlog.TimestampFieldName = "t"
	zlog.LevelFieldName = "l"
	zlog.MessageFieldName = "m"
	zlog.ErrorFieldName = "e"
	zlog.DurationFieldInteger = true
	zlog.DurationFieldUnit = time.Millisecond
	c := zlog.New(w).With()
	if withTimestamp { c.Timestamp() }

	return Logger{c.Logger()}
}

// Print sends a log event using debug level.
// Arguments are handled in the manner of fmt.Print.
// Note: this uses debug level only.
// func (l Logger)Print(i ...interface{}).

// Printf sends a log event using debug level.
// Arguments are handled in the manner of fmt.Printf.
// Note: this uses debug level only.
// func (l Logger) Printf(format string, i ...interface{}).

// With creates a logger Context.
// Additional context may be added to the Context.
// Once added, call Context.Logger() to get a sub logger with the context.
// func (l Logger) With() Context.

// Debug logs a new message with Debug level.
// Refer to logger_test for details of each usecases.
// func (l Logger) Debug() *Event.

// Info logs a new message with Info level.
// Refer to logger_test for details of each usecases.
// func (l Logger) Info()  *Event.

// Warn logs a new message with Warn level.
// Refer to logger_test for details of each usecases.
// func (l Logger) Warn() *Event.

// Error logs a new message with Error level.
// Refer to logger_test for details of each usecases.
// func (l Logger) Error() *Event.

// Fatal logs a new message with Fatal level.
// The os.Exit(1) function is called.
// Refer to logger_test for details of each usecases.
// func (l Logger) Fatal() *Event.

// Panic logs a new message with Panic level.
// The message is also sent to the panic function.
// Refer to logger_test for details of each usecases.
// func (l Logger) Panic() *Event.

// Level sets the level of the logger.
// Returns a new Sub Logger.
func (l Logger) Level(lvl LogLevel) Logger {
	return Logger{l.Logger.Level(zlog.Level(lvl))}
}


// SetGlobalLevel sets the global log level.
// Refer to logger_test for details of each usecases.
func SetGlobalLevel(level LogLevel) {
	zlog.SetGlobalLevel(zlog.Level(level))
}

// Enabled checks if a log is enabled.
// Refer to logger_test for details of each usecases.
// func (e *Event) Enabled() bool.

