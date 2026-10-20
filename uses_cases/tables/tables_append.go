package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func AppendTable(pdf_api *asposepdfcloud.PdfApiService, document string, pageNumber int32, outputDocument string, remoteFolder string) {
	// Append table to the document
	uploadFile(pdf_api, document)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	var new_table = InitializeTable()

	_, httpResponse, err := pdf_api.PostPageTables(document, pageNumber, []asposepdfcloud.Table{*new_table}, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("AppendTable()): Failed to append table to the document.")
	} else {
		fmt.Println("AppendTable(): table appended successfully on page", pageNumber, " to the document '"+document+"'.")
		downloadFile(pdf_api, document, outputDocument, "add_table_")
	}

}
