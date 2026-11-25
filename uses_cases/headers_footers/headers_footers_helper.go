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
	PDF_DOCUMENT  = "sample.pdf"
	PDF_OUTPUT    = "output_sample.pdf"
	IMAGE_FILE    = "sample.png"
	TEXT_FOOTER   = "NEW ADDED TEXT FOOTER"
	TEXT_HEADER   = "NEW ADDED TEXT HEADER"
	PAGE_NUMBER   = 1
)

func initPdfApi() *asposepdfcloud.PdfApiService {
	ClientId := "******"
	ClientSecret := "******"

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
func downloadFile(pdf_api *asposepdfcloud.PdfApiService, name string) {
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	result_data, httpResponse, err := pdf_api.DownloadFile(path.Join(REMOTE_FOLDER, name), args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println("File '" + name + "'successfully downloaded.")
		fileName := path.Join(LOCAL_FOLDER, PDF_OUTPUT)
		f, _ := os.Create(fileName)
		_, _ = f.Write(result_data)
	}
}
