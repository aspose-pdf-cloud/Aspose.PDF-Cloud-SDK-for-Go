package main

import (
	"fmt"
	"path/filepath"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func resizeAllPages(pdf_api *asposepdfcloud.PdfApiService, document_name string, htmlTempDoc string, width int, height int, outputDocument string, localFolder string, tempFolder string) {
	uploadFile(pdf_api, document_name)

	htmlTempPath := filepath.Join(tempFolder, htmlTempDoc)

	args := map[string]interface{}{
		"folder":       tempFolder,
		"documentType": string(asposepdfcloud.HtmlDocumentTypeXhtml),
		"outputFormat": string(asposepdfcloud.OutputFormatFolder),
	}

	_, response, err := pdf_api.PutPdfInStorageToHtml(document_name, htmlTempPath, args)

	if err != nil {
		fmt.Println(err.Error())
	} else if response.StatusCode < 200 || response.StatusCode > 299 {
		fmt.Println("resizePages(): Can't convert pdf to html!")
	} else {
		fmt.Println("resizePages(): temporary file '" + htmlTempDoc + "' succesfully creaated.")
	}

	args2 := map[string]interface{}{
		"dstFolder":    tempFolder,
		"htmlFileName": htmlTempDoc,
		"height":       float64(height),
		"width":        float64(width),
	}

	_, response, err = pdf_api.PutHtmlInStorageToPdf(outputDocument, htmlTempPath, args2)
	if err != nil {
		fmt.Println(err.Error())
	} else if response.StatusCode < 200 || response.StatusCode > 299 {
		fmt.Println("resizePages(): Can't convert html to pdf!")
	} else {
		fmt.Println("resizePages(): Pages successfully resized.")
		downloadFile(pdf_api, outputDocument, "resized_doc_")
	}
}
