package tablelayout

import (
	a "github.com/jpincas/htmlfunc/attributes"
	"github.com/jpincas/htmlfunc/css"
	h "github.com/jpincas/htmlfunc/html"
)

const (
	docSectionMargin       = 20
	devModeBackgroundColor = "blue"
)

func Page(devMode bool, attrs a.Attributes, sections ...h.Element) h.Element {
	if devMode {
		h.ElementsAppendAttrs(sections, a.Style(css.BackgroundColor(devModeBackgroundColor)))
	}

	return h.Table(
		a.Attrs4(
			a.CellPadding(0),
			a.CellSpacing(0),
			a.Width(css.WithUnits(100, css.Percent)),
			a.Style(
				css.PageBreakBefore(css.Always),
				css.BorderCollapse(css.Collapse),
			),
			attrs,
		),
		h.Tr(
			a.Attrs(),
			h.Td(
				attrs,
				h.Table(
					a.Attrs(
						a.Width(css.WithUnits(100, css.Percent)),
						a.CellPadding(0),
						a.CellSpacing(docSectionMargin),
					),
					sections...,
				),
			),
		),
	)
}
