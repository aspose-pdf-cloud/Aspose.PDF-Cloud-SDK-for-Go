package main

import (
	"fmt"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func convertPdfToWord(pdf_api *asposepdfcloud.PdfApiService, pdf_name string, doc_name string, remote_folder string) {
	uploadFile(pdf_api, pdf_name)

	outPath := path.Join(remote_folder, doc_name)

	args := map[string]interface{}{
		"folder": remote_folder,
		"format": string(asposepdfcloud.DocFormatDocX),
	}

	_, httpResponse, err := pdf_api.PutPdfInStorageToDoc(pdf_name, outPath, args)

	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		downloadFile(pdf_api, outPath, doc_name)
	}
}
