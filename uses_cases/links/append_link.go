package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func appendLink(pdf_api *asposepdfcloud.PdfApiService, document string, output_document string, page_num int32, link_action string, rect *asposepdfcloud.Rectangle, remote_folder string) {
	uploadFile(pdf_api, document)

	args := map[string]interface{}{
		"folder": remote_folder,
	}

	link := asposepdfcloud.Link{Href: link_action}

	link_annotation := asposepdfcloud.LinkAnnotation{
		Links:        []asposepdfcloud.Link{link},
		ActionType:   asposepdfcloud.LinkActionTypeGoToURIAction,
		Action:       link_action,
		Highlighting: asposepdfcloud.LinkHighlightingModeInvert,
		Color:        &asposepdfcloud.Color{A: 0xFF, R: 0xAA, G: 0x00, B: 0x00},
		Rect:         rect,
	}

	result, httpResponse, err := pdf_api.PostPageLinkAnnotations(
		document, page_num, []asposepdfcloud.LinkAnnotation{link_annotation}, args,
	)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println(result)

		downloadFile(pdf_api, document, output_document)
	}
}
