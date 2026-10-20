package main

import (
	"fmt"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func AppendStampImage(pdf_api *asposepdfcloud.PdfApiService, documentName string, outputDocument string, imageFileName string, width float64, height float64, remoteFolder string) {
	uploadFile(pdf_api, documentName)
	uploadFile(pdf_api, imageFileName)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	stamp := asposepdfcloud.ImageStamp{
		Background:          true,
		HorizontalAlignment: asposepdfcloud.HorizontalAlignmentCenter,
		VerticalAlignment:   asposepdfcloud.VerticalAlignmentCenter,
		Opacity:             1,
		Rotate:              asposepdfcloud.RotationNone,
		RotateAngle:         45,
		Width:               width,
		Height:              height,
		Zoom:                1,
		FileName:            path.Join(remoteFolder, imageFileName),
	}
	_, httpResponse, err := pdf_api.PostDocumentImageStamps(documentName, []asposepdfcloud.ImageStamp{stamp}, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("appendStampImage()): Failed to append image stamp to the document.")
	} else {
		fmt.Println("appendStampImage(): image stamp '" + imageFileName + "' appended successfully to the document '" + documentName + "'.")
		downloadFile(pdf_api, documentName, outputDocument, "add_image_stamp_")
	}
}
