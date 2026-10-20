package main

func main() {
	pdfApi := initPdfApi()

	extractDocumentBookmarks(pdfApi, PDF_DOCUMENT, REMOTE_FOLDER)
	extractBookmark(pdfApi, PDF_DOCUMENT, BOOKMARK_PATH, REMOTE_FOLDER)

	appendBookmark(pdfApi, PDF_DOCUMENT, BOOKMARK_PATH, BOOKMARK_TITLE, REMOTE_FOLDER)

	replaceBookmark(pdfApi, PDF_DOCUMENT, BOOKMARK_PATH, BOOKMARK_TITLE, REMOTE_FOLDER)

	removeBookmark(pdfApi, PDF_DOCUMENT, BOOKMARK_PATH, REMOTE_FOLDER)
}
