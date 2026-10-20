package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func ParseExtractTextBoxes(pdf_api *asposepdfcloud.PdfApiService, documentName string, localFolder string, remoteFolder string) {
	// Extract tables form the document
	uploadFile(pdf_api, documentName)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	result, httpResponse, err := pdf_api.GetDocumentTextBoxFields(documentName, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("ParseExtractTextBoxes(): Failed to extract text boxes from the document.")
	} else {
		if result.Fields == nil || len(result.Fields.List) == 0 {
			fmt.Println("ParseExtractTextBoxes(): Text boxes not found in the document.")
		} else {
			resultJson := "[\n"
			for _, t := range result.Fields.List {
				respTextBox, httpResponse, err := pdf_api.GetTextBoxField(documentName, t.FullName, args)
				if err != nil {
					fmt.Println(err.Error())
				} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
					fmt.Println("ParseExtractTextBoxes(): Failed to extract text box from the document.")
				} else {
					fmt.Println("TextBox", respTextBox.Field)
					jsTable, _ := json.Marshal(respTextBox.Field)
					resultJson += string(jsTable) + ",\n\n"
				}
			}
			resultJson += "]"
			fileName := path.Join(localFolder, ("parsed_taext_boxes_output.json"))
			f, _ := os.Create(fileName)
			_, _ = f.Write([]byte(resultJson))
			fmt.Println("File '" + fileName + "' successfully downloaded.")
		}
	}
}
