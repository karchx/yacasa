package main

import "github.com/maddalax/htmgo/framework/h"

func Index(ctx *h.RequestContext) *h.Page {
	return h.NewPage(
		h.Html(
			h.HxExtension(
				h.BaseExtensions(),
			),
			h.Head(
				h.Meta("viewport", "width=device-width, initial-scale=1"),
				h.Script("/public/htmgo.js"),
			),
			h.Body(
				h.Pf("Hello htmgo"),
			),
		),
	)
}
