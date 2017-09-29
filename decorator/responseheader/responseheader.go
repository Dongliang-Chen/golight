// Copyright 2017 The Golight Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package responseheader

import (
	"net/http"
	. "github.com/dlmc/golight/decorator"
)

type HeaderMap map[string]string


// CreateDecor creates a response header decorator that adds/sets http headers into http response
// hm : HeaderMap{"this":"that", "key":"value"}
// add: true to add the http headers to the associated key. It appends to any 
//            existing values associated with key. 
//	    false to set the associated key with the value. It replaces any existing
//            values associated with key.
func CreateDecor(hm HeaderMap, add bool) Decorator {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if add {
				for k,v := range hm {
					w.Header().Add(k, v)
				}
			} else {
				for k,v := range hm {
					w.Header().Set(k, v)
				}
			}
			next.ServeHTTP(w, r)
		})
	}		
}
