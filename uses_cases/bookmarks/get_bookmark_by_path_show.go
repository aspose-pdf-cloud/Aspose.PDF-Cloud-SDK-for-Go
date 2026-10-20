package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func extractBookmark(pdf_api *asposepdfcloud.PdfApiService, document_name string, bookmark_path string, remote_folder string) {
	uploadFile(pdf_api, document_name)

	args := map[string]interface{}{
		"folder": remote_folder,
	}
	result, httpResponse, err := pdf_api.GetBookmarks(document_name, bookmark_path, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println(result.Bookmarks)
	}
}
