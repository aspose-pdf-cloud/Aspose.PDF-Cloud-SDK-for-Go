package main

import (
	"fmt"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func ParseExtractFormsAsXML(pdf_api *asposepdfcloud.PdfApiService, documentName string, outputXMLName string, remoteFolder string) {
	// Extract Form fields from the document to XML file
	uploadFile(pdf_api, documentName)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	xmlPath := path.Join(remoteFolder, outputXMLName)

	_, httpResponse, err := pdf_api.PutExportFieldsFromPdfToXmlInStorage(documentName, xmlPath, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("ParseExtractFormsAsXML(): Failed to extract Form fields from the document.")
	} else {
		fmt.Println("ParseExtractFormsAsXML(): Forms fields successfully extracted from the document '" + documentName + "'.")
		downloadFile(pdf_api, outputXMLName, outputXMLName)
	}
}
