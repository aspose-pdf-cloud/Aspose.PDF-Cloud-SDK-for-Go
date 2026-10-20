package main

import (
	"fmt"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func converBmpToPdf(pdf_api *asposepdfcloud.PdfApiService, bmp_file string, pdf_name string, remote_folder string) {
	imageTemplateList := []asposepdfcloud.ImageTemplate{}

	imageTemplate := asposepdfcloud.ImageTemplate{
		ImagePath:    path.Join(remote_folder, bmp_file),
		ImageSrcType: asposepdfcloud.ImageSrcTypeCommon,
	}
	imageTemplateList = append(imageTemplateList, imageTemplate)

	uploadFile(pdf_api, bmp_file)

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
