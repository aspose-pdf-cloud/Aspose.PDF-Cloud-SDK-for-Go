package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v25"
)

const (
	REMOTE_FOLDER   = "Your_Temp_Pdf_Cloud"
	LOCAL_FOLDER    = "c:\\Samples"
	PDF_DOCUMENT    = "sample.pdf"
	PDF_OUTPUT_NAME = "output_sample.pdf"
	FIELD_NAME      = "Signature_1"
)

func initPdfApi() *asposepdfcloud.PdfApiService {
	// Initialize Credentials and create Pdf.Cloud service object
	AppSID := "YOUR_APP_SID" // Your Application SID
	AppKey := "YOUR_APP_KEY" // Your Application Key

	pdfApi := asposepdfcloud.NewPdfApiService(AppSID, AppKey, "")
	return pdfApi
}

func uploadFile(pdf_api *asposepdfcloud.PdfApiService, name string) {
	// Upload local file to the Pdf.Cloud folder
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	file, _ := os.Open(filepath.Join(LOCAL_FOLDER, name))
	fmt.Println("Download file'" + name + "' starting...")

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
	fmt.Println("Result file'" + fileName + "' successfully downloaded!")
}
