package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func ParseExtractTables(pdf_api *asposepdfcloud.PdfApiService, documentName string, localFolder string, remoteFolder string) {
	// Extract tables form the document
	uploadFile(pdf_api, documentName)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	result, httpResponse, err := pdf_api.GetDocumentTables(documentName, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("ExtractTables(): Failed to extract tables from the document.")
	} else {
		if result.Tables == nil || len(result.Tables.List) == 0 {
			fmt.Println("ExtractTables(): Tables not found in the document.")
		} else {
			resultJson := "[\n"
			for _, t := range result.Tables.List {
				respTable, httpResponse, err := pdf_api.GetTable(documentName, t.Id, args)
				if err != nil {
					fmt.Println(err.Error())
				} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
					fmt.Println("ExtractTables(): Failed to extract table from the document.")
				} else {
					fmt.Println("table", respTable.Table)
					jsTable, _ := json.Marshal(respTable.Table)
					resultJson += string(jsTable) + ",\n\n"
				}
			}
			resultJson += "]"
			fileName := path.Join(localFolder, ("parsed_tables_output.json"))
			f, _ := os.Create(fileName)
			_, _ = f.Write([]byte(resultJson))
			fmt.Println("File '" + fileName + "' successfully downloaded.")
		}
	}
}
