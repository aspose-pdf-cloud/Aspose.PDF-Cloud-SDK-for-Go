package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func replaceLink(pdf_api *asposepdfcloud.PdfApiService, document string, output_document string, link_id string, link_action string, remote_folder string) {
	uploadFile(pdf_api, document)

	args := map[string]interface{}{
		"folder": remote_folder,
	}

	result, httpResponse, err := pdf_api.GetLinkAnnotation(document, link_id, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		links := []asposepdfcloud.LinkAnnotation{*result.Link}
		showLinks(&links)

		link := asposepdfcloud.Link{Href: link_action}

		link_annotation := asposepdfcloud.LinkAnnotation{
			Links:        []asposepdfcloud.Link{link},
			ActionType:   asposepdfcloud.LinkActionTypeGoToURIAction,
			Action:       link_action,
			Highlighting: asposepdfcloud.LinkHighlightingModeInvert,
			Color:        &asposepdfcloud.Color{A: 0xFF, R: 0xAA, G: 0x00, B: 0x00},
			Rect:         result.Link.Rect,
		}

		result2, httpResponse, err := pdf_api.PutLinkAnnotation(PDF_DOCUMENT, result.Link.Id, link_annotation, args)
		if err != nil {
			fmt.Println(err.Error())
		} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
			fmt.Println("Unexpected error!")
		} else {
			fmt.Println(result2)

			downloadFile(pdf_api, document, output_document)
		}
	}
}
