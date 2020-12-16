package tablelayout

import (
	a "github.com/jpincas/htmlfunc/attributes"
	"github.com/jpincas/htmlfunc/css"
	h "github.com/jpincas/htmlfunc/html"
)

func Section(attrs a.Attributes, els ...h.Element) h.Element {
	return h.Tr(
		a.Attrs(),
		h.Td(
			a.Attrs1(
				a.Width(css.WithUnits(100, css.Percent)),
				attrs,
			),
			els...,
		),
	)
}
