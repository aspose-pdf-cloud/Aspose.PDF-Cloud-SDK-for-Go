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

// Represents graphics info.
type GraphInfo struct {
	// Gets or sets a float value that indicates the line width of the graph.
	LineWidth float64 `json:"LineWidth,omitempty"`
	// Gets or sets a Color object that indicates the color of the graph.
	Color *Color `json:"Color,omitempty"`
	// Gets or sets a dash array.
	DashArray []int32 `json:"DashArray,omitempty"`
	// Gets or sets a dash phase.
	DashPhase int32 `json:"DashPhase,omitempty"`
	// Gets or sets a Color object that indicates the fill color of the graph.
	FillColor *Color `json:"FillColor,omitempty"`
	// Gets or sets is border doubled.
	IsDoubled bool `json:"IsDoubled,omitempty"`
	// Gets or sets a float value that indicates the skew angle of the x-coordinate when transforming a coordinate system.
	SkewAngleX float64 `json:"SkewAngleX,omitempty"`
	// Gets or sets a float value that indicates the skew angle of the y-coordinate when transforming a coordinate system.
	SkewAngleY float64 `json:"SkewAngleY,omitempty"`
	// Gets or sets a float value that indicates the scaling rate of the x-coordinate when transforming a coordinate system.
	ScalingRateX float64 `json:"ScalingRateX,omitempty"`
	// Gets or sets a float value that indicates the scaling rate of the y-coordinate when transforming a coordinate system.
	ScalingRateY float64 `json:"ScalingRateY,omitempty"`
	// Gets or sets a float value that indicates the rotation angle of the coordinate system  when transforming a coordinate system.
	RotationAngle float64 `json:"RotationAngle,omitempty"`
}
