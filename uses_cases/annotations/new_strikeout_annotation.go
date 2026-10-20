package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func AppendStrikeoutAnnotation(pdf_api *asposepdfcloud.PdfApiService, document_name string, page_num int32, output_document string, prefix string, remote_folder string) {
	// Append a new strikeout text annotation to the PDF document.
	UploadFile(pdf_api, document_name)
	args := map[string]interface{}{
		"folder": remote_folder,
	}

	new_annotation := asposepdfcloud.StrikeOutAnnotation{
		Name:                "Highlight_Text_Annotation",
		Rect:                &asposepdfcloud.Rectangle{LLX: 100, LLY: 350, URX: 450, URY: 400},
		Flags:               []asposepdfcloud.AnnotationFlags{asposepdfcloud.AnnotationFlagsDefault},
		HorizontalAlignment: asposepdfcloud.HorizontalAlignmentLeft,
		VerticalAlignment:   asposepdfcloud.VerticalAlignmentTop,
		RichText:            NEW_SO_ANNOTATION_TEXT,
		Subject:             NEW_SO_ANNOTATION_SUBJECT,
		Contents:            NEW_SO_ANNOTATION_CONTENTS,
		Title:               NEW_SO_ANNOTATION_DESCRIPTION,
		ZIndex:              1,
		Color:               &asposepdfcloud.Color{A: 0xFF, R: 0x00, G: 0xFF, B: 0x00},
		QuadPoints: []asposepdfcloud.Point{
			{X: 10, Y: 10},
			{X: 20, Y: 10},
			{X: 10, Y: 20},
			{X: 10, Y: 10},
		},
		Modified: "03/27/2026 00:00:00.000 AM",
	}

	_, httpResponse, err := pdf_api.PostPageStrikeOutAnnotations(document_name, page_num, []asposepdfcloud.StrikeOutAnnotation{new_annotation}, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("AppendStrikeoutAnnotation(): Failed to append annotation to the document page.")
	} else {
		fmt.Println("AppendStrikeoutAnnotation(): annotation '" + NEW_SO_ANNOTATION_TEXT + "' added to the document '" + document_name + "'.")
		DownloadFile(pdf_api, document_name, output_document, prefix)
	}
}
