// Copyright 2017 The Golight Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logging

import (
	"github.com/dlmc/golight/decorator"
	"github.com/dlmc/golight/ghttp"
	log "github.com/rs/zerolog"
	"io"
	"time"
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

func init() {
	//zlog.TimeFieldFormat
	log.TimestampFieldName = "t"
	log.LevelFieldName = "l"
	log.MessageFieldName = "m"
	log.ErrorFieldName = "e"
	log.DurationFieldInteger = true
	log.DurationFieldUnit = time.Millisecond
}

func NewLogger(w io.Writer) log.Logger {
	return log.New(w)
}


func NewContext(w io.Writer) log.Context {
	return log.New(w).With()
}

func NewContextWithTimestamp(w io.Writer) log.Context {
	return log.New(w).With().Timestamp()
}

func NewLoggerWithTimestamp(w io.Writer) log.Logger {
	return log.New(w).With().Timestamp().Logger()
}


// SetGlobalLevel sets the global log level.
// Refer to logger_test for details of each usecases.
func SetGlobalLevel(level LogLevel) {
	log.SetGlobalLevel(log.Level(level))
}


// Internal int key
//var loggingKey = ghttp.GetNextCtxKey()


// GetLogger returns the Logger in the request Context
// Prior to call GetLogger, the request will have to be decorated 
// by the decor created by logger.CreateDecor
//func GetLogger(c ghttp.Ctx) log.Context {
//	return c.Value(loggingKey).(log.Context)
//}

// CreateDecor creates a decorator that adds the passed in Logger into the request context map
// for future use
func CreateDecor(lc log.Context) decorator.Decorator {
	return func(next ghttp.Handler) ghttp.Handler {
		return ghttp.HandlerFunc(func(c ghttp.Ctx, h *ghttp.Http) ghttp.Ctx{
			//c = ghttp.ChildCtx(c, loggingKey, lc)
			h.Log = lc
			return next.ServeHTTPWithCtx(c, h)
		})
	}		
}
