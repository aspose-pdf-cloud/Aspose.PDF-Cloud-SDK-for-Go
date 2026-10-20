package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

// Extract images watermarks from document and show info to console
func GetWatermarks(pdf_api *asposepdfcloud.PdfApiService, document string, remoteFolder string) {
	uploadFile(pdf_api, document)

	args := map[string]interface{}{
		"folder": remoteFolder,
	}

	result, httpResponse, err := pdf_api.GetPages(document, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("GetPages() - Unexpected error!")
	} else {
		for pageNum := range result.Pages.List {
			imageResult, httpResponse, err := pdf_api.GetImages(document, int32(pageNum+1), args)
			if err != nil {
				fmt.Println(err.Error())
			} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
				fmt.Println("Unexpected error!")
			} else {
				for _, image := range imageResult.Images.List {
					fmt.Println("Page", pageNum+1, "Image Id=", image.Id, "LinkId:", image.Links[0].Href, "Width:", (image.Rectangle.LLY - image.Rectangle.LLX), "Height:", (image.Rectangle.URY - image.Rectangle.URX))
				}
			}
		}
	}
}
