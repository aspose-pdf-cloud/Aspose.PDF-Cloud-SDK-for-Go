package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func ExtractPageInfo(pdf_api *asposepdfcloud.PdfApiService, document string, pageNumber int32, remoteFolder string) {
	// Get page information of the PDF document.
	uploadFile(pdf_api, document)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	result, httpResponse, err := pdf_api.GetPage(document, pageNumber, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println("Successfully extract page :", pageNumber, "Width :", result.Page.Rectangle.URX, "Height: ", result.Page.Rectangle.URY)

	}
}
