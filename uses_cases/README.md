#### annotations
- **[annotations/annotations_launch.go](annotations/annotations_launch.go)** – Orchestrates various PDF annotation operations including adding highlights, strikeouts, free text, underlines, and managing annotations.
  ```bash
  go run uses_cases/annotations/*
  ```
- *[annotations/delete_page_annotations.go](annotations/delete_page_annotations.go)* – Deletes all annotations from a specified page in a PDF document.
- *[annotations/delete_text_annotation.go](annotations/delete_text_annotation.go)* – Deletes a specific annotation by its ID and its associated popup from a PDF document.
- *[annotations/get_annotation_by_id.go](annotations/get_annotation_by_id.go)* – Retrieves and displays the details of a specific text annotation by its ID.
- *[annotations/get_annotations.go](annotations/get_annotations.go)* – Fetches and lists all annotations from a specified page in a PDF document.
- *[annotations/new_highlight_annotation.go](annotations/new_highlight_annotation.go)* – Adds a new highlight annotation with custom text, color, and position to a PDF page.
- *[annotations/new_strikeout_annotation.go](annotations/new_strikeout_annotation.go)* – Inserts a new strikeout text annotation into a specified page of a PDF document.
- *[annotations/new_text_annotation.go](annotations/new_text_annotation.go)* – Appends a new free text annotation with custom styling to a PDF page.
- *[annotations/new_underline_annotation.go](annotations/new_underline_annotation.go)* – Creates and adds a new underline annotation to a specified page in a PDF.
- *[annotations/replace_annotation.go](annotations/replace_annotation.go)* – Modifies the content and icon of an existing text annotation in a PDF document.

#### attachments
- **[attachments/attachments_launch.go](attachments/attachments_launch.go)** – Demonstrates PDF attachment operations including listing, extracting, and appending new attachments.
  ```bash
  go run uses_cases/attachments/*
  ```
- *[attachments/append_attachment.go](attachments/append_attachment.go)* – Appends a new file attachment to an existing PDF document.
- *[attachments/get_attachment_by_index_show_save.go](attachments/get_attachment_by_index_show_save.go)* – Retrieves a specific attachment by its index, displays its info, and saves it locally.
- *[attachments/get_attachments_show.go](attachments/get_attachments_show.go)* – Lists all attachments embedded within a PDF document.

#### bookmarks
- **[bookmarks/bookmarks_launch.go](bookmarks/bookmarks_launch.go)** – Executes a series of PDF bookmark operations including extraction, addition, replacement, and removal.
  ```bash
  go run uses_cases/bookmarks/*
  ```
- *[bookmarks/append_bookmark.go](bookmarks/append_bookmark.go)* – Adds a new bookmark with specific navigation properties to a PDF document.
- *[bookmarks/get_bookmark_by_path_show.go](bookmarks/get_bookmark_by_path_show.go)* – Extracts and displays the details of a bookmark found at a specific path.
- *[bookmarks/get_bookmarks_show.go](bookmarks/get_bookmarks_show.go)* – Extracts and lists all bookmarks from a PDF document.
- *[bookmarks/remove_bookmark.go](bookmarks/remove_bookmark.go)* – Deletes a bookmark located at a specified path from a PDF document.
- *[bookmarks/replace_bookmark.go](bookmarks/replace_bookmark.go)* – Replaces an existing bookmark's properties with new values at a given path.

#### change_layout
- **[change_layout/change_layout_launch.go](change_layout/change_layout_launch.go)** – Coordinates PDF layout modification tasks such as page rotation, cropping, and resizing.
  ```bash
  go run uses_cases/change_layout/*
  ```
- *[change_layout/crop_document_page.go](change_layout/crop_document_page.go)* – Crops a specified PDF page to defined dimensions by converting it to an image and creating a new PDF.
- *[change_layout/resize_document_all_pages.go](change_layout/resize_document_all_pages.go)* – Resizes all pages of a PDF document by converting it to HTML and back with new dimensions.
- *[change_layout/rotate_documents_pages.go](change_layout/rotate_documents_pages.go)* – Rotates a specified range of pages in a PDF document by a given angle.

#### compares
- **[compares/compares_launch.go](compares/compares_launch.go)** – Launches the process to compare two PDF documents and generate a result PDF.
  ```bash
  go run uses_cases/compares/*
  ```
- *[compares/compare_pdf_documents.go](compares/compare_pdf_documents.go)* – Compares two PDF documents and produces a new PDF highlighting the differences.

#### conversions
- **[conversions/conversions_launch.go](conversions/conversions_launch.go)** – Orchestrates a comprehensive suite of PDF conversion tasks to and from various formats (images, HTML, Office docs, etc.).
  ```bash
  go run uses_cases/conversions/*
  ```
- *[conversions/convertBmpToPdf.go](conversions/convertBmpToPdf.go)* – Converts a BMP image file into a PDF document.
- *[conversions/convertEpubToPdf.go](conversions/convertEpubToPdf.go)* – Converts an EPUB file to a PDF document.
- *[conversions/convertGifToPdf.go](conversions/convertGifToPdf.go)* – Converts a GIF image file into a PDF document.
- *[conversions/convertHtmlToPdf.go](conversions/convertHtmlToPdf.go)* – Converts an HTML file (with referenced resources) into a PDF document.
- *[conversions/convertImagesToPdf.go](conversions/convertImagesToPdf.go)* – Converts a list of image files into a single multi-page PDF document.
- *[conversions/convertJpegToPdf.go](conversions/convertJpegToPdf.go)* – Converts a JPEG image file into a PDF document.
- *[conversions/convertMdToPdf.go](conversions/convertMdToPdf.go)* – Converts a Markdown file into a PDF document.
- *[conversions/convertMhtmlToPdf.go](conversions/convertMhtmlToPdf.go)* – Converts an MHTML (web archive) file into a PDF document.
- *[conversions/convertPclToPdf.go](conversions/convertPclToPdf.go)* – Converts a PCL file into a PDF document.
- *[conversions/convertPdfToDoc.go](conversions/convertPdfToDoc.go)* – Converts a PDF document to the legacy DOC (Microsoft Word) format.
- *[conversions/convertPdfToDocx.go](conversions/convertPdfToDocx.go)* – Converts a PDF document to the modern DOCX (Microsoft Word) format.
- *[conversions/convertPdfToEpub.go](conversions/convertPdfToEpub.go)* – Converts a PDF document to the EPUB format.
- *[conversions/convertPdfToExcel.go](conversions/convertPdfToExcel.go)* – Converts a PDF document to the XLSX (Microsoft Excel) format.
- *[conversions/convertPdfToHtnl.go](conversions/convertPdfToHtnl.go)* – Converts a PDF document to a ZIP archive containing HTML files.
- *[conversions/convertPdfToMobiXml.go](conversions/convertPdfToMobiXml.go)* – Converts a PDF document to the MOBI XML format.
- *[conversions/convertPdfToPowerpoint.go](conversions/convertPdfToPowerpoint.go)* – Converts a PDF document to the PPTX (Microsoft PowerPoint) format.
- *[conversions/convertPdfToSvg.go](conversions/convertPdfToSvg.go)* – Converts a PDF document to a ZIP archive containing SVG files.
- *[conversions/convertPdfToTex.go](conversions/convertPdfToTex.go)* – Converts a PDF document to the TeX (LaTeX) format.
- *[conversions/convertPdfToTiff.go](conversions/convertPdfToTiff.go)* – Converts a PDF document to a multi-page TIFF image file.
- *[conversions/convertPdfToXls.go](conversions/convertPdfToXls.go)* – Converts a PDF document to the legacy XLS (Microsoft Excel) format.
- *[conversions/convertPdfToXml.go](conversions/convertPdfToXml.go)* – Converts a PDF document to an XML file.
- *[conversions/convertPdfToXps.go](conversions/convertPdfToXps.go)* – Converts a PDF document to the XPS format.
- *[conversions/convertPngToPdf.go](conversions/convertPngToPdf.go)* – Converts a PNG image file into a PDF document.
- *[conversions/convertPsToPdf.go](conversions/convertPsToPdf.go)* – Converts a PostScript (PS) file into a PDF document.

#### create_documents
- **[create_documents/createPdfWithConfig.go](create_documents/createPdfWithConfig.go)** – Demonstrates creating a PDF document with custom configuration (pages, properties, display settings) and a simple creation method.
  ```bash
  go run uses_cases/create_documents/*
  ```
- *[create_documents/createPdfSimple.go](create_documents/createPdfSimple.go)* – Creates a simple, empty PDF document with default settings.

#### encrypt_decrypt
- **[encrypt_decrypt/encrypt_decrypt_launch.go](encrypt_decrypt/encrypt_decrypt_launch.go)** – Manages PDF security operations including encryption, decryption, and password changes.
  ```bash
  go run uses_cases/encrypt_decrypt/*
  ```
- *[encrypt_decrypt/change_document_passwords.go](encrypt_decrypt/change_document_passwords.go)* – Changes both the user and owner passwords of an encrypted PDF document.
- *[encrypt_decrypt/decrypt_document.go](encrypt_decrypt/decrypt_document.go)* – Removes password protection from an encrypted PDF document.
- *[encrypt_decrypt/encrypt_document.go](encrypt_decrypt/encrypt_document.go)* – Applies encryption with specified passwords and algorithm to a PDF document.

#### headers_footers
- **[headers_footers/headers_footers_launch.go](headers_footers/headers_footers_launch.go)** – Executes operations to add both image and text headers and footers to different page ranges of a PDF.
  ```bash
  go run uses_cases/headers_footers/*
  ```
- *[headers_footers/append_image_footer.go](headers_footers/append_image_footer.go)* – Appends an image as a footer to specified pages of a PDF document.
- *[headers_footers/append_image_header.go](headers_footers/append_image_header.go)* – Appends an image as a header to specified pages of a PDF document.
- *[headers_footers/append_text_footer.go](headers_footers/append_text_footer.go)* – Appends a text string as a footer to specified pages of a PDF document.
- *[headers_footers/append_text_header.go](headers_footers/append_text_header.go)* – Appends a text string as a header to specified pages of a PDF document.

#### links
- **[links/links_launch.go](links/links_launch.go)** – Coordinates a workflow for managing hyperlinks in PDFs, including adding, retrieving, replacing, and deleting links.
  ```bash
  go run uses_cases/links/*
  ```
- *[links/append_link.go](links/append_link.go)* – Adds a new hyperlink annotation pointing to a URL on a specified PDF page.
- *[links/extract_link.go](links/extract_link.go)* – Extracts and displays the properties of a specific link annotation by its ID.
- *[links/extract_links.go](links/extract_links.go)* – Extracts and lists all link annotations from a specified page of a PDF document.
- *[links/remove_link.go](links/remove_link.go)* – Deletes a specific hyperlink annotation from a PDF document.
- *[links/replace_link.go](links/replace_link.go)* – Replaces the target URL and properties of an existing hyperlink annotation.

#### merge
- **[merge/mergePdfDocuments_launch.go](merge/mergePdfDocuments_launch.go)** – Launches the process to merge multiple PDF documents into a single output file.
  ```bash
  go run uses_cases/merge/*
  ```
- *[merge/mergeDocuments.go](merge/mergeDocuments.go)* – Merges a list of specified PDF documents into one combined PDF file.

#### pages
- **[pages/pages_launch.go](pages/pages_launch.go)** – Orchestrates various PDF page manipulation tasks including adding, moving, deleting, stamping, and extracting info.
  ```bash
  go run uses_cases/pages/*
  ```
- *[pages/page_save_as_image.go](pages/page_save_as_image.go)* – Converts a specified PDF page to a PNG image file.
- *[pages/pages_add.go](pages/pages_add.go)* – Appends a new, blank page to an existing PDF document.
- *[pages/pages_append_image_stamp.go](pages/pages_append_image_stamp.go)* – Adds an image stamp to a specified page of a PDF document.
- *[pages/pages_append_text_stamp.go](pages/pages_append_text_stamp.go)* – Adds a text stamp to a specified page of a PDF document.
- *[pages/pages_delete.go](pages/pages_delete.go)* – Deletes a specified page from a PDF document.
- *[pages/pages_extract_info.go](pages/pages_extract_info.go)* – Extracts and displays the dimensions and information of a specified PDF page.
- *[pages/pages_get_words_count.go](pages/pages_get_words_count.go)* – Calculates and reports the number of words on each page of a PDF document.
- *[pages/pages_move.go](pages/pages_move.go)* – Moves a page from one position to another within a PDF document.

#### parser
- **[parser/parser_launch.go](parser/parser_launch.go)** – Launches PDF parsing operations to extract forms, images, tables, and text boxes.
  ```bash
  go run uses_cases/parser/*
  ```
- *[parser/parse_get_forms_fdf.go](parser/parse_get_forms_fdf.go)* – Extracts interactive form field data from a PDF and saves it as an FDF file.
- *[parser/parse_get_forms_xml.go](parser/parse_get_forms_xml.go)* – Extracts interactive form field data from a PDF and saves it as an XML file.
- *[parser/parse_get_images.go](parser/parse_get_images.go)* – Extracts all images from a specified page of a PDF document and saves them as PNG files.
- *[parser/parse_get_tables.go](parser/parse_get_tables.go)* – Extracts all tables from a PDF document and saves their structured data as JSON.
- *[parser/parse_get_textboxes.go](parser/parse_get_textboxes.go)* – Extracts all text box form fields from a PDF document and saves their data as JSON.

#### stamps
- **[stamps/stamps_launcher.go](stamps/stamps_launcher.go)** – Executes operations to add and remove text and image stamps from PDF documents and pages.
  ```bash
  go run uses_cases/stamps/*
  ```
- *[stamps/stamps_append_image.go](stamps/stamps_append_image.go)* – Applies an image stamp to all pages of a PDF document.
- *[stamps/stamps_append_text.go](stamps/stamps_append_text.go)* – Applies a text stamp to all pages of a PDF document.
- *[stamps/stamps_delete.go](stamps/stamps_delete.go)* – Deletes all stamps from an entire PDF document.
- *[stamps/stamps_page_delete.go](stamps/stamps_page_delete.go)* – Deletes all stamps from a specific page of a PDF document.

#### tables
- **[tables/tables_launch.go](tables/tables_launch.go)** – Manages the lifecycle of tables in PDFs, including creation, extraction, replacement, and deletion.
  ```bash
  go run uses_cases/tables/*
  ```
- *[tables/tables_append.go](tables/tables_append.go)* – Appends a new, formatted table to a specified page of a PDF document.
- *[tables/tables_delete_all.go](tables/tables_delete_all.go)* – Deletes all tables from an entire PDF document.
- *[tables/tables_delete_by_id.go](tables/tables_delete_by_id.go)* – Deletes a specific table by its ID from a PDF document.
- *[tables/tables_delete_on_page.go](tables/tables_delete_on_page.go)* – Deletes all tables from a specific page of a PDF document.
- *[tables/tables_extract_all.go](tables/tables_extract_all.go)* – Extracts and lists information about all tables found in a PDF document.
- *[tables/tables_extract_by_id.go](tables/tables_extract_by_id.go)* – Extracts and displays the detailed structure of a specific table by its ID.
- *[tables/tables_extract_on_page.go](tables/tables_extract_on_page.go)* – Extracts and lists information about all tables found on a specific page of a PDF.
- *[tables/tables_replace.go](tables/tables_replace.go)* – Replaces an existing table in a PDF document with a newly defined table structure.

#### watermarks
- **[watermarks/watermarks_launch.go](watermarks/watermarks_launch.go)** – Coordinates operations to add, list, and remove image-based watermarks from PDF documents.
  ```bash
  go run uses_cases/watermarks/*
  ```
- *[watermarks/watermarks_extract.go](watermarks/watermarks_extract.go)* – Scans a PDF document and lists information about all image-based watermarks found.
- *[watermarks/watermarks_images_add.go](watermarks/watermarks_images_add.go)* – Adds an image as a semi-transparent watermark to all pages of a PDF document.
- *[watermarks/watermarks_images_delete.go](watermarks/watermarks_images_delete.go)* – Deletes a list of specific image watermarks by their IDs from a PDF document.
