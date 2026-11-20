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

// Represents a text state of a text
type TextState struct {
	// Gets or sets font size of the text.
	FontSize float64 `json:"FontSize"`
	// Gets or sets font name of the text.
	Font string `json:"Font,omitempty"`
	// Gets or sets foreground color of the text.
	ForegroundColor *Color `json:"ForegroundColor,omitempty"`
	// Sets background color of the text.
	BackgroundColor *Color `json:"BackgroundColor,omitempty"`
	// Sets font style of the text.
	FontStyle FontStyles `json:"FontStyle"`
	// Sets path of font file in storage.
	FontFile string `json:"FontFile,omitempty"`
	// Gets or sets underline of the text.
	Underline bool `json:"Underline,omitempty"`
	// Gets or sets strikeout of the text.
	StrikeOut bool `json:"StrikeOut,omitempty"`
	// Gets or sets superscript mode of the text.
	Superscript bool `json:"Superscript,omitempty"`
	// Gets or sets subscript mode of the text.
	Subscript bool `json:"Subscript,omitempty"`
}
