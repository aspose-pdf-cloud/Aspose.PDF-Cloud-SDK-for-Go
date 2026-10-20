package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func getAttachmentByIndex(pdf_api *asposepdfcloud.PdfApiService, document_name string, attachment_index int32, remote_folder string, local_folder string) {
	args := map[string]interface{}{
		"folder": remote_folder,
	}
	result, httpResponse, err := pdf_api.GetDocumentAttachmentByIndex(document_name, attachment_index, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println(result.Attachment)
		extractAttachmentByIndex(pdf_api, document_name, attachment_index, result.Attachment.Name, remote_folder, local_folder)
	}
}

func extractAttachmentByIndex(pdf_api *asposepdfcloud.PdfApiService, document_name string, attachment_index int32, attachment_name string, remote_folder string, local_folder string) {
	args := map[string]interface{}{
		"folder": remote_folder,
	}
	result, httpResponse, err := pdf_api.GetDownloadDocumentAttachmentByIndex(document_name, attachment_index, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		saveByteArrayToFile(local_folder, attachment_name, result)
	}
}
