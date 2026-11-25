package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

const (
	REMOTE_FOLDER             = "Your_Temp_Pdf_Cloud"
	LOCAL_FOLDER              = "test_data"
	PDF_DOCUMENT              = "sample.pdf"
	PDF_DOCUMENT_WITH_ATTACH  = "PdfWithEmbeddedFiles.pdf"
	PDF_OUTPUT                = "output_sample.pdf"
	NEW_ATTACHMENT_FILE       = "sample.pdf"
	NEW_ATTACHMENT_MIME       = "application/pdf"
	NEW_ATTACHMENT_DECRIPTION = "This is a sample attachment"
	PAGE_NUMBER               = 2
)

func initPdfApi() *asposepdfcloud.PdfApiService {
	// Initialize Credentials and create Pdf.Cloud service object
	ClientId := "******" // Your Application Client ID
	ClientSecret := "******" // Your Application Client Secret

	pdfApi := asposepdfcloud.NewPdfApiService(ClientId, ClientSecret, "")
	return pdfApi
}

func uploadFile(pdf_api *asposepdfcloud.PdfApiService, name string) {
	// Upload local file to the Pdf.Cloud folder
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}

	path := filepath.Join(LOCAL_FOLDER, name)
	abs, err := filepath.Abs(path)
	println(abs)
	file, err := os.Open(abs)
	if err != nil {

	}
	res, a, b := pdf_api.UploadFile(filepath.Join(REMOTE_FOLDER, name), file, args)
	if res.Errors != nil && a != nil && b != nil {

	}
}

func saveByteArrayToFile(local_folder string, file_name string, data []byte) {
	// Save byte array data to the local folder
	fileName := path.Join(local_folder, file_name)
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		size, err := f.Write(data)
		if err != nil || size == 0 {
			fmt.Println("Failures in downloading result!")
		} else {
			fmt.Println("Result file'" + fileName + "' successfully downloaded!")
		}
	}
}

func downloadFile(pdf_api *asposepdfcloud.PdfApiService, name string) {
	// Download modified Pdf document to local folder from the Pdf.Cloud folder
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	result_data, _, _ := pdf_api.DownloadFile(path.Join(REMOTE_FOLDER, name), args)
	saveByteArrayToFile(LOCAL_FOLDER, PDF_OUTPUT, result_data)
}
