package main

import (
	"fmt"
	"os"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func createPdfSimple() {
	LOCAL_FOLDER := "test_data"
	PDF_DOCUMENT := "output_created_simple.pdf"

	ClientId := "*******"
	ClientSecret := "*******"

	pdfApi := asposepdfcloud.NewPdfApiService(ClientId, ClientSecret, "")
	_, httpResponse, err := pdfApi.PutCreateDocument(PDF_DOCUMENT, nil)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		result_data, _, _ := pdfApi.DownloadFile(PDF_DOCUMENT, nil)
		fileName := path.Join(LOCAL_FOLDER, PDF_DOCUMENT)
		f, _ := os.Create(fileName)
		_, _ = f.Write(result_data)
	}
}
