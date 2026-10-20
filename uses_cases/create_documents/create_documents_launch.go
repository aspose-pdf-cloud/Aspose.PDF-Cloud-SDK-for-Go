package main

func main() {
	pdfApi := initPdfApi()

	createPdf(pdfApi, PDF_DOCUMENT, REMOTE_FOLDER)
	downloadFile(pdfApi, PDF_DOCUMENT)

	createPdfSimple(pdfApi, PDF_SIMPLE_OUTPUT, REMOTE_FOLDER)
}
