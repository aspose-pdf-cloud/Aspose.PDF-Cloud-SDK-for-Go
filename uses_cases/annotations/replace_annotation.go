package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func getAnnotation(pdf_api *asposepdfcloud.PdfApiService, document_name string, annotation_id string, remote_folder string) *asposepdfcloud.TextAnnotation {
	// Get annotation by Id in the PDF document.
	args := map[string]interface{}{
		"folder": remote_folder,
	}
	result, httpResponse, err := pdf_api.GetTextAnnotation(document_name, annotation_id, args)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("getAnnotation(): Failed to get annotation in the document.")
		return nil
	} else {
		fmt.Println("getAnnotation(): nnotation '" + annotation_id + "' successfully found '" + result.Annotation.Contents + "' in the document '" + document_name + "'.")
		return result.Annotation
	}
}

func ModifyAnnotation(pdf_api *asposepdfcloud.PdfApiService, document_name string, output_document string, annotation_id string, remote_folder string) {
	// Change annotation by Id in the PDF document.
	UploadFile(pdf_api, document_name)
	args := map[string]interface{}{
		"folder": remote_folder,
	}
	annotation := getAnnotation(pdf_api, document_name, annotation_id, remote_folder)
	annotation.Contents = REPLACED_CONTENT
	annotation.Icon = asposepdfcloud.TextIconStar

	_, httpResponse, err := pdf_api.PutTextAnnotation(document_name, annotation_id, *annotation, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("ModifyAnnotation(): Failed to modify annotation in the document.")
	} else {
		fmt.Println("ModifyAnnotation(): annotation '" + annotation.Id + "' successfully modified in the document '" + document_name + "'.")
		DownloadFile(pdf_api, document_name, output_document, "replaced_annotatiom_")
	}
}
