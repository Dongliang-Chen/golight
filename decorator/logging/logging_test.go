// Copyright 2017 The Golight Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logging_test

import (
	"testing"
	"bytes"
	"net/http"
	"io/ioutil"
	"net/http/httptest"
	"github.com/dlmc/golight/decorator"
	"github.com/dlmc/golight/ghttp"
	"github.com/dlmc/golight/decorator/logging"
)


//An Http request handler function
var th = ghttp.HandlerFunc(func(c ghttp.Ctx, h *ghttp.Http) ghttp.Ctx {
	qmap := h.Query
	str := "HandlerFunc " + qmap["a"][0] + qmap["b"][0] + qmap.Encode() + "\n"
	//Refer to logger_test for more log usecases
	//lg := logging.GetLogger(c)
	lg := h.Log
	lg = lg.Str("t", "t1")
	lg.Logger().Print("hello world", 23)
	lg.Logger().Error().Msg("s1")
	
	h.W.Write([]byte(str))
	return c
})

//Process test results
func tResult1(t *testing.T, res *http.Response, err error, want, wantlog string, out *bytes.Buffer) {
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
	
	gotlog :=  out.String()
	if gotlog != wantlog {
		t.Errorf("tResult failed got:  %v\nwant: %v", gotlog, wantlog)
	}
}


func TestLoggingDecorator(t *testing.T) {
	mux := http.NewServeMux()
	
	out := &bytes.Buffer{}
	lg := logging.CreateDecor(logging.NewContext(out))

	h := decorator.Decorate(th, lg)
	mux.Handle("/test", ghttp.Router{"GET":h})
	
	ts := httptest.NewServer(mux)
	defer ts.Close()
	
	res,err := http.Get(ts.URL+"/test?a=a1&b=b1&c")
	
	strRes := "HandlerFunc a1b1a=a1&b=b1&c=\n"
	logStrs := `{"l":"debug","t":"t1","m":"hello world23"}`+"\n"+`{"l":"error","t":"t1","m":"s1"}`+"\n"
	tResult1(t, res, err, strRes, logStrs, out)
}


