package main

import (
	"fmt"
	"os"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func SavePageAsPNG(pdf_api *asposepdfcloud.PdfApiService, document string, pageNumber int32, outputPNG string, remoteFolder string) {
	// Show page information of the PDF document.
	uploadFile(pdf_api, document)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	result, httpResponse, err := pdf_api.GetPageConvertToPng(document, pageNumber, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fileName := path.Join(LOCAL_FOLDER, outputPNG)
		f, _ := os.Create(fileName)
		_, _ = f.Write(result)
		fmt.Println("File '" + outputPNG + "' successfully downloaded.")
	}
}
