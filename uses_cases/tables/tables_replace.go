package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func ReplaceTable(pdf_api *asposepdfcloud.PdfApiService, document string, tableId string, outputDocument string, remoteFolder string) {
	// Replace table in the document
	uploadFile(pdf_api, document)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	var new_table = InitializeTable()

	_, httpResponse, err := pdf_api.PutTable(document, tableId, *new_table, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("ReplaceTable()): Failed to replace table in the document.")
	} else {
		fmt.Println("ReplaceTable(): table replaced successfully in the document '" + document + "'.")
		downloadFile(pdf_api, document, outputDocument, "replace_table_")
	}

}
