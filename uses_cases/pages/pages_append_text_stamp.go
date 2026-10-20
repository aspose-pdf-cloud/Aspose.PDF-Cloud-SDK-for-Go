package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func AppendPageStampText(pdf_api *asposepdfcloud.PdfApiService, documentName string, pageNumber int32, outputDocument string, text_value string, remoteFolder string) {
	// Append text stamp to page of the PDF document.
	uploadFile(pdf_api, documentName)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	textStamp := asposepdfcloud.TextStamp{
		Background:          true,
		HorizontalAlignment: asposepdfcloud.HorizontalAlignmentCenter,
		VerticalAlignment:   asposepdfcloud.VerticalAlignmentCenter,
		Opacity:             1,
		Rotate:              asposepdfcloud.RotationNone,
		RotateAngle:         30,
		Zoom:                1,
		TextAlignment:       asposepdfcloud.HorizontalAlignmentCenter,
		Value:               text_value,
		TextState:           &asposepdfcloud.TextState{FontSize: 14, FontStyle: asposepdfcloud.FontStylesBoldItalic, Font: "Arial", ForegroundColor: &asposepdfcloud.Color{A: 0xFF, R: 0xFF, G: 0x00, B: 0x00}},
	}

	_, httpResponse, err := pdf_api.PostPageTextStamps(documentName, pageNumber, []asposepdfcloud.TextStamp{textStamp}, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("appendStampText(): Failed to append text stamp to the document.")
	} else {
		fmt.Println("appendStampText(): text stamp '"+text_value+"' appended successfully on page", pageNumber, " to the document '"+documentName+"'.")
		downloadFile(pdf_api, documentName, outputDocument, "add_text_stamp_")
	}
}
