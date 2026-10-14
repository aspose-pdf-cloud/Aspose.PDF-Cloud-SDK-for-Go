package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v25"
)

func getFormFiled(pdf_api *asposepdfcloud.PdfApiService, document_name string, remote_folder string) {
	uploadFile(pdf_api, document_name)

	args := map[string]interface{}{
		"folder": remote_folder,
	}

	result, httpResponse, err := pdf_api.GetFields(document_name, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		for i, f := range result.Fields.List {
			fmt.Println("field > ", i, ": type: '"+string(f.Type_)+"', name: '"+f.Name+",' values: '"+f.Values[0]+"'")
		}
	}
}
