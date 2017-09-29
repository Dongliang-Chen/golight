// Copyright 2017 The Golight Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package queryparser_test

import (
	"testing"
	"net/http"
	"io/ioutil"
	"net/http/httptest"
	dc "github.com/dlmc/golight/decorator"
	qp "github.com/dlmc/golight/decorator/queryparser"
	rt "github.com/dlmc/golight/router"
)

//An Http request handler function
var th = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	qmap := qp.GetQueryValues(r)
	str := "HandlerFunc " + qmap["a"][0] + qmap["b"][0] + qmap.Encode() + "\n"
	w.Write([]byte(str))
})


//Simple Decorator
func tDecorator(tag string) dc.Decorator {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("before " + tag))
			h.ServeHTTP(w, r)
			w.Write([]byte("after " + tag))
		})
	}
}


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
		t.Errorf("tResult failed, got: \n%s\nwant: \n%s\n", got, want)
	}
}

func TestQueryparserDecorator(t *testing.T) {
	mux := http.NewServeMux()
	
	d1 := tDecorator("d1\n")
	d2 := tDecorator("d2\n")
	d3 := tDecorator("d3\n")
	qparser := qp.CreateDecor()

	h := dc.Decorate(th, d3, d2, d1, qparser)
	//same as h := qparser(d1(d2(d3(th))))

	mux.Handle("/test", rt.Router{"GET":h})
	
	ts := httptest.NewServer(mux)
	defer ts.Close()
	
	res,err := http.Get(ts.URL+"/test?a=a1&b=b1&c")
	
	strRes := "before d1\nbefore d2\nbefore d3\nHandlerFunc a1b1a=a1&b=b1&c=\nafter d3\nafter d2\nafter d1\n"	
	tResult(t, res, err, strRes)
}

