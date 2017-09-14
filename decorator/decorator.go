// Copyright 2017 The Golight Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decorator

import (
	"net/http"
)

// Decorator impliments the decorator pattern for http.Handler
// It's also commonly refered to as middleware / adopter pattern. 
// The idea behind is to write reuseable, moduler, and common code that used 
// across multiple APIs without polute the core logic for the API.
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
*/
// Refer to decorator_test for details of the use cases
type Decorator func(http.Handler) http.Handler


// Decorate decorates the http.Handler with a list of Decorators and 
// Returns the decorated http.Handler
// The Decorators are left associative. The order of the decorators are
// called as the following:
/*
	h := Decorate(hdl, d3, d2, d1)
	before d1
	before d2
	before d3
	HandlerFunc
	after d3
	after d2
	after d1
*/ 
// Refer to decorator_test for addtional info
func Decorate(h http.Handler, decorators ...Decorator) http.Handler {
	for _, d := range decorators {
		h = d(h)
	}
	//use reverse order, it is not good for multiple call to Decorate
	//see TestDecoratorDecorate
	//for i:=len(decorators)-1; i>=0; i-- {
	//	h=decorators[i](h)
	//}
	return h
}

// DecoratorChain is an array of Decorators
type DecoratorChain []Decorator

// Chain chains the Decorators into an array
// DecorateChain may be re-used to decorate different http.Handler
func Chain(decorators ...Decorator) DecoratorChain {
	return decorators
}

// Decorate an http.Handler with the given DecoratorChain
func (dc DecoratorChain) Decorate(h http.Handler) http.Handler {
	for _, d := range dc {
		h = d(h)
	}
	return h
}

// Append additional Decorators to the given DecoratorChain
func (dc DecoratorChain) Append(decorators ...Decorator) DecoratorChain {
		dc = append(dc, decorators...)
		return dc
}
