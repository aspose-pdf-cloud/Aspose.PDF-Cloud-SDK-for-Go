package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

const (
	REMOTE_FOLDER = "Your_Temp_Pdf_Cloud"
	LOCAL_FOLDER  = "test_data"
	PDF_DOCUMENT  = "PageNumberStamp.pdf"
	PDF_OUTPUT    = "output_sample.pdf"

	IMAGE_STAMP_FILE   = "sample.png"
	PAGE_NUMBER        = 2
	STAMP_TEXT         = "NEW TEXT STAMP"
	IMAGE_STAMP_LLY    = 800
	IMAGE_STAMP_WIDTH  = 24
	IMAGE_STAMP_HEIGHT = 24

	CLIENT_ID     = "****************" // Your Application Client ID
	CLIENT_SECRET = "****************" // Your Application Client Secret
)

func initPdfApi() *asposepdfcloud.PdfApiService {
	// Initialize Credentials and create Pdf.Cloud service object
	pdfApi := asposepdfcloud.NewPdfApiService(CLIENT_ID, CLIENT_SECRET, "")
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
func downloadFile(pdf_api *asposepdfcloud.PdfApiService, name string, output_name string, output_prefix string) {
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	result_data, _, _ := pdf_api.DownloadFile(path.Join(REMOTE_FOLDER, name), args)
	fileName := path.Join(LOCAL_FOLDER, output_prefix+output_name)
	f, _ := os.Create(fileName)
	_, _ = f.Write(result_data)
	fmt.Println("Result file'" + (output_prefix + output_name) + "' successfully downloaded!")
}
