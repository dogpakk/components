package table

import (
	"github.com/jpincas/htmlfunc/attributes"
	h "github.com/jpincas/htmlfunc/html"
)

// ConstructTable is for quickly constructing simple tables with global attributes
func ConstructTable(attrs attributes.Attributes, headerRow []string, rows []h.Elements) h.Element {
	return ComplexTable{
		GlobalAttrs: attrs,
		HeaderRow:   headerRow,
		Rows:        rows,
	}.Render()
}

type ComplexTable struct {
	FirstColumnIsTitle                     bool
	HeaderRow                              []string
	Rows                                   []h.Elements
	GlobalAttrs                            attributes.Attributes
	HeadAttrs, HeadRowAttrs, HeadCellAttrs attributes.Attributes
	LastRowAttrs, LastRowCellAttrs         attributes.Attributes
	FirstColumnCellAttrs                   attributes.Attributes
	BodyAttrs, BodyRowAttrs, BodyCellAttrs attributes.Attributes
	NthColumnCellAttrs                     map[int]attributes.Attributes
}

// ConstructComplexTable is for constructing more complex tables with inline attributes
// at every level.  Useful, for example, for tables in HTML emails
func (complexTable ComplexTable) Render() h.Element {
	headerCells := h.Els()
	for _, columnHeading := range complexTable.HeaderRow {
		headerCells = append(headerCells, h.Th(complexTable.HeadCellAttrs, h.Text(columnHeading)))
	}
	header := h.THead(complexTable.HeadAttrs, h.Tr(complexTable.HeadRowAttrs, headerCells...))

	bodyRows := h.Els()
	for i, row := range complexTable.Rows {
		tableCells := h.Els()
		var cellAttrs, rowAttrs attributes.Attributes

		if i < (len(complexTable.Rows)-1) || len(complexTable.Rows) == 1 {
			cellAttrs = complexTable.BodyCellAttrs
			rowAttrs = complexTable.BodyRowAttrs
		} else {
			cellAttrs = complexTable.LastRowCellAttrs
			rowAttrs = complexTable.LastRowAttrs
		}

		for i, cellValue := range row {
			cellAttrs_ := cellAttrs
			if i == 0 {
				cellAttrs_ = complexTable.FirstColumnCellAttrs
			} else if overrideAttrs, ok := complexTable.NthColumnCellAttrs[i]; ok {
				cellAttrs_ = overrideAttrs
			}

			tableCells = append(tableCells, h.Td(cellAttrs_, cellValue))
		}

		bodyRows = append(bodyRows, h.Tr(rowAttrs, tableCells...))
	}

	body := h.TBody(complexTable.BodyAttrs, bodyRows...)

	return h.Table(
		complexTable.GlobalAttrs,
		header,
		body,
	)
}
