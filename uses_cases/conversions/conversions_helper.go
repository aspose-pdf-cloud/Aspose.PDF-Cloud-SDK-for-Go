package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

// Set file names and folders
const (
	REMOTE_FOLDER = "Your_Temp_Pdf_Cloud"
	LOCAL_FOLDER  = "test_data"
	PDF_DOCUMENT  = "sample.pdf"

	JPEG_FILE = "Koala.jpg"
	BMP_FILE  = "sample.bmp"
	PNG_FILE  = "sample.png"
	GIF_FILE  = "sample.gif"

	PDF_BMP_OUTPUT  = "output_bmp.pdf"
	PDF_JPEG_OUTPUT = "output_jpeg.pdf"
	PDF_PNG_OUTPUT  = "output_png.pdf"
	PDF_GIF_OUTPUT  = "output_gif.pdf"

	HTML_FILE_NAME  = "HtmlExample1.html"
	HTML_ZIP_FILE   = "HtmlWithImage.zip"
	PDF_HTML_OUTPUT = "output_html.pdf"

	MHTML_FILE_NAME  = "MhtExample.mht"
	PDF_MHTML_OUTPUT = "output_mht.pdf"

	MD_FILE_NAME  = "sample.md"
	PDF_MD_OUTPUT = "output_md.pdf"

	EPUB_FILE_NAME  = "4pages.epub"
	PDF_EPUB_OUTPUT = "output_epub.pdf"

	PCL_FILE_NAME  = "template.pcl"
	PDF_PCL_OUTPUT = "output_pcl.pdf"

	PS_FILE_NAME  = "sample.ps"
	PDF_PS_OUTPUT = "output_ps.pdf"

	PDF_OUTPUT        = "output_sample.pdf"
	PDF_IMAGES_OUTPUT = "output_images.pdf"
	DOC_OUTPUT        = "output_sample.doc"
	DOCX_OUTPUT       = "output_sample.docx"
	EPUB_OUTPUT       = "output_sample.epub"
	HTML_ZIP_OUTPUT   = "output_sample_html.zip"
	TIFF_OUTPUT       = "output_sample.tiff"
	SVG_ZIP_OUTPUT    = "output_sample_svg.zip"
	XLS_OUTPUT        = "output_sample.xls"
	XLSX_OUTPUT       = "output_sample.xlsx"
	PPTX_OUTPUT       = "output_sample.pptx"
	XML_OUTPUT        = "output_sample.xml"
	MXML_OUTPUT       = "output_sample.mobi"
	TEX_OUTPUT        = "output_sample.tex"
	XPS_OUTPUT        = "output_sample.xps"
)

// Initialize credentials and Rest API service
func initPdfApi() *asposepdfcloud.PdfApiService {
	ClientId := "*****"
	ClientSecret := "*****"

	pdfApi := asposepdfcloud.NewPdfApiService(ClientId, ClientSecret, "")
	return pdfApi
}

// Upload local file to the remote folder with check errors
func uploadFile(pdf_api *asposepdfcloud.PdfApiService, name string) {
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	file, err := os.Open(filepath.Join(LOCAL_FOLDER, name))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		result, httpResponse, err := pdf_api.UploadFile(path.Join(REMOTE_FOLDER, name), file, args)
		if err != nil {
			fmt.Println(err.Error())
		} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
			fmt.Println("Unexpected error!")
		} else {
			fmt.Println(result)
		}
	}
}

// Download file from remote folder and save it locally with check errors
func downloadFile(pdf_api *asposepdfcloud.PdfApiService, name string, output_name string) {
	if len(strings.TrimSpace(output_name)) == 0 {
		output_name = name
	}
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	result_data, httpResponse, err := pdf_api.DownloadFile(name, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fileName := path.Join(LOCAL_FOLDER, output_name)
		f, _ := os.Create(fileName)
		_, _ = f.Write(result_data)
		fmt.Println("File '" + output_name + "'successfully downloaded.")
	}
}
