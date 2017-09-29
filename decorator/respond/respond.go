// Copyright 2017 The Golight Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package respond

import (
	"github.com/dlmc/golight/decorator"
	"github.com/dlmc/golight/ghttp"
	"encoding/json"
)

// CreateDecor creates a respond decorator that will send out http response using
// the content of h.Resp struct
func CreateDecor() decorator.Decorator {
	return func(next ghttp.Handler) ghttp.Handler {
		return ghttp.HandlerFunc(func(c ghttp.Ctx, h *ghttp.Http) ghttp.Ctx{
			c = next.ServeHTTPWithCtx(c, h)
			w := h.W
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(h.Resp.Code)
			//out, _ := json.Marshal(h.Resp)
			//w.Write(out)	   //will not with "\n" at the end
			json.NewEncoder(w).Encode(h.Resp)	 // will write "\n" at the end
			return c
		})
	}		
}

