# PdfApi

All URIs are relative to *https://api.aspose.cloud/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyFile**](PdfApi.md#CopyFile) | **Put** /pdf/storage/file/copy/{srcPath} | Copy file
[**CopyFolder**](PdfApi.md#CopyFolder) | **Put** /pdf/storage/folder/copy/{srcPath} | Copy folder
[**CreateFolder**](PdfApi.md#CreateFolder) | **Put** /pdf/storage/folder/{path} | Create the folder
[**DeleteAnnotation**](PdfApi.md#DeleteAnnotation) | **Delete** /pdf/{name}/annotations/{annotationId} | Delete document annotation by ID
[**DeleteBookmark**](PdfApi.md#DeleteBookmark) | **Delete** /pdf/{name}/bookmarks/bookmark/{bookmarkPath} | Delete document bookmark by ID.
[**DeleteDocumentAnnotations**](PdfApi.md#DeleteDocumentAnnotations) | **Delete** /pdf/{name}/annotations | Delete all annotations from the document
[**DeleteDocumentBookmarks**](PdfApi.md#DeleteDocumentBookmarks) | **Delete** /pdf/{name}/bookmarks/tree | Delete all document bookmarks.
[**DeleteDocumentLayer**](PdfApi.md#DeleteDocumentLayer) | **Delete** /pdf/{name}/layers | Remove document layer.
[**DeleteDocumentLinkAnnotations**](PdfApi.md#DeleteDocumentLinkAnnotations) | **Delete** /pdf/{name}/links | Delete all link annotations from the document
[**DeleteDocumentStamps**](PdfApi.md#DeleteDocumentStamps) | **Delete** /pdf/{name}/stamps | Delete all stamps from the document
[**DeleteDocumentTables**](PdfApi.md#DeleteDocumentTables) | **Delete** /pdf/{name}/tables | Delete all tables from the document
[**DeleteField**](PdfApi.md#DeleteField) | **Delete** /pdf/{name}/fields/{fieldName} | Delete document field by name.
[**DeleteFile**](PdfApi.md#DeleteFile) | **Delete** /pdf/storage/file/{path} | Delete file
[**DeleteFolder**](PdfApi.md#DeleteFolder) | **Delete** /pdf/storage/folder/{path} | Delete folder
[**DeleteImage**](PdfApi.md#DeleteImage) | **Delete** /pdf/{name}/images/{imageId} | Delete image from document page.
[**DeleteLinkAnnotation**](PdfApi.md#DeleteLinkAnnotation) | **Delete** /pdf/{name}/links/{linkId} | Delete document page link annotation by ID
[**DeletePage**](PdfApi.md#DeletePage) | **Delete** /pdf/{name}/pages/{pageNumber} | Delete document page by its number.
[**DeletePageAnnotations**](PdfApi.md#DeletePageAnnotations) | **Delete** /pdf/{name}/pages/{pageNumber}/annotations | Delete all annotations from the page
[**DeletePageLinkAnnotations**](PdfApi.md#DeletePageLinkAnnotations) | **Delete** /pdf/{name}/pages/{pageNumber}/links | Delete all link annotations from the page
[**DeletePageStamps**](PdfApi.md#DeletePageStamps) | **Delete** /pdf/{name}/pages/{pageNumber}/stamps | Delete all stamps from the page
[**DeletePageTables**](PdfApi.md#DeletePageTables) | **Delete** /pdf/{name}/pages/{pageNumber}/tables | Delete all tables from the page
[**DeleteProperties**](PdfApi.md#DeleteProperties) | **Delete** /pdf/{name}/documentproperties | Delete custom document properties.
[**DeleteProperty**](PdfApi.md#DeleteProperty) | **Delete** /pdf/{name}/documentproperties/{propertyName} | Delete document property.
[**DeleteStamp**](PdfApi.md#DeleteStamp) | **Delete** /pdf/{name}/stamps/{stampId} | Delete document stamp by ID
[**DeleteTable**](PdfApi.md#DeleteTable) | **Delete** /pdf/{name}/tables/{tableId} | Delete document table by ID
[**DownloadFile**](PdfApi.md#DownloadFile) | **Get** /pdf/storage/file/{path} | Download file
[**GetApsInStorageToPdf**](PdfApi.md#GetApsInStorageToPdf) | **Get** /pdf/create/aps | Convert APS file (located on storage) to PDF format and return resulting file in response. 
[**GetBookmark**](PdfApi.md#GetBookmark) | **Get** /pdf/{name}/bookmarks/bookmark/{bookmarkPath} | Read document bookmark.
[**GetBookmarks**](PdfApi.md#GetBookmarks) | **Get** /pdf/{name}/bookmarks/list/{bookmarkPath} | Read document bookmarks node list.
[**GetCaretAnnotation**](PdfApi.md#GetCaretAnnotation) | **Get** /pdf/{name}/annotations/caret/{annotationId} | Read document page caret annotation by ID.
[**GetCheckBoxField**](PdfApi.md#GetCheckBoxField) | **Get** /pdf/{name}/fields/checkbox/{fieldName} | Read document checkbox field by name.
[**GetCircleAnnotation**](PdfApi.md#GetCircleAnnotation) | **Get** /pdf/{name}/annotations/circle/{annotationId} | Read document page circle annotation by ID.
[**GetComboBoxField**](PdfApi.md#GetComboBoxField) | **Get** /pdf/{name}/fields/combobox/{fieldName} | Read document combobox field by name.
[**GetDiscUsage**](PdfApi.md#GetDiscUsage) | **Get** /pdf/storage/disc | Get disc usage
[**GetDocument**](PdfApi.md#GetDocument) | **Get** /pdf/{name} | Read common document info.
[**GetDocumentAnnotations**](PdfApi.md#GetDocumentAnnotations) | **Get** /pdf/{name}/annotations | Read document page annotations. Returns only FreeTextAnnotations, TextAnnotations, other annotations will implemented next releases.
[**GetDocumentAttachmentByIndex**](PdfApi.md#GetDocumentAttachmentByIndex) | **Get** /pdf/{name}/attachments/{attachmentIndex} | Read document attachment info by its index.
[**GetDocumentAttachments**](PdfApi.md#GetDocumentAttachments) | **Get** /pdf/{name}/attachments | Read document attachments info.
[**GetDocumentBookmarks**](PdfApi.md#GetDocumentBookmarks) | **Get** /pdf/{name}/bookmarks/tree | Read document bookmarks tree.
[**GetDocumentCaretAnnotations**](PdfApi.md#GetDocumentCaretAnnotations) | **Get** /pdf/{name}/annotations/caret | Read document caret annotations.
[**GetDocumentCheckBoxFields**](PdfApi.md#GetDocumentCheckBoxFields) | **Get** /pdf/{name}/fields/checkbox | Read document checkbox fields.
[**GetDocumentCircleAnnotations**](PdfApi.md#GetDocumentCircleAnnotations) | **Get** /pdf/{name}/annotations/circle | Read document circle annotations.
[**GetDocumentComboBoxFields**](PdfApi.md#GetDocumentComboBoxFields) | **Get** /pdf/{name}/fields/combobox | Read document combobox fields.
[**GetDocumentDisplayProperties**](PdfApi.md#GetDocumentDisplayProperties) | **Get** /pdf/{name}/displayproperties | Read document display properties.
[**GetDocumentFileAttachmentAnnotations**](PdfApi.md#GetDocumentFileAttachmentAnnotations) | **Get** /pdf/{name}/annotations/fileattachment | Read document FileAttachment annotations.
[**GetDocumentFreeTextAnnotations**](PdfApi.md#GetDocumentFreeTextAnnotations) | **Get** /pdf/{name}/annotations/freetext | Read document free text annotations.
[**GetDocumentHighlightAnnotations**](PdfApi.md#GetDocumentHighlightAnnotations) | **Get** /pdf/{name}/annotations/highlight | Read document highlight annotations.
[**GetDocumentInkAnnotations**](PdfApi.md#GetDocumentInkAnnotations) | **Get** /pdf/{name}/annotations/ink | Read document ink annotations.
[**GetDocumentLayers**](PdfApi.md#GetDocumentLayers) | **Get** /pdf/{name}/layers | Gets document layers.
[**GetDocumentLineAnnotations**](PdfApi.md#GetDocumentLineAnnotations) | **Get** /pdf/{name}/annotations/line | Read document line annotations.
[**GetDocumentListBoxFields**](PdfApi.md#GetDocumentListBoxFields) | **Get** /pdf/{name}/fields/listbox | Read document listbox fields.
[**GetDocumentMovieAnnotations**](PdfApi.md#GetDocumentMovieAnnotations) | **Get** /pdf/{name}/annotations/movie | Read document movie annotations.
[**GetDocumentPolyLineAnnotations**](PdfApi.md#GetDocumentPolyLineAnnotations) | **Get** /pdf/{name}/annotations/polyline | Read document polyline annotations.
[**GetDocumentPolygonAnnotations**](PdfApi.md#GetDocumentPolygonAnnotations) | **Get** /pdf/{name}/annotations/polygon | Read document polygon annotations.
[**GetDocumentPopupAnnotations**](PdfApi.md#GetDocumentPopupAnnotations) | **Get** /pdf/{name}/annotations/popup | Read document popup annotations.
[**GetDocumentPopupAnnotationsByParent**](PdfApi.md#GetDocumentPopupAnnotationsByParent) | **Get** /pdf/{name}/annotations/{annotationId}/popup | Read document popup annotations by parent id.
[**GetDocumentProperties**](PdfApi.md#GetDocumentProperties) | **Get** /pdf/{name}/documentproperties | Read document properties.
[**GetDocumentProperty**](PdfApi.md#GetDocumentProperty) | **Get** /pdf/{name}/documentproperties/{propertyName} | Read document property by name.
[**GetDocumentRadioButtonFields**](PdfApi.md#GetDocumentRadioButtonFields) | **Get** /pdf/{name}/fields/radiobutton | Read document radiobutton fields.
[**GetDocumentRedactionAnnotations**](PdfApi.md#GetDocumentRedactionAnnotations) | **Get** /pdf/{name}/annotations/redaction | Read document redaction annotations.
[**GetDocumentScreenAnnotations**](PdfApi.md#GetDocumentScreenAnnotations) | **Get** /pdf/{name}/annotations/screen | Read document screen annotations.
[**GetDocumentSignatureFields**](PdfApi.md#GetDocumentSignatureFields) | **Get** /pdf/{name}/fields/signature | Read document signature fields.
[**GetDocumentSoundAnnotations**](PdfApi.md#GetDocumentSoundAnnotations) | **Get** /pdf/{name}/annotations/sound | Read document sound annotations.
[**GetDocumentSquareAnnotations**](PdfApi.md#GetDocumentSquareAnnotations) | **Get** /pdf/{name}/annotations/square | Read document square annotations.
[**GetDocumentSquigglyAnnotations**](PdfApi.md#GetDocumentSquigglyAnnotations) | **Get** /pdf/{name}/annotations/squiggly | Read document squiggly annotations.
[**GetDocumentStampAnnotations**](PdfApi.md#GetDocumentStampAnnotations) | **Get** /pdf/{name}/annotations/stamp | Read document stamp annotations.
[**GetDocumentStamps**](PdfApi.md#GetDocumentStamps) | **Get** /pdf/{name}/stamps | Read document stamps.
[**GetDocumentStrikeOutAnnotations**](PdfApi.md#GetDocumentStrikeOutAnnotations) | **Get** /pdf/{name}/annotations/strikeout | Read document StrikeOut annotations.
[**GetDocumentTables**](PdfApi.md#GetDocumentTables) | **Get** /pdf/{name}/tables | Read document tables.
[**GetDocumentTextAnnotations**](PdfApi.md#GetDocumentTextAnnotations) | **Get** /pdf/{name}/annotations/text | Read document text annotations.
[**GetDocumentTextBoxFields**](PdfApi.md#GetDocumentTextBoxFields) | **Get** /pdf/{name}/fields/textbox | Read document text box fields.
[**GetDocumentUnderlineAnnotations**](PdfApi.md#GetDocumentUnderlineAnnotations) | **Get** /pdf/{name}/annotations/underline | Read document underline annotations.
[**GetDownloadDocumentAttachmentByIndex**](PdfApi.md#GetDownloadDocumentAttachmentByIndex) | **Get** /pdf/{name}/attachments/{attachmentIndex}/download | Download document attachment content by its index.
[**GetEpubInStorageToPdf**](PdfApi.md#GetEpubInStorageToPdf) | **Get** /pdf/create/epub | Convert EPUB file (located on storage) to PDF format and return resulting file in response. 
[**GetExportFieldsFromPdfToFdfInStorage**](PdfApi.md#GetExportFieldsFromPdfToFdfInStorage) | **Get** /pdf/{name}/export/fdf | Export fields from from PDF in storage to FDF file.
[**GetExportFieldsFromPdfToXfdfInStorage**](PdfApi.md#GetExportFieldsFromPdfToXfdfInStorage) | **Get** /pdf/{name}/export/xfdf | Export fields from from PDF in storage to XFDF file.
[**GetExportFieldsFromPdfToXmlInStorage**](PdfApi.md#GetExportFieldsFromPdfToXmlInStorage) | **Get** /pdf/{name}/export/xml | Export fields from from PDF in storage to XML file.
[**GetField**](PdfApi.md#GetField) | **Get** /pdf/{name}/fields/{fieldName} | Get document field by name.
[**GetFields**](PdfApi.md#GetFields) | **Get** /pdf/{name}/fields | Get document fields.
[**GetFileAttachmentAnnotation**](PdfApi.md#GetFileAttachmentAnnotation) | **Get** /pdf/{name}/annotations/fileattachment/{annotationId} | Read document page FileAttachment annotation by ID.
[**GetFileAttachmentAnnotationData**](PdfApi.md#GetFileAttachmentAnnotationData) | **Get** /pdf/{name}/annotations/fileattachment/{annotationId}/data | Read document page FileAttachment annotation by ID.
[**GetFileVersions**](PdfApi.md#GetFileVersions) | **Get** /pdf/storage/version/{path} | Get file versions
[**GetFilesList**](PdfApi.md#GetFilesList) | **Get** /pdf/storage/folder/{path} | Get all files and folders within a folder
[**GetFreeTextAnnotation**](PdfApi.md#GetFreeTextAnnotation) | **Get** /pdf/{name}/annotations/freetext/{annotationId} | Read document page free text annotation by ID.
[**GetHighlightAnnotation**](PdfApi.md#GetHighlightAnnotation) | **Get** /pdf/{name}/annotations/highlight/{annotationId} | Read document page highlight annotation by ID.
[**GetHtmlInStorageToPdf**](PdfApi.md#GetHtmlInStorageToPdf) | **Get** /pdf/create/html | Convert HTML file (located on storage) to PDF format and return resulting file in response. 
[**GetImage**](PdfApi.md#GetImage) | **Get** /pdf/{name}/images/{imageId} | Read document image by ID.
[**GetImageExtractAsGif**](PdfApi.md#GetImageExtractAsGif) | **Get** /pdf/{name}/images/{imageId}/extract/gif | Extract document image in GIF format
[**GetImageExtractAsJpeg**](PdfApi.md#GetImageExtractAsJpeg) | **Get** /pdf/{name}/images/{imageId}/extract/jpeg | Extract document image in JPEG format
[**GetImageExtractAsPng**](PdfApi.md#GetImageExtractAsPng) | **Get** /pdf/{name}/images/{imageId}/extract/png | Extract document image in PNG format
[**GetImageExtractAsTiff**](PdfApi.md#GetImageExtractAsTiff) | **Get** /pdf/{name}/images/{imageId}/extract/tiff | Extract document image in TIFF format
[**GetImages**](PdfApi.md#GetImages) | **Get** /pdf/{name}/pages/{pageNumber}/images | Read document images.
[**GetImagesExtractSvg**](PdfApi.md#GetImagesExtractSvg) | **Get** /pdf/{name}/pages/{pageNumber}/images/extract/svg | Extract SVG images from document page.
[**GetImportFieldsFromFdfInStorage**](PdfApi.md#GetImportFieldsFromFdfInStorage) | **Get** /pdf/{name}/import/fdf | Update fields from FDF file in storage.
[**GetImportFieldsFromXfdfInStorage**](PdfApi.md#GetImportFieldsFromXfdfInStorage) | **Get** /pdf/{name}/import/xfdf | Update fields from XFDF file in storage.
[**GetImportFieldsFromXmlInStorage**](PdfApi.md#GetImportFieldsFromXmlInStorage) | **Get** /pdf/{name}/import/xml | Import from XML file (located on storage) to PDF format and return resulting file in response. 
[**GetInkAnnotation**](PdfApi.md#GetInkAnnotation) | **Get** /pdf/{name}/annotations/ink/{annotationId} | Read document page ink annotation by ID.
[**GetLineAnnotation**](PdfApi.md#GetLineAnnotation) | **Get** /pdf/{name}/annotations/line/{annotationId} | Read document page line annotation by ID.
[**GetLinkAnnotation**](PdfApi.md#GetLinkAnnotation) | **Get** /pdf/{name}/links/{linkId} | Read document link annotation by ID.
[**GetListBoxField**](PdfApi.md#GetListBoxField) | **Get** /pdf/{name}/fields/listbox/{fieldName} | Read document listbox field by name.
[**GetMarkdownInStorageToPdf**](PdfApi.md#GetMarkdownInStorageToPdf) | **Get** /pdf/create/markdown | Convert MD file (located on storage) to PDF format and return resulting file in response. 
[**GetMhtInStorageToPdf**](PdfApi.md#GetMhtInStorageToPdf) | **Get** /pdf/create/mht | Convert MHT file (located on storage) to PDF format and return resulting file in response. 
[**GetMovieAnnotation**](PdfApi.md#GetMovieAnnotation) | **Get** /pdf/{name}/annotations/movie/{annotationId} | Read document page movie annotation by ID.
[**GetPage**](PdfApi.md#GetPage) | **Get** /pdf/{name}/pages/{pageNumber} | Read document page info.
[**GetPageAnnotations**](PdfApi.md#GetPageAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations | Read document page annotations. Returns only FreeTextAnnotations, TextAnnotations, other annotations will implemented next releases.
[**GetPageCaretAnnotations**](PdfApi.md#GetPageCaretAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/caret | Read document page caret annotations.
[**GetPageCheckBoxFields**](PdfApi.md#GetPageCheckBoxFields) | **Get** /pdf/{name}/page/{pageNumber}/fields/checkbox | Read document page checkbox fields.
[**GetPageCircleAnnotations**](PdfApi.md#GetPageCircleAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/circle | Read document page circle annotations.
[**GetPageComboBoxFields**](PdfApi.md#GetPageComboBoxFields) | **Get** /pdf/{name}/page/{pageNumber}/fields/combobox | Read document page combobox fields.
[**GetPageConvertToBmp**](PdfApi.md#GetPageConvertToBmp) | **Get** /pdf/{name}/pages/{pageNumber}/convert/bmp | Convert document page to Bmp image and return resulting file in response.
[**GetPageConvertToEmf**](PdfApi.md#GetPageConvertToEmf) | **Get** /pdf/{name}/pages/{pageNumber}/convert/emf | Convert document page to Emf image and return resulting file in response.
[**GetPageConvertToGif**](PdfApi.md#GetPageConvertToGif) | **Get** /pdf/{name}/pages/{pageNumber}/convert/gif | Convert document page to Gif image and return resulting file in response.
[**GetPageConvertToJpeg**](PdfApi.md#GetPageConvertToJpeg) | **Get** /pdf/{name}/pages/{pageNumber}/convert/jpeg | Convert document page to Jpeg image and return resulting file in response.
[**GetPageConvertToPng**](PdfApi.md#GetPageConvertToPng) | **Get** /pdf/{name}/pages/{pageNumber}/convert/png | Convert document page to Png image and return resulting file in response.
[**GetPageConvertToTiff**](PdfApi.md#GetPageConvertToTiff) | **Get** /pdf/{name}/pages/{pageNumber}/convert/tiff | Convert document page to Tiff image  and return resulting file in response.
[**GetPageFileAttachmentAnnotations**](PdfApi.md#GetPageFileAttachmentAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/fileattachment | Read document page FileAttachment annotations.
[**GetPageFreeTextAnnotations**](PdfApi.md#GetPageFreeTextAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/freetext | Read document page free text annotations.
[**GetPageHighlightAnnotations**](PdfApi.md#GetPageHighlightAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/highlight | Read document page highlight annotations.
[**GetPageInkAnnotations**](PdfApi.md#GetPageInkAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/ink | Read document page ink annotations.
[**GetPageLineAnnotations**](PdfApi.md#GetPageLineAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/line | Read document page line annotations.
[**GetPageLinkAnnotation**](PdfApi.md#GetPageLinkAnnotation) | **Get** /pdf/{name}/pages/{pageNumber}/links/{linkId} | Read document page link annotation by ID.
[**GetPageLinkAnnotations**](PdfApi.md#GetPageLinkAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/links | Read document page link annotations.
[**GetPageListBoxFields**](PdfApi.md#GetPageListBoxFields) | **Get** /pdf/{name}/page/{pageNumber}/fields/listbox | Read document page listbox fields.
[**GetPageMovieAnnotations**](PdfApi.md#GetPageMovieAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/movie | Read document page movie annotations.
[**GetPagePolyLineAnnotations**](PdfApi.md#GetPagePolyLineAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/polyline | Read document page polyline annotations.
[**GetPagePolygonAnnotations**](PdfApi.md#GetPagePolygonAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/polygon | Read document page polygon annotations.
[**GetPagePopupAnnotations**](PdfApi.md#GetPagePopupAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/popup | Read document page popup annotations.
[**GetPageRadioButtonFields**](PdfApi.md#GetPageRadioButtonFields) | **Get** /pdf/{name}/page/{pageNumber}/fields/radiobutton | Read document page radiobutton fields.
[**GetPageRedactionAnnotations**](PdfApi.md#GetPageRedactionAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/redaction | Read document page redaction annotations.
[**GetPageScreenAnnotations**](PdfApi.md#GetPageScreenAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/screen | Read document page screen annotations.
[**GetPageSignatureFields**](PdfApi.md#GetPageSignatureFields) | **Get** /pdf/{name}/page/{pageNumber}/fields/signature | Read document page signature fields.
[**GetPageSoundAnnotations**](PdfApi.md#GetPageSoundAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/sound | Read document page sound annotations.
[**GetPageSquareAnnotations**](PdfApi.md#GetPageSquareAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/square | Read document page square annotations.
[**GetPageSquigglyAnnotations**](PdfApi.md#GetPageSquigglyAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/squiggly | Read document page squiggly annotations.
[**GetPageStampAnnotations**](PdfApi.md#GetPageStampAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/stamp | Read document page stamp annotations.
[**GetPageStamps**](PdfApi.md#GetPageStamps) | **Get** /pdf/{name}/pages/{pageNumber}/stamps | Read page document stamps.
[**GetPageStrikeOutAnnotations**](PdfApi.md#GetPageStrikeOutAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/strikeout | Read document page StrikeOut annotations.
[**GetPageTables**](PdfApi.md#GetPageTables) | **Get** /pdf/{name}/pages/{pageNumber}/tables | Read document page tables.
[**GetPageText**](PdfApi.md#GetPageText) | **Get** /pdf/{name}/pages/{pageNumber}/text | Read page text items.
[**GetPageTextAnnotations**](PdfApi.md#GetPageTextAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/text | Read document page text annotations.
[**GetPageTextBoxFields**](PdfApi.md#GetPageTextBoxFields) | **Get** /pdf/{name}/page/{pageNumber}/fields/textbox | Read document page text box fields.
[**GetPageUnderlineAnnotations**](PdfApi.md#GetPageUnderlineAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/underline | Read document page underline annotations.
[**GetPages**](PdfApi.md#GetPages) | **Get** /pdf/{name}/pages | Read document pages info.
[**GetPclInStorageToPdf**](PdfApi.md#GetPclInStorageToPdf) | **Get** /pdf/create/pcl | Convert PCL file (located on storage) to PDF format and return resulting file in response. 
[**GetPdfAInStorageToPdf**](PdfApi.md#GetPdfAInStorageToPdf) | **Get** /pdf/create/pdfa | Convert PDFA file (located on storage) to PDF format and return resulting file in response. 
[**GetPdfInStorageToAps**](PdfApi.md#GetPdfInStorageToAps) | **Get** /pdf/{name}/convert/aps | Converts PDF document (located on storage) to APS format and returns resulting file in response content
[**GetPdfInStorageToDoc**](PdfApi.md#GetPdfInStorageToDoc) | **Get** /pdf/{name}/convert/doc | Converts PDF document (located on storage) to DOC format and returns resulting file in response content.
[**GetPdfInStorageToEpub**](PdfApi.md#GetPdfInStorageToEpub) | **Get** /pdf/{name}/convert/epub | Converts PDF document (located on storage) to EPUB format and returns resulting file in response content
[**GetPdfInStorageToHtml**](PdfApi.md#GetPdfInStorageToHtml) | **Get** /pdf/{name}/convert/html | Converts PDF document (located on storage) to Html format and returns resulting file in response content
[**GetPdfInStorageToMobiXml**](PdfApi.md#GetPdfInStorageToMobiXml) | **Get** /pdf/{name}/convert/mobixml | Converts PDF document (located on storage) to MOBIXML format and returns resulting ZIP archive file in response content.
[**GetPdfInStorageToPdfA**](PdfApi.md#GetPdfInStorageToPdfA) | **Get** /pdf/{name}/convert/pdfa | Converts PDF document (located on storage) to PdfA format and returns resulting file in response content
[**GetPdfInStorageToPptx**](PdfApi.md#GetPdfInStorageToPptx) | **Get** /pdf/{name}/convert/pptx | Converts PDF document (located on storage) to PPTX format and returns resulting file in response content
[**GetPdfInStorageToSvg**](PdfApi.md#GetPdfInStorageToSvg) | **Get** /pdf/{name}/convert/svg | Converts PDF document (located on storage) to SVG format and returns resulting file in response content
[**GetPdfInStorageToTeX**](PdfApi.md#GetPdfInStorageToTeX) | **Get** /pdf/{name}/convert/tex | Converts PDF document (located on storage) to TeX format and returns resulting file in response content
[**GetPdfInStorageToText**](PdfApi.md#GetPdfInStorageToText) | **Get** /pdf/{name}/convert/text | Converts PDF document (located on storage) to Text format and returns resulting file in response content
[**GetPdfInStorageToTiff**](PdfApi.md#GetPdfInStorageToTiff) | **Get** /pdf/{name}/convert/tiff | Converts PDF document (located on storage) to TIFF format and returns resulting file in response content
[**GetPdfInStorageToXls**](PdfApi.md#GetPdfInStorageToXls) | **Get** /pdf/{name}/convert/xls | Converts PDF document (located on storage) to XLS format and returns resulting file in response content
[**GetPdfInStorageToXlsx**](PdfApi.md#GetPdfInStorageToXlsx) | **Get** /pdf/{name}/convert/xlsx | Converts PDF document (located on storage) to XLSX format and returns resulting file in response content
[**GetPdfInStorageToXml**](PdfApi.md#GetPdfInStorageToXml) | **Get** /pdf/{name}/convert/xml | Converts PDF document (located on storage) to XML format and returns resulting file in response content
[**GetPdfInStorageToXps**](PdfApi.md#GetPdfInStorageToXps) | **Get** /pdf/{name}/convert/xps | Converts PDF document (located on storage) to XPS format and returns resulting file in response content
[**GetPolyLineAnnotation**](PdfApi.md#GetPolyLineAnnotation) | **Get** /pdf/{name}/annotations/polyline/{annotationId} | Read document page polyline annotation by ID.
[**GetPolygonAnnotation**](PdfApi.md#GetPolygonAnnotation) | **Get** /pdf/{name}/annotations/polygon/{annotationId} | Read document page polygon annotation by ID.
[**GetPopupAnnotation**](PdfApi.md#GetPopupAnnotation) | **Get** /pdf/{name}/annotations/popup/{annotationId} | Read document page popup annotation by ID.
[**GetPsInStorageToPdf**](PdfApi.md#GetPsInStorageToPdf) | **Get** /pdf/create/ps | Convert PS file (located on storage) to PDF format and return resulting file in response. 
[**GetRadioButtonField**](PdfApi.md#GetRadioButtonField) | **Get** /pdf/{name}/fields/radiobutton/{fieldName} | Read document RadioButton field by name.
[**GetRedactionAnnotation**](PdfApi.md#GetRedactionAnnotation) | **Get** /pdf/{name}/annotations/redaction/{annotationId} | Read document page redaction annotation by ID.
[**GetScreenAnnotation**](PdfApi.md#GetScreenAnnotation) | **Get** /pdf/{name}/annotations/screen/{annotationId} | Read document page screen annotation by ID.
[**GetScreenAnnotationData**](PdfApi.md#GetScreenAnnotationData) | **Get** /pdf/{name}/annotations/screen/{annotationId}/data | Read document page screen annotation by ID.
[**GetSignatureField**](PdfApi.md#GetSignatureField) | **Get** /pdf/{name}/fields/signature/{fieldName} | Read document signature field by name.
[**GetSoundAnnotation**](PdfApi.md#GetSoundAnnotation) | **Get** /pdf/{name}/annotations/sound/{annotationId} | Read document page sound annotation by ID.
[**GetSoundAnnotationData**](PdfApi.md#GetSoundAnnotationData) | **Get** /pdf/{name}/annotations/sound/{annotationId}/data | Read document page sound annotation by ID.
[**GetSquareAnnotation**](PdfApi.md#GetSquareAnnotation) | **Get** /pdf/{name}/annotations/square/{annotationId} | Read document page square annotation by ID.
[**GetSquigglyAnnotation**](PdfApi.md#GetSquigglyAnnotation) | **Get** /pdf/{name}/annotations/squiggly/{annotationId} | Read document page squiggly annotation by ID.
[**GetStampAnnotation**](PdfApi.md#GetStampAnnotation) | **Get** /pdf/{name}/annotations/stamp/{annotationId} | Read document page stamp annotation by ID.
[**GetStampAnnotationData**](PdfApi.md#GetStampAnnotationData) | **Get** /pdf/{name}/annotations/stamp/{annotationId}/data | Read document page stamp annotation by ID.
[**GetStrikeOutAnnotation**](PdfApi.md#GetStrikeOutAnnotation) | **Get** /pdf/{name}/annotations/strikeout/{annotationId} | Read document page StrikeOut annotation by ID.
[**GetSvgInStorageToPdf**](PdfApi.md#GetSvgInStorageToPdf) | **Get** /pdf/create/svg | Convert SVG file (located on storage) to PDF format and return resulting file in response. 
[**GetTable**](PdfApi.md#GetTable) | **Get** /pdf/{name}/tables/{tableId} | Read document page table by ID.
[**GetTeXInStorageToPdf**](PdfApi.md#GetTeXInStorageToPdf) | **Get** /pdf/create/tex | Convert TeX file (located on storage) to PDF format and return resulting file in response. 
[**GetText**](PdfApi.md#GetText) | **Get** /pdf/{name}/text | Read document text.
[**GetTextAnnotation**](PdfApi.md#GetTextAnnotation) | **Get** /pdf/{name}/annotations/text/{annotationId} | Read document page text annotation by ID.
[**GetTextBoxField**](PdfApi.md#GetTextBoxField) | **Get** /pdf/{name}/fields/textbox/{fieldName} | Read document text box field by name.
[**GetUnderlineAnnotation**](PdfApi.md#GetUnderlineAnnotation) | **Get** /pdf/{name}/annotations/underline/{annotationId} | Read document page underline annotation by ID.
[**GetVerifySignature**](PdfApi.md#GetVerifySignature) | **Get** /pdf/{name}/verifySignature | Verify signature document.
[**GetWebInStorageToPdf**](PdfApi.md#GetWebInStorageToPdf) | **Get** /pdf/create/web | Convert web page to PDF format and return resulting file in response. 
[**GetWordsPerPage**](PdfApi.md#GetWordsPerPage) | **Get** /pdf/{name}/pages/wordCount | Get number of words per document page.
[**GetXfaPdfInStorageToAcroForm**](PdfApi.md#GetXfaPdfInStorageToAcroForm) | **Get** /pdf/{name}/convert/xfatoacroform | Converts PDF document which contains XFA form (located on storage) to PDF with AcroForm and returns resulting file response content
[**GetXmlInStorageToPdf**](PdfApi.md#GetXmlInStorageToPdf) | **Get** /pdf/create/xml | Convert XML file (located on storage) to PDF format and return resulting file in response. 
[**GetXmpMetadataJson**](PdfApi.md#GetXmpMetadataJson) | **Get** /pdf/{name}/xmpmetadata/json | Gets document XMP Metadata as JSON.
[**GetXmpMetadataXml**](PdfApi.md#GetXmpMetadataXml) | **Get** /pdf/{name}/xmpmetadata/xml | Gets document XMP Metadata as XML file.
[**GetXpsInStorageToPdf**](PdfApi.md#GetXpsInStorageToPdf) | **Get** /pdf/create/xps | Convert XPS file (located on storage) to PDF format and return resulting file in response. 
[**GetXslFoInStorageToPdf**](PdfApi.md#GetXslFoInStorageToPdf) | **Get** /pdf/create/xslfo | Convert XslFo file (located on storage) to PDF format and return resulting file in response. 
[**MoveFile**](PdfApi.md#MoveFile) | **Put** /pdf/storage/file/move/{srcPath} | Move file
[**MoveFolder**](PdfApi.md#MoveFolder) | **Put** /pdf/storage/folder/move/{srcPath} | Move folder
[**ObjectExists**](PdfApi.md#ObjectExists) | **Get** /pdf/storage/exist/{path} | Check if file or folder exists
[**PostAddDocumentAttachment**](PdfApi.md#PostAddDocumentAttachment) | **Post** /pdf/{name}/attachments | Adds a file attachment to the PDF document.
[**PostAppendDocument**](PdfApi.md#PostAppendDocument) | **Post** /pdf/{name}/appendDocument | Append document to existing one.
[**PostBookmark**](PdfApi.md#PostBookmark) | **Post** /pdf/{name}/bookmarks/bookmark/{bookmarkPath} | Add document bookmarks.
[**PostChangePasswordDocumentInStorage**](PdfApi.md#PostChangePasswordDocumentInStorage) | **Post** /pdf/{name}/changepassword | Change document password in storage.
[**PostCheckBoxFields**](PdfApi.md#PostCheckBoxFields) | **Post** /pdf/{name}/fields/checkbox | Add document checkbox fields.
[**PostComboBoxFields**](PdfApi.md#PostComboBoxFields) | **Post** /pdf/{name}/fields/combobox | Add document combobox fields.
[**PostComparePdf**](PdfApi.md#PostComparePdf) | **Post** /pdf/compare | Compare two PDF documents.
[**PostCreateDocument**](PdfApi.md#PostCreateDocument) | **Post** /pdf/{name} | Create empty document.
[**PostCreateField**](PdfApi.md#PostCreateField) | **Post** /pdf/{name}/fields | Create field.
[**PostDecryptDocumentInStorage**](PdfApi.md#PostDecryptDocumentInStorage) | **Post** /pdf/{name}/decrypt | Decrypt document in storage.
[**PostDocumentImageFooter**](PdfApi.md#PostDocumentImageFooter) | **Post** /pdf/{name}/footer/image | Add document image footer.
[**PostDocumentImageHeader**](PdfApi.md#PostDocumentImageHeader) | **Post** /pdf/{name}/header/image | Add document image header.
[**PostDocumentImageStamps**](PdfApi.md#PostDocumentImageStamps) | **Post** /pdf/{name}/stamps/image | Add document pages image stamps.
[**PostDocumentImageStampsPageSpecified**](PdfApi.md#PostDocumentImageStampsPageSpecified) | **Post** /pdf/{name}/stamps/image/pagespecified | Add document image stamps to specified pages.
[**PostDocumentPageNumberStamps**](PdfApi.md#PostDocumentPageNumberStamps) | **Post** /pdf/{name}/stamps/pagenumber | Add document page number stamps.
[**PostDocumentPagesCrop**](PdfApi.md#PostDocumentPagesCrop) | **Post** /pdf/{name}/crop | Crop PDF document pages.
[**PostDocumentPagesResize**](PdfApi.md#PostDocumentPagesResize) | **Post** /pdf/{name}/resize | Rsize PDF document.
[**PostDocumentPagesRotate**](PdfApi.md#PostDocumentPagesRotate) | **Post** /pdf/{name}/rotate | Rotate PDF document.
[**PostDocumentTextFooter**](PdfApi.md#PostDocumentTextFooter) | **Post** /pdf/{name}/footer/text | Add document text footer.
[**PostDocumentTextHeader**](PdfApi.md#PostDocumentTextHeader) | **Post** /pdf/{name}/header/text | Add document text header.
[**PostDocumentTextReplace**](PdfApi.md#PostDocumentTextReplace) | **Post** /pdf/{name}/text/replace | Document&#39;s replace text method.
[**PostDocumentTextStamps**](PdfApi.md#PostDocumentTextStamps) | **Post** /pdf/{name}/stamps/text | Add document pages text stamps.
[**PostDocumentTextStampsPageSpecified**](PdfApi.md#PostDocumentTextStampsPageSpecified) | **Post** /pdf/{name}/stamps/text/pagespecified | Add document text stamps to specified pages.
[**PostEncryptDocumentInStorage**](PdfApi.md#PostEncryptDocumentInStorage) | **Post** /pdf/{name}/encrypt | Encrypt document in storage.
[**PostFlattenDocument**](PdfApi.md#PostFlattenDocument) | **Post** /pdf/{name}/flatten | Flatten the document.
[**PostHtmlToPdf**](PdfApi.md#PostHtmlToPdf) | **Post** /pdf/create/html | Convert HTML file (zip archive in request content) to PDF format and return resulting file in response. 
[**PostImportFieldsFromFdf**](PdfApi.md#PostImportFieldsFromFdf) | **Post** /pdf/{name}/import/fdf | Update fields from FDF file in request.
[**PostImportFieldsFromXfdf**](PdfApi.md#PostImportFieldsFromXfdf) | **Post** /pdf/{name}/import/xfdf | Update fields from XFDF file in request.
[**PostImportFieldsFromXml**](PdfApi.md#PostImportFieldsFromXml) | **Post** /pdf/{name}/import/xml | Update fields from XML file in request.
[**PostInsertImage**](PdfApi.md#PostInsertImage) | **Post** /pdf/{name}/pages/{pageNumber}/images | Insert image to document page.
[**PostListBoxFields**](PdfApi.md#PostListBoxFields) | **Post** /pdf/{name}/fields/listbox | Add document listbox fields.
[**PostMovePage**](PdfApi.md#PostMovePage) | **Post** /pdf/{name}/pages/{pageNumber}/movePage | Move page to new position.
[**PostOptimizeDocument**](PdfApi.md#PostOptimizeDocument) | **Post** /pdf/{name}/optimize | Optimize document.
[**PostOrganizeDocument**](PdfApi.md#PostOrganizeDocument) | **Post** /pdf/{name}/organize | Merge selected pages of a document.
[**PostOrganizeDocuments**](PdfApi.md#PostOrganizeDocuments) | **Post** /pdf/organize | Merge selected pages of different documents.
[**PostPageCaretAnnotations**](PdfApi.md#PostPageCaretAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/caret | Add document page caret annotations.
[**PostPageCertify**](PdfApi.md#PostPageCertify) | **Post** /pdf/{name}/pages/{pageNumber}/certify | Certify document page.
[**PostPageCircleAnnotations**](PdfApi.md#PostPageCircleAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/circle | Add document page circle annotations.
[**PostPageFileAttachmentAnnotations**](PdfApi.md#PostPageFileAttachmentAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/fileattachment | Add document page FileAttachment annotations.
[**PostPageFreeTextAnnotations**](PdfApi.md#PostPageFreeTextAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/freetext | Add document page free text annotations.
[**PostPageHighlightAnnotations**](PdfApi.md#PostPageHighlightAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/highlight | Add document page highlight annotations.
[**PostPageImageStamps**](PdfApi.md#PostPageImageStamps) | **Post** /pdf/{name}/pages/{pageNumber}/stamps/image | Add document page image stamps.
[**PostPageInkAnnotations**](PdfApi.md#PostPageInkAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/ink | Add document page ink annotations.
[**PostPageLineAnnotations**](PdfApi.md#PostPageLineAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/line | Add document page line annotations.
[**PostPageLinkAnnotations**](PdfApi.md#PostPageLinkAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/links | Add document page link annotations.
[**PostPageMovieAnnotations**](PdfApi.md#PostPageMovieAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/movie | Add document page movie annotations.
[**PostPagePdfPageStamps**](PdfApi.md#PostPagePdfPageStamps) | **Post** /pdf/{name}/pages/{pageNumber}/stamps/pdfpage | Add document pdf page stamps.
[**PostPagePolyLineAnnotations**](PdfApi.md#PostPagePolyLineAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/polyline | Add document page polyline annotations.
[**PostPagePolygonAnnotations**](PdfApi.md#PostPagePolygonAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/polygon | Add document page polygon annotations.
[**PostPageRedactionAnnotations**](PdfApi.md#PostPageRedactionAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/redaction | Add document page redaction annotations.
[**PostPageScreenAnnotations**](PdfApi.md#PostPageScreenAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/screen | Add document page screen annotations.
[**PostPageSoundAnnotations**](PdfApi.md#PostPageSoundAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/sound | Add document page sound annotations.
[**PostPageSquareAnnotations**](PdfApi.md#PostPageSquareAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/square | Add document page square annotations.
[**PostPageSquigglyAnnotations**](PdfApi.md#PostPageSquigglyAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/squiggly | Add document page squiggly annotations.
[**PostPageStampAnnotations**](PdfApi.md#PostPageStampAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/stamp | Add document page stamp annotations.
[**PostPageStrikeOutAnnotations**](PdfApi.md#PostPageStrikeOutAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/strikeout | Add document page StrikeOut annotations.
[**PostPageTables**](PdfApi.md#PostPageTables) | **Post** /pdf/{name}/pages/{pageNumber}/tables | Add document page tables.
[**PostPageTextAnnotations**](PdfApi.md#PostPageTextAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/text | Add document page text annotations.
[**PostPageTextReplace**](PdfApi.md#PostPageTextReplace) | **Post** /pdf/{name}/pages/{pageNumber}/text/replace | Page&#39;s replace text method.
[**PostPageTextStamps**](PdfApi.md#PostPageTextStamps) | **Post** /pdf/{name}/pages/{pageNumber}/stamps/text | Add document page text stamps.
[**PostPageUnderlineAnnotations**](PdfApi.md#PostPageUnderlineAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/underline | Add document page underline annotations.
[**PostPdfToXlsx**](PdfApi.md#PostPdfToXlsx) | **Post** /pdf/convert/xlsx | Converts PDF document (in request content) to XLSX format and uploads and returns resulting file in response content.
[**PostPopupAnnotation**](PdfApi.md#PostPopupAnnotation) | **Post** /pdf/{name}/annotations/{annotationId}/popup | Add document popup annotations.
[**PostRadioButtonFields**](PdfApi.md#PostRadioButtonFields) | **Post** /pdf/{name}/fields/radiobutton | Add document RadioButton fields.
[**PostSignDocument**](PdfApi.md#PostSignDocument) | **Post** /pdf/{name}/sign | Sign document.
[**PostSignPage**](PdfApi.md#PostSignPage) | **Post** /pdf/{name}/pages/{pageNumber}/sign | Sign page.
[**PostSignatureField**](PdfApi.md#PostSignatureField) | **Post** /pdf/{name}/fields/signature | Add document signature field.
[**PostSplitDocument**](PdfApi.md#PostSplitDocument) | **Post** /pdf/{name}/split | Split document to parts.
[**PostSplitRangePdfDocument**](PdfApi.md#PostSplitRangePdfDocument) | **Post** /pdf/{name}/splitrangepdf | Split document into ranges.
[**PostTextBoxFields**](PdfApi.md#PostTextBoxFields) | **Post** /pdf/{name}/fields/textbox | Add document text box fields.
[**PostXmpMetadata**](PdfApi.md#PostXmpMetadata) | **Post** /pdf/{name}/xmpmetadata | Add or remove XMP Metadata properties.
[**PutAddNewPage**](PdfApi.md#PutAddNewPage) | **Put** /pdf/{name}/pages | Add new page to end of the document.
[**PutAddText**](PdfApi.md#PutAddText) | **Put** /pdf/{name}/pages/{pageNumber}/text | Add text to PDF document page.
[**PutAnnotationsFlatten**](PdfApi.md#PutAnnotationsFlatten) | **Put** /pdf/{name}/annotations/flatten | Flattens the annotations of the specified types
[**PutApsInStorageToPdf**](PdfApi.md#PutApsInStorageToPdf) | **Put** /pdf/{name}/create/aps | Convert APS file (located on storage) to PDF format and upload resulting file to storage. 
[**PutBookmark**](PdfApi.md#PutBookmark) | **Put** /pdf/{name}/bookmarks/bookmark/{bookmarkPath} | Update document bookmark.
[**PutCaretAnnotation**](PdfApi.md#PutCaretAnnotation) | **Put** /pdf/{name}/annotations/caret/{annotationId} | Replace document caret annotation
[**PutChangePasswordDocument**](PdfApi.md#PutChangePasswordDocument) | **Put** /pdf/changepassword | Change document password from content.
[**PutCheckBoxField**](PdfApi.md#PutCheckBoxField) | **Put** /pdf/{name}/fields/checkbox/{fieldName} | Replace document checkbox field
[**PutCircleAnnotation**](PdfApi.md#PutCircleAnnotation) | **Put** /pdf/{name}/annotations/circle/{annotationId} | Replace document circle annotation
[**PutComboBoxField**](PdfApi.md#PutComboBoxField) | **Put** /pdf/{name}/fields/combobox/{fieldName} | Replace document combobox field
[**PutCreateDocument**](PdfApi.md#PutCreateDocument) | **Put** /pdf/{name} | Create empty document.
[**PutCreatePdfFromLayer**](PdfApi.md#PutCreatePdfFromLayer) | **Put** /pdf/{name}/layers | Create a separate PDF from a PDF Layer and upload resulting file to storage.
[**PutDecryptDocument**](PdfApi.md#PutDecryptDocument) | **Put** /pdf/decrypt | Decrypt document from content.
[**PutDocumentDisplayProperties**](PdfApi.md#PutDocumentDisplayProperties) | **Put** /pdf/{name}/displayproperties | Update document display properties.
[**PutEncryptDocument**](PdfApi.md#PutEncryptDocument) | **Put** /pdf/encrypt | Encrypt document from content.
[**PutEpubInStorageToPdf**](PdfApi.md#PutEpubInStorageToPdf) | **Put** /pdf/{name}/create/epub | Convert EPUB file (located on storage) to PDF format and upload resulting file to storage. 
[**PutExportFieldsFromPdfToFdfInStorage**](PdfApi.md#PutExportFieldsFromPdfToFdfInStorage) | **Put** /pdf/{name}/export/fdf | Export fields from from PDF in storage to FDF file in storage.
[**PutExportFieldsFromPdfToXfdfInStorage**](PdfApi.md#PutExportFieldsFromPdfToXfdfInStorage) | **Put** /pdf/{name}/export/xfdf | Export fields from from PDF in storage to XFDF file in storage.
[**PutExportFieldsFromPdfToXmlInStorage**](PdfApi.md#PutExportFieldsFromPdfToXmlInStorage) | **Put** /pdf/{name}/export/xml | Export fields from from PDF in storage to XML file in storage.
[**PutFieldsFlatten**](PdfApi.md#PutFieldsFlatten) | **Put** /pdf/{name}/fields/flatten | Flatten form fields in document.
[**PutFileAttachmentAnnotation**](PdfApi.md#PutFileAttachmentAnnotation) | **Put** /pdf/{name}/annotations/fileattachment/{annotationId} | Replace document FileAttachment annotation
[**PutFileAttachmentAnnotationDataExtract**](PdfApi.md#PutFileAttachmentAnnotationDataExtract) | **Put** /pdf/{name}/annotations/fileattachment/{annotationId}/data/extract | Extract document FileAttachment annotation content to storage
[**PutFreeTextAnnotation**](PdfApi.md#PutFreeTextAnnotation) | **Put** /pdf/{name}/annotations/freetext/{annotationId} | Replace document free text annotation
[**PutHighlightAnnotation**](PdfApi.md#PutHighlightAnnotation) | **Put** /pdf/{name}/annotations/highlight/{annotationId} | Replace document highlight annotation
[**PutHtmlInStorageToPdf**](PdfApi.md#PutHtmlInStorageToPdf) | **Put** /pdf/{name}/create/html | Convert HTML file (located on storage) to PDF format and upload resulting file to storage. 
[**PutImageExtractAsGif**](PdfApi.md#PutImageExtractAsGif) | **Put** /pdf/{name}/images/{imageId}/extract/gif | Extract document image in GIF format to folder
[**PutImageExtractAsJpeg**](PdfApi.md#PutImageExtractAsJpeg) | **Put** /pdf/{name}/images/{imageId}/extract/jpeg | Extract document image in JPEG format to folder
[**PutImageExtractAsPng**](PdfApi.md#PutImageExtractAsPng) | **Put** /pdf/{name}/images/{imageId}/extract/png | Extract document image in PNG format to folder
[**PutImageExtractAsTiff**](PdfApi.md#PutImageExtractAsTiff) | **Put** /pdf/{name}/images/{imageId}/extract/tiff | Extract document image in TIFF format to folder
[**PutImageInStorageToPdf**](PdfApi.md#PutImageInStorageToPdf) | **Put** /pdf/{name}/create/images | Convert image file (located on storage) to PDF format and upload resulting file to storage. 
[**PutImagesExtractAsGif**](PdfApi.md#PutImagesExtractAsGif) | **Put** /pdf/{name}/pages/{pageNumber}/images/extract/gif | Extract document images in GIF format to folder.
[**PutImagesExtractAsJpeg**](PdfApi.md#PutImagesExtractAsJpeg) | **Put** /pdf/{name}/pages/{pageNumber}/images/extract/jpeg | Extract document images in JPEG format to folder.
[**PutImagesExtractAsPng**](PdfApi.md#PutImagesExtractAsPng) | **Put** /pdf/{name}/pages/{pageNumber}/images/extract/png | Extract document images in PNG format to folder.
[**PutImagesExtractAsTiff**](PdfApi.md#PutImagesExtractAsTiff) | **Put** /pdf/{name}/pages/{pageNumber}/images/extract/tiff | Extract document images in TIFF format to folder.
[**PutImportFieldsFromFdfInStorage**](PdfApi.md#PutImportFieldsFromFdfInStorage) | **Put** /pdf/{name}/import/fdf | Update fields from FDF file in storage.
[**PutImportFieldsFromXfdfInStorage**](PdfApi.md#PutImportFieldsFromXfdfInStorage) | **Put** /pdf/{name}/import/xfdf | Update fields from XFDF file in storage.
[**PutImportFieldsFromXmlInStorage**](PdfApi.md#PutImportFieldsFromXmlInStorage) | **Put** /pdf/{name}/import/xml | Update fields from XML file in storage.
[**PutInkAnnotation**](PdfApi.md#PutInkAnnotation) | **Put** /pdf/{name}/annotations/ink/{annotationId} | Replace document ink annotation
[**PutLineAnnotation**](PdfApi.md#PutLineAnnotation) | **Put** /pdf/{name}/annotations/line/{annotationId} | Replace document line annotation
[**PutLinkAnnotation**](PdfApi.md#PutLinkAnnotation) | **Put** /pdf/{name}/links/{linkId} | Replace document page link annotations
[**PutListBoxField**](PdfApi.md#PutListBoxField) | **Put** /pdf/{name}/fields/listbox/{fieldName} | Replace document listbox field
[**PutMarkdownInStorageToPdf**](PdfApi.md#PutMarkdownInStorageToPdf) | **Put** /pdf/{name}/create/markdown | Convert MD file (located on storage) to PDF format and upload resulting file to storage. 
[**PutMergeDocuments**](PdfApi.md#PutMergeDocuments) | **Put** /pdf/{name}/merge | Merge a list of documents.
[**PutMhtInStorageToPdf**](PdfApi.md#PutMhtInStorageToPdf) | **Put** /pdf/{name}/create/mht | Convert MHT file (located on storage) to PDF format and upload resulting file to storage. 
[**PutMovieAnnotation**](PdfApi.md#PutMovieAnnotation) | **Put** /pdf/{name}/annotations/movie/{annotationId} | Replace document movie annotation
[**PutPageAddStamp**](PdfApi.md#PutPageAddStamp) | **Put** /pdf/{name}/pages/{pageNumber}/stamp | Add page stamp.
[**PutPageConvertToBmp**](PdfApi.md#PutPageConvertToBmp) | **Put** /pdf/{name}/pages/{pageNumber}/convert/bmp | Convert document page to bmp image and upload resulting file to storage.
[**PutPageConvertToEmf**](PdfApi.md#PutPageConvertToEmf) | **Put** /pdf/{name}/pages/{pageNumber}/convert/emf | Convert document page to emf image and upload resulting file to storage.
[**PutPageConvertToGif**](PdfApi.md#PutPageConvertToGif) | **Put** /pdf/{name}/pages/{pageNumber}/convert/gif | Convert document page to gif image and upload resulting file to storage.
[**PutPageConvertToJpeg**](PdfApi.md#PutPageConvertToJpeg) | **Put** /pdf/{name}/pages/{pageNumber}/convert/jpeg | Convert document page to Jpeg image and upload resulting file to storage.
[**PutPageConvertToPng**](PdfApi.md#PutPageConvertToPng) | **Put** /pdf/{name}/pages/{pageNumber}/convert/png | Convert document page to png image and upload resulting file to storage.
[**PutPageConvertToTiff**](PdfApi.md#PutPageConvertToTiff) | **Put** /pdf/{name}/pages/{pageNumber}/convert/tiff | Convert document page to Tiff image and upload resulting file to storage.
[**PutPclInStorageToPdf**](PdfApi.md#PutPclInStorageToPdf) | **Put** /pdf/{name}/create/pcl | Convert PCL file (located on storage) to PDF format and upload resulting file to storage. 
[**PutPdfAInStorageToPdf**](PdfApi.md#PutPdfAInStorageToPdf) | **Put** /pdf/{name}/create/pdfa | Convert PDFA file (located on storage) to PDF format and upload resulting file to storage. 
[**PutPdfInRequestToAps**](PdfApi.md#PutPdfInRequestToAps) | **Put** /pdf/convert/aps | Converts PDF document (in request content) to APS format and uploads resulting file to storage.
[**PutPdfInRequestToDoc**](PdfApi.md#PutPdfInRequestToDoc) | **Put** /pdf/convert/doc | Converts PDF document (in request content) to DOC format and uploads resulting file to storage.
[**PutPdfInRequestToEpub**](PdfApi.md#PutPdfInRequestToEpub) | **Put** /pdf/convert/epub | Converts PDF document (in request content) to EPUB format and uploads resulting file to storage.
[**PutPdfInRequestToHtml**](PdfApi.md#PutPdfInRequestToHtml) | **Put** /pdf/convert/html | Converts PDF document (in request content) to Html format and uploads resulting file to storage.
[**PutPdfInRequestToMobiXml**](PdfApi.md#PutPdfInRequestToMobiXml) | **Put** /pdf/convert/mobixml | Converts PDF document (in request content) to MOBIXML format and uploads resulting ZIP archive file to storage.
[**PutPdfInRequestToPdfA**](PdfApi.md#PutPdfInRequestToPdfA) | **Put** /pdf/convert/pdfa | Converts PDF document (in request content) to PdfA format and uploads resulting file to storage.
[**PutPdfInRequestToPptx**](PdfApi.md#PutPdfInRequestToPptx) | **Put** /pdf/convert/pptx | Converts PDF document (in request content) to PPTX format and uploads resulting file to storage.
[**PutPdfInRequestToSvg**](PdfApi.md#PutPdfInRequestToSvg) | **Put** /pdf/convert/svg | Converts PDF document (in request content) to SVG format and uploads resulting file to storage.
[**PutPdfInRequestToTeX**](PdfApi.md#PutPdfInRequestToTeX) | **Put** /pdf/convert/tex | Converts PDF document (in request content) to TeX format and uploads resulting file to storage.
[**PutPdfInRequestToTiff**](PdfApi.md#PutPdfInRequestToTiff) | **Put** /pdf/convert/tiff | Converts PDF document (in request content) to TIFF format and uploads resulting file to storage.
[**PutPdfInRequestToXls**](PdfApi.md#PutPdfInRequestToXls) | **Put** /pdf/convert/xls | Converts PDF document (in request content) to XLS format and uploads resulting file to storage.
[**PutPdfInRequestToXlsx**](PdfApi.md#PutPdfInRequestToXlsx) | **Put** /pdf/convert/xlsx | Converts PDF document (in request content) to XLSX format and uploads resulting file to storage.
[**PutPdfInRequestToXml**](PdfApi.md#PutPdfInRequestToXml) | **Put** /pdf/convert/xml | Converts PDF document (in request content) to XML format and uploads resulting file to storage.
[**PutPdfInRequestToXps**](PdfApi.md#PutPdfInRequestToXps) | **Put** /pdf/convert/xps | Converts PDF document (in request content) to XPS format and uploads resulting file to storage.
[**PutPdfInStorageToAps**](PdfApi.md#PutPdfInStorageToAps) | **Put** /pdf/{name}/convert/aps | Converts PDF document (located on storage) to APS format and uploads resulting file to storage
[**PutPdfInStorageToDoc**](PdfApi.md#PutPdfInStorageToDoc) | **Put** /pdf/{name}/convert/doc | Converts PDF document (located on storage) to DOC format and uploads resulting file to storage.
[**PutPdfInStorageToEpub**](PdfApi.md#PutPdfInStorageToEpub) | **Put** /pdf/{name}/convert/epub | Converts PDF document (located on storage) to EPUB format and uploads resulting file to storage
[**PutPdfInStorageToHtml**](PdfApi.md#PutPdfInStorageToHtml) | **Put** /pdf/{name}/convert/html | Converts PDF document (located on storage) to Html format and uploads resulting file to storage
[**PutPdfInStorageToMobiXml**](PdfApi.md#PutPdfInStorageToMobiXml) | **Put** /pdf/{name}/convert/mobixml | Converts PDF document (located on storage) to MOBIXML format and uploads resulting ZIP archive file to storage
[**PutPdfInStorageToPdfA**](PdfApi.md#PutPdfInStorageToPdfA) | **Put** /pdf/{name}/convert/pdfa | Converts PDF document (located on storage) to PdfA format and uploads resulting file to storage
[**PutPdfInStorageToPptx**](PdfApi.md#PutPdfInStorageToPptx) | **Put** /pdf/{name}/convert/pptx | Converts PDF document (located on storage) to PPTX format and uploads resulting file to storage
[**PutPdfInStorageToSvg**](PdfApi.md#PutPdfInStorageToSvg) | **Put** /pdf/{name}/convert/svg | Converts PDF document (located on storage) to SVG format and uploads resulting file to storage
[**PutPdfInStorageToTeX**](PdfApi.md#PutPdfInStorageToTeX) | **Put** /pdf/{name}/convert/tex | Converts PDF document (located on storage) to TeX format and uploads resulting file to storage
[**PutPdfInStorageToTiff**](PdfApi.md#PutPdfInStorageToTiff) | **Put** /pdf/{name}/convert/tiff | Converts PDF document (located on storage) to TIFF format and uploads resulting file to storage
[**PutPdfInStorageToXls**](PdfApi.md#PutPdfInStorageToXls) | **Put** /pdf/{name}/convert/xls | Converts PDF document (located on storage) to XLS format and uploads resulting file to storage
[**PutPdfInStorageToXlsx**](PdfApi.md#PutPdfInStorageToXlsx) | **Put** /pdf/{name}/convert/xlsx | Converts PDF document (located on storage) to XLSX format and uploads resulting file to storage
[**PutPdfInStorageToXml**](PdfApi.md#PutPdfInStorageToXml) | **Put** /pdf/{name}/convert/xml | Converts PDF document (located on storage) to XML format and uploads resulting file to storage
[**PutPdfInStorageToXps**](PdfApi.md#PutPdfInStorageToXps) | **Put** /pdf/{name}/convert/xps | Converts PDF document (located on storage) to XPS format and uploads resulting file to storage
[**PutPolyLineAnnotation**](PdfApi.md#PutPolyLineAnnotation) | **Put** /pdf/{name}/annotations/polyline/{annotationId} | Replace document polyline annotation
[**PutPolygonAnnotation**](PdfApi.md#PutPolygonAnnotation) | **Put** /pdf/{name}/annotations/polygon/{annotationId} | Replace document polygon annotation
[**PutPopupAnnotation**](PdfApi.md#PutPopupAnnotation) | **Put** /pdf/{name}/annotations/popup/{annotationId} | Replace document popup annotation
[**PutPrivileges**](PdfApi.md#PutPrivileges) | **Put** /pdf/{name}/privileges | Update privilege document.
[**PutPsInStorageToPdf**](PdfApi.md#PutPsInStorageToPdf) | **Put** /pdf/{name}/create/ps | Convert PS file (located on storage) to PDF format and upload resulting file to storage. 
[**PutRadioButtonField**](PdfApi.md#PutRadioButtonField) | **Put** /pdf/{name}/fields/radiobutton/{fieldName} | Replace document RadioButton field
[**PutRedactionAnnotation**](PdfApi.md#PutRedactionAnnotation) | **Put** /pdf/{name}/annotations/redaction/{annotationId} | Replace document redaction annotation
[**PutReplaceImage**](PdfApi.md#PutReplaceImage) | **Put** /pdf/{name}/images/{imageId} | Replace document image.
[**PutReplaceMultipleImage**](PdfApi.md#PutReplaceMultipleImage) | **Put** /pdf/{name}/images/replace | Replace document multiple image.
[**PutScreenAnnotation**](PdfApi.md#PutScreenAnnotation) | **Put** /pdf/{name}/annotations/screen/{annotationId} | Replace document screen annotation
[**PutScreenAnnotationDataExtract**](PdfApi.md#PutScreenAnnotationDataExtract) | **Put** /pdf/{name}/annotations/screen/{annotationId}/data/extract | Extract document screen annotation content to storage
[**PutSearchableDocument**](PdfApi.md#PutSearchableDocument) | **Put** /pdf/{name}/ocr | Create searchable PDF document. Generate OCR layer for images in input PDF document.
[**PutSetProperty**](PdfApi.md#PutSetProperty) | **Put** /pdf/{name}/documentproperties/{propertyName} | Add/update document property.
[**PutSignatureField**](PdfApi.md#PutSignatureField) | **Put** /pdf/{name}/fields/signature/{fieldName} | Replace document signature field.
[**PutSoundAnnotation**](PdfApi.md#PutSoundAnnotation) | **Put** /pdf/{name}/annotations/sound/{annotationId} | Replace document sound annotation
[**PutSoundAnnotationDataExtract**](PdfApi.md#PutSoundAnnotationDataExtract) | **Put** /pdf/{name}/annotations/sound/{annotationId}/data/extract | Extract document sound annotation content to storage
[**PutSquareAnnotation**](PdfApi.md#PutSquareAnnotation) | **Put** /pdf/{name}/annotations/square/{annotationId} | Replace document square annotation
[**PutSquigglyAnnotation**](PdfApi.md#PutSquigglyAnnotation) | **Put** /pdf/{name}/annotations/squiggly/{annotationId} | Replace document squiggly annotation
[**PutStampAnnotation**](PdfApi.md#PutStampAnnotation) | **Put** /pdf/{name}/annotations/stamp/{annotationId} | Replace document stamp annotation
[**PutStampAnnotationDataExtract**](PdfApi.md#PutStampAnnotationDataExtract) | **Put** /pdf/{name}/annotations/stamp/{annotationId}/data/extract | Extract document stamp annotation content to storage
[**PutStrikeOutAnnotation**](PdfApi.md#PutStrikeOutAnnotation) | **Put** /pdf/{name}/annotations/strikeout/{annotationId} | Replace document StrikeOut annotation
[**PutSvgInStorageToPdf**](PdfApi.md#PutSvgInStorageToPdf) | **Put** /pdf/{name}/create/svg | Convert SVG file (located on storage) to PDF format and upload resulting file to storage. 
[**PutTable**](PdfApi.md#PutTable) | **Put** /pdf/{name}/tables/{tableId} | Replace document page table.
[**PutTeXInStorageToPdf**](PdfApi.md#PutTeXInStorageToPdf) | **Put** /pdf/{name}/create/tex | Convert TeX file (located on storage) to PDF format and upload resulting file to storage. 
[**PutTextAnnotation**](PdfApi.md#PutTextAnnotation) | **Put** /pdf/{name}/annotations/text/{annotationId} | Replace document text annotation
[**PutTextBoxField**](PdfApi.md#PutTextBoxField) | **Put** /pdf/{name}/fields/textbox/{fieldName} | Replace document text box field
[**PutUnderlineAnnotation**](PdfApi.md#PutUnderlineAnnotation) | **Put** /pdf/{name}/annotations/underline/{annotationId} | Replace document underline annotation
[**PutUpdateField**](PdfApi.md#PutUpdateField) | **Put** /pdf/{name}/fields/{fieldName} | Update field.
[**PutUpdateFields**](PdfApi.md#PutUpdateFields) | **Put** /pdf/{name}/fields | Update fields.
[**PutWebInStorageToPdf**](PdfApi.md#PutWebInStorageToPdf) | **Put** /pdf/{name}/create/web | Convert web page to PDF format and upload resulting file to storage. 
[**PutXfaPdfInRequestToAcroForm**](PdfApi.md#PutXfaPdfInRequestToAcroForm) | **Put** /pdf/convert/xfatoacroform | Converts PDF document which contains XFA form (in request content) to PDF with AcroForm and uploads resulting file to storage.
[**PutXfaPdfInStorageToAcroForm**](PdfApi.md#PutXfaPdfInStorageToAcroForm) | **Put** /pdf/{name}/convert/xfatoacroform | Converts PDF document which contains XFA form (located on storage) to PDF with AcroForm and uploads resulting file to storage
[**PutXmlInStorageToPdf**](PdfApi.md#PutXmlInStorageToPdf) | **Put** /pdf/{name}/create/xml | Convert XML file (located on storage) to PDF format and upload resulting file to storage. 
[**PutXpsInStorageToPdf**](PdfApi.md#PutXpsInStorageToPdf) | **Put** /pdf/{name}/create/xps | Convert XPS file (located on storage) to PDF format and upload resulting file to storage. 
[**PutXslFoInStorageToPdf**](PdfApi.md#PutXslFoInStorageToPdf) | **Put** /pdf/{name}/create/xslfo | Convert XslFo file (located on storage) to PDF format and upload resulting file to storage. 
[**StorageExists**](PdfApi.md#StorageExists) | **Get** /pdf/storage/{storageName}/exist | Check if storage exists
[**UploadFile**](PdfApi.md#UploadFile) | **Put** /pdf/storage/file/{path} | Upload file


# **CopyFile**
> CopyFile(srcPath, destPath, optional)
Copy file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Source file path e.g. &#39;/folder/file.ext&#39; | 
  **destPath** | **string**| Destination file path | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Source file path e.g. &#39;/folder/file.ext&#39; | 
 **destPath** | **string**| Destination file path | 
 **srcStorageName** | **string**| Source storage name | 
 **destStorageName** | **string**| Destination storage name | 
 **versionId** | **string**| File version ID to copy | 

### Return type

 (empty response body)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CopyFolder**
> CopyFolder(srcPath, destPath, optional)
Copy folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Source folder path e.g. &#39;/src&#39; | 
  **destPath** | **string**| Destination folder path e.g. &#39;/dst&#39; | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Source folder path e.g. &#39;/src&#39; | 
 **destPath** | **string**| Destination folder path e.g. &#39;/dst&#39; | 
 **srcStorageName** | **string**| Source storage name | 
 **destStorageName** | **string**| Destination storage name | 

### Return type

 (empty response body)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFolder**
> CreateFolder(path, optional)
Create the folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| Folder path to create e.g. &#39;folder_1/folder_2/&#39; | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| Folder path to create e.g. &#39;folder_1/folder_2/&#39; | 
 **storageName** | **string**| Storage name | 

### Return type

 (empty response body)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAnnotation**
> AsposeResponse DeleteAnnotation(name, annotationId, optional)
Delete document annotation by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBookmark**
> AsposeResponse DeleteBookmark(name, bookmarkPath, optional)
Delete document bookmark by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **bookmarkPath** | **string**| The bookmark path. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **bookmarkPath** | **string**| The bookmark path. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDocumentAnnotations**
> AsposeResponse DeleteDocumentAnnotations(name, optional)
Delete all annotations from the document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDocumentBookmarks**
> AsposeResponse DeleteDocumentBookmarks(name, optional)
Delete all document bookmarks.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDocumentLayer**
> AsposeResponse DeleteDocumentLayer(name, pageNumber, layerId, optional)
Remove document layer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| Layer page. | 
  **layerId** | **string**| Layer Id. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| Layer page. | 
 **layerId** | **string**| Layer Id. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **passBase64** | **string**| The password (Base64). | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDocumentLinkAnnotations**
> AsposeResponse DeleteDocumentLinkAnnotations(name, optional)
Delete all link annotations from the document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDocumentStamps**
> AsposeResponse DeleteDocumentStamps(name, optional)
Delete all stamps from the document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| Base64 encoded password. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDocumentTables**
> AsposeResponse DeleteDocumentTables(name, optional)
Delete all tables from the document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteField**
> AsposeResponse DeleteField(name, fieldName, optional)
Delete document field by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fieldName** | **string**| The field name/ | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fieldName** | **string**| The field name/ | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFile**
> DeleteFile(path, optional)
Delete file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| File path e.g. &#39;/folder/file.ext&#39; | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| File path e.g. &#39;/folder/file.ext&#39; | 
 **storageName** | **string**| Storage name | 
 **versionId** | **string**| File version ID to delete | 

### Return type

 (empty response body)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFolder**
> DeleteFolder(path, optional)
Delete folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| Folder path e.g. &#39;/folder&#39; | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| Folder path e.g. &#39;/folder&#39; | 
 **storageName** | **string**| Storage name | 
 **recursive** | **bool**| Enable to delete folders, subfolders and files | [default to false]

### Return type

 (empty response body)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteImage**
> AsposeResponse DeleteImage(name, imageId, optional)
Delete image from document page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **imageId** | **string**| Image ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageId** | **string**| Image ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLinkAnnotation**
> AsposeResponse DeleteLinkAnnotation(name, linkId, optional)
Delete document page link annotation by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **linkId** | **string**| The link ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **linkId** | **string**| The link ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePage**
> AsposeResponse DeletePage(name, pageNumber, optional)
Delete document page by its number.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePageAnnotations**
> AsposeResponse DeletePageAnnotations(name, pageNumber, optional)
Delete all annotations from the page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePageLinkAnnotations**
> AsposeResponse DeletePageLinkAnnotations(name, pageNumber, optional)
Delete all link annotations from the page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePageStamps**
> AsposeResponse DeletePageStamps(name, pageNumber, optional)
Delete all stamps from the page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| Base64 encoded password. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePageTables**
> AsposeResponse DeletePageTables(name, pageNumber, optional)
Delete all tables from the page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProperties**
> AsposeResponse DeleteProperties(name, optional)
Delete custom document properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProperty**
> AsposeResponse DeleteProperty(name, propertyName, optional)
Delete document property.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **propertyName** | **string**| Property name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **propertyName** | **string**| Property name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStamp**
> AsposeResponse DeleteStamp(name, stampId, optional)
Delete document stamp by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **stampId** | **string**| The stamp ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **stampId** | **string**| The stamp ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| Base64 encoded password. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTable**
> AsposeResponse DeleteTable(name, tableId, optional)
Delete document table by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **tableId** | **string**| The table ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **tableId** | **string**| The table ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadFile**
> []byte DownloadFile(path, optional)
Download file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| File path e.g. &#39;/folder/file.ext&#39; | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| File path e.g. &#39;/folder/file.ext&#39; | 
 **storageName** | **string**| Storage name | 
 **versionId** | **string**| File version ID to download | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApsInStorageToPdf**
> []byte GetApsInStorageToPdf(srcPath, optional)
Convert APS file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xps) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xps) | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBookmark**
> BookmarkResponse GetBookmark(name, bookmarkPath, optional)
Read document bookmark.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **bookmarkPath** | **string**| The bookmark path. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **bookmarkPath** | **string**| The bookmark path. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**BookmarkResponse**](BookmarkResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBookmarks**
> BookmarksResponse GetBookmarks(name, bookmarkPath, optional)
Read document bookmarks node list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **bookmarkPath** | **string**| The bookmark path. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **bookmarkPath** | **string**| The bookmark path. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**BookmarksResponse**](BookmarksResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCaretAnnotation**
> CaretAnnotationResponse GetCaretAnnotation(name, annotationId, optional)
Read document page caret annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**CaretAnnotationResponse**](CaretAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCheckBoxField**
> CheckBoxFieldResponse GetCheckBoxField(name, fieldName, optional)
Read document checkbox field by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fieldName** | **string**| The field name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fieldName** | **string**| The field name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**CheckBoxFieldResponse**](CheckBoxFieldResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCircleAnnotation**
> CircleAnnotationResponse GetCircleAnnotation(name, annotationId, optional)
Read document page circle annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**CircleAnnotationResponse**](CircleAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComboBoxField**
> ComboBoxFieldResponse GetComboBoxField(name, fieldName, optional)
Read document combobox field by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fieldName** | **string**| The field name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fieldName** | **string**| The field name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**ComboBoxFieldResponse**](ComboBoxFieldResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDiscUsage**
> DiscUsage GetDiscUsage(optional)
Get disc usage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storageName** | **string**| Storage name | 

### Return type

[**DiscUsage**](DiscUsage.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocument**
> DocumentResponse GetDocument(name, optional)
Read common document info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| Base64 encoded password. | 

### Return type

[**DocumentResponse**](DocumentResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentAnnotations**
> AnnotationsInfoResponse GetDocumentAnnotations(name, optional)
Read document page annotations. Returns only FreeTextAnnotations, TextAnnotations, other annotations will implemented next releases.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AnnotationsInfoResponse**](AnnotationsInfoResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentAttachmentByIndex**
> AttachmentResponse GetDocumentAttachmentByIndex(name, attachmentIndex, optional)
Read document attachment info by its index.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **attachmentIndex** | **int32**| The attachment index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **attachmentIndex** | **int32**| The attachment index. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AttachmentResponse**](AttachmentResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentAttachments**
> AttachmentsResponse GetDocumentAttachments(name, optional)
Read document attachments info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AttachmentsResponse**](AttachmentsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentBookmarks**
> BookmarksResponse GetDocumentBookmarks(name, optional)
Read document bookmarks tree.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**BookmarksResponse**](BookmarksResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentCaretAnnotations**
> CaretAnnotationsResponse GetDocumentCaretAnnotations(name, optional)
Read document caret annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**CaretAnnotationsResponse**](CaretAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentCheckBoxFields**
> CheckBoxFieldsResponse GetDocumentCheckBoxFields(name, optional)
Read document checkbox fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**CheckBoxFieldsResponse**](CheckBoxFieldsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentCircleAnnotations**
> CircleAnnotationsResponse GetDocumentCircleAnnotations(name, optional)
Read document circle annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**CircleAnnotationsResponse**](CircleAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentComboBoxFields**
> ComboBoxFieldsResponse GetDocumentComboBoxFields(name, optional)
Read document combobox fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**ComboBoxFieldsResponse**](ComboBoxFieldsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentDisplayProperties**
> DisplayPropertiesResponse GetDocumentDisplayProperties(name, optional)
Read document display properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**|  | 
 **storage** | **string**|  | 
 **folder** | **string**|  | 

### Return type

[**DisplayPropertiesResponse**](DisplayPropertiesResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentFileAttachmentAnnotations**
> FileAttachmentAnnotationsResponse GetDocumentFileAttachmentAnnotations(name, optional)
Read document FileAttachment annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**FileAttachmentAnnotationsResponse**](FileAttachmentAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentFreeTextAnnotations**
> FreeTextAnnotationsResponse GetDocumentFreeTextAnnotations(name, optional)
Read document free text annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**FreeTextAnnotationsResponse**](FreeTextAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentHighlightAnnotations**
> HighlightAnnotationsResponse GetDocumentHighlightAnnotations(name, optional)
Read document highlight annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**HighlightAnnotationsResponse**](HighlightAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentInkAnnotations**
> InkAnnotationsResponse GetDocumentInkAnnotations(name, optional)
Read document ink annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**InkAnnotationsResponse**](InkAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentLayers**
> DocumentLayers GetDocumentLayers(name, optional)
Gets document layers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **passBase64** | **string**| The password (Base64). | 

### Return type

[**DocumentLayers**](DocumentLayers.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentLineAnnotations**
> LineAnnotationsResponse GetDocumentLineAnnotations(name, optional)
Read document line annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**LineAnnotationsResponse**](LineAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentListBoxFields**
> ListBoxFieldsResponse GetDocumentListBoxFields(name, optional)
Read document listbox fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**ListBoxFieldsResponse**](ListBoxFieldsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentMovieAnnotations**
> MovieAnnotationsResponse GetDocumentMovieAnnotations(name, optional)
Read document movie annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**MovieAnnotationsResponse**](MovieAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentPolyLineAnnotations**
> PolyLineAnnotationsResponse GetDocumentPolyLineAnnotations(name, optional)
Read document polyline annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PolyLineAnnotationsResponse**](PolyLineAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentPolygonAnnotations**
> PolygonAnnotationsResponse GetDocumentPolygonAnnotations(name, optional)
Read document polygon annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PolygonAnnotationsResponse**](PolygonAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentPopupAnnotations**
> PopupAnnotationsResponse GetDocumentPopupAnnotations(name, optional)
Read document popup annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PopupAnnotationsResponse**](PopupAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentPopupAnnotationsByParent**
> PopupAnnotationsResponse GetDocumentPopupAnnotationsByParent(name, annotationId, optional)
Read document popup annotations by parent id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The parent annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The parent annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PopupAnnotationsResponse**](PopupAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentProperties**
> DocumentPropertiesResponse GetDocumentProperties(name, optional)
Read document properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**DocumentPropertiesResponse**](DocumentPropertiesResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentProperty**
> DocumentPropertyResponse GetDocumentProperty(name, propertyName, optional)
Read document property by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **propertyName** | **string**| Property name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **propertyName** | **string**| Property name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**DocumentPropertyResponse**](DocumentPropertyResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentRadioButtonFields**
> RadioButtonFieldsResponse GetDocumentRadioButtonFields(name, optional)
Read document radiobutton fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**RadioButtonFieldsResponse**](RadioButtonFieldsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentRedactionAnnotations**
> RedactionAnnotationsResponse GetDocumentRedactionAnnotations(name, optional)
Read document redaction annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**RedactionAnnotationsResponse**](RedactionAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentScreenAnnotations**
> ScreenAnnotationsResponse GetDocumentScreenAnnotations(name, optional)
Read document screen annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**ScreenAnnotationsResponse**](ScreenAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentSignatureFields**
> SignatureFieldsResponse GetDocumentSignatureFields(name, optional)
Read document signature fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SignatureFieldsResponse**](SignatureFieldsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentSoundAnnotations**
> SoundAnnotationsResponse GetDocumentSoundAnnotations(name, optional)
Read document sound annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SoundAnnotationsResponse**](SoundAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentSquareAnnotations**
> SquareAnnotationsResponse GetDocumentSquareAnnotations(name, optional)
Read document square annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SquareAnnotationsResponse**](SquareAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentSquigglyAnnotations**
> SquigglyAnnotationsResponse GetDocumentSquigglyAnnotations(name, optional)
Read document squiggly annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SquigglyAnnotationsResponse**](SquigglyAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentStampAnnotations**
> StampAnnotationsResponse GetDocumentStampAnnotations(name, optional)
Read document stamp annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**StampAnnotationsResponse**](StampAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentStamps**
> StampsInfoResponse GetDocumentStamps(name, optional)
Read document stamps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**StampsInfoResponse**](StampsInfoResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentStrikeOutAnnotations**
> StrikeOutAnnotationsResponse GetDocumentStrikeOutAnnotations(name, optional)
Read document StrikeOut annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**StrikeOutAnnotationsResponse**](StrikeOutAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentTables**
> TablesRecognizedResponse GetDocumentTables(name, optional)
Read document tables.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**|  | 
 **storage** | **string**|  | 
 **folder** | **string**|  | 

### Return type

[**TablesRecognizedResponse**](TablesRecognizedResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentTextAnnotations**
> TextAnnotationsResponse GetDocumentTextAnnotations(name, optional)
Read document text annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**TextAnnotationsResponse**](TextAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentTextBoxFields**
> TextBoxFieldsResponse GetDocumentTextBoxFields(name, optional)
Read document text box fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**TextBoxFieldsResponse**](TextBoxFieldsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentUnderlineAnnotations**
> UnderlineAnnotationsResponse GetDocumentUnderlineAnnotations(name, optional)
Read document underline annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**UnderlineAnnotationsResponse**](UnderlineAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDownloadDocumentAttachmentByIndex**
> []byte GetDownloadDocumentAttachmentByIndex(name, attachmentIndex, optional)
Download document attachment content by its index.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **attachmentIndex** | **int32**| The attachment index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **attachmentIndex** | **int32**| The attachment index. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEpubInStorageToPdf**
> []byte GetEpubInStorageToPdf(srcPath, optional)
Convert EPUB file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.epub) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.epub) | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExportFieldsFromPdfToFdfInStorage**
> []byte GetExportFieldsFromPdfToFdfInStorage(name, optional)
Export fields from from PDF in storage to FDF file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExportFieldsFromPdfToXfdfInStorage**
> []byte GetExportFieldsFromPdfToXfdfInStorage(name, optional)
Export fields from from PDF in storage to XFDF file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExportFieldsFromPdfToXmlInStorage**
> []byte GetExportFieldsFromPdfToXmlInStorage(name, optional)
Export fields from from PDF in storage to XML file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetField**
> FieldResponse GetField(name, fieldName, optional)
Get document field by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fieldName** | **string**| The field name (name should be encoded). | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fieldName** | **string**| The field name (name should be encoded). | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**FieldResponse**](FieldResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFields**
> FieldsResponse GetFields(name, optional)
Get document fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**FieldsResponse**](FieldsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileAttachmentAnnotation**
> FileAttachmentAnnotationResponse GetFileAttachmentAnnotation(name, annotationId, optional)
Read document page FileAttachment annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**FileAttachmentAnnotationResponse**](FileAttachmentAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileAttachmentAnnotationData**
> []byte GetFileAttachmentAnnotationData(name, annotationId, optional)
Read document page FileAttachment annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileVersions**
> FileVersions GetFileVersions(path, optional)
Get file versions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| File path e.g. &#39;/file.ext&#39; | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| File path e.g. &#39;/file.ext&#39; | 
 **storageName** | **string**| Storage name | 

### Return type

[**FileVersions**](FileVersions.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFilesList**
> FilesList GetFilesList(path, optional)
Get all files and folders within a folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| Folder path e.g. &#39;/folder&#39; | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| Folder path e.g. &#39;/folder&#39; | 
 **storageName** | **string**| Storage name | 

### Return type

[**FilesList**](FilesList.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFreeTextAnnotation**
> FreeTextAnnotationResponse GetFreeTextAnnotation(name, annotationId, optional)
Read document page free text annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**FreeTextAnnotationResponse**](FreeTextAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHighlightAnnotation**
> HighlightAnnotationResponse GetHighlightAnnotation(name, annotationId, optional)
Read document page highlight annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**HighlightAnnotationResponse**](HighlightAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHtmlInStorageToPdf**
> []byte GetHtmlInStorageToPdf(srcPath, optional)
Convert HTML file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.zip) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.zip) | 
 **htmlFileName** | **string**| Name of HTML file in ZIP. | 
 **height** | **float64**| Page height | 
 **width** | **float64**| Page width | 
 **isLandscape** | **bool**| Is page landscaped | 
 **marginLeft** | **float64**| Page margin left | 
 **marginBottom** | **float64**| Page margin bottom | 
 **marginRight** | **float64**| Page margin right | 
 **marginTop** | **float64**| Page margin top | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImage**
> ImageResponse GetImage(name, imageId, optional)
Read document image by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **imageId** | **string**| Image ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageId** | **string**| Image ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**ImageResponse**](ImageResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImageExtractAsGif**
> []byte GetImageExtractAsGif(name, imageId, optional)
Extract document image in GIF format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **imageId** | **string**| Image ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageId** | **string**| Image ID. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImageExtractAsJpeg**
> []byte GetImageExtractAsJpeg(name, imageId, optional)
Extract document image in JPEG format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **imageId** | **string**| Image ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageId** | **string**| Image ID. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImageExtractAsPng**
> []byte GetImageExtractAsPng(name, imageId, optional)
Extract document image in PNG format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **imageId** | **string**| Image ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageId** | **string**| Image ID. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImageExtractAsTiff**
> []byte GetImageExtractAsTiff(name, imageId, optional)
Extract document image in TIFF format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **imageId** | **string**| Image ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageId** | **string**| Image ID. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImages**
> ImagesResponse GetImages(name, pageNumber, optional)
Read document images.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**ImagesResponse**](ImagesResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImagesExtractSvg**
> SvgImages GetImagesExtractSvg(name, pageNumber, optional)
Extract SVG images from document page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **passBase64** | **string**| The password (Base64). | 

### Return type

[**SvgImages**](SvgImages.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImportFieldsFromFdfInStorage**
> []byte GetImportFieldsFromFdfInStorage(name, fdfFilePath, optional)
Update fields from FDF file in storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fdfFilePath** | **string**| The Fdf file path. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fdfFilePath** | **string**| The Fdf file path. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImportFieldsFromXfdfInStorage**
> []byte GetImportFieldsFromXfdfInStorage(name, xfdfFilePath, optional)
Update fields from XFDF file in storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **xfdfFilePath** | **string**| The XFDF file path. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **xfdfFilePath** | **string**| The XFDF file path. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImportFieldsFromXmlInStorage**
> []byte GetImportFieldsFromXmlInStorage(name, xmlFilePath, optional)
Import from XML file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **xmlFilePath** | **string**| Full source filename (ex. /folder1/folder2/template.xml) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **xmlFilePath** | **string**| Full source filename (ex. /folder1/folder2/template.xml) | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInkAnnotation**
> InkAnnotationResponse GetInkAnnotation(name, annotationId, optional)
Read document page ink annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**InkAnnotationResponse**](InkAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLineAnnotation**
> LineAnnotationResponse GetLineAnnotation(name, annotationId, optional)
Read document page line annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**LineAnnotationResponse**](LineAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLinkAnnotation**
> LinkAnnotationResponse GetLinkAnnotation(name, linkId, optional)
Read document link annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **linkId** | **string**| The link ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **linkId** | **string**| The link ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**LinkAnnotationResponse**](LinkAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListBoxField**
> ListBoxFieldResponse GetListBoxField(name, fieldName, optional)
Read document listbox field by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fieldName** | **string**| The field name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fieldName** | **string**| The field name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**ListBoxFieldResponse**](ListBoxFieldResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarkdownInStorageToPdf**
> []byte GetMarkdownInStorageToPdf(srcPath, optional)
Convert MD file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.md) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.md) | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMhtInStorageToPdf**
> []byte GetMhtInStorageToPdf(srcPath, optional)
Convert MHT file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.mht) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.mht) | 
 **height** | **float64**| Page height | 
 **width** | **float64**| Page width | 
 **marginLeft** | **float64**| Page margin left | 
 **marginBottom** | **float64**| Page margin bottom | 
 **marginRight** | **float64**| Page margin right | 
 **marginTop** | **float64**| Page margin top | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMovieAnnotation**
> MovieAnnotationResponse GetMovieAnnotation(name, annotationId, optional)
Read document page movie annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**MovieAnnotationResponse**](MovieAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPage**
> DocumentPageResponse GetPage(name, pageNumber, optional)
Read document page info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**DocumentPageResponse**](DocumentPageResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageAnnotations**
> AnnotationsInfoResponse GetPageAnnotations(name, pageNumber, optional)
Read document page annotations. Returns only FreeTextAnnotations, TextAnnotations, other annotations will implemented next releases.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AnnotationsInfoResponse**](AnnotationsInfoResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageCaretAnnotations**
> CaretAnnotationsResponse GetPageCaretAnnotations(name, pageNumber, optional)
Read document page caret annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**CaretAnnotationsResponse**](CaretAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageCheckBoxFields**
> CheckBoxFieldsResponse GetPageCheckBoxFields(name, pageNumber, optional)
Read document page checkbox fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**CheckBoxFieldsResponse**](CheckBoxFieldsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageCircleAnnotations**
> CircleAnnotationsResponse GetPageCircleAnnotations(name, pageNumber, optional)
Read document page circle annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**CircleAnnotationsResponse**](CircleAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageComboBoxFields**
> ComboBoxFieldsResponse GetPageComboBoxFields(name, pageNumber, optional)
Read document page combobox fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**ComboBoxFieldsResponse**](ComboBoxFieldsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageConvertToBmp**
> []byte GetPageConvertToBmp(name, pageNumber, optional)
Convert document page to Bmp image and return resulting file in response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageConvertToEmf**
> []byte GetPageConvertToEmf(name, pageNumber, optional)
Convert document page to Emf image and return resulting file in response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageConvertToGif**
> []byte GetPageConvertToGif(name, pageNumber, optional)
Convert document page to Gif image and return resulting file in response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageConvertToJpeg**
> []byte GetPageConvertToJpeg(name, pageNumber, optional)
Convert document page to Jpeg image and return resulting file in response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageConvertToPng**
> []byte GetPageConvertToPng(name, pageNumber, optional)
Convert document page to Png image and return resulting file in response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageConvertToTiff**
> []byte GetPageConvertToTiff(name, pageNumber, optional)
Convert document page to Tiff image  and return resulting file in response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageFileAttachmentAnnotations**
> FileAttachmentAnnotationsResponse GetPageFileAttachmentAnnotations(name, pageNumber, optional)
Read document page FileAttachment annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**FileAttachmentAnnotationsResponse**](FileAttachmentAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageFreeTextAnnotations**
> FreeTextAnnotationsResponse GetPageFreeTextAnnotations(name, pageNumber, optional)
Read document page free text annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**FreeTextAnnotationsResponse**](FreeTextAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageHighlightAnnotations**
> HighlightAnnotationsResponse GetPageHighlightAnnotations(name, pageNumber, optional)
Read document page highlight annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**HighlightAnnotationsResponse**](HighlightAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageInkAnnotations**
> InkAnnotationsResponse GetPageInkAnnotations(name, pageNumber, optional)
Read document page ink annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**InkAnnotationsResponse**](InkAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageLineAnnotations**
> LineAnnotationsResponse GetPageLineAnnotations(name, pageNumber, optional)
Read document page line annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**LineAnnotationsResponse**](LineAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageLinkAnnotation**
> LinkAnnotationResponse GetPageLinkAnnotation(name, pageNumber, linkId, optional)
Read document page link annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **linkId** | **string**| The link ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **linkId** | **string**| The link ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**LinkAnnotationResponse**](LinkAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageLinkAnnotations**
> LinkAnnotationsResponse GetPageLinkAnnotations(name, pageNumber, optional)
Read document page link annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**LinkAnnotationsResponse**](LinkAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageListBoxFields**
> ListBoxFieldsResponse GetPageListBoxFields(name, pageNumber, optional)
Read document page listbox fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**ListBoxFieldsResponse**](ListBoxFieldsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageMovieAnnotations**
> MovieAnnotationsResponse GetPageMovieAnnotations(name, pageNumber, optional)
Read document page movie annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**MovieAnnotationsResponse**](MovieAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagePolyLineAnnotations**
> PolyLineAnnotationsResponse GetPagePolyLineAnnotations(name, pageNumber, optional)
Read document page polyline annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PolyLineAnnotationsResponse**](PolyLineAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagePolygonAnnotations**
> PolygonAnnotationsResponse GetPagePolygonAnnotations(name, pageNumber, optional)
Read document page polygon annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PolygonAnnotationsResponse**](PolygonAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagePopupAnnotations**
> PopupAnnotationsResponse GetPagePopupAnnotations(name, pageNumber, optional)
Read document page popup annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PopupAnnotationsResponse**](PopupAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageRadioButtonFields**
> RadioButtonFieldsResponse GetPageRadioButtonFields(name, pageNumber, optional)
Read document page radiobutton fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**RadioButtonFieldsResponse**](RadioButtonFieldsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageRedactionAnnotations**
> RedactionAnnotationsResponse GetPageRedactionAnnotations(name, pageNumber, optional)
Read document page redaction annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**RedactionAnnotationsResponse**](RedactionAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageScreenAnnotations**
> ScreenAnnotationsResponse GetPageScreenAnnotations(name, pageNumber, optional)
Read document page screen annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**ScreenAnnotationsResponse**](ScreenAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageSignatureFields**
> SignatureFieldsResponse GetPageSignatureFields(name, pageNumber, optional)
Read document page signature fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SignatureFieldsResponse**](SignatureFieldsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageSoundAnnotations**
> SoundAnnotationsResponse GetPageSoundAnnotations(name, pageNumber, optional)
Read document page sound annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SoundAnnotationsResponse**](SoundAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageSquareAnnotations**
> SquareAnnotationsResponse GetPageSquareAnnotations(name, pageNumber, optional)
Read document page square annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SquareAnnotationsResponse**](SquareAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageSquigglyAnnotations**
> SquigglyAnnotationsResponse GetPageSquigglyAnnotations(name, pageNumber, optional)
Read document page squiggly annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SquigglyAnnotationsResponse**](SquigglyAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageStampAnnotations**
> StampAnnotationsResponse GetPageStampAnnotations(name, pageNumber, optional)
Read document page stamp annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**StampAnnotationsResponse**](StampAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageStamps**
> StampsInfoResponse GetPageStamps(name, pageNumber, optional)
Read page document stamps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| Base64 encoded password. | 

### Return type

[**StampsInfoResponse**](StampsInfoResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageStrikeOutAnnotations**
> StrikeOutAnnotationsResponse GetPageStrikeOutAnnotations(name, pageNumber, optional)
Read document page StrikeOut annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**StrikeOutAnnotationsResponse**](StrikeOutAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageTables**
> TablesRecognizedResponse GetPageTables(name, pageNumber, optional)
Read document page tables.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**|  | 
  **pageNumber** | **int32**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**|  | 
 **pageNumber** | **int32**|  | 
 **storage** | **string**|  | 
 **folder** | **string**|  | 

### Return type

[**TablesRecognizedResponse**](TablesRecognizedResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageText**
> TextRectsResponse GetPageText(name, pageNumber, lLX, lLY, uRX, uRY, optional)
Read page text items.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| Number of page (starting from 1). | 
  **lLX** | **float64**| X-coordinate of lower - left corner. | 
  **lLY** | **float64**| Y - coordinate of lower-left corner. | 
  **uRX** | **float64**| X - coordinate of upper-right corner. | 
  **uRY** | **float64**| Y - coordinate of upper-right corner. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| Number of page (starting from 1). | 
 **lLX** | **float64**| X-coordinate of lower - left corner. | 
 **lLY** | **float64**| Y - coordinate of lower-left corner. | 
 **uRX** | **float64**| X - coordinate of upper-right corner. | 
 **uRY** | **float64**| Y - coordinate of upper-right corner. | 
 **format** | [**[]string**](string.md)| List of formats for search. | 
 **regex** | **string**| Formats are specified as a regular expression. | 
 **splitRects** | **bool**| Split result fragments (default is true). | [default to true]
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**TextRectsResponse**](TextRectsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageTextAnnotations**
> TextAnnotationsResponse GetPageTextAnnotations(name, pageNumber, optional)
Read document page text annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**TextAnnotationsResponse**](TextAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageTextBoxFields**
> TextBoxFieldsResponse GetPageTextBoxFields(name, pageNumber, optional)
Read document page text box fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**TextBoxFieldsResponse**](TextBoxFieldsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageUnderlineAnnotations**
> UnderlineAnnotationsResponse GetPageUnderlineAnnotations(name, pageNumber, optional)
Read document page underline annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**UnderlineAnnotationsResponse**](UnderlineAnnotationsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPages**
> DocumentPagesResponse GetPages(name, optional)
Read document pages info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**DocumentPagesResponse**](DocumentPagesResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPclInStorageToPdf**
> []byte GetPclInStorageToPdf(srcPath, optional)
Convert PCL file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.pcl) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.pcl) | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfAInStorageToPdf**
> []byte GetPdfAInStorageToPdf(srcPath, optional)
Convert PDFA file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.pdf) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.pdf) | 
 **dontOptimize** | **bool**| If set, document resources will not be optimized. | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToAps**
> []byte GetPdfInStorageToAps(name, optional)
Converts PDF document (located on storage) to APS format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToDoc**
> []byte GetPdfInStorageToDoc(name, optional)
Converts PDF document (located on storage) to DOC format and returns resulting file in response content.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **addReturnToLineEnd** | **bool**| Add return to line end. | 
 **format** | **string**| Allows to specify .doc or .docx file format. | 
 **imageResolutionX** | **int32**| Image resolution X. | 
 **imageResolutionY** | **int32**| Image resolution Y. | 
 **maxDistanceBetweenTextLines** | **float64**| Max distance between text lines. | 
 **mode** | **string**| Allows to control how a PDF document is converted into a word processing document. | 
 **recognizeBullets** | **bool**| Recognize bullets. | 
 **relativeHorizontalProximity** | **float64**| Relative horizontal proximity. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToEpub**
> []byte GetPdfInStorageToEpub(name, optional)
Converts PDF document (located on storage) to EPUB format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **contentRecognitionMode** | **string**| Property tunes conversion for this or that desirable method of recognition of content. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToHtml**
> []byte GetPdfInStorageToHtml(name, optional)
Converts PDF document (located on storage) to Html format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **additionalMarginWidthInPoints** | **int32**| Defines width of margin that will be forcibly left around that output HTML-areas. | 
 **compressSvgGraphicsIfAny** | **bool**| The flag that indicates whether found SVG graphics(if any) will be compressed(zipped) into SVGZ format during saving. | 
 **convertMarkedContentToLayers** | **bool**| If attribute ConvertMarkedContentToLayers set to true then an all elements inside a PDF marked content (layer) will be put into an HTML div with &quot;data-pdflayer&quot; attribute specifying a layer name. This layer name will be extracted from optional properties of PDF marked content. If this attribute is false (by default) then no any layers will be created from PDF marked content. | 
 **defaultFontName** | **string**| Specifies the name of an installed font which is used to substitute any document font that is not embedded and not installed in the system. If null then default substitution font is used. | 
 **documentType** | **string**| Result document type. | 
 **fixedLayout** | **bool**| The value indicating whether that HTML is created as fixed layout. | 
 **imageResolution** | **int32**| Resolution for image rendering. | 
 **minimalLineWidth** | **int32**| This attribute sets minimal width of graphic path line. If thickness of line is less than 1px Adobe Acrobat rounds it to this value. So this attribute can be used to emulate this behavior for HTML browsers. | 
 **preventGlyphsGrouping** | **bool**| This attribute switch on the mode when text glyphs will not be grouped into words and strings This mode allows to keep maximum precision during positioning of glyphs on the page and it can be used for conversion documents with music notes or glyphs that should be placed separately each other. This parameter will be applied to document only when the value of FixedLayout attribute is true. | 
 **splitCssIntoPages** | **bool**| When multipage-mode selected(i.e &#39;SplitIntoPages&#39; is &#39;true&#39;), then this attribute defines whether should be created separate CSS-file for each result HTML page. | 
 **splitIntoPages** | **bool**| The flag that indicates whether each page of source document will be converted into it&#39;s own target HTML document, i.e whether result HTML will be splitted into several HTML-pages. | 
 **useZOrder** | **bool**| If attribute UseZORder set to true, graphics and text are added to resultant HTML document accordingly Z-order in original PDF document. If this attribute is false all graphics is put as single layer which may cause some unnecessary effects for overlapped objects. | 
 **antialiasingProcessing** | **string**| The parameter defines required antialiasing measures during conversion of compound background images from PDF to HTML. | 
 **cssClassNamesPrefix** | **string**| When PDFtoHTML converter generates result CSSs, CSS class names (something like &quot;.stl_01 {}&quot; ... &quot;.stl_NN {}) are generated and used in result CSS. This property allows forcibly set class name prefix. | 
 **explicitListOfSavedPages** | [**[]int32**](int32.md)| With this property You can explicitely define what pages of document should be converted. Pages in this list must have 1-based numbers. I.e. valid numbers of pages must be taken from range (1...[NumberOfPagesInConvertedDocument]) Order of appearing of pages in this list does not affect their order in result HTML page(s) - in result pages allways will go in order in which they are present in source PDF. | 
 **fontEncodingStrategy** | **string**| Defines encoding special rule to tune PDF decoding for current document. | 
 **fontSavingMode** | **string**| Defines font saving mode that will be used during saving of PDF to desirable format. | 
 **htmlMarkupGenerationMode** | **string**| Sometimes specific reqirments to generation of HTML markup are present. This parameter defines HTML preparing modes that can be used during conversion of PDF to HTML to match such specific requirments. | 
 **lettersPositioningMethod** | **string**| The mode of positioning of letters in words in result HTML. | 
 **pagesFlowTypeDependsOnViewersScreenSize** | **bool**| If attribute &#39;SplitOnPages&#x3D;false&#39;, than whole HTML representing all input PDF pages will be put into one big result HTML file. This flag defines whether result HTML will be generated in such way that flow of areas that represent PDF pages in result HTML will depend on screen resolution of viewer. | 
 **partsEmbeddingMode** | **string**| It defines whether referenced files (HTML, Fonts,Images, CSSes) will be embedded into main HTML file or will be generated as apart binary entities. | 
 **rasterImagesSavingMode** | **string**| Converted PDF can contain raster images This parameter defines how they should be handled during conversion of PDF to HTML. | 
 **removeEmptyAreasOnTopAndBottom** | **bool**| Defines whether in created HTML will be removed top and bottom empty area without any content (if any). | 
 **saveShadowedTextsAsTransparentTexts** | **bool**| Pdf can contain texts that are shadowed by another elements (f.e. by images) but can be selected to clipboard in Acrobat Reader (usually it happen when document contains images and OCRed texts extracted from it). This settings tells to converter whether we need save such texts as transparent selectable texts in result HTML to mimic behaviour of Acrobat Reader (othervise such texts are usually saved as hidden, not available for copying to clipboard). | 
 **saveTransparentTexts** | **bool**| Pdf can contain transparent texts that can be selected to clipboard (usually it happen when document contains images and OCRed texts extracted from it). This settings tells to converter whether we need save such texts as transparent selectable texts in result HTML. | 
 **specialFolderForAllImages** | **string**| The path to directory to which must be saved any images if they are encountered during saving of document as HTML. If parameter is empty or null then image files(if any) wil be saved together with other files linked to HTML It does not affect anything if CustomImageSavingStrategy property was successfully used to process relevant image file. | 
 **specialFolderForSvgImages** | **string**| The path to directory to which must be saved only SVG-images if they are encountered during saving of document as HTML. If parameter is empty or null then SVG files(if any) wil be saved together with other image-files (near to output file) or in special folder for images (if it specified in SpecialImagesFolderIfAny option). It does not affect anything if CustomImageSavingStrategy property was successfully used to process relevant image file. | 
 **trySaveTextUnderliningAndStrikeoutingInCss** | **bool**| PDF itself does not contain underlining markers for texts. It emulated with line situated under text. This option allows converter try guess that this or that line is a text&#39;s underlining and put this info into CSS instead of drawing of underlining graphically. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **flowLayoutParagraphFullWidth** | **bool**| This attribute specifies full width paragraph text for Flow mode, FixedLayout &#x3D; false. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToMobiXml**
> []byte GetPdfInStorageToMobiXml(name, optional)
Converts PDF document (located on storage) to MOBIXML format and returns resulting ZIP archive file in response content.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToPdfA**
> []byte GetPdfInStorageToPdfA(name, type_, optional)
Converts PDF document (located on storage) to PdfA format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **type_** | **string**| Type of PdfA format. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **type_** | **string**| Type of PdfA format. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToPptx**
> []byte GetPdfInStorageToPptx(name, optional)
Converts PDF document (located on storage) to PPTX format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **separateImages** | **bool**| Separate images. | 
 **slidesAsImages** | **bool**| Slides as images. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| Base64 encoded password. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToSvg**
> []byte GetPdfInStorageToSvg(name, optional)
Converts PDF document (located on storage) to SVG format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **compressOutputToZipArchive** | **bool**| Specifies whether output will be created as one zip-archive. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToTeX**
> []byte GetPdfInStorageToTeX(name, optional)
Converts PDF document (located on storage) to TeX format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToText**
> []byte GetPdfInStorageToText(name, optional)
Converts PDF document (located on storage) to Text format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToTiff**
> []byte GetPdfInStorageToTiff(name, optional)
Converts PDF document (located on storage) to TIFF format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **brightness** | **float64**| Image brightness. | 
 **compression** | **string**| Tiff compression. Possible values are: LZW, CCITT4, CCITT3, RLE, None. | 
 **colorDepth** | **string**| Image color depth. Possible valuse are: Default, Format8bpp, Format4bpp, Format1bpp. | 
 **leftMargin** | **int32**| Left image margin. | 
 **rightMargin** | **int32**| Right image margin. | 
 **topMargin** | **int32**| Top image margin. | 
 **bottomMargin** | **int32**| Bottom image margin. | 
 **orientation** | **string**| Image orientation. Possible values are: None, Landscape, Portait. | 
 **skipBlankPages** | **bool**| Skip blank pages flag. | 
 **width** | **int32**| Image width. | 
 **height** | **int32**| Image height. | 
 **xResolution** | **int32**| Horizontal resolution. | 
 **yResolution** | **int32**| Vertical resolution. | 
 **pageIndex** | **int32**| Start page to export. | 
 **pageCount** | **int32**| Number of pages to export. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToXls**
> []byte GetPdfInStorageToXls(name, optional)
Converts PDF document (located on storage) to XLS format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **insertBlankColumnAtFirst** | **bool**| Insert blank column at first | 
 **minimizeTheNumberOfWorksheets** | **bool**| Minimize the number of worksheets | 
 **scaleFactor** | **float64**| Scale factor (Obsolete) | 
 **uniformWorksheets** | **bool**| Uniform worksheets | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToXlsx**
> []byte GetPdfInStorageToXlsx(name, optional)
Converts PDF document (located on storage) to XLSX format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **insertBlankColumnAtFirst** | **bool**| Insert blank column at first | 
 **minimizeTheNumberOfWorksheets** | **bool**| Minimize the number of worksheets | 
 **scaleFactor** | **float64**| Scale factor (Obsolete) | 
 **uniformWorksheets** | **bool**| Uniform worksheets | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToXml**
> []byte GetPdfInStorageToXml(name, optional)
Converts PDF document (located on storage) to XML format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToXps**
> []byte GetPdfInStorageToXps(name, optional)
Converts PDF document (located on storage) to XPS format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolyLineAnnotation**
> PolyLineAnnotationResponse GetPolyLineAnnotation(name, annotationId, optional)
Read document page polyline annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PolyLineAnnotationResponse**](PolyLineAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolygonAnnotation**
> PolygonAnnotationResponse GetPolygonAnnotation(name, annotationId, optional)
Read document page polygon annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PolygonAnnotationResponse**](PolygonAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPopupAnnotation**
> PopupAnnotationResponse GetPopupAnnotation(name, annotationId, optional)
Read document page popup annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PopupAnnotationResponse**](PopupAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPsInStorageToPdf**
> []byte GetPsInStorageToPdf(srcPath, optional)
Convert PS file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.ps) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.ps) | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRadioButtonField**
> RadioButtonFieldResponse GetRadioButtonField(name, fieldName, optional)
Read document RadioButton field by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fieldName** | **string**| The field name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fieldName** | **string**| The field name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**RadioButtonFieldResponse**](RadioButtonFieldResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRedactionAnnotation**
> RedactionAnnotationResponse GetRedactionAnnotation(name, annotationId, optional)
Read document page redaction annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**RedactionAnnotationResponse**](RedactionAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScreenAnnotation**
> ScreenAnnotationResponse GetScreenAnnotation(name, annotationId, optional)
Read document page screen annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**ScreenAnnotationResponse**](ScreenAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScreenAnnotationData**
> []byte GetScreenAnnotationData(name, annotationId, optional)
Read document page screen annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSignatureField**
> SignatureFieldResponse GetSignatureField(name, fieldName, optional)
Read document signature field by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fieldName** | **string**| The field name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fieldName** | **string**| The field name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SignatureFieldResponse**](SignatureFieldResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSoundAnnotation**
> SoundAnnotationResponse GetSoundAnnotation(name, annotationId, optional)
Read document page sound annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SoundAnnotationResponse**](SoundAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSoundAnnotationData**
> []byte GetSoundAnnotationData(name, annotationId, optional)
Read document page sound annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSquareAnnotation**
> SquareAnnotationResponse GetSquareAnnotation(name, annotationId, optional)
Read document page square annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SquareAnnotationResponse**](SquareAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSquigglyAnnotation**
> SquigglyAnnotationResponse GetSquigglyAnnotation(name, annotationId, optional)
Read document page squiggly annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SquigglyAnnotationResponse**](SquigglyAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStampAnnotation**
> StampAnnotationResponse GetStampAnnotation(name, annotationId, optional)
Read document page stamp annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**StampAnnotationResponse**](StampAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStampAnnotationData**
> []byte GetStampAnnotationData(name, annotationId, optional)
Read document page stamp annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStrikeOutAnnotation**
> StrikeOutAnnotationResponse GetStrikeOutAnnotation(name, annotationId, optional)
Read document page StrikeOut annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**StrikeOutAnnotationResponse**](StrikeOutAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSvgInStorageToPdf**
> []byte GetSvgInStorageToPdf(srcPath, optional)
Convert SVG file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.svg) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.svg) | 
 **adjustPageSize** | **bool**| Adjust page size | 
 **height** | **float64**| Page height | 
 **width** | **float64**| Page width | 
 **isLandscape** | **bool**| Is page landscaped | 
 **marginLeft** | **float64**| Page margin left | 
 **marginBottom** | **float64**| Page margin bottom | 
 **marginRight** | **float64**| Page margin right | 
 **marginTop** | **float64**| Page margin top | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTable**
> TableRecognizedResponse GetTable(name, tableId, optional)
Read document page table by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **tableId** | **string**| The table ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **tableId** | **string**| The table ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**TableRecognizedResponse**](TableRecognizedResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeXInStorageToPdf**
> []byte GetTeXInStorageToPdf(srcPath, optional)
Convert TeX file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.tex) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.tex) | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetText**
> TextRectsResponse GetText(name, lLX, lLY, uRX, uRY, optional)
Read document text.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **lLX** | **float64**| X-coordinate of lower - left corner. | 
  **lLY** | **float64**| Y - coordinate of lower-left corner. | 
  **uRX** | **float64**| X - coordinate of upper-right corner. | 
  **uRY** | **float64**| Y - coordinate of upper-right corner. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **lLX** | **float64**| X-coordinate of lower - left corner. | 
 **lLY** | **float64**| Y - coordinate of lower-left corner. | 
 **uRX** | **float64**| X - coordinate of upper-right corner. | 
 **uRY** | **float64**| Y - coordinate of upper-right corner. | 
 **format** | [**[]string**](string.md)| List of formats for search. | 
 **regex** | **string**| Formats are specified as a regular expression. | 
 **splitRects** | **bool**| Split result fragments (default is true). | [default to true]
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**TextRectsResponse**](TextRectsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTextAnnotation**
> TextAnnotationResponse GetTextAnnotation(name, annotationId, optional)
Read document page text annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**TextAnnotationResponse**](TextAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTextBoxField**
> TextBoxFieldResponse GetTextBoxField(name, fieldName, optional)
Read document text box field by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fieldName** | **string**| The field name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fieldName** | **string**| The field name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**TextBoxFieldResponse**](TextBoxFieldResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUnderlineAnnotation**
> UnderlineAnnotationResponse GetUnderlineAnnotation(name, annotationId, optional)
Read document page underline annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**UnderlineAnnotationResponse**](UnderlineAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVerifySignature**
> SignatureVerifyResponse GetVerifySignature(name, signName, optional)
Verify signature document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **signName** | **string**| Sign name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **signName** | **string**| Sign name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SignatureVerifyResponse**](SignatureVerifyResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebInStorageToPdf**
> []byte GetWebInStorageToPdf(url, optional)
Convert web page to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string**| Source url | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string**| Source url | 
 **height** | **float64**| Page height | 
 **width** | **float64**| Page width | 
 **isLandscape** | **bool**| Is page landscaped | 
 **marginLeft** | **float64**| Page margin left | 
 **marginBottom** | **float64**| Page margin bottom | 
 **marginRight** | **float64**| Page margin right | 
 **marginTop** | **float64**| Page margin top | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWordsPerPage**
> WordCountResponse GetWordsPerPage(name, optional)
Get number of words per document page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**WordCountResponse**](WordCountResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetXfaPdfInStorageToAcroForm**
> []byte GetXfaPdfInStorageToAcroForm(name, optional)
Converts PDF document which contains XFA form (located on storage) to PDF with AcroForm and returns resulting file response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetXmlInStorageToPdf**
> []byte GetXmlInStorageToPdf(srcPath, optional)
Convert XML file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xml) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xml) | 
 **xslFilePath** | **string**| Full XSL source filename (ex. /folder1/folder2/template.xsl) | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetXmpMetadataJson**
> XmpMetadata GetXmpMetadataJson(name, optional)
Gets document XMP Metadata as JSON.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **passBase64** | **string**| The password (Base64). | 

### Return type

[**XmpMetadata**](XmpMetadata.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetXmpMetadataXml**
> []byte GetXmpMetadataXml(name, optional)
Gets document XMP Metadata as XML file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **passBase64** | **string**| The password (Base64). | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetXpsInStorageToPdf**
> []byte GetXpsInStorageToPdf(srcPath, optional)
Convert XPS file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xps) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xps) | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetXslFoInStorageToPdf**
> []byte GetXslFoInStorageToPdf(srcPath, optional)
Convert XslFo file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xslfo) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xslfo) | 
 **storage** | **string**| The document storage. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveFile**
> MoveFile(srcPath, destPath, optional)
Move file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Source file path e.g. &#39;/src.ext&#39; | 
  **destPath** | **string**| Destination file path e.g. &#39;/dest.ext&#39; | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Source file path e.g. &#39;/src.ext&#39; | 
 **destPath** | **string**| Destination file path e.g. &#39;/dest.ext&#39; | 
 **srcStorageName** | **string**| Source storage name | 
 **destStorageName** | **string**| Destination storage name | 
 **versionId** | **string**| File version ID to move | 

### Return type

 (empty response body)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveFolder**
> MoveFolder(srcPath, destPath, optional)
Move folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Folder path to move e.g. &#39;/folder&#39; | 
  **destPath** | **string**| Destination folder path to move to e.g &#39;/dst&#39; | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Folder path to move e.g. &#39;/folder&#39; | 
 **destPath** | **string**| Destination folder path to move to e.g &#39;/dst&#39; | 
 **srcStorageName** | **string**| Source storage name | 
 **destStorageName** | **string**| Destination storage name | 

### Return type

 (empty response body)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ObjectExists**
> ObjectExist ObjectExists(path, optional)
Check if file or folder exists

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| File or folder path e.g. &#39;/file.ext&#39; or &#39;/folder&#39; | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| File or folder path e.g. &#39;/file.ext&#39; or &#39;/folder&#39; | 
 **storageName** | **string**| Storage name | 
 **versionId** | **string**| File version ID | 

### Return type

[**ObjectExist**](ObjectExist.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAddDocumentAttachment**
> AttachmentsResponse PostAddDocumentAttachment(name, attachmentInfo, optional)
Adds a file attachment to the PDF document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **attachmentInfo** | [**AttachmentInfo**](AttachmentInfo.md)| AttachmentInfoAttachmentInfo instance. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **attachmentInfo** | [**AttachmentInfo**](AttachmentInfo.md)| AttachmentInfoAttachmentInfo instance. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AttachmentsResponse**](AttachmentsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAppendDocument**
> DocumentResponse PostAppendDocument(name, appendFile, optional)
Append document to existing one.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The original document name. | 
  **appendFile** | **string**| Append file server path. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The original document name. | 
 **appendFile** | **string**| Append file server path. | 
 **startPage** | **int32**| Appending start page. | [default to 0]
 **endPage** | **int32**| Appending end page. | [default to 0]
 **storage** | **string**| The documents storage. | 
 **folder** | **string**| The original document folder. | 

### Return type

[**DocumentResponse**](DocumentResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostBookmark**
> BookmarksResponse PostBookmark(name, bookmarkPath, bookmarks, optional)
Add document bookmarks.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **bookmarkPath** | **string**| The parent bookmark path. Specify an empty string when adding a bookmark to the root. | 
  **bookmarks** | [**[]Bookmark**](Bookmark.md)| The array of bookmark. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **bookmarkPath** | **string**| The parent bookmark path. Specify an empty string when adding a bookmark to the root. | 
 **bookmarks** | [**[]Bookmark**](Bookmark.md)| The array of bookmark. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**BookmarksResponse**](BookmarksResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostChangePasswordDocumentInStorage**
> AsposeResponse PostChangePasswordDocumentInStorage(name, ownerPassword, newUserPassword, newOwnerPassword, optional)
Change document password in storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Document name. | 
  **ownerPassword** | **string**| Owner password (encrypted Base64). | 
  **newUserPassword** | **string**| New user password (encrypted Base64). | 
  **newOwnerPassword** | **string**| New owner password (encrypted Base64). | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Document name. | 
 **ownerPassword** | **string**| Owner password (encrypted Base64). | 
 **newUserPassword** | **string**| New user password (encrypted Base64). | 
 **newOwnerPassword** | **string**| New owner password (encrypted Base64). | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCheckBoxFields**
> AsposeResponse PostCheckBoxFields(name, fields, optional)
Add document checkbox fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fields** | [**[]CheckBoxField**](CheckBoxField.md)| The array of field. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fields** | [**[]CheckBoxField**](CheckBoxField.md)| The array of field. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostComboBoxFields**
> AsposeResponse PostComboBoxFields(name, fields, optional)
Add document combobox fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fields** | [**[]ComboBoxField**](ComboBoxField.md)| The array of field. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fields** | [**[]ComboBoxField**](ComboBoxField.md)| The array of field. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostComparePdf**
> AsposeResponse PostComparePdf(path1, path2, outPath, optional)
Compare two PDF documents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path1** | **string**| Path to first PDF document. | 
  **path2** | **string**| Path to second PDF document. | 
  **outPath** | **string**| Full filename of the resulting document. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path1** | **string**| Path to first PDF document. | 
 **path2** | **string**| Path to second PDF document. | 
 **outPath** | **string**| Full filename of the resulting document. | 
 **storage** | **string**| The documents storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCreateDocument**
> DocumentResponse PostCreateDocument(name, documentConfig, optional)
Create empty document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The new document name. | 
  **documentConfig** | [**DocumentConfig**](DocumentConfig.md)| The document config for new document. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The new document name. | 
 **documentConfig** | [**DocumentConfig**](DocumentConfig.md)| The document config for new document. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The new document folder. | 

### Return type

[**DocumentResponse**](DocumentResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCreateField**
> AsposeResponse PostCreateField(name, page, field, optional)
Create field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **page** | **int32**| Document page number. | 
  **field** | [**Field**](Field.md)| Field with the field data. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **page** | **int32**| Document page number. | 
 **field** | [**Field**](Field.md)| Field with the field data. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDecryptDocumentInStorage**
> AsposeResponse PostDecryptDocumentInStorage(name, password, optional)
Decrypt document in storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Document name. | 
  **password** | **string**| The password (encrypted Base64). | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Document name. | 
 **password** | **string**| The password (encrypted Base64). | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDocumentImageFooter**
> AsposeResponse PostDocumentImageFooter(name, imageFooter, optional)
Add document image footer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **imageFooter** | [**ImageFooter**](ImageFooter.md)| The image footer. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageFooter** | [**ImageFooter**](ImageFooter.md)| The image footer. | 
 **startPageNumber** | **int32**| The start page number. | 
 **endPageNumber** | **int32**| The end page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDocumentImageHeader**
> AsposeResponse PostDocumentImageHeader(name, imageHeader, optional)
Add document image header.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **imageHeader** | [**ImageHeader**](ImageHeader.md)| The image header. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageHeader** | [**ImageHeader**](ImageHeader.md)| The image header. | 
 **startPageNumber** | **int32**| The start page number. | 
 **endPageNumber** | **int32**| The end page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDocumentImageStamps**
> AsposeResponse PostDocumentImageStamps(name, stamps, optional)
Add document pages image stamps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **stamps** | [**[]ImageStamp**](ImageStamp.md)| The array of stamp. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **stamps** | [**[]ImageStamp**](ImageStamp.md)| The array of stamp. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| Base64 encoded password. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDocumentImageStampsPageSpecified**
> AsposeResponse PostDocumentImageStampsPageSpecified(name, stamps, optional)
Add document image stamps to specified pages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **stamps** | [**[]ImageStampPageSpecified**](ImageStampPageSpecified.md)| The array of stamps. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **stamps** | [**[]ImageStampPageSpecified**](ImageStampPageSpecified.md)| The array of stamps. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| Base64 encoded password. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDocumentPageNumberStamps**
> AsposeResponse PostDocumentPageNumberStamps(name, stamp, optional)
Add document page number stamps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **stamp** | [**PageNumberStamp**](PageNumberStamp.md)| The stamp. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **stamp** | [**PageNumberStamp**](PageNumberStamp.md)| The stamp. | 
 **startPageNumber** | **int32**| The start page number. | 
 **endPageNumber** | **int32**| The end page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| Base64 encoded password. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDocumentPagesCrop**
> AsposeResponse PostDocumentPagesCrop(name, pages, rect, optional)
Crop PDF document pages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pages** | **string**| Comma separated list of pages and page ranges. (Example: 1,3-5,8) | 
  **rect** | [**Rectangle**](Rectangle.md)| Rectangle of document area. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pages** | **string**| Comma separated list of pages and page ranges. (Example: 1,3-5,8) | 
 **rect** | [**Rectangle**](Rectangle.md)| Rectangle of document area. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| Base64 encoded password. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDocumentPagesResize**
> AsposeResponse PostDocumentPagesResize(name, height, width, pages, optional)
Rsize PDF document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **height** | **float64**| Page Height | 
  **width** | **float64**| Page width | 
  **pages** | **string**| Comma separated list of pages and page ranges. (Example: 1,3-5,8) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **height** | **float64**| Page Height | 
 **width** | **float64**| Page width | 
 **pages** | **string**| Comma separated list of pages and page ranges. (Example: 1,3-5,8) | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| Base64 encoded password. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDocumentPagesRotate**
> AsposeResponse PostDocumentPagesRotate(name, rotationAngle, pages, optional)
Rotate PDF document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **rotationAngle** | **string**| Rotation Angle (CKW). Can be 90, 180, 270. | 
  **pages** | **string**| Comma separated list of pages and page ranges. (Example: 1,3-5,8) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **rotationAngle** | **string**| Rotation Angle (CKW). Can be 90, 180, 270. | 
 **pages** | **string**| Comma separated list of pages and page ranges. (Example: 1,3-5,8) | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| Base64 encoded password. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDocumentTextFooter**
> AsposeResponse PostDocumentTextFooter(name, textFooter, optional)
Add document text footer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **textFooter** | [**TextFooter**](TextFooter.md)| The text footer. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **textFooter** | [**TextFooter**](TextFooter.md)| The text footer. | 
 **startPageNumber** | **int32**| The start page number. | 
 **endPageNumber** | **int32**| The end page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDocumentTextHeader**
> AsposeResponse PostDocumentTextHeader(name, textHeader, optional)
Add document text header.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **textHeader** | [**TextHeader**](TextHeader.md)| The text header. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **textHeader** | [**TextHeader**](TextHeader.md)| The text header. | 
 **startPageNumber** | **int32**| The start page number. | 
 **endPageNumber** | **int32**| The end page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDocumentTextReplace**
> TextReplaceResponse PostDocumentTextReplace(name, textReplace, optional)
Document's replace text method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**|  | 
  **textReplace** | [**TextReplaceListRequest**](TextReplaceListRequest.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**|  | 
 **textReplace** | [**TextReplaceListRequest**](TextReplaceListRequest.md)|  | 
 **storage** | **string**|  | 
 **folder** | **string**|  | 

### Return type

[**TextReplaceResponse**](TextReplaceResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDocumentTextStamps**
> AsposeResponse PostDocumentTextStamps(name, stamps, optional)
Add document pages text stamps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **stamps** | [**[]TextStamp**](TextStamp.md)| The array of stamp. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **stamps** | [**[]TextStamp**](TextStamp.md)| The array of stamp. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| Base64 encoded password. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDocumentTextStampsPageSpecified**
> AsposeResponse PostDocumentTextStampsPageSpecified(name, stamps, optional)
Add document text stamps to specified pages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **stamps** | [**[]TextStampPageSpecified**](TextStampPageSpecified.md)| The array of stamps. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **stamps** | [**[]TextStampPageSpecified**](TextStampPageSpecified.md)| The array of stamps. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| Base64 encoded password. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostEncryptDocumentInStorage**
> AsposeResponse PostEncryptDocumentInStorage(name, userPassword, ownerPassword, cryptoAlgorithm, optional)
Encrypt document in storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Document name. | 
  **userPassword** | **string**| User password (encrypted Base64). | 
  **ownerPassword** | **string**| Owner password (encrypted Base64). | 
  **cryptoAlgorithm** | **string**| Cryptographic algorithm, see CryptoAlgorithm for details. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Document name. | 
 **userPassword** | **string**| User password (encrypted Base64). | 
 **ownerPassword** | **string**| Owner password (encrypted Base64). | 
 **cryptoAlgorithm** | **string**| Cryptographic algorithm, see CryptoAlgorithm for details. | 
 **permissionsFlags** | [**[]PermissionsFlags**](PermissionsFlags.md)| Array of document permissions, see PermissionsFlags for details. | 
 **usePdf20** | **bool**| Support for revision 6 (Extension 8). | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFlattenDocument**
> AsposeResponse PostFlattenDocument(name, optional)
Flatten the document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **updateAppearances** | **bool**| If set, all field appearances will be regenerated before flattening. This option may help if field is incorrectly flattened. This option may decrease performance.. | 
 **callEvents** | **bool**| If set, formatting and other JavaScript events will be called. | 
 **hideButtons** | **bool**| If set, buttons will be removed from flattened document. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostHtmlToPdf**
> []byte PostHtmlToPdf(optional)
Convert HTML file (zip archive in request content) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **htmlFileName** | **string**| Name of HTML file in ZIP. | 
 **height** | **float64**| Page height | 
 **width** | **float64**| Page width | 
 **isLandscape** | **bool**| Is page landscaped | 
 **marginLeft** | **float64**| Page margin left | 
 **marginBottom** | **float64**| Page margin bottom | 
 **marginRight** | **float64**| Page margin right | 
 **marginTop** | **float64**| Page margin top | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostImportFieldsFromFdf**
> AsposeResponse PostImportFieldsFromFdf(name, optional)
Update fields from FDF file in request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **fdfData** | ***os.File**| Fdf file. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostImportFieldsFromXfdf**
> AsposeResponse PostImportFieldsFromXfdf(name, optional)
Update fields from XFDF file in request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **xfdfData** | ***os.File**| Xfdf file. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostImportFieldsFromXml**
> AsposeResponse PostImportFieldsFromXml(name, optional)
Update fields from XML file in request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **xmlData** | ***os.File**| Xml file. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostInsertImage**
> AsposeResponse PostInsertImage(name, pageNumber, llx, lly, urx, ury, optional)
Insert image to document page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **llx** | **float64**| Coordinate lower left X. | 
  **lly** | **float64**| Coordinate lower left Y. | 
  **urx** | **float64**| Coordinate upper right X. | 
  **ury** | **float64**| Coordinate upper right Y. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **llx** | **float64**| Coordinate lower left X. | 
 **lly** | **float64**| Coordinate lower left Y. | 
 **urx** | **float64**| Coordinate upper right X. | 
 **ury** | **float64**| Coordinate upper right Y. | 
 **imageFilePath** | **string**| Path to image file if specified. Request content is used otherwise. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **image** | ***os.File**| Image file. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostListBoxFields**
> AsposeResponse PostListBoxFields(name, fields, optional)
Add document listbox fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fields** | [**[]ListBoxField**](ListBoxField.md)| The array of field. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fields** | [**[]ListBoxField**](ListBoxField.md)| The array of field. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMovePage**
> AsposeResponse PostMovePage(name, pageNumber, newIndex, optional)
Move page to new position.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **newIndex** | **int32**| The new page position/index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **newIndex** | **int32**| The new page position/index. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostOptimizeDocument**
> AsposeResponse PostOptimizeDocument(name, options, optional)
Optimize document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **options** | [**OptimizeOptions**](OptimizeOptions.md)| The optimization options. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **options** | [**OptimizeOptions**](OptimizeOptions.md)| The optimization options. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostOrganizeDocument**
> AsposeResponse PostOrganizeDocument(name, pages, outPath, optional)
Merge selected pages of a document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The original document name. | 
  **pages** | **string**| 1-based page numbers of the source document that make up the resulting document. | 
  **outPath** | **string**| Full filename of the resulting document. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The original document name. | 
 **pages** | **string**| 1-based page numbers of the source document that make up the resulting document. | 
 **outPath** | **string**| Full filename of the resulting document. | 
 **storage** | **string**| The documents storage. | 
 **folder** | **string**| The source document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostOrganizeDocuments**
> AsposeResponse PostOrganizeDocuments(organizeDocuments, outPath, optional)
Merge selected pages of different documents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizeDocuments** | [**OrganizeDocumentRequest**](OrganizeDocumentRequest.md)| Array of OrganizeDocumentData to make up the resulting document. | 
  **outPath** | **string**| Full filename of the resulting document. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizeDocuments** | [**OrganizeDocumentRequest**](OrganizeDocumentRequest.md)| Array of OrganizeDocumentData to make up the resulting document. | 
 **outPath** | **string**| Full filename of the resulting document. | 
 **storage** | **string**| The documents storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageCaretAnnotations**
> AsposeResponse PostPageCaretAnnotations(name, pageNumber, annotations, optional)
Add document page caret annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]CaretAnnotation**](CaretAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]CaretAnnotation**](CaretAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageCertify**
> AsposeResponse PostPageCertify(name, pageNumber, sign, docMdpAccessPermissionType, optional)
Certify document page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **sign** | [**Signature**](Signature.md)| Signature object containing signature data. | 
  **docMdpAccessPermissionType** | **string**| The access permissions granted for this document. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **sign** | [**Signature**](Signature.md)| Signature object containing signature data. | 
 **docMdpAccessPermissionType** | **string**| The access permissions granted for this document. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageCircleAnnotations**
> AsposeResponse PostPageCircleAnnotations(name, pageNumber, annotations, optional)
Add document page circle annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]CircleAnnotation**](CircleAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]CircleAnnotation**](CircleAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageFileAttachmentAnnotations**
> AsposeResponse PostPageFileAttachmentAnnotations(name, pageNumber, annotations, optional)
Add document page FileAttachment annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]FileAttachmentAnnotation**](FileAttachmentAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]FileAttachmentAnnotation**](FileAttachmentAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageFreeTextAnnotations**
> AsposeResponse PostPageFreeTextAnnotations(name, pageNumber, annotations, optional)
Add document page free text annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]FreeTextAnnotation**](FreeTextAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]FreeTextAnnotation**](FreeTextAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageHighlightAnnotations**
> AsposeResponse PostPageHighlightAnnotations(name, pageNumber, annotations, optional)
Add document page highlight annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]HighlightAnnotation**](HighlightAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]HighlightAnnotation**](HighlightAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageImageStamps**
> AsposeResponse PostPageImageStamps(name, pageNumber, stamps, optional)
Add document page image stamps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **stamps** | [**[]ImageStamp**](ImageStamp.md)| The array of stamp. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **stamps** | [**[]ImageStamp**](ImageStamp.md)| The array of stamp. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| Base64 encoded password. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageInkAnnotations**
> AsposeResponse PostPageInkAnnotations(name, pageNumber, annotations, optional)
Add document page ink annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]InkAnnotation**](InkAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]InkAnnotation**](InkAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageLineAnnotations**
> AsposeResponse PostPageLineAnnotations(name, pageNumber, annotations, optional)
Add document page line annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]LineAnnotation**](LineAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]LineAnnotation**](LineAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageLinkAnnotations**
> AsposeResponse PostPageLinkAnnotations(name, pageNumber, links, optional)
Add document page link annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **links** | [**[]LinkAnnotation**](LinkAnnotation.md)| Array of link annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **links** | [**[]LinkAnnotation**](LinkAnnotation.md)| Array of link annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageMovieAnnotations**
> AsposeResponse PostPageMovieAnnotations(name, pageNumber, annotations, optional)
Add document page movie annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]MovieAnnotation**](MovieAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]MovieAnnotation**](MovieAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagePdfPageStamps**
> AsposeResponse PostPagePdfPageStamps(name, pageNumber, stamps, optional)
Add document pdf page stamps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **stamps** | [**[]PdfPageStamp**](PdfPageStamp.md)| The array of stamp. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **stamps** | [**[]PdfPageStamp**](PdfPageStamp.md)| The array of stamp. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| Base64 encoded password. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagePolyLineAnnotations**
> AsposeResponse PostPagePolyLineAnnotations(name, pageNumber, annotations, optional)
Add document page polyline annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]PolyLineAnnotation**](PolyLineAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]PolyLineAnnotation**](PolyLineAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagePolygonAnnotations**
> AsposeResponse PostPagePolygonAnnotations(name, pageNumber, annotations, optional)
Add document page polygon annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]PolygonAnnotation**](PolygonAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]PolygonAnnotation**](PolygonAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageRedactionAnnotations**
> AsposeResponse PostPageRedactionAnnotations(name, pageNumber, annotations, optional)
Add document page redaction annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]RedactionAnnotation**](RedactionAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]RedactionAnnotation**](RedactionAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **apply** | **bool**| Apply redaction immediately after adding. | [default to false]

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageScreenAnnotations**
> AsposeResponse PostPageScreenAnnotations(name, pageNumber, annotations, optional)
Add document page screen annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]ScreenAnnotation**](ScreenAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]ScreenAnnotation**](ScreenAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageSoundAnnotations**
> AsposeResponse PostPageSoundAnnotations(name, pageNumber, annotations, optional)
Add document page sound annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]SoundAnnotation**](SoundAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]SoundAnnotation**](SoundAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageSquareAnnotations**
> AsposeResponse PostPageSquareAnnotations(name, pageNumber, annotations, optional)
Add document page square annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]SquareAnnotation**](SquareAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]SquareAnnotation**](SquareAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageSquigglyAnnotations**
> AsposeResponse PostPageSquigglyAnnotations(name, pageNumber, annotations, optional)
Add document page squiggly annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]SquigglyAnnotation**](SquigglyAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]SquigglyAnnotation**](SquigglyAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageStampAnnotations**
> AsposeResponse PostPageStampAnnotations(name, pageNumber, annotations, optional)
Add document page stamp annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]StampAnnotation**](StampAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]StampAnnotation**](StampAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageStrikeOutAnnotations**
> AsposeResponse PostPageStrikeOutAnnotations(name, pageNumber, annotations, optional)
Add document page StrikeOut annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]StrikeOutAnnotation**](StrikeOutAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]StrikeOutAnnotation**](StrikeOutAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageTables**
> AsposeResponse PostPageTables(name, pageNumber, tables, optional)
Add document page tables.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **tables** | [**[]Table**](Table.md)| The array of table. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **tables** | [**[]Table**](Table.md)| The array of table. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageTextAnnotations**
> AsposeResponse PostPageTextAnnotations(name, pageNumber, annotations, optional)
Add document page text annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]TextAnnotation**](TextAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]TextAnnotation**](TextAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageTextReplace**
> TextReplaceResponse PostPageTextReplace(name, pageNumber, textReplaceListRequest, optional)
Page's replace text method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**|  | 
  **pageNumber** | **int32**|  | 
  **textReplaceListRequest** | [**TextReplaceListRequest**](TextReplaceListRequest.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**|  | 
 **pageNumber** | **int32**|  | 
 **textReplaceListRequest** | [**TextReplaceListRequest**](TextReplaceListRequest.md)|  | 
 **storage** | **string**|  | 
 **folder** | **string**|  | 

### Return type

[**TextReplaceResponse**](TextReplaceResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageTextStamps**
> AsposeResponse PostPageTextStamps(name, pageNumber, stamps, optional)
Add document page text stamps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **stamps** | [**[]TextStamp**](TextStamp.md)| The array of stamp. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **stamps** | [**[]TextStamp**](TextStamp.md)| The array of stamp. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| Base64 encoded password. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageUnderlineAnnotations**
> AsposeResponse PostPageUnderlineAnnotations(name, pageNumber, annotations, optional)
Add document page underline annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]UnderlineAnnotation**](UnderlineAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]UnderlineAnnotation**](UnderlineAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPdfToXlsx**
> []byte PostPdfToXlsx(optional)
Converts PDF document (in request content) to XLSX format and uploads and returns resulting file in response content.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insertBlankColumnAtFirst** | **bool**| Insert blank column at first | 
 **minimizeTheNumberOfWorksheets** | **bool**| Minimize the number of worksheets | 
 **uniformWorksheets** | **bool**| Uniform worksheets | 
 **password** | **string**| The password (Base64). | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

**[]byte**

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPopupAnnotation**
> AsposeResponse PostPopupAnnotation(name, annotationId, annotation, optional)
Add document popup annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The parent annotation ID. | 
  **annotation** | [**PopupAnnotation**](PopupAnnotation.md)| The annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The parent annotation ID. | 
 **annotation** | [**PopupAnnotation**](PopupAnnotation.md)| The annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostRadioButtonFields**
> AsposeResponse PostRadioButtonFields(name, fields, optional)
Add document RadioButton fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fields** | [**[]RadioButtonField**](RadioButtonField.md)| The array of field. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fields** | [**[]RadioButtonField**](RadioButtonField.md)| The array of field. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSignDocument**
> AsposeResponse PostSignDocument(name, sign, optional)
Sign document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **sign** | [**Signature**](Signature.md)| Signature object containing signature data. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **sign** | [**Signature**](Signature.md)| Signature object containing signature data. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSignPage**
> AsposeResponse PostSignPage(name, pageNumber, sign, optional)
Sign page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **sign** | [**Signature**](Signature.md)| Signature object containing signature data. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **sign** | [**Signature**](Signature.md)| Signature object containing signature data. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSignatureField**
> AsposeResponse PostSignatureField(name, field, optional)
Add document signature field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **field** | [**SignatureField**](SignatureField.md)| The field. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **field** | [**SignatureField**](SignatureField.md)| The field. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSplitDocument**
> SplitResultResponse PostSplitDocument(name, optional)
Split document to parts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Document name. | 
 **format** | **string**| Resulting documents format. | 
 **from** | **int32**| Start page if defined. | 
 **to** | **int32**| End page if defined. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| Base64 encoded password. | 

### Return type

[**SplitResultResponse**](SplitResultResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSplitRangePdfDocument**
> SplitResultResponse PostSplitRangePdfDocument(name, options, optional)
Split document into ranges.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **options** | [**SplitRangePdfOptions**](SplitRangePdfOptions.md)| The splitting options. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **options** | [**SplitRangePdfOptions**](SplitRangePdfOptions.md)| The splitting options. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| Base64 encoded password. | 

### Return type

[**SplitResultResponse**](SplitResultResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTextBoxFields**
> AsposeResponse PostTextBoxFields(name, fields, optional)
Add document text box fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fields** | [**[]TextBoxField**](TextBoxField.md)| The array of field. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fields** | [**[]TextBoxField**](TextBoxField.md)| The array of field. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostXmpMetadata**
> AsposeResponse PostXmpMetadata(name, metadata, optional)
Add or remove XMP Metadata properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **metadata** | [**XmpMetadata**](XmpMetadata.md)| XmpMetadata instance. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **metadata** | [**XmpMetadata**](XmpMetadata.md)| XmpMetadata instance. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **passBase64** | **string**| The password (Base64). | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAddNewPage**
> DocumentPagesResponse PutAddNewPage(name, optional)
Add new page to end of the document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**DocumentPagesResponse**](DocumentPagesResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAddText**
> AsposeResponse PutAddText(name, pageNumber, paragraph, optional)
Add text to PDF document page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| Number of page (starting from 1). | 
  **paragraph** | [**Paragraph**](Paragraph.md)| Paragraph data. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| Number of page (starting from 1). | 
 **paragraph** | [**Paragraph**](Paragraph.md)| Paragraph data. | 
 **folder** | **string**| Document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAnnotationsFlatten**
> AsposeResponse PutAnnotationsFlatten(name, optional)
Flattens the annotations of the specified types

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **startPage** | **int32**| The start page number. | 
 **endPage** | **int32**| The end page number. | 
 **annotationTypes** | [**[]AnnotationType**](AnnotationType.md)| Array of annotation types. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApsInStorageToPdf**
> AsposeResponse PutApsInStorageToPdf(name, srcPath, optional)
Convert APS file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xps) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xps) | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutBookmark**
> BookmarkResponse PutBookmark(name, bookmarkPath, bookmark, optional)
Update document bookmark.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **bookmarkPath** | **string**| The bookmark path. | 
  **bookmark** | [**Bookmark**](Bookmark.md)| The bookmark. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **bookmarkPath** | **string**| The bookmark path. | 
 **bookmark** | [**Bookmark**](Bookmark.md)| The bookmark. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**BookmarkResponse**](BookmarkResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCaretAnnotation**
> CaretAnnotationResponse PutCaretAnnotation(name, annotationId, annotation, optional)
Replace document caret annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**CaretAnnotation**](CaretAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**CaretAnnotation**](CaretAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**CaretAnnotationResponse**](CaretAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutChangePasswordDocument**
> AsposeResponse PutChangePasswordDocument(outPath, ownerPassword, newUserPassword, newOwnerPassword, optional)
Change document password from content.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.doc) | 
  **ownerPassword** | **string**| Owner password (encrypted Base64). | 
  **newUserPassword** | **string**| New user password (encrypted Base64). | 
  **newOwnerPassword** | **string**| New owner password (encrypted Base64). | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.doc) | 
 **ownerPassword** | **string**| Owner password (encrypted Base64). | 
 **newUserPassword** | **string**| New user password (encrypted Base64). | 
 **newOwnerPassword** | **string**| New owner password (encrypted Base64). | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be changed password. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCheckBoxField**
> CheckBoxFieldResponse PutCheckBoxField(name, fieldName, field, optional)
Replace document checkbox field

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fieldName** | **string**| The field name. | 
  **field** | [**CheckBoxField**](CheckBoxField.md)| The field. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fieldName** | **string**| The field name. | 
 **field** | [**CheckBoxField**](CheckBoxField.md)| The field. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**CheckBoxFieldResponse**](CheckBoxFieldResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCircleAnnotation**
> CircleAnnotationResponse PutCircleAnnotation(name, annotationId, annotation, optional)
Replace document circle annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**CircleAnnotation**](CircleAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**CircleAnnotation**](CircleAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**CircleAnnotationResponse**](CircleAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutComboBoxField**
> ComboBoxFieldResponse PutComboBoxField(name, fieldName, field, optional)
Replace document combobox field

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fieldName** | **string**| The field name. | 
  **field** | [**ComboBoxField**](ComboBoxField.md)| The field. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fieldName** | **string**| The field name. | 
 **field** | [**ComboBoxField**](ComboBoxField.md)| The field. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**ComboBoxFieldResponse**](ComboBoxFieldResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCreateDocument**
> DocumentResponse PutCreateDocument(name, optional)
Create empty document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The new document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The new document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The new document folder. | 

### Return type

[**DocumentResponse**](DocumentResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCreatePdfFromLayer**
> AsposeResponse PutCreatePdfFromLayer(name, pageNumber, outPath, layerId, optional)
Create a separate PDF from a PDF Layer and upload resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **outPath** | **string**| The out path of result image. | 
  **layerId** | **string**| Layer Id. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **outPath** | **string**| The out path of result image. | 
 **layerId** | **string**| Layer Id. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **passBase64** | **string**| The password (Base64). | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutDecryptDocument**
> AsposeResponse PutDecryptDocument(outPath, password, optional)
Decrypt document from content.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.doc) | 
  **password** | **string**| The password (encrypted Base64). | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.doc) | 
 **password** | **string**| The password (encrypted Base64). | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be decrypted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutDocumentDisplayProperties**
> DisplayPropertiesResponse PutDocumentDisplayProperties(name, displayProperties, optional)
Update document display properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **displayProperties** | [**DisplayProperties**](DisplayProperties.md)| The display properties. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **displayProperties** | [**DisplayProperties**](DisplayProperties.md)| The display properties. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**DisplayPropertiesResponse**](DisplayPropertiesResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutEncryptDocument**
> AsposeResponse PutEncryptDocument(outPath, userPassword, ownerPassword, cryptoAlgorithm, optional)
Encrypt document from content.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.doc) | 
  **userPassword** | **string**| User password (encrypted Base64). | 
  **ownerPassword** | **string**| Owner password (encrypted Base64). | 
  **cryptoAlgorithm** | **string**| Cryptographic algorithm, see CryptoAlgorithm for details. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.doc) | 
 **userPassword** | **string**| User password (encrypted Base64). | 
 **ownerPassword** | **string**| Owner password (encrypted Base64). | 
 **cryptoAlgorithm** | **string**| Cryptographic algorithm, see CryptoAlgorithm for details. | 
 **permissionsFlags** | [**[]PermissionsFlags**](PermissionsFlags.md)| Array of document permissions, see PermissionsFlags for details. | 
 **usePdf20** | **bool**| Support for revision 6 (Extension 8). | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be encrypted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutEpubInStorageToPdf**
> AsposeResponse PutEpubInStorageToPdf(name, srcPath, optional)
Convert EPUB file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.epub) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.epub) | 
 **storage** | **string**| The document storage. | 
 **dstFolder** | **string**| The destination document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutExportFieldsFromPdfToFdfInStorage**
> AsposeResponse PutExportFieldsFromPdfToFdfInStorage(name, fdfOutputFilePath, optional)
Export fields from from PDF in storage to FDF file in storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fdfOutputFilePath** | **string**| The output Fdf file path. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fdfOutputFilePath** | **string**| The output Fdf file path. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutExportFieldsFromPdfToXfdfInStorage**
> AsposeResponse PutExportFieldsFromPdfToXfdfInStorage(name, xfdfOutputFilePath, optional)
Export fields from from PDF in storage to XFDF file in storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **xfdfOutputFilePath** | **string**| The output xfdf file path. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **xfdfOutputFilePath** | **string**| The output xfdf file path. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutExportFieldsFromPdfToXmlInStorage**
> AsposeResponse PutExportFieldsFromPdfToXmlInStorage(name, xmlOutputFilePath, optional)
Export fields from from PDF in storage to XML file in storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **xmlOutputFilePath** | **string**| The output xml file path. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **xmlOutputFilePath** | **string**| The output xml file path. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFieldsFlatten**
> AsposeResponse PutFieldsFlatten(name, optional)
Flatten form fields in document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFileAttachmentAnnotation**
> FileAttachmentAnnotationResponse PutFileAttachmentAnnotation(name, annotationId, annotation, optional)
Replace document FileAttachment annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**FileAttachmentAnnotation**](FileAttachmentAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**FileAttachmentAnnotation**](FileAttachmentAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**FileAttachmentAnnotationResponse**](FileAttachmentAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFileAttachmentAnnotationDataExtract**
> AsposeResponse PutFileAttachmentAnnotationDataExtract(name, annotationId, optional)
Extract document FileAttachment annotation content to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **outFolder** | **string**| The output folder. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFreeTextAnnotation**
> FreeTextAnnotationResponse PutFreeTextAnnotation(name, annotationId, annotation, optional)
Replace document free text annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**FreeTextAnnotation**](FreeTextAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**FreeTextAnnotation**](FreeTextAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**FreeTextAnnotationResponse**](FreeTextAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutHighlightAnnotation**
> HighlightAnnotationResponse PutHighlightAnnotation(name, annotationId, annotation, optional)
Replace document highlight annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**HighlightAnnotation**](HighlightAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**HighlightAnnotation**](HighlightAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**HighlightAnnotationResponse**](HighlightAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutHtmlInStorageToPdf**
> AsposeResponse PutHtmlInStorageToPdf(name, srcPath, optional)
Convert HTML file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.zip) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.zip) | 
 **htmlFileName** | **string**| Name of HTML file in ZIP. | 
 **height** | **float64**| Page height | 
 **width** | **float64**| Page width | 
 **isLandscape** | **bool**| Is page landscaped | 
 **marginLeft** | **float64**| Page margin left | 
 **marginBottom** | **float64**| Page margin bottom | 
 **marginRight** | **float64**| Page margin right | 
 **marginTop** | **float64**| Page margin top | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImageExtractAsGif**
> AsposeResponse PutImageExtractAsGif(name, imageId, optional)
Extract document image in GIF format to folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **imageId** | **string**| Image ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageId** | **string**| Image ID. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **destFolder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImageExtractAsJpeg**
> AsposeResponse PutImageExtractAsJpeg(name, imageId, optional)
Extract document image in JPEG format to folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **imageId** | **string**| Image ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageId** | **string**| Image ID. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **destFolder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImageExtractAsPng**
> AsposeResponse PutImageExtractAsPng(name, imageId, optional)
Extract document image in PNG format to folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **imageId** | **string**| Image ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageId** | **string**| Image ID. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **destFolder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImageExtractAsTiff**
> AsposeResponse PutImageExtractAsTiff(name, imageId, optional)
Extract document image in TIFF format to folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **imageId** | **string**| Image ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageId** | **string**| Image ID. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **destFolder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImageInStorageToPdf**
> AsposeResponse PutImageInStorageToPdf(name, imageTemplates, optional)
Convert image file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **imageTemplates** | [**ImageTemplatesRequest**](ImageTemplatesRequest.md)| ImageTemplatesRequestImage templates | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageTemplates** | [**ImageTemplatesRequest**](ImageTemplatesRequest.md)| ImageTemplatesRequestImage templates | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImagesExtractAsGif**
> AsposeResponse PutImagesExtractAsGif(name, pageNumber, optional)
Extract document images in GIF format to folder.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **destFolder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImagesExtractAsJpeg**
> AsposeResponse PutImagesExtractAsJpeg(name, pageNumber, optional)
Extract document images in JPEG format to folder.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **destFolder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImagesExtractAsPng**
> AsposeResponse PutImagesExtractAsPng(name, pageNumber, optional)
Extract document images in PNG format to folder.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **destFolder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImagesExtractAsTiff**
> AsposeResponse PutImagesExtractAsTiff(name, pageNumber, optional)
Extract document images in TIFF format to folder.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **destFolder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImportFieldsFromFdfInStorage**
> AsposeResponse PutImportFieldsFromFdfInStorage(name, fdfFilePath, optional)
Update fields from FDF file in storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fdfFilePath** | **string**| The Fdf file path. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fdfFilePath** | **string**| The Fdf file path. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImportFieldsFromXfdfInStorage**
> AsposeResponse PutImportFieldsFromXfdfInStorage(name, xfdfFilePath, optional)
Update fields from XFDF file in storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **xfdfFilePath** | **string**| The XFDF file path. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **xfdfFilePath** | **string**| The XFDF file path. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImportFieldsFromXmlInStorage**
> AsposeResponse PutImportFieldsFromXmlInStorage(name, xmlFilePath, optional)
Update fields from XML file in storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **xmlFilePath** | **string**| Full source filename (ex. /folder1/folder2/template.xml) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **xmlFilePath** | **string**| Full source filename (ex. /folder1/folder2/template.xml) | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutInkAnnotation**
> InkAnnotationResponse PutInkAnnotation(name, annotationId, annotation, optional)
Replace document ink annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**InkAnnotation**](InkAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**InkAnnotation**](InkAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**InkAnnotationResponse**](InkAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutLineAnnotation**
> LineAnnotationResponse PutLineAnnotation(name, annotationId, annotation, optional)
Replace document line annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**LineAnnotation**](LineAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**LineAnnotation**](LineAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**LineAnnotationResponse**](LineAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutLinkAnnotation**
> LinkAnnotationResponse PutLinkAnnotation(name, linkId, link, optional)
Replace document page link annotations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **linkId** | **string**| The link ID. | 
  **link** | [**LinkAnnotation**](LinkAnnotation.md)| Link annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **linkId** | **string**| The link ID. | 
 **link** | [**LinkAnnotation**](LinkAnnotation.md)| Link annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**LinkAnnotationResponse**](LinkAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutListBoxField**
> ListBoxFieldResponse PutListBoxField(name, fieldName, field, optional)
Replace document listbox field

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fieldName** | **string**| The field name. | 
  **field** | [**ListBoxField**](ListBoxField.md)| The field. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fieldName** | **string**| The field name. | 
 **field** | [**ListBoxField**](ListBoxField.md)| The field. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**ListBoxFieldResponse**](ListBoxFieldResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutMarkdownInStorageToPdf**
> AsposeResponse PutMarkdownInStorageToPdf(name, srcPath, optional)
Convert MD file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.md) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.md) | 
 **storage** | **string**| The document storage. | 
 **dstFolder** | **string**| The destination document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutMergeDocuments**
> DocumentResponse PutMergeDocuments(name, mergeDocuments, optional)
Merge a list of documents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Resulting document name. | 
  **mergeDocuments** | [**MergeDocuments**](MergeDocuments.md)| MergeDocuments with a list of documents. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Resulting document name. | 
 **mergeDocuments** | [**MergeDocuments**](MergeDocuments.md)| MergeDocuments with a list of documents. | 
 **storage** | **string**| Resulting document storage. | 
 **folder** | **string**| Resulting document folder. | 

### Return type

[**DocumentResponse**](DocumentResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutMhtInStorageToPdf**
> AsposeResponse PutMhtInStorageToPdf(name, srcPath, optional)
Convert MHT file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.mht) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.mht) | 
 **height** | **float64**| Page height | 
 **width** | **float64**| Page width | 
 **marginLeft** | **float64**| Page margin left | 
 **marginBottom** | **float64**| Page margin bottom | 
 **marginRight** | **float64**| Page margin right | 
 **marginTop** | **float64**| Page margin top | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutMovieAnnotation**
> MovieAnnotationResponse PutMovieAnnotation(name, annotationId, annotation, optional)
Replace document movie annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**MovieAnnotation**](MovieAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**MovieAnnotation**](MovieAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**MovieAnnotationResponse**](MovieAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPageAddStamp**
> AsposeResponse PutPageAddStamp(name, pageNumber, stamp, optional)
Add page stamp.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **stamp** | [**Stamp**](Stamp.md)| Stamp with data. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **stamp** | [**Stamp**](Stamp.md)| Stamp with data. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPageConvertToBmp**
> AsposeResponse PutPageConvertToBmp(name, pageNumber, outPath, optional)
Convert document page to bmp image and upload resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **outPath** | **string**| The out path of result image. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **outPath** | **string**| The out path of result image. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPageConvertToEmf**
> AsposeResponse PutPageConvertToEmf(name, pageNumber, outPath, optional)
Convert document page to emf image and upload resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **outPath** | **string**| The out path of result image. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **outPath** | **string**| The out path of result image. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPageConvertToGif**
> AsposeResponse PutPageConvertToGif(name, pageNumber, outPath, optional)
Convert document page to gif image and upload resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **outPath** | **string**| The out path of result image. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **outPath** | **string**| The out path of result image. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPageConvertToJpeg**
> AsposeResponse PutPageConvertToJpeg(name, pageNumber, outPath, optional)
Convert document page to Jpeg image and upload resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **outPath** | **string**| The out path of result image. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **outPath** | **string**| The out path of result image. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPageConvertToPng**
> AsposeResponse PutPageConvertToPng(name, pageNumber, outPath, optional)
Convert document page to png image and upload resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **outPath** | **string**| The out path of result image. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **outPath** | **string**| The out path of result image. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPageConvertToTiff**
> AsposeResponse PutPageConvertToTiff(name, pageNumber, outPath, optional)
Convert document page to Tiff image and upload resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **outPath** | **string**| The out path of result image. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **outPath** | **string**| The out path of result image. | 
 **width** | **int32**| The converted image width. | [default to 0]
 **height** | **int32**| The converted image height. | [default to 0]
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPclInStorageToPdf**
> AsposeResponse PutPclInStorageToPdf(name, srcPath, optional)
Convert PCL file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.pcl) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.pcl) | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfAInStorageToPdf**
> AsposeResponse PutPdfAInStorageToPdf(name, srcPath, optional)
Convert PDFA file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.pdf) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.pdf) | 
 **dstFolder** | **string**| The destination document folder. | 
 **dontOptimize** | **bool**| If set, document resources will not be optimized. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToAps**
> AsposeResponse PutPdfInRequestToAps(outPath, optional)
Converts PDF document (in request content) to APS format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.aps) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.aps) | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToDoc**
> AsposeResponse PutPdfInRequestToDoc(outPath, optional)
Converts PDF document (in request content) to DOC format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.doc) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.doc) | 
 **addReturnToLineEnd** | **bool**| Add return to line end. | 
 **format** | **string**| Allows to specify .doc or .docx file format. | 
 **imageResolutionX** | **int32**| Image resolution X. | 
 **imageResolutionY** | **int32**| Image resolution Y. | 
 **maxDistanceBetweenTextLines** | **float64**| Max distance between text lines. | 
 **mode** | **string**| Allows to control how a PDF document is converted into a word processing document. | 
 **recognizeBullets** | **bool**| Recognize bullets. | 
 **relativeHorizontalProximity** | **float64**| Relative horizontal proximity. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToEpub**
> AsposeResponse PutPdfInRequestToEpub(outPath, optional)
Converts PDF document (in request content) to EPUB format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.epub) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.epub) | 
 **contentRecognitionMode** | **string**| Property tunes conversion for this or that desirable method of recognition of content. | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToHtml**
> AsposeResponse PutPdfInRequestToHtml(outPath, optional)
Converts PDF document (in request content) to Html format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.html) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.html) | 
 **additionalMarginWidthInPoints** | **int32**| Defines width of margin that will be forcibly left around that output HTML-areas. | 
 **compressSvgGraphicsIfAny** | **bool**| The flag that indicates whether found SVG graphics(if any) will be compressed(zipped) into SVGZ format during saving. | 
 **convertMarkedContentToLayers** | **bool**| If attribute ConvertMarkedContentToLayers set to true then an all elements inside a PDF marked content (layer) will be put into an HTML div with &quot;data-pdflayer&quot; attribute specifying a layer name. This layer name will be extracted from optional properties of PDF marked content. If this attribute is false (by default) then no any layers will be created from PDF marked content. | 
 **defaultFontName** | **string**| Specifies the name of an installed font which is used to substitute any document font that is not embedded and not installed in the system. If null then default substitution font is used. | 
 **documentType** | **string**| Result document type. | 
 **fixedLayout** | **bool**| The value indicating whether that HTML is created as fixed layout. | 
 **imageResolution** | **int32**| Resolution for image rendering. | 
 **minimalLineWidth** | **int32**| This attribute sets minimal width of graphic path line. If thickness of line is less than 1px Adobe Acrobat rounds it to this value. So this attribute can be used to emulate this behavior for HTML browsers. | 
 **preventGlyphsGrouping** | **bool**| This attribute switch on the mode when text glyphs will not be grouped into words and strings This mode allows to keep maximum precision during positioning of glyphs on the page and it can be used for conversion documents with music notes or glyphs that should be placed separately each other. This parameter will be applied to document only when the value of FixedLayout attribute is true. | 
 **splitCssIntoPages** | **bool**| When multipage-mode selected(i.e &#39;SplitIntoPages&#39; is &#39;true&#39;), then this attribute defines whether should be created separate CSS-file for each result HTML page. | 
 **splitIntoPages** | **bool**| The flag that indicates whether each page of source document will be converted into it&#39;s own target HTML document, i.e whether result HTML will be splitted into several HTML-pages. | 
 **useZOrder** | **bool**| If attribute UseZORder set to true, graphics and text are added to resultant HTML document accordingly Z-order in original PDF document. If this attribute is false all graphics is put as single layer which may cause some unnecessary effects for overlapped objects. | 
 **antialiasingProcessing** | **string**| The parameter defines required antialiasing measures during conversion of compound background images from PDF to HTML. | 
 **cssClassNamesPrefix** | **string**| When PDFtoHTML converter generates result CSSs, CSS class names (something like &quot;.stl_01 {}&quot; ... &quot;.stl_NN {}) are generated and used in result CSS. This property allows forcibly set class name prefix. | 
 **explicitListOfSavedPages** | [**[]int32**](int32.md)| With this property You can explicitely define what pages of document should be converted. Pages in this list must have 1-based numbers. I.e. valid numbers of pages must be taken from range (1...[NumberOfPagesInConvertedDocument]) Order of appearing of pages in this list does not affect their order in result HTML page(s) - in result pages allways will go in order in which they are present in source PDF. | 
 **fontEncodingStrategy** | **string**| Defines encoding special rule to tune PDF decoding for current document. | 
 **fontSavingMode** | **string**| Defines font saving mode that will be used during saving of PDF to desirable format. | 
 **htmlMarkupGenerationMode** | **string**| Sometimes specific reqirments to generation of HTML markup are present. This parameter defines HTML preparing modes that can be used during conversion of PDF to HTML to match such specific requirments. | 
 **lettersPositioningMethod** | **string**| The mode of positioning of letters in words in result HTML. | 
 **pagesFlowTypeDependsOnViewersScreenSize** | **bool**| If attribute &#39;SplitOnPages&#x3D;false&#39;, than whole HTML representing all input PDF pages will be put into one big result HTML file. This flag defines whether result HTML will be generated in such way that flow of areas that represent PDF pages in result HTML will depend on screen resolution of viewer. | 
 **partsEmbeddingMode** | **string**| It defines whether referenced files (HTML, Fonts,Images, CSSes) will be embedded into main HTML file or will be generated as apart binary entities. | 
 **rasterImagesSavingMode** | **string**| Converted PDF can contain raster images This parameter defines how they should be handled during conversion of PDF to HTML. | 
 **removeEmptyAreasOnTopAndBottom** | **bool**| Defines whether in created HTML will be removed top and bottom empty area without any content (if any). | 
 **saveShadowedTextsAsTransparentTexts** | **bool**| Pdf can contain texts that are shadowed by another elements (f.e. by images) but can be selected to clipboard in Acrobat Reader (usually it happen when document contains images and OCRed texts extracted from it). This settings tells to converter whether we need save such texts as transparent selectable texts in result HTML to mimic behaviour of Acrobat Reader (othervise such texts are usually saved as hidden, not available for copying to clipboard). | 
 **saveTransparentTexts** | **bool**| Pdf can contain transparent texts that can be selected to clipboard (usually it happen when document contains images and OCRed texts extracted from it). This settings tells to converter whether we need save such texts as transparent selectable texts in result HTML. | 
 **specialFolderForAllImages** | **string**| The path to directory to which must be saved any images if they are encountered during saving of document as HTML. If parameter is empty or null then image files(if any) wil be saved together with other files linked to HTML It does not affect anything if CustomImageSavingStrategy property was successfully used to process relevant image file. | 
 **specialFolderForSvgImages** | **string**| The path to directory to which must be saved only SVG-images if they are encountered during saving of document as HTML. If parameter is empty or null then SVG files(if any) wil be saved together with other image-files (near to output file) or in special folder for images (if it specified in SpecialImagesFolderIfAny option). It does not affect anything if CustomImageSavingStrategy property was successfully used to process relevant image file. | 
 **trySaveTextUnderliningAndStrikeoutingInCss** | **bool**| PDF itself does not contain underlining markers for texts. It emulated with line situated under text. This option allows converter try guess that this or that line is a text&#39;s underlining and put this info into CSS instead of drawing of underlining graphically. | 
 **storage** | **string**| The document storage. | 
 **flowLayoutParagraphFullWidth** | **bool**| This attribute specifies full width paragraph text for Flow mode, FixedLayout &#x3D; false. | 
 **outputFormat** | **string**| This parameter determines the type of result. The zip archive will be uploaded into the storage if the parameter is &quot;Zip&quot; or missed (default). In case of the &quot;Folder&quot;, all resulting files and directories will be uploaded to the folder of the resulting HTML file. | [default to Zip]
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToMobiXml**
> AsposeResponse PutPdfInRequestToMobiXml(outPath, optional)
Converts PDF document (in request content) to MOBIXML format and uploads resulting ZIP archive file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.mobixml) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.mobixml) | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToPdfA**
> AsposeResponse PutPdfInRequestToPdfA(outPath, type_, optional)
Converts PDF document (in request content) to PdfA format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pdf) | 
  **type_** | **string**| Type of PdfA format. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pdf) | 
 **type_** | **string**| Type of PdfA format. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToPptx**
> AsposeResponse PutPdfInRequestToPptx(outPath, optional)
Converts PDF document (in request content) to PPTX format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pptx) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pptx) | 
 **separateImages** | **bool**| Separate images. | 
 **slidesAsImages** | **bool**| Slides as images. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| Base64 encoded password. | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToSvg**
> AsposeResponse PutPdfInRequestToSvg(outPath, optional)
Converts PDF document (in request content) to SVG format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.svg) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.svg) | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToTeX**
> AsposeResponse PutPdfInRequestToTeX(outPath, optional)
Converts PDF document (in request content) to TeX format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.tex) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.tex) | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToTiff**
> AsposeResponse PutPdfInRequestToTiff(outPath, optional)
Converts PDF document (in request content) to TIFF format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.tiff) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.tiff) | 
 **brightness** | **float64**| Image brightness. | 
 **compression** | **string**| Tiff compression. Possible values are: LZW, CCITT4, CCITT3, RLE, None. | 
 **colorDepth** | **string**| Image color depth. Possible valuse are: Default, Format8bpp, Format4bpp, Format1bpp. | 
 **leftMargin** | **int32**| Left image margin. | 
 **rightMargin** | **int32**| Right image margin. | 
 **topMargin** | **int32**| Top image margin. | 
 **bottomMargin** | **int32**| Bottom image margin. | 
 **orientation** | **string**| Image orientation. Possible values are: None, Landscape, Portait. | 
 **skipBlankPages** | **bool**| Skip blank pages flag. | 
 **width** | **int32**| Image width. | 
 **height** | **int32**| Image height. | 
 **xResolution** | **int32**| Horizontal resolution. | 
 **yResolution** | **int32**| Vertical resolution. | 
 **pageIndex** | **int32**| Start page to export. | 
 **pageCount** | **int32**| Number of pages to export. | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToXls**
> AsposeResponse PutPdfInRequestToXls(outPath, optional)
Converts PDF document (in request content) to XLS format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xls) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xls) | 
 **insertBlankColumnAtFirst** | **bool**| Insert blank column at first | 
 **minimizeTheNumberOfWorksheets** | **bool**| Minimize the number of worksheets | 
 **scaleFactor** | **float64**| Scale factor (Obsolete) | 
 **uniformWorksheets** | **bool**| Uniform worksheets | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToXlsx**
> AsposeResponse PutPdfInRequestToXlsx(outPath, optional)
Converts PDF document (in request content) to XLSX format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xlsx) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xlsx) | 
 **insertBlankColumnAtFirst** | **bool**| Insert blank column at first | 
 **minimizeTheNumberOfWorksheets** | **bool**| Minimize the number of worksheets | 
 **scaleFactor** | **float64**| Scale factor (Obsolete) | 
 **uniformWorksheets** | **bool**| Uniform worksheets | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToXml**
> AsposeResponse PutPdfInRequestToXml(outPath, optional)
Converts PDF document (in request content) to XML format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xml) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xml) | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToXps**
> AsposeResponse PutPdfInRequestToXps(outPath, optional)
Converts PDF document (in request content) to XPS format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xps) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xps) | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToAps**
> AsposeResponse PutPdfInStorageToAps(name, outPath, optional)
Converts PDF document (located on storage) to APS format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.aps) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.aps) | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToDoc**
> AsposeResponse PutPdfInStorageToDoc(name, outPath, optional)
Converts PDF document (located on storage) to DOC format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.doc) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.doc) | 
 **addReturnToLineEnd** | **bool**| Add return to line end. | 
 **format** | **string**| Allows to specify .doc or .docx file format. | 
 **imageResolutionX** | **int32**| Image resolution X. | 
 **imageResolutionY** | **int32**| Image resolution Y. | 
 **maxDistanceBetweenTextLines** | **float64**| Max distance between text lines. | 
 **mode** | **string**| Allows to control how a PDF document is converted into a word processing document. | 
 **recognizeBullets** | **bool**| Recognize bullets. | 
 **relativeHorizontalProximity** | **float64**| Relative horizontal proximity. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToEpub**
> AsposeResponse PutPdfInStorageToEpub(name, outPath, optional)
Converts PDF document (located on storage) to EPUB format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.epub) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.epub) | 
 **contentRecognitionMode** | **string**| Property tunes conversion for this or that desirable method of recognition of content. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToHtml**
> AsposeResponse PutPdfInStorageToHtml(name, outPath, optional)
Converts PDF document (located on storage) to Html format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.html) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.html) | 
 **additionalMarginWidthInPoints** | **int32**| Defines width of margin that will be forcibly left around that output HTML-areas. | 
 **compressSvgGraphicsIfAny** | **bool**| The flag that indicates whether found SVG graphics(if any) will be compressed(zipped) into SVGZ format during saving. | 
 **convertMarkedContentToLayers** | **bool**| If attribute ConvertMarkedContentToLayers set to true then an all elements inside a PDF marked content (layer) will be put into an HTML div with &quot;data-pdflayer&quot; attribute specifying a layer name. This layer name will be extracted from optional properties of PDF marked content. If this attribute is false (by default) then no any layers will be created from PDF marked content. | 
 **defaultFontName** | **string**| Specifies the name of an installed font which is used to substitute any document font that is not embedded and not installed in the system. If null then default substitution font is used. | 
 **documentType** | **string**| Result document type. | 
 **fixedLayout** | **bool**| The value indicating whether that HTML is created as fixed layout. | 
 **imageResolution** | **int32**| Resolution for image rendering. | 
 **minimalLineWidth** | **int32**| This attribute sets minimal width of graphic path line. If thickness of line is less than 1px Adobe Acrobat rounds it to this value. So this attribute can be used to emulate this behavior for HTML browsers. | 
 **preventGlyphsGrouping** | **bool**| This attribute switch on the mode when text glyphs will not be grouped into words and strings This mode allows to keep maximum precision during positioning of glyphs on the page and it can be used for conversion documents with music notes or glyphs that should be placed separately each other. This parameter will be applied to document only when the value of FixedLayout attribute is true. | 
 **splitCssIntoPages** | **bool**| When multipage-mode selected(i.e &#39;SplitIntoPages&#39; is &#39;true&#39;), then this attribute defines whether should be created separate CSS-file for each result HTML page. | 
 **splitIntoPages** | **bool**| The flag that indicates whether each page of source document will be converted into it&#39;s own target HTML document, i.e whether result HTML will be splitted into several HTML-pages. | 
 **useZOrder** | **bool**| If attribute UseZORder set to true, graphics and text are added to resultant HTML document accordingly Z-order in original PDF document. If this attribute is false all graphics is put as single layer which may cause some unnecessary effects for overlapped objects. | 
 **antialiasingProcessing** | **string**| The parameter defines required antialiasing measures during conversion of compound background images from PDF to HTML. | 
 **cssClassNamesPrefix** | **string**| When PDFtoHTML converter generates result CSSs, CSS class names (something like &quot;.stl_01 {}&quot; ... &quot;.stl_NN {}) are generated and used in result CSS. This property allows forcibly set class name prefix. | 
 **explicitListOfSavedPages** | [**[]int32**](int32.md)| With this property You can explicitely define what pages of document should be converted. Pages in this list must have 1-based numbers. I.e. valid numbers of pages must be taken from range (1...[NumberOfPagesInConvertedDocument]) Order of appearing of pages in this list does not affect their order in result HTML page(s) - in result pages allways will go in order in which they are present in source PDF. | 
 **fontEncodingStrategy** | **string**| Defines encoding special rule to tune PDF decoding for current document. | 
 **fontSavingMode** | **string**| Defines font saving mode that will be used during saving of PDF to desirable format. | 
 **htmlMarkupGenerationMode** | **string**| Sometimes specific reqirments to generation of HTML markup are present. This parameter defines HTML preparing modes that can be used during conversion of PDF to HTML to match such specific requirments. | 
 **lettersPositioningMethod** | **string**| The mode of positioning of letters in words in result HTML. | 
 **pagesFlowTypeDependsOnViewersScreenSize** | **bool**| If attribute &#39;SplitOnPages&#x3D;false&#39;, than whole HTML representing all input PDF pages will be put into one big result HTML file. This flag defines whether result HTML will be generated in such way that flow of areas that represent PDF pages in result HTML will depend on screen resolution of viewer. | 
 **partsEmbeddingMode** | **string**| It defines whether referenced files (HTML, Fonts,Images, CSSes) will be embedded into main HTML file or will be generated as apart binary entities. | 
 **rasterImagesSavingMode** | **string**| Converted PDF can contain raster images This parameter defines how they should be handled during conversion of PDF to HTML. | 
 **removeEmptyAreasOnTopAndBottom** | **bool**| Defines whether in created HTML will be removed top and bottom empty area without any content (if any). | 
 **saveShadowedTextsAsTransparentTexts** | **bool**| Pdf can contain texts that are shadowed by another elements (f.e. by images) but can be selected to clipboard in Acrobat Reader (usually it happen when document contains images and OCRed texts extracted from it). This settings tells to converter whether we need save such texts as transparent selectable texts in result HTML to mimic behaviour of Acrobat Reader (othervise such texts are usually saved as hidden, not available for copying to clipboard). | 
 **saveTransparentTexts** | **bool**| Pdf can contain transparent texts that can be selected to clipboard (usually it happen when document contains images and OCRed texts extracted from it). This settings tells to converter whether we need save such texts as transparent selectable texts in result HTML. | 
 **specialFolderForAllImages** | **string**| The path to directory to which must be saved any images if they are encountered during saving of document as HTML. If parameter is empty or null then image files(if any) wil be saved together with other files linked to HTML It does not affect anything if CustomImageSavingStrategy property was successfully used to process relevant image file. | 
 **specialFolderForSvgImages** | **string**| The path to directory to which must be saved only SVG-images if they are encountered during saving of document as HTML. If parameter is empty or null then SVG files(if any) wil be saved together with other image-files (near to output file) or in special folder for images (if it specified in SpecialImagesFolderIfAny option). It does not affect anything if CustomImageSavingStrategy property was successfully used to process relevant image file. | 
 **trySaveTextUnderliningAndStrikeoutingInCss** | **bool**| PDF itself does not contain underlining markers for texts. It emulated with line situated under text. This option allows converter try guess that this or that line is a text&#39;s underlining and put this info into CSS instead of drawing of underlining graphically. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **flowLayoutParagraphFullWidth** | **bool**| This attribute specifies full width paragraph text for Flow mode, FixedLayout &#x3D; false. | 
 **outputFormat** | **string**| This parameter determines the type of result. The zip archive will be uploaded into the storage if the parameter is &quot;Zip&quot; or missed (default). In case of the &quot;Folder&quot;, all resulting files and directories will be uploaded to the folder of the resulting HTML file. | [default to Zip]

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToMobiXml**
> AsposeResponse PutPdfInStorageToMobiXml(name, outPath, optional)
Converts PDF document (located on storage) to MOBIXML format and uploads resulting ZIP archive file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.mobixml) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.mobixml) | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToPdfA**
> AsposeResponse PutPdfInStorageToPdfA(name, outPath, type_, optional)
Converts PDF document (located on storage) to PdfA format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pdf) | 
  **type_** | **string**| Type of PdfA format. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pdf) | 
 **type_** | **string**| Type of PdfA format. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToPptx**
> AsposeResponse PutPdfInStorageToPptx(name, outPath, optional)
Converts PDF document (located on storage) to PPTX format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pptx) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pptx) | 
 **separateImages** | **bool**| Separate images. | 
 **slidesAsImages** | **bool**| Slides as images. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| Base64 encoded password. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToSvg**
> AsposeResponse PutPdfInStorageToSvg(name, outPath, optional)
Converts PDF document (located on storage) to SVG format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.svg) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.svg) | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToTeX**
> AsposeResponse PutPdfInStorageToTeX(name, outPath, optional)
Converts PDF document (located on storage) to TeX format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.tex) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.tex) | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToTiff**
> AsposeResponse PutPdfInStorageToTiff(name, outPath, optional)
Converts PDF document (located on storage) to TIFF format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.tiff) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.tiff) | 
 **brightness** | **float64**| Image brightness. | 
 **compression** | **string**| Tiff compression. Possible values are: LZW, CCITT4, CCITT3, RLE, None. | 
 **colorDepth** | **string**| Image color depth. Possible valuse are: Default, Format8bpp, Format4bpp, Format1bpp. | 
 **leftMargin** | **int32**| Left image margin. | 
 **rightMargin** | **int32**| Right image margin. | 
 **topMargin** | **int32**| Top image margin. | 
 **bottomMargin** | **int32**| Bottom image margin. | 
 **orientation** | **string**| Image orientation. Possible values are: None, Landscape, Portait. | 
 **skipBlankPages** | **bool**| Skip blank pages flag. | 
 **width** | **int32**| Image width. | 
 **height** | **int32**| Image height. | 
 **xResolution** | **int32**| Horizontal resolution. | 
 **yResolution** | **int32**| Vertical resolution. | 
 **pageIndex** | **int32**| Start page to export. | 
 **pageCount** | **int32**| Number of pages to export. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToXls**
> AsposeResponse PutPdfInStorageToXls(name, outPath, optional)
Converts PDF document (located on storage) to XLS format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xls) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xls) | 
 **insertBlankColumnAtFirst** | **bool**| Insert blank column at first | 
 **minimizeTheNumberOfWorksheets** | **bool**| Minimize the number of worksheets | 
 **scaleFactor** | **float64**| Scale factor (Obsolete) | 
 **uniformWorksheets** | **bool**| Uniform worksheets | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToXlsx**
> AsposeResponse PutPdfInStorageToXlsx(name, outPath, optional)
Converts PDF document (located on storage) to XLSX format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xlsx) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xlsx) | 
 **insertBlankColumnAtFirst** | **bool**| Insert blank column at first | 
 **minimizeTheNumberOfWorksheets** | **bool**| Minimize the number of worksheets | 
 **scaleFactor** | **float64**| Scale factor (Obsolete) | 
 **uniformWorksheets** | **bool**| Uniform worksheets | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToXml**
> AsposeResponse PutPdfInStorageToXml(name, outPath, optional)
Converts PDF document (located on storage) to XML format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xml) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xml) | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToXps**
> AsposeResponse PutPdfInStorageToXps(name, outPath, optional)
Converts PDF document (located on storage) to XPS format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xps) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xps) | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPolyLineAnnotation**
> PolyLineAnnotationResponse PutPolyLineAnnotation(name, annotationId, annotation, optional)
Replace document polyline annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**PolyLineAnnotation**](PolyLineAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**PolyLineAnnotation**](PolyLineAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PolyLineAnnotationResponse**](PolyLineAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPolygonAnnotation**
> PolygonAnnotationResponse PutPolygonAnnotation(name, annotationId, annotation, optional)
Replace document polygon annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**PolygonAnnotation**](PolygonAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**PolygonAnnotation**](PolygonAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PolygonAnnotationResponse**](PolygonAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPopupAnnotation**
> PopupAnnotationResponse PutPopupAnnotation(name, annotationId, annotation, optional)
Replace document popup annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**PopupAnnotation**](PopupAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**PopupAnnotation**](PopupAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PopupAnnotationResponse**](PopupAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPrivileges**
> AsposeResponse PutPrivileges(name, privileges, optional)
Update privilege document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **privileges** | [**DocumentPrivilege**](DocumentPrivilege.md)| Document privileges. DocumentPrivilege | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **privileges** | [**DocumentPrivilege**](DocumentPrivilege.md)| Document privileges. DocumentPrivilege | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPsInStorageToPdf**
> AsposeResponse PutPsInStorageToPdf(name, srcPath, optional)
Convert PS file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.ps) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.ps) | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutRadioButtonField**
> RadioButtonFieldResponse PutRadioButtonField(name, fieldName, field, optional)
Replace document RadioButton field

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fieldName** | **string**| The field name. | 
  **field** | [**RadioButtonField**](RadioButtonField.md)| The field. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fieldName** | **string**| The field name. | 
 **field** | [**RadioButtonField**](RadioButtonField.md)| The field. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**RadioButtonFieldResponse**](RadioButtonFieldResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutRedactionAnnotation**
> RedactionAnnotationResponse PutRedactionAnnotation(name, annotationId, annotation, optional)
Replace document redaction annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**RedactionAnnotation**](RedactionAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**RedactionAnnotation**](RedactionAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **apply** | **bool**| Apply redaction immediately after adding. | [default to false]

### Return type

[**RedactionAnnotationResponse**](RedactionAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutReplaceImage**
> ImageResponse PutReplaceImage(name, imageId, optional)
Replace document image.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **imageId** | **string**| The image ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageId** | **string**| The image ID. | 
 **imageFilePath** | **string**| Path to image file if specified. Request content is used otherwise. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **image** | ***os.File**| Image file. | 

### Return type

[**ImageResponse**](ImageResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutReplaceMultipleImage**
> ImagesResponse PutReplaceMultipleImage(name, imageIds, optional)
Replace document multiple image.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **imageIds** | [**[]string**](string.md)| The image IDs. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageIds** | [**[]string**](string.md)| The image IDs. | 
 **imageFilePath** | **string**| Path to image file if specified. Request content is used otherwise. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **image** | ***os.File**| Image file. | 

### Return type

[**ImagesResponse**](ImagesResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutScreenAnnotation**
> ScreenAnnotationResponse PutScreenAnnotation(name, annotationId, annotation, optional)
Replace document screen annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**ScreenAnnotation**](ScreenAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**ScreenAnnotation**](ScreenAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**ScreenAnnotationResponse**](ScreenAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutScreenAnnotationDataExtract**
> AsposeResponse PutScreenAnnotationDataExtract(name, annotationId, outFilePath, optional)
Extract document screen annotation content to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **outFilePath** | **string**| The output file path. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **outFilePath** | **string**| The output file path. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutSearchableDocument**
> AsposeResponse PutSearchableDocument(name, optional)
Create searchable PDF document. Generate OCR layer for images in input PDF document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **lang** | **string**| language for OCR engine. Possible values: eng, ara, bel, ben, bul, ces, dan, deu, ell, fin, fra, heb, hin, ind, isl, ita, jpn, kor, nld, nor, pol, por, ron, rus, spa, swe, tha, tur, ukr, vie, chi_sim, chi_tra or thier combination e.g. eng,rus  | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutSetProperty**
> DocumentPropertyResponse PutSetProperty(name, propertyName, value, optional)
Add/update document property.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **propertyName** | **string**| Property name. | 
  **value** | **string**| Property value. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **propertyName** | **string**| Property name. | 
 **value** | **string**| Property value. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **password** | **string**| The password (Base64). | 

### Return type

[**DocumentPropertyResponse**](DocumentPropertyResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutSignatureField**
> SignatureFieldResponse PutSignatureField(name, fieldName, field, optional)
Replace document signature field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fieldName** | **string**| The field name. | 
  **field** | [**SignatureField**](SignatureField.md)| The field. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fieldName** | **string**| The field name. | 
 **field** | [**SignatureField**](SignatureField.md)| The field. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SignatureFieldResponse**](SignatureFieldResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutSoundAnnotation**
> SoundAnnotationResponse PutSoundAnnotation(name, annotationId, annotation, optional)
Replace document sound annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**SoundAnnotation**](SoundAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**SoundAnnotation**](SoundAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SoundAnnotationResponse**](SoundAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutSoundAnnotationDataExtract**
> AsposeResponse PutSoundAnnotationDataExtract(name, annotationId, outFilePath, optional)
Extract document sound annotation content to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **outFilePath** | **string**| The output file path. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **outFilePath** | **string**| The output file path. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutSquareAnnotation**
> SquareAnnotationResponse PutSquareAnnotation(name, annotationId, annotation, optional)
Replace document square annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**SquareAnnotation**](SquareAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**SquareAnnotation**](SquareAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SquareAnnotationResponse**](SquareAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutSquigglyAnnotation**
> SquigglyAnnotationResponse PutSquigglyAnnotation(name, annotationId, annotation, optional)
Replace document squiggly annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**SquigglyAnnotation**](SquigglyAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**SquigglyAnnotation**](SquigglyAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SquigglyAnnotationResponse**](SquigglyAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutStampAnnotation**
> StampAnnotationResponse PutStampAnnotation(name, annotationId, annotation, optional)
Replace document stamp annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**StampAnnotation**](StampAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**StampAnnotation**](StampAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**StampAnnotationResponse**](StampAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutStampAnnotationDataExtract**
> AsposeResponse PutStampAnnotationDataExtract(name, annotationId, outFilePath, optional)
Extract document stamp annotation content to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **outFilePath** | **string**| The output file path. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **outFilePath** | **string**| The output file path. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutStrikeOutAnnotation**
> StrikeOutAnnotationResponse PutStrikeOutAnnotation(name, annotationId, annotation, optional)
Replace document StrikeOut annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**StrikeOutAnnotation**](StrikeOutAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**StrikeOutAnnotation**](StrikeOutAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**StrikeOutAnnotationResponse**](StrikeOutAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutSvgInStorageToPdf**
> AsposeResponse PutSvgInStorageToPdf(name, srcPath, optional)
Convert SVG file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.svg) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.svg) | 
 **adjustPageSize** | **bool**| Adjust page size | 
 **height** | **float64**| Page height | 
 **width** | **float64**| Page width | 
 **isLandscape** | **bool**| Is page landscaped | 
 **marginLeft** | **float64**| Page margin left | 
 **marginBottom** | **float64**| Page margin bottom | 
 **marginRight** | **float64**| Page margin right | 
 **marginTop** | **float64**| Page margin top | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutTable**
> AsposeResponse PutTable(name, tableId, table, optional)
Replace document page table.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **tableId** | **string**| The table ID. | 
  **table** | [**Table**](Table.md)| The table. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **tableId** | **string**| The table ID. | 
 **table** | [**Table**](Table.md)| The table. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutTeXInStorageToPdf**
> AsposeResponse PutTeXInStorageToPdf(name, srcPath, optional)
Convert TeX file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.tex) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.tex) | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutTextAnnotation**
> TextAnnotationResponse PutTextAnnotation(name, annotationId, annotation, optional)
Replace document text annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**TextAnnotation**](TextAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**TextAnnotation**](TextAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**TextAnnotationResponse**](TextAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutTextBoxField**
> TextBoxFieldResponse PutTextBoxField(name, fieldName, field, optional)
Replace document text box field

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fieldName** | **string**| The field name. | 
  **field** | [**TextBoxField**](TextBoxField.md)| The field. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fieldName** | **string**| The field name. | 
 **field** | [**TextBoxField**](TextBoxField.md)| The field. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**TextBoxFieldResponse**](TextBoxFieldResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutUnderlineAnnotation**
> UnderlineAnnotationResponse PutUnderlineAnnotation(name, annotationId, annotation, optional)
Replace document underline annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**UnderlineAnnotation**](UnderlineAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**UnderlineAnnotation**](UnderlineAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**UnderlineAnnotationResponse**](UnderlineAnnotationResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutUpdateField**
> FieldResponse PutUpdateField(name, fieldName, field, optional)
Update field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fieldName** | **string**| The name of a field to be updated. | 
  **field** | [**Field**](Field.md)| Field with the field data. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fieldName** | **string**| The name of a field to be updated. | 
 **field** | [**Field**](Field.md)| Field with the field data. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**FieldResponse**](FieldResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutUpdateFields**
> FieldsResponse PutUpdateFields(name, fields, optional)
Update fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **fields** | [**Fields**](Fields.md)| Fields with the fields data. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fields** | [**Fields**](Fields.md)| Fields with the fields data. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**FieldsResponse**](FieldsResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutWebInStorageToPdf**
> AsposeResponse PutWebInStorageToPdf(name, url, optional)
Convert web page to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **url** | **string**| Source url | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **url** | **string**| Source url | 
 **height** | **float64**| Page height | 
 **width** | **float64**| Page width | 
 **isLandscape** | **bool**| Is page landscaped | 
 **marginLeft** | **float64**| Page margin left | 
 **marginBottom** | **float64**| Page margin bottom | 
 **marginRight** | **float64**| Page margin right | 
 **marginTop** | **float64**| Page margin top | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutXfaPdfInRequestToAcroForm**
> AsposeResponse PutXfaPdfInRequestToAcroForm(outPath, optional)
Converts PDF document which contains XFA form (in request content) to PDF with AcroForm and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pdf) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pdf) | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutXfaPdfInStorageToAcroForm**
> AsposeResponse PutXfaPdfInStorageToAcroForm(name, outPath, optional)
Converts PDF document which contains XFA form (located on storage) to PDF with AcroForm and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pdf) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pdf) | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutXmlInStorageToPdf**
> AsposeResponse PutXmlInStorageToPdf(name, srcPath, optional)
Convert XML file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xml) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xml) | 
 **xslFilePath** | **string**| Full XSL source filename (ex. /folder1/folder2/template.xsl) | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutXpsInStorageToPdf**
> AsposeResponse PutXpsInStorageToPdf(name, srcPath, optional)
Convert XPS file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xps) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xps) | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutXslFoInStorageToPdf**
> AsposeResponse PutXslFoInStorageToPdf(name, srcPath, optional)
Convert XslFo file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xpsfo) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xpsfo) | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StorageExists**
> StorageExist StorageExists(storageName)
Check if storage exists

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storageName** | **string**| Storage name | 

### Return type

[**StorageExist**](StorageExist.md)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadFile**
> FilesUploadResult UploadFile(path, file, optional)
Upload file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| Path where to upload including filename and extension e.g. /file.ext or /Folder 1/file.ext             If the content is multipart and path does not contains the file name it tries to get them from filename parameter             from Content-Disposition header.              | 
  **file** | ***os.File**| File to upload | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| Path where to upload including filename and extension e.g. /file.ext or /Folder 1/file.ext             If the content is multipart and path does not contains the file name it tries to get them from filename parameter             from Content-Disposition header.              | 
 **file** | ***os.File**| File to upload | 
 **storageName** | **string**| Storage name | 

### Return type

[**FilesUploadResult**](FilesUploadResult.md)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

