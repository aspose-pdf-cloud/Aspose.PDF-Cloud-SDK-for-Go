package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func getLinks(pdf_api *asposepdfcloud.PdfApiService, document string, page_num int32, remote_folder string) {
	uploadFile(pdf_api, document)

	args := map[string]interface{}{
		"folder": remote_folder,
	}

	result, httpResponse, err := pdf_api.GetPageLinkAnnotations(document, page_num, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		showLinks(&result.Links.List)
	}
}
