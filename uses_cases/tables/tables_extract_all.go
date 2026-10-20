package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func ExtractTables(pdf_api *asposepdfcloud.PdfApiService, document string, remoteFolder string) {
	// Extract all tables form the document and show table info
	uploadFile(pdf_api, document)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	result, httpResponse, err := pdf_api.GetDocumentTables(document, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("ExtractTables(): Failed to extract tablev from the document.")
	} else {
		if result.Tables == nil || len(result.Tables.List) == 0 {
			fmt.Println("ExtractTables(): Tables not found in the document.")
		} else {
			for i, t := range result.Tables.List {
				fmt.Println("table", i, " => id: '"+t.Id+"', page: '", t.PageNum, "', rows: '", len(t.RowList), "', columns: '", len(t.RowList[0].CellList), "'")
			}
		}
	}

}
