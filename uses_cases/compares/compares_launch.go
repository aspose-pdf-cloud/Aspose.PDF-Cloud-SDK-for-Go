package main

func main() {
	pdfApi := initPdfApi()

	СomparePdf(pdfApi, PDF_DOCUMENT_1, PDF_DOCUMENT_2, PDF_OUTPUT, REMOTE_FOLDER)
}
