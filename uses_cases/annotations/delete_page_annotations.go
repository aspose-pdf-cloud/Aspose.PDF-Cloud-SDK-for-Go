package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func DeletePageAnnotations(pdf_api *asposepdfcloud.PdfApiService, document_name string, page_num int32, output_name string) {
	// Delete annotation from the PDF document.
	UploadFile(pdf_api, document_name)
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}

	_, httpResponse, err := pdf_api.DeletePageAnnotations(document_name, page_num, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("DeletePageAnnotations(): Failed to delete annotation from the document.")
	} else {
		fmt.Println("DeletePageAnnotations(): annotations on page '", page_num, "' deleted from the document '"+document_name+"'.")
		DownloadFile(pdf_api, document_name, output_name, "del_page_annotations_")
	}
}
