# Table
Represents a table that can be added to the page.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**[]Link**](Link.md) | Link to the document. | [optional] [default to null]
**Alignment** | [***HorizontalAlignment**](HorizontalAlignment.md) | Gets HorizontalAlignment of the table alignment. | [optional] [default to null]
**HorizontalAlignment** | [***HorizontalAlignment**](HorizontalAlignment.md) | Gets HorizontalAlignment of the table alignment. | [optional] [default to null]
**VerticalAlignment** | [***VerticalAlignment**](VerticalAlignment.md) | Gets VerticalAlignment of the annotation. | [optional] [default to null]
**Top** | **float64** | Gets or sets the table top coordinate. | [optional] [default to null]
**Left** | **float64** | Gets or sets the table left coordinate. | [optional] [default to null]
**DefaultCellTextState** | [***TextState**](TextState.md) | Gets or sets the default cell text state. | [optional] [default to null]
**DefaultCellPadding** | [***MarginInfo**](MarginInfo.md) | Gets or sets the default cell padding. | [optional] [default to null]
**Border** | [***BorderInfo**](BorderInfo.md) | Gets or sets the border. | [optional] [default to null]
**Margin** | [***MarginInfo**](MarginInfo.md) | Gets or sets a outer margin for paragraph (for pdf generation) | [optional] [default to null]
**Rows** | [**[]Row**](Row.md) | Sets the rows of the table. | [default to null]
**DefaultColumnWidth** | **string** | Gets default cell border; | [optional] [default to null]
**DefaultCellBorder** | [***BorderInfo**](BorderInfo.md) | Gets default cell border; | [optional] [default to null]
**Broken** | [***TableBroken**](TableBroken.md) | Gets or sets table vertial broken; | [optional] [default to null]
**ColumnWidths** | **string** | Gets the column widths of the table. | [optional] [default to null]
**RepeatingRowsCount** | **int32** | Gets the first rows count repeated for several pages | [optional] [default to null]
**RepeatingColumnsCount** | **int32** | Gets or sets the maximum columns count for table | [optional] [default to null]
**RepeatingRowsStyle** | [***TextState**](TextState.md) | Gets the style for repeating rows | [optional] [default to null]
**CornerStyle** | [***BorderCornerStyle**](BorderCornerStyle.md) | Gets or sets the styles of the border corners | [optional] [default to null]
**BreakText** | [***TextRect**](TextRect.md) | Gets or sets break text for table | [optional] [default to null]
**BackgroundColor** | [***Color**](Color.md) | Gets or sets table background color | [optional] [default to null]
**IsBordersIncluded** | **bool** | Gets or sets border included in column widhts. | [optional] [default to null]
**ColumnAdjustment** | [***ColumnAdjustment**](ColumnAdjustment.md) | Gets or sets the table column adjustment. | [optional] [default to null]
**ZIndex** | **int32** | Gets ZIndex of the annotation. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)[[View Source]](../table.go)


