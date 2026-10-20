package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

func cropDocumentPage(pdf_api *asposepdfcloud.PdfApiService, document_name string, pageNumber int, llx int, lly int, width int, height int, outputDocument string, localFolder string, tempFolder string) {
	uploadFile(pdf_api, document_name)

	getPageInfo(pdf_api, document_name, pageNumber, tempFolder)

	imageFile := extractPdfPage(pdf_api, document_name, pageNumber, int(CROP_PAGE_WIDTH), int(CROP_PAGE_HEIGHT), localFolder, tempFolder)
	newPdf := createPdfDocument(pdf_api, outputDocument, width, height, tempFolder)
	if newPdf.Code != 200 {
		fmt.Println("cropPage(): Failed to create new PDF document!")
	} else {
		response := insertPageAsImage(pdf_api, outputDocument, imageFile, llx, lly, tempFolder)
		if response.Code == 200 {
			fmt.Println("cropPage(): Page successfully cropped.")
			downloadFile(pdf_api, outputDocument, "cropped_")
		} else {
			fmt.Println("cropPage(): Can't crop pdf document page!")
		}
	}
}
