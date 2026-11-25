package main

func main() {
	pdfApi := initPdfApi(PDF_API_CLIENT_ID, PDF_API_CLIENT_SECRET)

	СomparePdf(pdfApi, PDF_DOCUMENT_1, PDF_DOCUMENT_2, PDF_OUTPUT, REMOTE_FOLDER)
}
