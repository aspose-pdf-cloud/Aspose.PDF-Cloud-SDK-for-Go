package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

// Delete images watermarks from document
func DeleteImageWatermarks(pdf_api *asposepdfcloud.PdfApiService, document string, watermarkIds []string, outputDocument string, remoteFolder string) {
	uploadFile(pdf_api, document)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	for _, watermarkId := range watermarkIds {
		_, httpResponse, err := pdf_api.DeleteImage(document, watermarkId, args)

		if err != nil {
			fmt.Println(err.Error())
		} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
			fmt.Println("DeleteImageWatermarks()): Failed to delete image from the document.")
		} else {
			fmt.Println("DeleteImageWatermarks(): image '" + watermarkId + "' successfully deleted from the document '" + document + "'.")
		}
	}

	downloadFile(pdf_api, document, outputDocument, "del_image_stamp_")
}
