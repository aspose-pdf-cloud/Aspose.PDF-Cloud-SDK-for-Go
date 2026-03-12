/**
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

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// To DOC
func TestGetPdfInStorageToDoc(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPdfInStorageToDoc(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPdfInStorageToDoc - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPdfInStorageToDoc(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	resFileName := "result.doc"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInStorageToDoc(name, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInStorageToDoc - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutPdfInRequestToDoc(t *testing.T) {
	name := "4pages.pdf"
	resFileName := "result.doc"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name)
	if err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
		"file":   file,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInRequestToDoc(outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInRequestToDoc - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// To PDFA
func TestGetPdfInStorageToPdfA(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPdfInStorageToPdfA(name, string(PdfATypePDFA1A), args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPdfInStorageToPdfA - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPdfInStorageToPdfA(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	resFileName := "result.pdf"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInStorageToPdfA(name, outPath, string(PdfATypePDFA1A), args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInStorageToPdfA - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutPdfInRequestToPdfA(t *testing.T) {
	name := "4pages.pdf"
	resFileName := "result.pdf"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name)
	if err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
		"file":   file,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInRequestToPdfA(outPath, string(PdfATypePDFA1A), args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInRequestToPdfA - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// To TIFF
func TestGetPdfInStorageToTiff(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPdfInStorageToTiff(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPdfInStorageToTiff - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPdfInStorageToTiff(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	resFileName := "result.tiff"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInStorageToTiff(name, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInStorageToTiff - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutPdfInRequestToTiff(t *testing.T) {
	name := "4pages.pdf"
	resFileName := "result.tiff"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name)
	if err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
		"file":   file,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInRequestToTiff(outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInRequestToTiff - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// To SVG
func TestGetPdfInStorageToSvg(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPdfInStorageToSvg(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPdfInStorageToSvg - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPdfInStorageToSvg(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	resFileName := "result.svg"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInStorageToSvg(name, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInStorageToSvg - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutPdfInRequestToSvg(t *testing.T) {
	name := "4pages.pdf"
	resFileName := "result.svg"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name)
	if err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
		"file":   file,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInRequestToSvg(outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInRequestToSvg - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// To XPS
func TestGetPdfInStorageToXps(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPdfInStorageToXps(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPdfInStorageToXps - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPdfInStorageToXps(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	resFileName := "result.xps"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInStorageToXps(name, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInStorageToXps - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutPdfInRequestToXps(t *testing.T) {
	name := "4pages.pdf"
	resFileName := "result.xps"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name)
	if err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
		"file":   file,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInRequestToXps(outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInRequestToXps - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// To XLS
func TestGetPdfInStorageToXls(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPdfInStorageToXls(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPdfInStorageToXls - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPdfInStorageToXls(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	resFileName := "result.xls"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInStorageToXls(name, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInStorageToXls - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutPdfInRequestToXls(t *testing.T) {
	name := "4pages.pdf"
	resFileName := "result.xls"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name)
	if err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
		"file":   file,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInRequestToXls(outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInRequestToXls - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// To XLS
func TestGetPdfInStorageToXlsx(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPdfInStorageToXlsx(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPdfInStorageToXlsx - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPdfInStorageToXlsx(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	resFileName := "result.xlsx"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInStorageToXlsx(name, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInStorageToXlsx - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutPdfInRequestToXlsx(t *testing.T) {
	name := "4pages.pdf"
	resFileName := "result.xlsx"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name)
	if err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
		"file":   file,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInRequestToXlsx(outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInRequestToXlsx - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostPdfToXlsx(t *testing.T) {
	name := "4pages.pdf"
	file, err := os.Open(filepath.Join(GetBaseTest().localTestDataFolder, name))
	if err != nil {
		t.Error(err)
	}
	args := map[string]interface{}{
		"file": file,
	}
	response, httpResponse, err := GetBaseTest().PdfAPI.PostPdfToXlsx(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostPdfToXlsx - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

// To HTML
func TestGetPdfInStorageToHtml(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPdfInStorageToHtml(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPdfInStorageToHtml - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPdfInStorageToHtml(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	resFileName := "result.html"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	args := map[string]interface{}{
		"folder":       GetBaseTest().remoteFolder,
		"outputFormat": fmt.Sprintf("%v", OutputFormatFolder),
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInStorageToHtml(name, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInStorageToHtml - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutPdfInRequestToHtml(t *testing.T) {
	name := "4pages.pdf"
	resFileName := "result.zip"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name)
	if err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
		"file":   file,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInRequestToHtml(outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInRequestToHtml - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// To EPUB
func TestGetPdfInStorageToEpub(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPdfInStorageToEpub(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPdfInStorageToEpub - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPdfInStorageToEpub(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	resFileName := "result.epub"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInStorageToEpub(name, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInStorageToEpub - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutPdfInRequestToEpub(t *testing.T) {
	name := "4pages.pdf"
	resFileName := "result.epub"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name)
	if err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
		"file":   file,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInRequestToEpub(outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInRequestToEpub - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// To PPTX
func TestGetPdfInStorageToPptx(t *testing.T) {
	name := "4pages.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}
	response, httpResponse, err := GetBaseTest().PdfAPI.GetPdfInStorageToPptx(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPdfInStorageToPptx - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestGetPdfInStorageToPptxWithPassword(t *testing.T) {
	name := "4pagesEncrypted.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	args := map[string]interface{}{
		"folder":   GetBaseTest().remoteFolder,
		"password": base64.StdEncoding.EncodeToString([]byte("user $^Password!&")),
	}
	response, httpResponse, err := GetBaseTest().PdfAPI.GetPdfInStorageToPptx(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPdfInStorageToPptxWithPassword - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPdfInStorageToPptx(t *testing.T) {
	name := "4pages.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	resFileName := "result.pptx"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}
	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInStorageToPptx(name, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInStorageToPptx - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutPdfInStorageToPptxWithPassword(t *testing.T) {
	name := "4pagesEncrypted.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	resFileName := "result.pptx"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName
	args := map[string]interface{}{
		"folder":   GetBaseTest().remoteFolder,
		"password": base64.StdEncoding.EncodeToString([]byte("user $^Password!&")),
	}
	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInStorageToPptx(name, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInStorageToPptxWithPassword - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutPdfInRequestToPptx(t *testing.T) {
	name := "4pages.pdf"
	resFileName := "result.pptx"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName
	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name)
	if err != nil {
		t.Error(err)
	}
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
		"file":   file,
	}
	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInRequestToPptx(outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInRequestToPptx - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutPdfInRequestToPptxWithPassword(t *testing.T) {
	name := "4pagesEncrypted.pdf"
	resFileName := "result.pptx"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName
	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name)
	if err != nil {
		t.Error(err)
	}
	args := map[string]interface{}{
		"folder":   GetBaseTest().remoteFolder,
		"file":     file,
		"password": base64.StdEncoding.EncodeToString([]byte("user $^Password!&")),
	}
	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInRequestToPptx(outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInRequestToPptxWithPassword - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// To TeX
func TestGetPdfInStorageToTeX(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPdfInStorageToTeX(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPdfInStorageToTeX - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPdfInStorageToTeX(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	resFileName := "result.tex"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInStorageToTeX(name, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInStorageToTeX - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutPdfInRequestToTeX(t *testing.T) {
	name := "4pages.pdf"
	resFileName := "result.tex"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name)
	if err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
		"file":   file,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInRequestToTeX(outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInRequestToTeX - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// To  MOBI XML
func TestGetPdfInStorageToMobiXml(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPdfInStorageToMobiXml(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPdfInStorageToMobiXml - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPdfInStorageToMobiXml(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	resFileName := "result.mobi"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInStorageToMobiXml(name, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInStorageToMobiXml - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutPdfInRequestToMobiXml(t *testing.T) {
	name := "4pages.pdf"
	resFileName := "result.mobi"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name)
	if err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
		"file":   file,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInRequestToMobiXml(outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInRequestToMobiXml - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// To Aps
func TestGetPdfInStorageToMobiAps(t *testing.T) {

	name := "5pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPdfInStorageToAps(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPdfInStorageToAps - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPdfInStorageToAps(t *testing.T) {

	name := "5pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	resFileName := "result.xml"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInStorageToAps(name, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInStorageToAps - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutPdfInRequestToAps(t *testing.T) {
	name := "5pages.pdf"
	resFileName := "result.xml"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name)
	if err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
		"file":   file,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInRequestToAps(outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInRequestToAps - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// PDF XFA To PDF ACRO FORM
func TestGetXfaPdfInStorageToAcroForm(t *testing.T) {

	name := "PdfWithXfaForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetXfaPdfInStorageToAcroForm(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetXfaPdfInStorageToAcroForm - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutXfaPdfInStorageToAcroForm(t *testing.T) {

	name := "PdfWithXfaForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	resFileName := "result.pdf"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutXfaPdfInStorageToAcroForm(name, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutXfaPdfInStorageToAcroForm - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutXfaPdfInRequestToAcroForm(t *testing.T) {
	name := "PdfWithXfaForm.pdf"
	resFileName := "result.pdf"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name)
	if err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
		"file":   file,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutXfaPdfInRequestToAcroForm(outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutXfaPdfInRequestToAcroForm - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// To  XML
func TestGetPdfInStorageToXml(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPdfInStorageToXml(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPdfInStorageToXml - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPdfInStorageToXml(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	resFileName := "result.xml"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInStorageToXml(name, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInStorageToXml - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutPdfInRequestToXml(t *testing.T) {
	name := "4pages.pdf"
	resFileName := "result.xml"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName

	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name)
	if err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
		"file":   file,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInRequestToXml(outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInRequestToXml - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// To Text
func TestGetPdfInStorageToText(t *testing.T) {
	name := "4pages.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}
	response, httpResponse, err := GetBaseTest().PdfAPI.GetPdfInStorageToText(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tGetPdfInStorageToText - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}
