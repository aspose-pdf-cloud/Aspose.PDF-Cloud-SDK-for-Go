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

// Represents text paragraphs as multiline text object.
type Paragraph struct {
	// Line spacing mode.
	LineSpacing LineSpacing `json:"LineSpacing,omitempty"`
	// Word wrap mode.
	WrapMode WrapMode `json:"WrapMode,omitempty"`
	// Horizontal alignment for the text inside paragraph's rectangle.
	HorizontalAlignment TextHorizontalAlignment `json:"HorizontalAlignment,omitempty"`
	// Left margin.
	LeftMargin float64 `json:"LeftMargin,omitempty"`
	// Right margin.
	RightMargin float64 `json:"RightMargin,omitempty"`
	// Top margin.
	TopMargin float64 `json:"TopMargin,omitempty"`
	// Bottom margin.
	BottomMargin float64 `json:"BottomMargin,omitempty"`
	// Rectangle of the paragraph.
	Rectangle *Rectangle `json:"Rectangle,omitempty"`
	// Rotation angle in degrees.
	Rotation float64 `json:"Rotation,omitempty"`
	// Subsequent lines indent value.
	SubsequentLinesIndent float64 `json:"SubsequentLinesIndent,omitempty"`
	// Vertical alignment for the text inside paragraph's rectangle
	VerticalAlignment VerticalAlignment `json:"VerticalAlignment,omitempty"`
	// An array of text lines.
	Lines []TextLine `json:"Lines"`
}
