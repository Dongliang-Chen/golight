// Copyright 2017 The Golight Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package queryparser

import (
	"net/http"
	. "github.com/dlmc/golight/decorator"
	"github.com/dlmc/golight/ctx"
	"net/url"
)

// Internal int key
var querymapKey = ctx.GetNextCtxMapKey()


// GetQueryValues returns a non-nil map containing all the valid query parameters found
// Prior to call GetQueryValues, the request will have to be decorated 
// by the decor created by requestquery.CreateDecor
func GetQueryValues(r *http.Request) url.Values {
	return ctx.GetCtxMap(r)[querymapKey].(url.Values)
}

// CreateDecor creates a decorator that parses the URL-encoded query string into 
// a map listing the values specified for each key. The map is stored in the request context
// for future use
// There are two wa
func CreateDecor() Decorator {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx.GetCtxMap(r)[querymapKey] = r.URL.Query()
			next.ServeHTTP(w, r)
		})
	}		
}
