package main

import (
	"fmt"
	"os"
	"path"
)

func createPdfSimple() {
	LOCAL_FOLDER := "test_data"
	PDF_DOCUMENT := "output_created_simple.pdf"

	pdfApi := initPdfApi()
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
