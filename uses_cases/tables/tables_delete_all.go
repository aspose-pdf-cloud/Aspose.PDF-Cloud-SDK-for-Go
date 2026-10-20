package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func DeleteTables(pdf_api *asposepdfcloud.PdfApiService, document string, outputDocument string, remoteFolder string) {
	// Delete tables from the document
	uploadFile(pdf_api, document)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	_, httpResponse, err := pdf_api.DeleteDocumentTables(document, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("DeleteTables()): Failed to delete tables from the document.")
	} else {
		fmt.Println("DeleteTables(): table deleted successfully from the document '" + document + "'.")
		downloadFile(pdf_api, document, outputDocument, "del_tables_")
	}

}
