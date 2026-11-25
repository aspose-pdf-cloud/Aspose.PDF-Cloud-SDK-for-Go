package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

const (
	REMOTE_FOLDER = "Your_Temp_Pdf_Cloud"
	LOCAL_FOLDER  = "test_data"
	PDF_DOCUMENT  = "PdfWithScreenAnnotations.pdf"
	PDF_OUTPUT    = "output_sample.pdf"

	ROTATE_PAGES_ANGLE = asposepdfcloud.Rotationon90
	ROTATE_PAGES       = "1-2"

	CROP_PAGE_TEMP_FILE             = "sample_temp_file.png"
	CROP_LOCAL_RESULT_DOCUMENT_NAME = "output_sample.pdf"
	CROP_PAGE_NUMBER                = 1
	CROP_HEIGHT                     = 400
	CROP_WIDTH                      = 300
	CROP_LLX                        = 100
	CROP_LLY                        = 200

	RESIZE_PDF_HTML_FILE        = "sample_temp_file.html"
	RESIZE_RESULT_DOCUMENT_NAME = "output_sample.pdf"
	RESIZE_PAGE_NUMBER          = 2
	RESIZE_NEW_PAGE_WIDTH       = 800
	RESIZE_NEW_PAGE_HEIGHT      = 400
)

var (
	CROP_PAGE_WIDTH  = float64(0)
	CROP_PAGE_HEIGHT = float64(0)
)

func initPdfApi() *asposepdfcloud.PdfApiService {
	// Initialize Credentials and create Pdf.Cloud service object
	ClientId := "******"     // Your Application Client ID
	ClientSecret := "******" // Your Application Client Secret

	pdfApi := asposepdfcloud.NewPdfApiService(ClientId, ClientSecret, "")
	return pdfApi
}

func uploadFile(pdf_api *asposepdfcloud.PdfApiService, name string) {
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	file, _ := os.Open(filepath.Join(LOCAL_FOLDER, name))
	_, _, _ = pdf_api.UploadFile(filepath.Join(REMOTE_FOLDER, name), file, args)
	fmt.Println("File '" + name + "' successfully uploaded!")
}
func downloadFile(pdf_api *asposepdfcloud.PdfApiService, name string, output_prefix string) {
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	result_data, _, _ := pdf_api.DownloadFile(path.Join(REMOTE_FOLDER, name), args)
	fileName := path.Join(LOCAL_FOLDER, output_prefix+PDF_OUTPUT)
	f, _ := os.Create(fileName)
	_, _ = f.Write(result_data)
	fmt.Println("Result file'" + fileName + "' successfully downloaded!")
}
func getPageInfo(pdf_api *asposepdfcloud.PdfApiService, document_name string, pageNumber int, tempFolder string) {
	args := map[string]interface{}{
		"folder": tempFolder,
	}
	page, response, err := pdf_api.GetPage(document_name, int32(pageNumber), args)

	if response.StatusCode == 200 {
		showPages([]*asposepdfcloud.Page{page.Page}, "page")
		CROP_PAGE_HEIGHT = page.Page.Rectangle.URY - page.Page.Rectangle.LLY
		CROP_PAGE_WIDTH = page.Page.Rectangle.URX - page.Page.Rectangle.LLX
	} else {
		fmt.Println("Unexpected error : can't get pages!!! Error: " + err.Error())
	}
}

func showPages(pages []*asposepdfcloud.Page, prefix string) {
	if len(pages) > 0 {
		for _, page := range pages {
			fmt.Println(prefix + " => id: '" + strconv.FormatInt(int64(page.Id), 10) + "', lLx: '" + strconv.FormatFloat(page.Rectangle.LLX, 'f', -1, 64) + "', lLY: '" + strconv.FormatFloat(page.Rectangle.LLY, 'f', -1, 64) + "', uRX: '" + strconv.FormatFloat(page.Rectangle.URX, 'f', -1, 64) + "', uRY: '" + strconv.FormatFloat(page.Rectangle.URY, 'f', -1, 64) + "'")
		}
	} else {
		fmt.Println("showPages() error: array of pages is empty!")
	}
}

func extractPdfPage(pdf_api *asposepdfcloud.PdfApiService, document_name string, pageNumber int, width int, height int, localFolder string, tempFolder string) string {
	args := map[string]interface{}{
		"folder": tempFolder,
		"width":  int32(width),
		"height": int32(height),
	}

	image, response, err := pdf_api.GetPageConvertToPng(document_name, int32(pageNumber), args)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	} else if response.StatusCode != 200 {
		fmt.Println("extractPdfPage(): Faild to convert page to image!")
		return ""
	}
	imageFile := document_name + ".png"

	filePath := filepath.Join(localFolder, imageFile)
	os.WriteFile(filePath, image, 0777)

	uploadFile(pdf_api, imageFile)

	fmt.Println("Page #" + strconv.FormatInt(int64(pageNumber), 10) + " extracted as image.")
	return imageFile
}

func createPdfDocument(pdf_api *asposepdfcloud.PdfApiService, document_name string, width int, height int, tempFolder string) *asposepdfcloud.DocumentResponse {
	args := map[string]interface{}{
		"folder": tempFolder,
	}

	config := asposepdfcloud.DocumentConfig{
		PagesCount: 1,
		DocumentProperties: &asposepdfcloud.DocumentProperties{
			List: []asposepdfcloud.DocumentProperty{
				{
					BuiltIn: false,
					Name:    "prop1",
					Value:   "Val1",
				},
			},
		},
		DisplayProperties: &asposepdfcloud.DisplayProperties{
			CenterWindow: true,
			HideMenuBar:  true,
		},
		DefaultPageConfig: &asposepdfcloud.DefaultPageConfig{
			Height: float64(height),
			Width:  float64(width),
		},
	}

	result, httpResponse, err := pdf_api.PostCreateDocument(document_name, config, args)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error on create new document!")
		return nil
	} else {
		return &result
	}
}

func insertPageAsImage(pdf_api *asposepdfcloud.PdfApiService, document_name string, imageFileValue string, llx int, lly int, tempFolder string) *asposepdfcloud.AsposeResponse {
	args := map[string]interface{}{
		"folder": tempFolder,
	}

	stamp := asposepdfcloud.ImageStamp{
		Background:          true,
		HorizontalAlignment: asposepdfcloud.HorizontalAlignmentNone,
		VerticalAlignment:   asposepdfcloud.VerticalAlignmentNone,
		Opacity:             1,
		Rotate:              asposepdfcloud.RotationNone,
		RotateAngle:         0,
		XIndent:             -float64(llx),
		YIndent:             -float64(lly),
		Zoom:                1,
		FileName:            path.Join(tempFolder, imageFileValue),
	}
	stamps := []asposepdfcloud.ImageStamp{stamp}

	result, response, err := pdf_api.PostPageImageStamps(document_name, 1, stamps, args)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	} else if response.StatusCode < 200 || response.StatusCode > 299 {
		fmt.Println("Unexpected error on inserting image to the document!")
		return nil
	} else {
		return &result
	}
}
