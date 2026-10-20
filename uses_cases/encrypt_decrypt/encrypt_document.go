package main

import (
	"encoding/base64"
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func encryptDocument(pdf_api *asposepdfcloud.PdfApiService, document string, output_document string, userPassword string, ownerPassword string, encrypt_algorithm string, remote_folder string) {
	uploadFile(pdf_api, document)

	args := map[string]interface{}{
		"folder": remote_folder,
	}

	result, httpResponse, err := pdf_api.PostEncryptDocumentInStorage(document,
		base64.StdEncoding.EncodeToString([]byte(userPassword)),
		base64.StdEncoding.EncodeToString([]byte(ownerPassword)),
		encrypt_algorithm,
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
