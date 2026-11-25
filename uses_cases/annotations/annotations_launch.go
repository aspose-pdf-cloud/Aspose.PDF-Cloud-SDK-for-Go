package main

func main() {
	pdfApi := InitPdfApi()

	AppendHighlightAnnotation(pdfApi, PDF_DOCUMENT_NAME, PAGE_NUMBER, PDF_OUTPUT, "add_highlight_annotation_", REMOTE_FOLDER)

	AppendStrikeoutAnnotation(pdfApi, PDF_DOCUMENT_NAME, PAGE_NUMBER, PDF_OUTPUT, "add_strikeout_annotation_", REMOTE_FOLDER)

	AppendTextAnnotation(pdfApi, PDF_DOCUMENT_NAME, PAGE_NUMBER, PDF_OUTPUT, "add_freetext_annotation_", REMOTE_FOLDER)

	AppendUnderlineAnnotation(pdfApi, PDF_DOCUMENT_NAME, PAGE_NUMBER, PDF_OUTPUT, "add_underline_annotation_", REMOTE_FOLDER)

	annotation_id := RequestAnnotations(pdfApi, PDF_DOCUMENT_NAME, PAGE_NUMBER, REMOTE_FOLDER)
	RequestAnnotationById(pdfApi, PDF_DOCUMENT_NAME, annotation_id, REMOTE_FOLDER)

	DeleteAnnotation(pdfApi, PDF_DOCUMENT_NAME, ANNOTATION_ID, PDF_OUTPUT)

	DeletePageAnnotations(pdfApi, PDF_DOCUMENT_NAME, PAGE_NUMBER, PDF_OUTPUT)

	ModifyAnnotation(pdfApi, PDF_DOCUMENT_NAME, PDF_OUTPUT, ANNOTATION_ID, REMOTE_FOLDER)
}
