package main

func main() {
	pdfApi := initPdfApi()

	mergeDocuments(pdfApi, PDF_DOCUMENT, REMOTE_FOLDER)
}
