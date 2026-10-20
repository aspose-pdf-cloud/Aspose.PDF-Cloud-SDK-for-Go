package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func removeLink(pdf_api *asposepdfcloud.PdfApiService, document string, output_document string, link_id string, remote_folder string) {
	uploadFile(pdf_api, document)

	getLink(pdf_api, document, link_id, remote_folder)

	args := map[string]interface{}{
		"folder": remote_folder,
	}

	result, httpResponse, err := pdf_api.DeleteLinkAnnotation(document, link_id, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println(result)
		downloadFile(pdf_api, document, output_document)
	}
}
