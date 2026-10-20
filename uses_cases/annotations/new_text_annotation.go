package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func AppendTextAnnotation(pdf_api *asposepdfcloud.PdfApiService, document_name string, page_num int32, output_document string, prefix string, remote_folder string) {
	// Append a new free text annotation to the PDF document.
	UploadFile(pdf_api, document_name)
	args := map[string]interface{}{
		"folder": remote_folder,
	}

	text_style := asposepdfcloud.TextStyle{
		FontSize:        20,
		Font:            "Arial",
		ForegroundColor: &asposepdfcloud.Color{A: 0xFF, R: 0x00, G: 0xFF, B: 0x00},
		BackgroundColor: &asposepdfcloud.Color{A: 0xFF, R: 0xFF, G: 0x00, B: 0x00},
	}

	new_annotation := asposepdfcloud.FreeTextAnnotation{
		Rect:                &asposepdfcloud.Rectangle{LLX: 100, LLY: 350, URX: 450, URY: 400},
		TextStyle:           &text_style,
		Name:                "Free Text Annotation",
		Flags:               []asposepdfcloud.AnnotationFlags{asposepdfcloud.AnnotationFlagsDefault},
		HorizontalAlignment: asposepdfcloud.HorizontalAlignmentCenter,
		Intent:              asposepdfcloud.FreeTextIntentFreeTextTypeWriter,
		RichText:            NEW_FT_ANNOTATION_TEXT,
		Subject:             NEW_FT_ANNOTATION_SUBJECT,
		Contents:            NEW_FT_ANNOTATION_CONTENTS,
		Title:               NEW_FT_ANNOTATION_DESCRIPTION,
		ZIndex:              1,
		Justification:       asposepdfcloud.JustificationCenter,
	}

	_, httpResponse, err := pdf_api.PostPageFreeTextAnnotations(document_name, page_num, []asposepdfcloud.FreeTextAnnotation{new_annotation}, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("AppendTextAnnotation(): Failed to append annotation to the document page.")
	} else {
		fmt.Println("AppendTextAnnotation(): annotation '" + NEW_FT_ANNOTATION_TEXT + "' added to the document '" + document_name + "'.")
		DownloadFile(pdf_api, document_name, output_document, prefix)
	}
}
