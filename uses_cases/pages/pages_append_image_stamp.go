package main

import (
	"fmt"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func AppendPageStampImage(pdf_api *asposepdfcloud.PdfApiService, documentName string, outputDocument string, imageFileName string, pageNumber int32, width float64, height float64, remoteFolder string) {
	// Append image stamp to page of the PDF document.
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
	_, httpResponse, err := pdf_api.PostPageImageStamps(documentName, pageNumber, []asposepdfcloud.ImageStamp{stamp}, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("appendStampImage()): Failed to append image stamp to the document.")
	} else {
		fmt.Println("appendStampImage(): image stamp '"+imageFileName+"' appended successfully pn page", pageNumber, " to the document '"+documentName+"'.")
		downloadFile(pdf_api, documentName, outputDocument, "add_image_stamp_")
	}
}
