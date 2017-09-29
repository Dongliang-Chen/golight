
package responseheader_test

import (
	"testing"
	"net/http"
	"reflect"
	"net/http/httptest"
	dc "github.com/dlmc/golight/decorator"
	rh "github.com/dlmc/golight/decorator/responseheader"
//	"fmt"

)

//An Http request handler function
var th = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HandlerFunc\n"))
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
func tResult(t *testing.T, h http.Handler, wantBody string, wantH1, wantH2 []string) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	h.ServeHTTP(w, r)

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
	dh1 := rh.CreateDecor(rh.HeaderMap{"K1":"v1a", "K2":"v2a"}, true)  //true to add
	dh2 := rh.CreateDecor(rh.HeaderMap{"K1":"v1b", "K2":"v2b"}, true)
	dh3 := rh.CreateDecor(rh.HeaderMap{"K2":"v2c"}, false) //false to set / replace "k2"

	h := dc.Decorate(th, dh3, dh2, dh1, d3, d2, d1)
	
	strRes := "before d1\nbefore d2\nbefore d3\nHandlerFunc\nafter d3\nafter d2\nafter d1\n"
	tResult(t, h, strRes, []string{"v1a", "v1b"}, []string{"v2c"})
}

