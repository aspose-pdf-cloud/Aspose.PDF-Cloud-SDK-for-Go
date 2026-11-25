package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

const (
	REMOTE_FOLDER   = "Your_Temp_Pdf_Cloud"
	LOCAL_FOLDER    = "test_data"
	PDF_DOCUMENT    = "sample.pdf"
	PDF_OUTPUT      = "output_pages.pdf"
	PNG_PAGE_OUTPUT = "output_sample_page.png"
	PAGE_NUMBER     = 2

	STAMP_TEXT = "NEW TEXT STAMP"

	IMAGE_STAMP_FILE   = "sample.png"
	IMAGE_STAMP_LLY    = 800
	IMAGE_STAMP_WIDTH  = 24
	IMAGE_STAMP_HEIGHT = 24

	API_CLIENT_ID     = "****************" // Your Application Client ID
	API_CLIENT_SECRET = "****************" // Your Application Client Secret
)

func initPdfApi(clientId string, clientSecret string) *asposepdfcloud.PdfApiService {
	ClientId := clientId
	ClientSecret := clientSecret

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
		_, httpResponse, err := pdf_api.UploadFile(path.Join(REMOTE_FOLDER, name), file, args)
		if err != nil {
			fmt.Println(err.Error())
		} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
			fmt.Println("Unexpected error!")
		} else {
			fmt.Println("File '" + name + " ' succeddfully uplouaded.")
		}
	}
}

// Download file from remote folder and save it locally with check errors
func downloadFile(pdf_api *asposepdfcloud.PdfApiService, name string, output_name string, prefix string) {
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	result_data, httpResponse, err := pdf_api.DownloadFile(path.Join(REMOTE_FOLDER, name), args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fileName := path.Join(LOCAL_FOLDER, prefix+output_name)
		f, _ := os.Create(fileName)
		_, _ = f.Write(result_data)
		fmt.Println("File '" + prefix + fileName + "' successfully downloaded.")
	}
}
