// Copyright 2017 The Golight Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ghttp

import (
	"net/http"
	"net/url"
	"context"
	"sort"
	"strings"
	log "github.com/rs/zerolog"
)


type Response struct {
	Code int       		`json:"code"`
	Message string 		`json:"message,omitempty"`
	Data interface{}    `json:"data,omitempty"`
}

type Http struct {
	Resp  Response
	Query url.Values
	W http.ResponseWriter
	R *http.Request
	Log log.Context			//nil - use logging.Decor to assign the logger
}	

// Internal int key
var ctxKeyIndex int = 0

// GetNextCtxKey returns next available integer key for the Ctx
func GetNextCtxKey() int {
	//Do not expect the integer to wrap around ever happen
	ctxKeyIndex++
	return ctxKeyIndex
}


type Ctx interface {
	context.Context
}

type Handler interface {
	ServeHTTPWithCtx(Ctx, *Http) Ctx
}

type HandlerFunc func(Ctx, *Http) Ctx

func (hf HandlerFunc) ServeHTTPWithCtx(c Ctx, h *Http) Ctx {
    return hf(c, h)
}

func ChildCtx(parent Ctx, key, val interface{}) Ctx {
	return context.WithValue(parent, key, val)
}



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
//type Router map[string]http.Handler
type Router map[string]Handler


func (rt Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h:=rt[r.Method]; h != nil {
		hp := &Http{W:w, R:r, Query:r.URL.Query()}
		h.ServeHTTPWithCtx(context.Background(), hp)
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

/*
type valueCtx struct {
	Context
	key, val interface{}
}

type key int
const requestIDKey key = 0

type ResponseStruct struct {
	Code int       		`json:"code"`
	Message string 		`json:"message,omitempty"`
	Data interface{}    `json:"data,omitempty"`
}

func (rs *ResponseStruct) xx(w http.ResponseWriter){
	resp, _ := json.Marshal(rs)
	w.WriteHeader(d.Code)
	w.Write(resp)
}



func newContextWithRequestID(ctx context.Context, req *http.Request) context.Context {
    return context.WithValue(ctx, requestIDKey, req.Header.Get("X-Request-ID"))
}

func requestIDFromContext(ctx context.Context) string {
    return ctx.Value(requestIDKey).(string)
}

ctx := context.Background()
ctx = newContextWithRequestID(ctx, req)
*/