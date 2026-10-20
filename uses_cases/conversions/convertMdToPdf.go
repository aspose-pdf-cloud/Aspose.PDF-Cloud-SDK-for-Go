package main

import (
	"fmt"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func convertMarkdownToPdf(pdf_api *asposepdfcloud.PdfApiService, md_document string, pdf_name string, remote_folder string) {
	uploadFile(pdf_api, md_document)

	args := map[string]interface{}{
		"folder": remote_folder,
	}

	src_path := path.Join(remote_folder, md_document)

	_, httpResponse, err := pdf_api.PutMarkdownInStorageToPdf(pdf_name, src_path, args)

	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		downloadFile(pdf_api, pdf_name, "")
	}
}
