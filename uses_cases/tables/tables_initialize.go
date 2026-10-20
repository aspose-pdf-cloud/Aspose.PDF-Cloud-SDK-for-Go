package main

import (
	"strconv"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func InitializeTable() *asposepdfcloud.Table {
	// Initialize new table
	num_of_cols := 5
	num_of_rows := 5

	header_text_state := asposepdfcloud.TextState{
		Font:            "Arial Bold",
		FontSize:        11,
		ForegroundColor: &asposepdfcloud.Color{A: 255, R: 255, G: 255, B: 255},
		FontStyle:       asposepdfcloud.FontStylesBold,
	}

	common_text_state := asposepdfcloud.TextState{
		Font:            "Arial Bold",
		FontSize:        11,
		ForegroundColor: &asposepdfcloud.Color{A: 255, R: 112, G: 112, B: 112},
		FontStyle:       asposepdfcloud.FontStylesRegular,
	}

	col_widths := ""
	for col_index := 0; col_index < num_of_cols; col_index++ {
		col_widths += " 70"
	}

	var table_rows = []asposepdfcloud.Row{}

	border_table_border := asposepdfcloud.GraphInfo{
		Color:     &asposepdfcloud.Color{A: 255, R: 0, G: 255, B: 0},
		LineWidth: 0.5,
	}

	for row_index := 0; row_index < num_of_rows; row_index++ {
		var row_cells = []asposepdfcloud.Cell{}

		for col_index := 0; col_index < num_of_cols; col_index++ {
			cell := asposepdfcloud.Cell{DefaultCellTextState: &common_text_state}

			if row_index == 0 {
				// header cells
				cell.BackgroundColor = &asposepdfcloud.Color{A: 255, R: 128, G: 128, B: 128}
				cell.DefaultCellTextState = &header_text_state
			} else {
				cell.BackgroundColor = &asposepdfcloud.Color{A: 255, R: 255, G: 255, B: 255}
			}

			text_rect := asposepdfcloud.TextRect{}
			if row_index == 0 {
				text_rect.Text = "header #" + strconv.Itoa(col_index)
			} else {
				text_rect.Text = "value '" + strconv.Itoa(row_index) + "','" + strconv.Itoa(col_index) + "'"
			}

			cell.Paragraphs = append(cell.Paragraphs, text_rect)

			row_cells = append(row_cells, cell)
		}

		row := asposepdfcloud.Row{Cells: row_cells}

		table_rows = append(table_rows, row)
	}

	table := asposepdfcloud.Table{Left: 150, Top: 250, ColumnWidths: col_widths, Rows: table_rows}

	table.DefaultCellBorder = &asposepdfcloud.BorderInfo{
		Top:                 &border_table_border,
		Right:               &border_table_border,
		Bottom:              &border_table_border,
		Left:                &border_table_border,
		RoundedBorderRadius: 2,
	}

	return &table
}
