// Copyright 2017 The Golight Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ghttp_test

import (
	"testing"
	"net/http"
	"io/ioutil"
	"net/http/httptest"
	"github.com/dlmc/golight/ghttp"
)


//Get Http request handler function
var getHandler = ghttp.HandlerFunc(func(c ghttp.Ctx, h *ghttp.Http) ghttp.Ctx{
	c = ghttp.ChildCtx(c, "b", " 2")
	str := "GetHandler" + c.Value("b").(string)
	h.W.Write([]byte(str))
	return c
})

//Post http request handler function
var postHandler = ghttp.HandlerFunc(func(c ghttp.Ctx, h *ghttp.Http) ghttp.Ctx {
	c = ghttp.ChildCtx(c, "b", " 2")
	str := "PostHandler" + c.Value("b").(string)
	h.W.Write([]byte(str))
	return c
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
	http.Handle("/test", ghttp.Router{"GET":getHandler, "POST":postHandler})
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
	mux.Handle("/test", ghttp.Router{"GET":getHandler, "POST":postHandler})
	
	//use http.ListenAndServe(":3000", mux) for real http server
	ts := httptest.NewServer(mux)
	defer ts.Close()
	url := ts.URL+"/test"

	res,err := http.Get(url)
	tResult(t, res, err, "GetHandler 2")

	res,err = http.Post(url, "text/html; charset=utf-8", nil)
	tResult(t, res, err, "PostHandler 2")
}



//An Http request handler function
var th = ghttp.HandlerFunc(func(c ghttp.Ctx, h *ghttp.Http) ghttp.Ctx{
	qmap := h.Query
	str := "HandlerFunc " + qmap["a"][0] + qmap["b"][0] + qmap.Encode() + "\n"
	h.W.Write([]byte(str))
	return c
})



//Process test results
func tResultQuery(t *testing.T, res *http.Response, err error, want string) {
	if err != nil {
		t.Errorf("tResult failed", err)
	}
	got, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Errorf("tResult failed", err)
	}
	if want != string(got) {
		t.Errorf("tResult failed, got: \n%s\nwant: \n%s\n", got, want)
	}
}

func TestQuery(t *testing.T) {
	mux := http.NewServeMux()
	
	mux.Handle("/test", ghttp.Router{"GET":th})
	
	ts := httptest.NewServer(mux)
	defer ts.Close()
	
	res,err := http.Get(ts.URL+"/test?a=a1&b=b1&c")
	
	strRes := "HandlerFunc a1b1a=a1&b=b1&c=\n"	
	tResultQuery(t, res, err, strRes)
}


