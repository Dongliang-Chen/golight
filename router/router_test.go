// Copyright 2017 The Golight Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package router_test

import (
	"testing"
	"net/http"
	"io/ioutil"
	"net/http/httptest"
	"github.com/dlmc/golight/ctx"
	rt "github.com/dlmc/golight/router"
)


//Get Http request handler function
var getHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	m := ctx.GetCtxMap(r)
	m["b"]=" 2"
	str := "GetHandler" + m["b"].(string)
		
	w.Write([]byte(str))
})

//Post http request handler function
var postHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	m := ctx.GetCtxMap(r)
	m["b"]=" 2"
	str := "PostHandler" + m["b"].(string)
	w.Write([]byte(str))
})


//Process test results
func tResult(t *testing.T, res *http.Response, err error, want string) {
	if err != nil {
		t.Errorf("tResult failed", err)
	}
	got, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Errorf("tResult failed", err)
	}
	if want != string(got) {
		t.Errorf("tResult failed, got: %s, want: %s", got, want)
	}
}

//Simple test case for Router
func TestRouterWithDefaultServeMux(t *testing.T) {
	http.Handle("/test", rt.Router{"GET":getHandler, "POST":postHandler})
	ts := httptest.NewServer(http.DefaultServeMux)
	defer ts.Close()
	url := ts.URL+"/test"

	res,err := http.Get(ts.URL)
	tResult(t, res, err, "404 page not found"+"\n")

	res,err = http.Head(url)
	tResult(t, res, err, "")

	res,err = http.Get(url)
	tResult(t, res, err, "GetHandler 2")

	res,err = http.Post(url, "text/html; charset=utf-8", nil)
	tResult(t, res, err, "PostHandler 2")
}

//Simple test case for Router
func TestRouterWithNewMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/test", rt.Router{"GET":getHandler, "POST":postHandler})
	
	//use http.ListenAndServe(":3000", mux) for real http server
	ts := httptest.NewServer(mux)
	defer ts.Close()
	url := ts.URL+"/test"

	res,err := http.Get(url)
	tResult(t, res, err, "GetHandler 2")

	res,err = http.Post(url, "text/html; charset=utf-8", nil)
	tResult(t, res, err, "PostHandler 2")
}

