package main

import (
	"fmt"
	"os"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func ParseExtractImages(pdf_api *asposepdfcloud.PdfApiService, documentName string, pageNumber int32, localFolder string, remoteFolder string) {
	// Extract Images from the page of PDF document
	uploadFile(pdf_api, documentName)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	respImages, httpResponse, err := pdf_api.GetImages(documentName, pageNumber, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("ParseExtractImages(): Failed to extract images from the page of document.")
	} else {
		for _, image := range respImages.Images.List {

			response, httpResponse, err := pdf_api.GetImageExtractAsPng(documentName, image.Id, args)

			if err != nil {
				fmt.Println(err.Error())
			} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
				fmt.Println("ParseExtractImages(): Failed to extract image.")
			} else {
				fmt.Println("ParseExtractImages(): Images'" + image.Id + "' successfully extracted from the page of document.")

				fileName := path.Join(localFolder, (image.Id + ".png"))
				f, _ := os.Create(fileName)
				_, _ = f.Write(response)
				fmt.Println("File '" + fileName + "' successfully downloaded.")
			}
		}
	}
}
