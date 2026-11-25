package main

import (
	"fmt"
	"os"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

const (
	REMOTE_FOLDER = "Your_Temp_Pdf_Cloud"
	LOCAL_FOLDER  = "test_data"
	PDF_DOCUMENT  = "output_created_config.pdf"
	PAGE_WIDTH    = 590
	PAGE_HEIGHT   = 894
	PAGES_COUNT   = 5
)

func initPdfApi() *asposepdfcloud.PdfApiService {
	ClientId := "******"
	ClientSecret := "******"

	pdfApi := asposepdfcloud.NewPdfApiService(ClientId, ClientSecret, "")
	return pdfApi
}

func downloadFile(pdf_api *asposepdfcloud.PdfApiService, name string) {
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	result_data, _, _ := pdf_api.DownloadFile(path.Join(REMOTE_FOLDER, name), args)
	fileName := path.Join(LOCAL_FOLDER, PDF_DOCUMENT)
	f, _ := os.Create(fileName)
	_, _ = f.Write(result_data)
}
func createPdf(pdf_api *asposepdfcloud.PdfApiService, pdf_name string, remote_folder string) {
	args := map[string]interface{}{
		"folder": remote_folder,
	}

	config := asposepdfcloud.DocumentConfig{
		PagesCount: PAGES_COUNT,
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
			Height: PAGE_HEIGHT,
			Width:  PAGE_WIDTH,
		},
	}

	result, httpResponse, err := pdf_api.PostCreateDocument(pdf_name, config, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println(result)
	}
}

func main() {
	pdfApi := initPdfApi()
	createPdf(pdfApi, PDF_DOCUMENT, REMOTE_FOLDER)
	downloadFile(pdfApi, PDF_DOCUMENT)

	createPdfSimple()
}
