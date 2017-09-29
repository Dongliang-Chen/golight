// Copyright 2017 The Golight Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logging

import (
	"net/http"
	. "github.com/dlmc/golight/decorator"
	"github.com/dlmc/golight/ctx"
	//"github.com/dlmc/golight/logger"
	log "github.com/rs/zerolog"
)

// Internal int key
var loggingKey = ctx.GetNextCtxMapKey()


// GetLogger returns the Logger in the request Context
// Prior to call GetLogger, the request will have to be decorated 
// by the decor created by logger.CreateDecor
func GetLogger(r *http.Request) log.Context {
	return *ctx.GetCtxMap(r)[loggingKey].(*log.Context)
}

// CreateDecor creates a decorator that adds the passed in Logger into the request context map
// for future use
func CreateDecor(lc *log.Context) Decorator {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx.GetCtxMap(r)[loggingKey] = lc
			next.ServeHTTP(w, r)
		})
	}		
}
