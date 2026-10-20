package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func DeleteDocumentStamps(pdf_api *asposepdfcloud.PdfApiService, documentName string, outputDocument string, remoteFolder string) {
	uploadFile(pdf_api, documentName)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	_, httpResponse, err := pdf_api.DeleteDocumentStamps(documentName, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("DeleteDocumentStamps(): Failed to delete stamps in the document.")
	} else {
		fmt.Println("DeleteDocumentStamps(): all stamps successfully deleted in the document '" + documentName + "'.")
		downloadFile(pdf_api, documentName, outputDocument, "del_doc_stamps_")
	}
}
