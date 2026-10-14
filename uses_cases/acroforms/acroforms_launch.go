package main

func main() {
	pdfApi := initPdfApi()

	appendFormFiled(pdfApi, PDF_DOCUMENT, PDF_OUTPUT_NAME, REMOTE_FOLDER)

	getFormFiled(pdfApi, PDF_DOCUMENT, REMOTE_FOLDER)

	deleteFormFiled(pdfApi, PDF_DOCUMENT, PDF_OUTPUT_NAME, FIELD_NAME, REMOTE_FOLDER)

	setFormFiled(pdfApi, PDF_DOCUMENT, PDF_OUTPUT_NAME, FIELD_NAME, REMOTE_FOLDER)

	updateFormFiled(pdfApi, PDF_DOCUMENT, PDF_OUTPUT_NAME, FIELD_NAME, REMOTE_FOLDER)
}
