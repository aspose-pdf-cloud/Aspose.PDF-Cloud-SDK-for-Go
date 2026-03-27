package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v25"
)

func appendFormFiled(pdf_api *asposepdfcloud.PdfApiService, document_name string, output_name string, remote_folder string) {
	uploadFile(pdf_api, document_name)

	args := map[string]interface{}{
		"folder": remote_folder,
	}

	textBox := asposepdfcloud.TextBoxField{
		PageIndex:   1,
		PartialName: "EMail",
		Rect:        &asposepdfcloud.Rectangle{LLX: 100, LLY: 100, URX: 180, URY: 120},
		Value:       "aspose-pdf-cloud@example.com",
		Border: &asposepdfcloud.Border{
			Width: 5,
			Dash:  &asposepdfcloud.Dash{On: 1, Off: 1},
		},
	}

	_, httpResponse, err := pdf_api.PutTextBoxField(document_name, "EMail", textBox, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		downloadFile(pdf_api, document_name, output_name, "add_form_")
	}
}
