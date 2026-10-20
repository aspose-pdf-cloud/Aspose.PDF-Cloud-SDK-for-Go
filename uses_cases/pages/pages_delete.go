package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func DeletePage(pdf_api *asposepdfcloud.PdfApiService, document string, pageNumber int32, outputDocument string, remoteFolder string) {
	// Delete page of the PDF document.
	uploadFile(pdf_api, document)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	_, httpResponse, err := pdf_api.DeletePage(document, pageNumber, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println("Successfully delete ", pageNumber, " page.")
		downloadFile(pdf_api, document, outputDocument, "delete_")
	}
}
