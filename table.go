/*
*
* Copyright (c) 2026 Aspose.PDF Cloud
* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:
* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.
* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
* OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
* SOFTWARE.
*
*/

package asposepdfcloud

// Represents a table that can be added to the page.
type Table struct {
	// Link to the document.
	Links []Link `json:"Links,omitempty"`
	// Gets HorizontalAlignment of the table alignment.
	Alignment HorizontalAlignment `json:"Alignment,omitempty"`
	// Gets HorizontalAlignment of the table alignment.
	HorizontalAlignment HorizontalAlignment `json:"HorizontalAlignment,omitempty"`
	// Gets VerticalAlignment of the annotation.
	VerticalAlignment VerticalAlignment `json:"VerticalAlignment,omitempty"`
	// Gets or sets the table top coordinate.
	Top float64 `json:"Top,omitempty"`
	// Gets or sets the table left coordinate.
	Left float64 `json:"Left,omitempty"`
	// Gets or sets the default cell text state.
	DefaultCellTextState *TextState `json:"DefaultCellTextState,omitempty"`
	// Gets or sets the default cell padding.
	DefaultCellPadding *MarginInfo `json:"DefaultCellPadding,omitempty"`
	// Gets or sets the border.
	Border *BorderInfo `json:"Border,omitempty"`
	// Gets or sets a outer margin for paragraph (for pdf generation)
	Margin *MarginInfo `json:"Margin,omitempty"`
	// Sets the rows of the table.
	Rows []Row `json:"Rows"`
	// Gets default cell border;
	DefaultColumnWidth string `json:"DefaultColumnWidth,omitempty"`
	// Gets default cell border;
	DefaultCellBorder *BorderInfo `json:"DefaultCellBorder,omitempty"`
	// Gets or sets table vertial broken;
	Broken TableBroken `json:"Broken,omitempty"`
	// Gets the column widths of the table.
	ColumnWidths string `json:"ColumnWidths,omitempty"`
	// Gets the first rows count repeated for several pages
	RepeatingRowsCount int32 `json:"RepeatingRowsCount,omitempty"`
	// Gets or sets the maximum columns count for table
	RepeatingColumnsCount int32 `json:"RepeatingColumnsCount,omitempty"`
	// Gets the style for repeating rows
	RepeatingRowsStyle *TextState `json:"RepeatingRowsStyle,omitempty"`
	// Gets or sets the styles of the border corners
	CornerStyle BorderCornerStyle `json:"CornerStyle,omitempty"`
	// Gets or sets break text for table
	BreakText *TextRect `json:"BreakText,omitempty"`
	// Gets or sets table background color
	BackgroundColor *Color `json:"BackgroundColor,omitempty"`
	// Gets or sets border included in column widhts.
	IsBordersIncluded bool `json:"IsBordersIncluded,omitempty"`
	// Gets or sets the table column adjustment.
	ColumnAdjustment ColumnAdjustment `json:"ColumnAdjustment,omitempty"`
	// Gets ZIndex of the annotation.
	ZIndex int32 `json:"ZIndex,omitempty"`
}
