package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func ExtractTablesOnPage(pdf_api *asposepdfcloud.PdfApiService, document string, pageNumber int32, remoteFolder string) {
	// Extract table form the document and show table info
	uploadFile(pdf_api, document)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	result, httpResponse, err := pdf_api.GetPageTables(document, pageNumber, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("ExtractTablesOnPage(): Failed to extract table from the document.")
	} else {
		if result.Tables == nil || len(result.Tables.List) == 0 {
			fmt.Println("ExtractTablesOnPage(): Tables not found on ", pageNumber, "page in the document.")
		} else {
			for i, t := range result.Tables.List {
				fmt.Println("table", i, " => id: '"+t.Id+"', page: '", t.PageNum, "', rows: '", len(t.RowList), "', columns: '", len(t.RowList[0].CellList), "'")
			}
		}
	}

}
