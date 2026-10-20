package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func GetPagesWordsCount(pdf_api *asposepdfcloud.PdfApiService, document string, remoteFolder string) {
	// Get page information of the PDF document.
	uploadFile(pdf_api, document)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	result, httpResponse, err := pdf_api.GetWordsPerPage(document, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		for _, p := range result.WordsPerPage.List {
			fmt.Println("Page ", p.PageNumber, " has ", p.Count, " words.")
		}
	}
}
