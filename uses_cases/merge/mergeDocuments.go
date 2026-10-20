package main

import (
	"fmt"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func mergeDocuments(pdf_api *asposepdfcloud.PdfApiService, output_name string, remote_folder string) {
	names := []string{PDF_DOCUMENT, PDF_DOCUMENT_2, PDF_DOCUMENT_3}

	mergeDocuments := asposepdfcloud.MergeDocuments{
		List: []string{},
	}

	for _, name := range names {
		uploadFile(pdf_api, name)
		mergeDocuments.List = append(mergeDocuments.List, path.Join(remote_folder, name))
	}

	args := map[string]interface{}{
		"folder": remote_folder,
	}

	_, httpResponse, err := pdf_api.PutMergeDocuments(output_name, mergeDocuments, args)

	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		downloadFile(pdf_api, output_name)
	}
}
