package email

import (
	a "github.com/jpincas/htmlfunc/attributes"
	"github.com/jpincas/htmlfunc/css"
	h "github.com/jpincas/htmlfunc/html"
)

const (
	mainMargin    = 0
	sectionMargin = 10
	mainWidth     = 600
)

var (
	white       = css.RGB(255, 255, 255)
	neutralGrey = css.RGB(240, 240, 240)
)

// Html is a somewhat opinionated attempt at standardising a wrapper that works well across clients and can be resused for all emails
// It incorporates the head and body declarations and works its way down to a table wrapper consisting of a single
// td which then sort of acts as the starting point for building a custom layout,
// i.e. the attributes and elements you pass in end up in that final td wrapper
func Html(attrs a.Attributes, sections ...h.Element) h.Element {
	return h.Html(
		a.Attrs(a.Xmlns("http://www.w3.org/1999/xhtml")),
		h.Head(
			a.Attrs(),
			h.Meta(a.Attrs(
				a.Name("viewport"),
				a.Content("width=device-width, initial-scale=1.0"),
			)),
		),

		// We immedatiely wrap the contents in a top-level table
		// This is good practice because email clients routinely strip
		// attributes from the BODY tag - so we apply any top level
		// attributes to this wrapper table, its tr and its td
		h.Body(
			a.Attrs(
				a.Style(
					css.Margin(css.Zero),
					css.BackgroundColor(neutralGrey),
				),
			),
			h.Table(
				a.Attrs(
					a.CellPadding(mainMargin),
					a.Width(css.NoUnits(mainWidth)),
					a.Style(
						css.BorderCollapse(css.Collapse),
						css.MarginLeft(css.Auto),
						css.MarginRight(css.Auto),
					),
				),
				h.Tr(
					a.Attrs(),
					h.Td(
						a.Attrs1(
							a.Width(css.WithUnits(100, css.Percent)),
							attrs,
						),
						h.Table(
							a.Attrs(
								a.CellPadding(sectionMargin),
								a.CellSpacing(0),
								a.Style(
									css.Margin(css.Zero),
									css.BackgroundColor(white),
								),
							),
							sections...,
						),
					),
				),
			),
		),
	)
}

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
