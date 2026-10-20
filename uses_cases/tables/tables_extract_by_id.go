package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func ExtractTableById(pdf_api *asposepdfcloud.PdfApiService, document string, tableId string, remoteFolder string) {
	// Extract table form the document and show table info
	uploadFile(pdf_api, document)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	result, httpResponse, err := pdf_api.GetTable(document, tableId, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("ExtractTableById(): Failed to extract table from the document.")
	} else {
		fmt.Println("foud => id: '"+result.Table.Id+"', page: '", result.Table.PageNum, "', rows: '", len(result.Table.RowList), "', columns: '", len(result.Table.RowList[0].CellList), "'")
	}

}
