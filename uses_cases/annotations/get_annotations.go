package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func RequestAnnotations(pdf_api *asposepdfcloud.PdfApiService, document_name string, page_num int32, remote_folder string) string {
	// Get annotations from the page in the PDF document.
	UploadFile(pdf_api, document_name)
	args := map[string]interface{}{
		"folder": remote_folder,
	}
	annotation_result := ""
	result, httpResponse, err := pdf_api.GetPageAnnotations(document_name, page_num, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("RequestAnnotations(): Failed to delete annotation from the document page.")
	} else {
		for _, annotation := range result.Annotations.List {
			if annotation.AnnotationType == asposepdfcloud.AnnotationTypeText {
				fmt.Println("RequestAnnotations(): annotation id=", annotation.Id, " with '"+annotation.Contents+"' content get from the document '"+document_name+"' on ", annotation.PageIndex, " page.")
				annotation_result := annotation.Id
				return annotation_result
			}
		}
	}
	fmt.Println("RequestAnnotations(): Failed to get annotation in the document.")
	return annotation_result
}
