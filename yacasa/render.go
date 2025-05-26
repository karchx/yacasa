package main

import (
	"net/http"

	"github.com/maddalax/htmgo/framework/h"
)

func RenderPage(req *http.Request, w http.ResponseWriter, page func(ctx *h.RequestContext) *h.Page) {
	ctx := h.RequestContext{
		Request:  req,
		Response: w,
	}
	h.HtmlView(w, page(&ctx))
}
