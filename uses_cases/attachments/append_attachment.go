package main

import (
	"fmt"
	"path/filepath"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func appendAttachment(pdf_api *asposepdfcloud.PdfApiService, document_name string, attachment_path string, description string, mime_type string, remote_folder string) {
	args := map[string]interface{}{
		"folder": remote_folder,
	}

	attachment := asposepdfcloud.AttachmentInfo{
		Path:        filepath.Join(remote_folder, attachment_path),
		Description: description,
		Name:        attachment_path,
		MimeType:    mime_type,
	}

	result, httpResponse, err := pdf_api.PostAddDocumentAttachment(document_name, attachment, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println(result.Attachments.Links)
	}
}
