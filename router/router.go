// Copyright 2017 The Golight Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package router

import (
	"net/http"
)

// Router impliments http handler registration
// Usage example:
// http.Handle("/query", router.Router{"GET":getHandler, "POST":postHandler})
// In case there is no handler found, http status code 501 is returned
// Refer to https://golang.org/src/net/http/method.go for supported methods
//          https://golang.org/pkg/net/http/ for overall http information

type Router map[string]http.Handler

func (rt Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler:=rt[r.Method]; handler != nil {
		handler.ServeHTTP(w,r)
	} else {
		//http.Error(w, http.StatusText(501), 501)
		allow := []string{}
		for k := range rt {
			allow = append(allow, k)
		}
		sort.Strings(allow)
		w.Header().Set("Allow", strings.Join(allow, ", "))
		if req.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), 
				http.StatusMethodNotAllowed)
		}
	}
}
