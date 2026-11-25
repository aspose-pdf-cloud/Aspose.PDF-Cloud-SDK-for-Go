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
	PDF_DOCUMENT    = "PdfWithTable.pdf"
	PDF_OUTPUT      = "output_pages.pdf"
	PNG_PAGE_OUTPUT = "output_sample_page.png"
	PAGE_NUMBER     = 1
	TABLE_ID        = "GE5TIMJ3HEYCYOBTFQ2TANZMG43TA"

	CLIENT_ID     = "*************" // Your Application Client ID
	CLIENT_SECRET = "*************" // Your Application Client Secret
)

func initPdfApi() *asposepdfcloud.PdfApiService {
	pdfApi := asposepdfcloud.NewPdfApiService(CLIENT_ID, CLIENT_SECRET, "")
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
