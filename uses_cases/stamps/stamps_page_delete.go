package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func DeletePageStamps(pdf_api *asposepdfcloud.PdfApiService, documentName string, pageNumber int32, outputDocument string, remoteFolder string) {
	uploadFile(pdf_api, documentName)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	_, httpResponse, err := pdf_api.DeletePageStamps(documentName, pageNumber, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("DeletDeletePageStampseStampText(): Failed to delete stamps in the document.")
	} else {
		fmt.Println("DeletePageStamps(): stamps successfully deleted on ", pageNumber, "page in the document '"+documentName+"'.")
		downloadFile(pdf_api, documentName, outputDocument, "del_page_stamps_")
	}
}
