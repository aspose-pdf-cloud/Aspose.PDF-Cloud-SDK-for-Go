package main

import (
	"encoding/base64"
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func decryptDocument(pdf_api *asposepdfcloud.PdfApiService, document string, output_document string, password string, remote_folder string) {
	uploadFile(pdf_api, document)

	args := map[string]interface{}{
		"folder": remote_folder,
	}

	result, httpResponse, err := pdf_api.PostDecryptDocumentInStorage(document,
		base64.StdEncoding.EncodeToString([]byte(password)),
		args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println(result)

		downloadFile(pdf_api, document, output_document)
	}
}
