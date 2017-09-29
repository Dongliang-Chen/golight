
package decorator_test

import (
	"testing"
	"net/http"
	"reflect"
	"fmt"
	"io"
	"io/ioutil"
	"net/http/httptest"
	"github.com/dlmc/golight/decorator"
	"github.com/dlmc/golight/ghttp"
)

//An Http request handler function
var th = ghttp.HandlerFunc(func(c ghttp.Ctx, h *ghttp.Http) ghttp.Ctx {
	h.W.Write([]byte("HandlerFunc\n"))
	return c
})

//Another http request handler function
var th1 = ghttp.HandlerFunc(func(c ghttp.Ctx, h *ghttp.Http) ghttp.Ctx {
	h.W.Write([]byte("HandlerFunc1\n"))
	return c
})

//Simple Decorator
func tDecorator(tag string) decorator.Decorator {
	return func(hdl ghttp.Handler) ghttp.Handler {
		return ghttp.HandlerFunc(func(c ghttp.Ctx, h *ghttp.Http) ghttp.Ctx {
			h.W.Write([]byte("before " + tag))
			c = hdl.ServeHTTPWithCtx(c, h)
			h.W.Write([]byte("after " + tag))
			return c
		})
	}
}


//Compare two functions
func fEqual(f1, f2 interface{}) bool {
	val1 := reflect.ValueOf(f1)
	val2 := reflect.ValueOf(f2)
	return val1.Pointer() == val2.Pointer()
}

//Process test results
func tResult(t *testing.T, h ghttp.Handler, want string) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	h.ServeHTTPWithCtx(nil,  &ghttp.Http{W:w, R:r})

	if got := w.Body.String(); want != got {
		t.Errorf("Decorate handler failed, got: %s, want: %s", got, want)
	}
}

//Simple test case for Chain length
func TestDecoratorChainLen(t *testing.T) {
	d1 := func(h ghttp.Handler) ghttp.Handler { return nil }
	d2 := func(h ghttp.Handler) ghttp.Handler { return nil }
	c := decorator.Chain(d1, d2)

	if len(c) != 2 {
		t.Errorf("Chain len failed %d", len(c))		
	}
	if !fEqual(c[0], d1) || !fEqual(c[1], d2) {
		t.Errorf("Decorator Chain failed")
	}
}

//Simple test case for DecoratorChain with single item
func TestDecoratorChainSingleDecortor(t *testing.T) {
	d1 := tDecorator("d1\n")

	d := decorator.Chain(d1)
	h := d.Decorate(th)
	
	strRes := "before d1\nHandlerFunc\nafter d1\n"
	tResult(t, h, strRes)
}

//Simple test case for DecoratorChain with multi items
func TestDecoratorChainMultiDecortors(t *testing.T) {
	d1 := tDecorator("d1\n")
	d2 := tDecorator("d2\n")
	d3 := tDecorator("d3\n")

	d := decorator.Chain(d3, d2, d1)	
	h := d.Decorate(th)
	
	strRes := "before d1\nbefore d2\nbefore d3\nHandlerFunc\nafter d3\nafter d2\nafter d1\n"
	tResult(t, h, strRes)
}

//Simple test case for DecoratorChain append and reuse
func TestDecoratorChainAppend(t *testing.T) {
	d1 := tDecorator("d1\n")
	d2 := tDecorator("d2\n")
	d3 := tDecorator("d3\n")

	d := decorator.Chain(d3)
	decorator := d.Append(d2, d1)	
	h := decorator.Decorate(th)
	h1 := decorator.Decorate(th1)
	
	strRes := "before d1\nbefore d2\nbefore d3\nHandlerFunc\nafter d3\nafter d2\nafter d1\n"
	tResult(t, h, strRes)
	
	strRes1 := "before d1\nbefore d2\nbefore d3\nHandlerFunc1\nafter d3\nafter d2\nafter d1\n"
	tResult(t, h1, strRes1)
	
	h = d.Decorate(th)
	strRes = "before d3\nHandlerFunc\nafter d3\n"
	tResult(t, h, strRes)
	
}

//Simple test case for Decorate with single item
func TestDecoratorDecorateSingleDecorator(t *testing.T) {
	d1 := tDecorator("d1\n")

	h := decorator.Decorate(th, d1)
	
	strRes := "before d1\nHandlerFunc\nafter d1\n"
	tResult(t, h, strRes)
}

//Simple test case for Decorate with multiple items
func TestDecoratorDecorateMultipleDecorators(t *testing.T) {
	d1 := tDecorator("d1\n")
	d2 := tDecorator("d2\n")
	d3 := tDecorator("d3\n")

	h := decorator.Decorate(th, d3, d2, d1)
	
	strRes := "before d1\nbefore d2\nbefore d3\nHandlerFunc\nafter d3\nafter d2\nafter d1\n"
	tResult(t, h, strRes)
}

//Simple test case for multiple Decorate calls
func TestDecoratorDecorateMultipleCalls(t *testing.T) {
	d1 := tDecorator("d1\n")
	d2 := tDecorator("d2\n")
	d3 := tDecorator("d3\n")

	h := decorator.Decorate(th, d3, d2)
	h = decorator.Decorate(h, d1)
	
	strRes := "before d1\nbefore d2\nbefore d3\nHandlerFunc\nafter d3\nafter d2\nafter d1\n"
	tResult(t, h, strRes)
}

//Simple test case for differnt Decorate level
func TestDecoratorDecorateDifferentLevel(t *testing.T) {
	d1 := tDecorator("d1\n")
	d2 := tDecorator("d2\n")
	d3 := tDecorator("d3\n")

	h := decorator.Decorate(th, d3, d2)
	h1 := decorator.Decorate(h, d1)
	
	strRes := "before d2\nbefore d3\nHandlerFunc\nafter d3\nafter d2\n"
	tResult(t, h, strRes)
	
	strRes = "before d1\nbefore d2\nbefore d3\nHandlerFunc\nafter d3\nafter d2\nafter d1\n"
	tResult(t, h1, strRes)
}

func tResultRouter(t *testing.T, res *http.Response, err error, want string) {
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

func getRouter() ghttp.Router {
	return ghttp.Router{
		"GET":&StoreDelete{},
	}
}

type StoreDelete struct {
}

func (s *StoreDelete) ServeHTTPWithCtx(c ghttp.Ctx, h *ghttp.Http) ghttp.Ctx {
	h.W.Write([]byte("HandlerFunc\n"))
	return c
}

func TestDecorateRouter(t *testing.T) {
	mux := http.NewServeMux()
	
	d1 := tDecorator("d1\n")
	d2 := tDecorator("d2\n")
	d3 := tDecorator("d3\n")

	h := decorator.DecorateRouter(getRouter(), d3, d2, d1)
	mux.Handle("/test", h)
	
	ts := httptest.NewServer(mux)
	defer ts.Close()
	
	res,err := http.Get(ts.URL+"/test")
	
	strRes := "before d1\nbefore d2\nbefore d3\nHandlerFunc\nafter d3\nafter d2\nafter d1\n"	
	tResultRouter(t, res, err, strRes)
}


func ExampleHttpTest() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
	// Output:
	// 200
	// text/html; charset=utf-8
	// <html><body>Hello World!</body></html>
}

func ExampleHttpServer() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `Hello, client`)
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		fmt.Printf("Decorator Chain failed", err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Printf("Decorator Chain failed", err)
	}

	fmt.Printf("%s", greeting)
	// Output: Hello, client

}

