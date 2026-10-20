package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func replaceBookmark(pdf_api *asposepdfcloud.PdfApiService, document_name string, bookmark_path string, title string, remote_folder string) {
	uploadFile(pdf_api, document_name)

	args := map[string]interface{}{
		"folder": remote_folder,
	}

	bookmark := asposepdfcloud.Bookmark{
		Action:            "GoTo",
		Bold:              true,
		Italic:            false,
		Title:             title,
		PageDisplay:       "XYZ",
		PageDisplayBottom: 10,
		PageDisplayLeft:   10,
		PageDisplayRight:  10,
		PageDisplayTop:    10,
		PageDisplayZoom:   2,
		PageNumber:        1,
		Color:             &asposepdfcloud.Color{A: 0x00, R: 0x00, G: 0xFF, B: 0x00},
	}

	_, httpResponse, err := pdf_api.PutBookmark(document_name, bookmark_path, bookmark, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		downloadFile(pdf_api, document_name, "replaced_bookmark_")
	}
}
