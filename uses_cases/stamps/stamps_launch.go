package main

func main() {
	pdfApi := initPdfApi()

	AppendStampText(pdfApi, PDF_DOCUMENT, PDF_OUTPUT, STAMP_TEXT, REMOTE_FOLDER)

	AppendStampImage(pdfApi, PDF_DOCUMENT, PDF_OUTPUT, IMAGE_STAMP_FILE, IMAGE_STAMP_WIDTH, IMAGE_STAMP_HEIGHT, REMOTE_FOLDER)

	DeleteDocumentStamps(pdfApi, PDF_DOCUMENT, PDF_OUTPUT, REMOTE_FOLDER)

	DeletePageStamps(pdfApi, PDF_DOCUMENT, PAGE_NUMBER, PDF_OUTPUT, REMOTE_FOLDER)
}
