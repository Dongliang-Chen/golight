
package header_test

import (
	"testing"
	"net/http"
	"reflect"
	"net/http/httptest"
	"github.com/dlmc/golight/decorator"
	"github.com/dlmc/golight/ghttp"
	"github.com/dlmc/golight/decorator/header"
//	"fmt"

)

//An Http request handler function
var th = ghttp.HandlerFunc(func(c ghttp.Ctx, h *ghttp.Http) ghttp.Ctx{
	h.W.Write([]byte("HandlerFunc\n"))
	return c
})


//Simple Decorator
func tDecorator(tag string) decorator.Decorator {
	return func(hdl ghttp.Handler) ghttp.Handler {
		return ghttp.HandlerFunc(func(c ghttp.Ctx, h *ghttp.Http) ghttp.Ctx{
			w := h.W
			w.Write([]byte("before " + tag))
			c = hdl.ServeHTTPWithCtx(c, h)
			w.Write([]byte("after " + tag))
			return c
		})
	}
}

//Process test results
func tResult(t *testing.T, h ghttp.Handler, wantBody string, wantH1, wantH2 []string) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	h.ServeHTTPWithCtx(nil, &ghttp.Http{W:w, R:r})

	if got := w.Body.String(); wantBody != got {
		t.Errorf("Responseheader handler failed, got: %s, want: %s", got, wantBody)
	}

	hd := w.Header()
	//access the header map directly, notice the case change to the key
	if !reflect.DeepEqual(hd["K1"], wantH1) {
		t.Errorf("Responseheader handler failed, got: %s, want: %s", hd["K1"], wantH1)
	}
	if !reflect.DeepEqual(hd["K2"], wantH2) {
		t.Errorf("Responseheader handler failed, got: %s, want: %s", hd["K2"], wantH2)
	}
	
	//access with Get method
	if got := hd.Get("K1"); got != wantH1[0] {
		t.Errorf("Responseheader handler failed, got: %s, want: %s", got, wantH1[0])		
	}
	
}

//Simple test case for Decorate with multiple items
func TestResponseheaderDecorator(t *testing.T) {
	d1 := tDecorator("d1\n")
	d2 := tDecorator("d2\n")
	d3 := tDecorator("d3\n")
	dh1 := header.CreateDecor(header.HeaderMap{"K1":"v1a", "K2":"v2a"}, true)  //true to add
	dh2 := header.CreateDecor(header.HeaderMap{"K1":"v1b", "K2":"v2b"}, true)
	dh3 := header.CreateDecor(header.HeaderMap{"K2":"v2c"}, false) //false to set / replace "k2"

	h := decorator.Decorate(th, dh3, dh2, dh1, d3, d2, d1)
	
	strRes := "before d1\nbefore d2\nbefore d3\nHandlerFunc\nafter d3\nafter d2\nafter d1\n"
	tResult(t, h, strRes, []string{"v1a", "v1b"}, []string{"v2c"})
}

