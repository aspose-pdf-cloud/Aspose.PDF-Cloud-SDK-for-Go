package main

import (
	"fmt"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func converImagesToPdf(pdf_api *asposepdfcloud.PdfApiService, images []string, pdf_name string, remote_folder string) {
	imageTemplateList := []asposepdfcloud.ImageTemplate{}

	for _, fileName := range images {
		imageTemplate := asposepdfcloud.ImageTemplate{
			ImagePath:    path.Join(remote_folder, fileName),
			ImageSrcType: asposepdfcloud.ImageSrcTypeCommon,
		}
		imageTemplateList = append(imageTemplateList, imageTemplate)

		uploadFile(pdf_api, fileName)
	}

	imageTemplatesRequest := asposepdfcloud.ImageTemplatesRequest{
		IsOCR:      true,
		OCRLangs:   "eng",
		ImagesList: imageTemplateList,
	}

	args := map[string]interface{}{
		"folder": remote_folder,
	}

	_, httpResponse, err := pdf_api.PutImageInStorageToPdf(pdf_name, imageTemplatesRequest, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		downloadFile(pdf_api, pdf_name, "")
	}
}
