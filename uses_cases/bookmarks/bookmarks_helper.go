package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

const (
	REMOTE_FOLDER  = "Your_Temp_Pdf_Cloud"
	LOCAL_FOLDER   = "test_data"
	PDF_DOCUMENT   = "PdfWithBookmarks.pdf"
	PDF_OUTPUT     = "output_sample.pdf"
	BOOKMARK_TITLE = "NEW Bookmark Title XYZ"
	BOOKMARK_PATH  = "/5"
)

func initPdfApi() *asposepdfcloud.PdfApiService {
	// Initialize Credentials and create Pdf.Cloud service object
	content, err := os.ReadFile("settings/credentials.json")
	if err != nil {
		panic(err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(content, &m); err != nil {
		panic(err)
	}
	pdfApi := asposepdfcloud.NewPdfApiService(m["client_id"].(string), m["client_secret"].(string), "")
	return pdfApi
}

func uploadFile(pdf_api *asposepdfcloud.PdfApiService, name string) {
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	file, _ := os.Open(filepath.Join(LOCAL_FOLDER, name))
	_, _, _ = pdf_api.UploadFile(filepath.Join(REMOTE_FOLDER, name), file, args)
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
