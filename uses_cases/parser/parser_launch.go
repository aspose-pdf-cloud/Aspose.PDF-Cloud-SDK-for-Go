package main

func main() {
	pdfApi := initPdfApi()

	ParseExtractFormsAsXML(pdfApi, PDF_DOCUMENT, XML_OUTPUT_FILE, REMOTE_FOLDER)
	ParseExtractFormsAsFDF(pdfApi, PDF_DOCUMENT, FDF_OUTPUT_FILE, REMOTE_FOLDER)

	ParseExtractImages(pdfApi, PDF_DOCUMENT, PAGE_NUMBER, LOCAL_FOLDER, REMOTE_FOLDER)
	ParseExtractTables(pdfApi, PDF_DOCUMENT, LOCAL_FOLDER, REMOTE_FOLDER)
	ParseExtractTextBoxes(pdfApi, PDF_DOCUMENT, LOCAL_FOLDER, REMOTE_FOLDER)

}
