package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func DeleteAnnotation(pdf_api *asposepdfcloud.PdfApiService, document_name string, annotation_id string, output_name string) {
	// Delete annotation from the PDF document.
	UploadFile(pdf_api, document_name)
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}

	_, httpResponse, err := pdf_api.DeleteAnnotation(document_name, annotation_id, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("DeleteAnnotation(): Failed to delete annotation from the document page.")
	} else {
		DeletePopupAnnotations(pdf_api, document_name, annotation_id)
		fmt.Println("DdeleteAnnotation(): annotation '" + annotation_id + "' deleted from the document '" + document_name + "'.")
		DownloadFile(pdf_api, document_name, output_name, "del_annotation_")
	}
}
