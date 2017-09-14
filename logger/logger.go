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
// Log.AddKV() and Log.Debug...Panic() functions
type LogObj map[string]interface{}

// Level defines log levels.
type LogLevel uint8

const (
	// LogDebug defines debug log level.
	LogDebug LogLevel = iota
	// LogInfo defines info log level.
	LogInfo
	// LogWarn defines warn log level.
	LogWarn
	// LogError defines error log level.
	LogError
	// LogFatal defines fatal log level.
	LogFatal
	// LogPanic defines panic log level.
	LogPanic
	// LogDisabled disables the logger.
	LogDisabled
)

type Logger struct {
	z zlog.Logger
	c zlog.Context
}

// New creates a new Logger using the io.Writer argument
// If withTimestamp is set to true, the logger will append the timestamp to the log message
// On creation, the key, value store is empty
func New(w io.Writer, withTimestamp bool) *Logger {
	//zlog.TimeFieldFormat
	zlog.TimestampFieldName = "t"
	zlog.LevelFieldName = "l"
	zlog.MessageFieldName = "m"
	zlog.ErrorFieldName = "e"
	zlog.DurationFieldInteger = true
	zlog.DurationFieldUnit = time.Millisecond
	c := zlog.New(w).With()
	if withTimestamp { c.Timestamp() }

	return &Logger{z:c.Logger(), c:c}
}

// Print sends a log event using debug level
// Arguments are handled in the manner of fmt.Print.
// Note: this uses debug level only
func (l *Logger)Print(i ...interface{}) {
	l.z.Print(i...)
}

// Printf sends a log event using debug level
// Arguments are handled in the manner of fmt.Printf.
// Note: this uses debug level only
func (l *Logger) Printf(format string, i ...interface{}) {
	l.z.Printf(format, i...)
}

// Sublogger creates a sub-logger, which inherits all the property from its parent
func (l *Logger)Sublogger() *Logger {
	c := l.z.With()
	return &Logger{z:c.Logger(), c:c}
}

// AddKV adds the key/val pair to the logger's context and 
// Returns the same logger, making chained AddKV call possible
// Note: key/val does not need to be unique
/*
//  Calling:
	log.AddKV(k1,v1)
	log.AddKV(k2,v2)
	log.AddKV(k2,v3)
	is equavilent to call:
	log.AddKV(k1,v1).AddKV(k2,v2).AddKV(k2,v3)	
*/
// Example using LogObj
// log.AddKV("key", LogObj{"a":"aa","b":19,"d":19,"g":LogObj{"c":"c","d":[]int{1,2,3}}})	
// Refer to logger_test for details of each usecases
func (l *Logger) AddKV(key string, val interface{}) *Logger {
	//may just use interface, but json.Marchal uses reflection
	switch v:=val.(type) {
	case string:
		l.c = l.c.Str(key, v)
	case []string:
		l.c = l.c.Strs(key, v)
	case bool:
		l.c = l.c.Bool(key, v)
	case []bool:
		l.c = l.c.Bools(key, v)
	case int:
		l.c = l.c.Int(key, v)
	case []int:
		l.c = l.c.Ints(key, v)
	case int8:
		l.c = l.c.Int8(key, v)
	case []int8:
		l.c = l.c.Ints8(key, v)
	case int16:
		l.c = l.c.Int16(key, v)
	case []int16:
		l.c = l.c.Ints16(key, v)
	case int32:
		l.c = l.c.Int32(key, v)
	case []int32:
		l.c = l.c.Ints32(key, v)
	case int64:
		l.c = l.c.Int64(key, v)
	case []int64:
		l.c = l.c.Ints64(key, v)
	case uint:
		l.c = l.c.Uint(key, v)
	case []uint:
		l.c = l.c.Uints(key, v)
	case uint8:
		l.c = l.c.Uint8(key, v)
	case []uint8:
		l.c = l.c.Uints8(key, v)
	case uint16:
		l.c = l.c.Uint16(key, v)
	case []uint16:
		l.c = l.c.Uints16(key, v)
	case uint32:
		l.c = l.c.Uint32(key, v)
	case []uint32:
		l.c = l.c.Uints32(key, v)
	case uint64:
		l.c = l.c.Uint64(key, v)
	case []uint64:
		l.c = l.c.Uints64(key, v)
	case float32:
		l.c = l.c.Float32(key, v)
	case []float32:
		l.c = l.c.Floats32(key, v)
	case float64:
		l.c = l.c.Float64(key, v)
	case []float64:
		l.c = l.c.Floats64(key, v)
	case time.Time:
		l.c = l.c.Time(key, v)
	case []time.Time:
		l.c = l.c.Times(key, v)
	case time.Duration:
		l.c = l.c.Dur(key, v)	
	case []time.Duration:
		l.c = l.c.Durs(key, v)
	default:
		l.c = l.c.Interface(key, v)	
	}
	l.z = l.c.Logger()
	return l
}

func eventLog(e *zlog.Event, key string, i interface{}) {
	if !e.Enabled() {
		return
	}

	switch v:=i.(type) {
	case string:
		e.Str(key, v)
	case []string:
		e.Strs(key, v)
	case bool:
		e.Bool(key, v)
	case []bool:
		e.Bools(key, v)
	case int:
		e.Int(key, v)
	case []int:
		e.Ints(key, v)
	case int8:
		e.Int8(key, v)
	case []int8:
		e.Ints8(key, v)
	case int16:
		e.Int16(key, v)
	case []int16:
		e.Ints16(key, v)
	case int32:
		e.Int32(key, v)
	case []int32:
		e.Ints32(key, v)
	case int64:
		e.Int64(key, v)
	case []int64:
		e.Ints64(key, v)
	case uint:
		e.Uint(key, v)
	case []uint:
		e.Uints(key, v)
	case uint8:
		e.Uint8(key, v)
	case []uint8:
		e.Uints8(key, v)
	case uint16:
		e.Uint16(key, v)
	case []uint16:
		e.Uints16(key, v)
	case uint32:
		e.Uint32(key, v)
	case []uint32:
		e.Uints32(key, v)
	case uint64:
		e.Uint64(key, v)
	case []uint64:
		e.Uints64(key, v)
	case float32:
		e.Float32(key, v)
	case []float32:
		e.Floats32(key, v)
	case float64:
		e.Float64(key, v)
	case []float64:
		e.Floats64(key, v)
	case time.Time:
		e.Time(key, v)
	case []time.Time:
		e.Times(key, v)
	case time.Duration:
		e.Dur(key, v)
	case []time.Duration:
		e.Durs(key, v)
	default:
		e.Interface(key, v)
	}
	e.Msg("")
}

// Debug logs a new message with Debug level.
// Example using LogObj
// log.Debug("key", LogObj{"a":"aa","b":19,"d":19,"g":LogObj{"c":"c","d":[]int{1,2,3}}})	
// Refer to logger_test for details of each usecases
func (l *Logger) Debug(key string, i interface{}) {
	eventLog(l.z.Debug(), key, i)
}

// Info logs a new message with Info level.
// Example using LogObj
// log.Info("key", LogObj{"a":"aa","b":19,"d":19,"g":LogObj{"c":"c","d":[]int{1,2,3}}})	
// Refer to logger_test for details of each usecases
func (l *Logger) Info(key string, i interface{}) {
	eventLog(l.z.Info(), key, i)
}

// Warn logs a new message with Warn level.
// Example using LogObj
// log.Warn("key", LogObj{"a":"aa","b":19,"d":19,"g":LogObj{"c":"c","d":[]int{1,2,3}}})	
// Refer to logger_test for details of each usecases
func (l *Logger) Warn(key string, i interface{}) {
	eventLog(l.z.Warn(), key, i)
}

// Error logs a new message with Error level.
// Example using LogObj
// log.Error("key", LogObj{"a":"aa","b":19,"d":19,"g":LogObj{"c":"c","d":[]int{1,2,3}}})	
// Refer to logger_test for details of each usecases
func (l *Logger) Error(key string, i interface{}) {
	eventLog(l.z.Error(), key, i)
}

// Fatal logs a new message with Fatal level.
// The os.Exit(1) function is called.
// Example using LogObj
// log.Fatal("key", LogObj{"a":"aa","b":19,"d":19,"g":LogObj{"c":"c","d":[]int{1,2,3}}})	
// Refer to logger_test for details of each usecases
func (l *Logger) Fatal(key string, i interface{}) {
	eventLog(l.z.Fatal(), key, i)
}

// Panic logs a new message with Panic level.
// The message is also sent to the panic function.
// Example using LogObj
// log.Panic("key", LogObj{"a":"aa","b":19,"d":19,"g":LogObj{"c":"c","d":[]int{1,2,3}}})	
// Refer to logger_test for details of each usecases
func (l *Logger) Panic(key string, i interface{}) {
	eventLog(l.z.Panic(), key, i)
}

// SetLogLevel sets the global log level
// Refer to logger_test for details of each usecases
func (l *Logger)SetLogLevel(level LogLevel) {
	zlog.SetGlobalLevel(zlog.Level(level))
}

// Enabled checks if a log level is enabled
// Refer to logger_test for details of each usecases
func (l *Logger) Enabled(level LogLevel) bool {
	return l.z.WithLevel(zlog.Level(level)).Enabled()
}

