package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func AppendNewPage(pdf_api *asposepdfcloud.PdfApiService, document string, outputDocument string, remoteFolder string) {
	// Append page to the PDF document.
	uploadFile(pdf_api, document)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	result, httpResponse, err := pdf_api.PutAddNewPage(document, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println("Successfully added new page. ", result.Pages.List)
		downloadFile(pdf_api, document, outputDocument, "append_")
	}
}
