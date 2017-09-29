// Copyright 2017 The Golight Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package respond_test

import (
	"testing"
	"net/http"
	"io/ioutil"
	"net/http/httptest"
	"github.com/dlmc/golight/decorator"
	"github.com/dlmc/golight/ghttp"
	"github.com/dlmc/golight/decorator/respond"
//	"encoding/json"
)


//An Http request handler function
var th = ghttp.HandlerFunc(func(c ghttp.Ctx, h *ghttp.Http) ghttp.Ctx {
	qmap := h.Query
	//str := "HandlerFunc " + qmap["a"][0] + qmap["b"][0] + qmap.Encode() + "\n"
	//str,_ := json.Marshal(qmap)
	r := &h.Resp	
	//r.Data = string(str)
	r.Data = qmap
	r.Code = http.StatusOK
	r.Message = http.StatusText(r.Code)
	return c
})

//Process test results
func tResult1(t *testing.T, res *http.Response, err error, want string) {
	if err != nil {
		t.Errorf("tResult failed", err)
	}
	got, err := ioutil.ReadAll(res.Body)
	//fmt.Printf("%+v\n", res)
	res.Body.Close()
	if err != nil {
		t.Errorf("tResult failed", err)
	}
	if want != string(got) {
		t.Errorf("tResult failed, got: %s, want: %s", got, want)
	}	
}


func TestLoggingDecorator(t *testing.T) {
	mux := http.NewServeMux()
	
	lg := respond.CreateDecor()

	h := decorator.Decorate(th, lg)
	mux.Handle("/test", ghttp.Router{"GET":h})
	
	ts := httptest.NewServer(mux)
	defer ts.Close()
	
	res,err := http.Get(ts.URL+"/test?a=a1&b=b1&c")
	
	strRes := `{"code":200,"message":"OK","data":{"a":["a1"],"b":["b1"],"c":[""]}}`+"\n"
	tResult1(t, res, err, strRes)
}


