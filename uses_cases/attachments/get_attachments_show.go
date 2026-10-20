package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func getDocumentAttachments(pdf_api *asposepdfcloud.PdfApiService, document_name string, remote_folder string) {
	args := map[string]interface{}{
		"folder": remote_folder,
	}
	result, httpResponse, err := pdf_api.GetDocumentAttachments(document_name, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println(result.Attachments.Links)
	}
}
