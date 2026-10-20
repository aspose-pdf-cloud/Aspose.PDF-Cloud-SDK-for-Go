package main

import (
	"fmt"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

// Append image watermark to a document
func AppendNewWatermark(pdf_api *asposepdfcloud.PdfApiService, document string, imageFileName string, opacity float64, rotate float64, xPos float64, yPos float64, width float64, height float64, outputDocument string, remoteFolder string) {
	uploadFile(pdf_api, document)
	uploadFile(pdf_api, imageFileName)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	image := asposepdfcloud.ImageStamp{
		Background:  true,
		Opacity:     opacity,
		Rotate:      asposepdfcloud.RotationNone,
		RotateAngle: rotate,
		XIndent:     xPos,
		YIndent:     yPos,
		Width:       width,
		Height:      height,
		Zoom:        1,
		FileName:    path.Join(remoteFolder, imageFileName),
	}

	_, httpResponse, err := pdf_api.PostDocumentImageStamps(document, []asposepdfcloud.ImageStamp{image}, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println("Successfully added image watermark. ")
		downloadFile(pdf_api, document, outputDocument, "add_watermark_")
	}
}
