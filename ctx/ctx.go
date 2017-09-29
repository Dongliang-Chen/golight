// Copyright 2017 The Golight Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ctx

import (
	"net/http"
	"context"
)

// CtxMap holds the ephemeral data for the http.Request, it exists during the lifetime
// of the request
type CtxMap map[interface{}]interface{}

// Internal int key
var ctxMapKeyIndex int = 0
var ctxmapKey = ctxMapKeyIndex

// GetNextCtxMapKey returns next available integer key for the CtxMap
func GetNextCtxMapKey() int {
	//Do not expect the integer to wrap around ever happen
	ctxMapKeyIndex++
	return ctxMapKeyIndex
}


// InitRequestCtxMap returns a shallow copy of r with its context changed
// to a copy of its current ctx in which an empty map is associated with an internal 
// key. The map can then be rertieved with GetCtxMap(r *http.Request)
// Note: RequestCtx is per request-scoped
//       This function shall be normally called once at the very early stage of request 
//       processing, like in the request dispatching routing
func InitRequestCtxMap(r *http.Request) *http.Request {
	//Better way to handle this?
	//r.WithContext will copy the whole http.Request
	return r.WithContext(context.WithValue(r.Context(), ctxmapKey, CtxMap{}))
}

// GetCtxMap returns the data structure CtxMap
// CtxMap exists during the lifetime of the request and it's specific to the current request
func GetCtxMap(r *http.Request) CtxMap {
	return r.Context().Value(ctxmapKey).(CtxMap)
}

// GetCtx returns the request context.Context
// See also GetCtxMap
func GetCtx(r *http.Request) context.Context {
	return r.Context()
}
