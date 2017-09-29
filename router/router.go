// Copyright 2017 The Golight Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package router

import (
	"net/http"
	"sort"
	"strings"
	"github.com/dlmc/golight/ctx"

)

// Router impliments http handler registration
// Usage example:
// http.Handle("/query", router.Router{"GET":getHandler, "POST":postHandler})
// In case there is no handler found, http status code 501 is returned
// Refer to https://golang.org/src/net/http/method.go for supported methods
//          https://golang.org/pkg/net/http/ for overall http information
//
// Below is a simple HTTP server:
/*
//Get Http request handler function
var getHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetTestResponse"))
})

//Post http request handler function
var postHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PostTestResponse"))
})

//Put http request handler
type PutHandler struct {
}
func (ph *PutHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	//decode request
	//process request
	//encode response
	//write encoded response
	w.Write([]byte("PutTestResponse"))
}
func main() {
	mux := http.NewServeMux()
	mux.Handle("/test", Router{"GET":getHandler, "POST":postHandler, "PUT":PutHanlder{}})
	log.Println("Listening on port 8080...")
	log.Println(http.ListenAndServe(":8080", mux))
}
*/
type Router map[string]http.Handler


func (rt Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler:=rt[r.Method]; handler != nil {
		handler.ServeHTTP(w,ctx.InitRequestCtxMap(r))
	} else {
		//http.Error(w, http.StatusText(501), 501)
		allow := []string{}
		for k := range rt {
			allow = append(allow, k)
		}
		sort.Strings(allow)
		w.Header().Set("Allow", strings.Join(allow, ", "))
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), 
				http.StatusMethodNotAllowed)
		}
	}
}
