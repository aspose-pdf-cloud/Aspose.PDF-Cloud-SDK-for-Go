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
	PDF_OUTPUT    = "output_pages.pdf"

	PDF_DOCUMENT_WITH_WATERMARKS = "sample_watermarks.pdf"

	PAGE_NUMBER = 2

	IMAGE_FILE    = "sample.png"
	IMAGE_OPACITY = 0.4
	IMAGE_ROTATE  = 45
	IMAGE_X       = 200
	IMAGE_Y       = 400
	IMAGE_WIDTH   = 64
	IMAGE_HEIGHT  = 64

	CLIENT_ID     = "************" // Your Application Client ID
	CLIENT_SECRET = "************" // Your Application Client Secret
)

var (
	IMAGE_IDS_DEL = []string{
		"GE5TKMBQGE4TWMRQGAWDIMBQFQZDSMJMGQ4TC",
		"GE5TGOJVGQYDWNJQGQWDOMRQFQ2TANJMG4ZDC",
		"GI5TKMBQGE4TWMRQGAWDIMBQFQZDSMJMGQ4TC",
		"GI5TGOJVGQYDWNJQGQWDOMRQFQ2TANJMG4ZDC",
		"GM5TKMBQGE4TWMRQGAWDIMBQFQZDSMJMGQ4TC",
		"GM5TGOJVGQYDWNJQGQWDOMRQFQ2TANJMG4ZDC",
		"GQ5TKMBQGE4TWMRQGAWDIMBQFQZDSMJMGQ4TC",
		"GQ5TGOJVGQYDWNJQGQWDOMRQFQ2TANJMG4ZDC",
	}
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
