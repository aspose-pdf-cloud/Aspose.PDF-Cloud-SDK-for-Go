package main

import (
	"fmt"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func ParseExtractFormsAsFDF(pdf_api *asposepdfcloud.PdfApiService, documentName string, outputFDFName string, remoteFolder string) {
	// Extract Form fields from the document to FDF file
	uploadFile(pdf_api, documentName)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	fdfPath := path.Join(remoteFolder, outputFDFName)

	_, httpResponse, err := pdf_api.PutExportFieldsFromPdfToFdfInStorage(documentName, fdfPath, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("ParseExtractFormsAsFDF(): Failed to extract Form fields from the document.")
	} else {
		fmt.Println("ParseExtractFormsAsFDF(): Forms fields successfully extracted from the document '" + documentName + "'.")
		downloadFile(pdf_api, outputFDFName, outputFDFName)
	}
}
