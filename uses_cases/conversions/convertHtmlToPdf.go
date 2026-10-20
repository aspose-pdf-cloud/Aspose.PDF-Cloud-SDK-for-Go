package main

import (
	"fmt"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func convertHtmlToPdf(pdf_api *asposepdfcloud.PdfApiService, html_document string, html_zip string, pdf_name string, remote_folder string) {
	uploadFile(pdf_api, html_document)

	args := map[string]interface{}{
		"folder":       remote_folder,
		"height":       float64(850),
		"width":        float64(600),
		"htmlFileName": html_document,
	}

	src_path := path.Join(remote_folder, html_document)

	_, httpResponse, err := pdf_api.PutHtmlInStorageToPdf(pdf_name, src_path, args)

	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		downloadFile(pdf_api, pdf_name, "")
	}
}
