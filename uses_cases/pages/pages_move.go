package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func MovePage(pdf_api *asposepdfcloud.PdfApiService, document string, pageNumber int32, newPageNumber int32, outputDocument string, remoteFolder string) {
	// Move page to new position in the PDF document.
	uploadFile(pdf_api, document)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	_, httpResponse, err := pdf_api.PostMovePage(document, pageNumber, newPageNumber, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println("Successfully moe page fron ", pageNumber, "to", newPageNumber, "position.")
		downloadFile(pdf_api, document, outputDocument, "move_")
	}
}
