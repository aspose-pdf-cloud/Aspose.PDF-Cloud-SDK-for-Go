package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func RequestAnnotationById(pdf_api *asposepdfcloud.PdfApiService, document_name string, annotation_id string, remote_folder string) {
	// Get annotation by Id in the PDF document.
	UploadFile(pdf_api, document_name)
	args := map[string]interface{}{
		"folder": remote_folder,
	}
	result, httpResponse, err := pdf_api.GetTextAnnotation(document_name, annotation_id, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("RequestAnnotationById(): Failed to request text annotation from the document page.")
	} else {
		fmt.Println("RequestAnnotationById(): annotation '" + annotation_id + "' successfully found '" + result.Annotation.Contents + "' in the document '" + document_name + "'.")
	}
}
