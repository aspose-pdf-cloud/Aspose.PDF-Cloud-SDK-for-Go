package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func appendTextFooter(pdf_api *asposepdfcloud.PdfApiService, document_name string, text string, startPage int32, end_page int32, remote_folder string) {
	uploadFile(pdf_api, document_name)

	args := map[string]interface{}{
		"folder":          remote_folder,
		"startPageNumber": startPage,
		"endPageNumber":   end_page,
	}

	footer := asposepdfcloud.TextFooter{
		Background:          true,
		LeftMargin:          1,
		RightMargin:         2,
		HorizontalAlignment: asposepdfcloud.HorizontalAlignmentLeft,
		Opacity:             1,
		Rotate:              asposepdfcloud.RotationNone,
		RotateAngle:         0,
		XIndent:             0,
		YIndent:             0,
		Zoom:                1,
		TextAlignment:       asposepdfcloud.HorizontalAlignmentCenter,
		Value:               text,
	}

	_, httpResponse, err := pdf_api.PostDocumentTextFooter(document_name, footer, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		downloadFile(pdf_api, document_name)
	}
}
