package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v25"
)

func setFormFiled(pdf_api *asposepdfcloud.PdfApiService, document_name string, output_name string, field_name string, remote_folder string) {
	uploadFile(pdf_api, document_name)

	args := map[string]interface{}{
		"folder": remote_folder,
	}

	field := asposepdfcloud.Field{
		Name:   "EMail",
		Type_:  asposepdfcloud.FieldTypeText,
		Values: []string{"aspose-pdf-cloud@example.com"},
	}

	_, httpResponse, err := pdf_api.PutUpdateField(document_name, field_name, field, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		downloadFile(pdf_api, document_name, output_name, "upd_form_")
	}
}
