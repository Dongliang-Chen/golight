// Copyright 2017 The Golight Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decorator

import (
	"net/http"
	"github.com/dlmc/golight/router"

)

// Having been researching the middleware / decorator pattern...
// especially when it comes to handler request handling ctx...

// Can we add a context to the Decorator so that it can carry the context tree?
// At the top most level of the decorator in the http request handling chain, we should
// use the default context from context package and then at each decorator level, if 
// new context is needed, we can chain the context to the existing context and passes it
// down to the decorator at the next level.
// Upon the returning from a low level of the decorator chain to the next level, the context
// restores to the context for that specific level
// Yes, the answer is possible by changing the Decorator from a func to a struct and then 
// attaching method similar to 
// func (* Decorator)ServeHTTP(c context.Context, w http.ResponseWriter, r *http.Request)
// So the trade off is to pass the ctx around instead of store the ctx map in the reuqest obj

// Decorator impliments the decorator pattern for http.Handler
// It's also commonly refered to as middleware / adopter pattern. 
// The idea behind is to write reuseable, moduler, and common code that used 
// across multiple APIs without polute the core logic for the API. The other key
// design consideration is to make it request-scoped such that it will have no 
// side effect once the request it processed
// Usage example:
/*
// Using the following pattern to create a decorator
// generic Decorator pattern
func GenericDecorator(args ...interface{}) Decorator {
	return func(next http.Handler) http.Handler {
		return	http.HandlerFunc(func(w	http.ResponseWriter,	r	*http.Request)	{
			//... logic before
			next.ServeHTTP(w, r)	
			//... logic after
		})
	}
}
// ExampleDecorator is an Decorator that sets an HTTP header.
func ExampleDecorator(key, value string) Decorator {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header.Add(key, value)
			next.ServeHTTP(w, r)
		})
	}
}
// Another way to implement the ExampleDecorator.
// Note: key/value is shared across all requests.
// This pattern is good to implement across request handling 
// like performance monitoring etc.
type HeaderDecor struct {
	 http.Handler
	 key,value string
}
func (m *HeaderDecor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header.Add(m.key, m.value)
	m.Handler.ServeHTTP(w, r)
}
func ExampleDecorator(key, value string) Decorator {
	return func(next http.Handler) http.Handler {
		return &MyDecor{next,key, value}
	}		
}
*/

// Decordator is a func that takes an http.Handler and returns an http.Handler.
// Refer to decorator_test for details of the use cases.
type Decorator func(http.Handler) http.Handler 


// Decorate decorates the http.Handler with a list of Decorators and 
// returns the decorated http.Handler.
// Each of the decorators pasted in is a func that takes an http.Hanlder and
// returns an http.Handler.
// The idea is to chain the http.Handler h with the decorators so that the decorator
// funcion is called before h.
// The Decorators are left associative. The order of the decorators are
// called as the following (from right / last to the left / first):
/*
	h := Decorate(hdl, d3, d2, d1)
	//same as
	// h:= d1(d2(d3(hdl)))
	before d1
	before d2
	before d3
	HandlerFunc
	after d3
	after d2
	after d1
*/ 
// Refer to decorator_test for addtional info.
func Decorate(h http.Handler, decorators ...Decorator) http.Handler {
	for _, d := range decorators {
		h = d(h)
	}
	return h
}

// DecorateRouter decorates each http.Handler in the router with the decorators list
// and returns the decorated router.
func DecorateRouter(r router.Router, decorators ...Decorator) router.Router {
	for k,v := range r {
		r[k] = Decorate(v, decorators...)
	}
	return r
}

// DecoratorChain is an array of Decorators.
type DecoratorChain []Decorator

// Chain chains the Decorators into an array.
// DecorateChain may be re-used to decorate different http.Handler.
func Chain(decorators ...Decorator) DecoratorChain {
	return decorators
}

// Decorate decortoes an http.Handler with the given DecoratorChain
// and returns the decorated http Handler.
func (dc DecoratorChain) Decorate(h http.Handler) http.Handler {
	return Decorate(h, dc...)
}

// DecorateRouter decorates each http.Handler in the router with the 
// given DecoratorChain and returns the decorated router.
func (dc DecoratorChain) DecorateRouter(r router.Router) router.Router {
	return DecorateRouter(r, dc...)
}

// Append additional Decorators to the given DecoratorChain.
func (dc DecoratorChain) Append(decorators ...Decorator) DecoratorChain {
		dc = append(dc, decorators...)
		return dc
}
