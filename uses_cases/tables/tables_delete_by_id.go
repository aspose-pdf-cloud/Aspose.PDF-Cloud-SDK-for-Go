package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func DeleteTableById(pdf_api *asposepdfcloud.PdfApiService, document string, tableId string, outputDocument string, remoteFolder string) {
	// Delete table from the document
	uploadFile(pdf_api, document)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	_, httpResponse, err := pdf_api.DeleteTable(document, tableId, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("DeleteTableById()): Failed to delete table from the document.")
	} else {
		fmt.Println("DeleteTableById(): table deleted successfully in the document '" + document + "'.")
		downloadFile(pdf_api, document, outputDocument, "del_table_")
	}

}
