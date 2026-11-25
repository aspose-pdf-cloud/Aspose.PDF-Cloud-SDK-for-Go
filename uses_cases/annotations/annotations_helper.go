package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

const (
	LOCAL_FOLDER      = "test_data"
	REMOTE_FOLDER     = "Your_Temp_Pdf_Cloud"
	PDF_DOCUMENT_NAME = "PdfWithAnnotations.pdf"
	PDF_OUTPUT        = "output_sample.pdf"
	PAGE_NUMBER       = 2

	ANNOTATION_ID = "GI5TAOZRGU3CYNZSGEWDCNZWFQ3TGOI"

	NEW_HL_ANNOTATION_TEXT        = "NEW HIGHLIGHT TEXT ANNOTATION"
	NEW_HL_ANNOTATION_DESCRIPTION = "This is a sample highlight annotation"
	NEW_HL_ANNOTATION_SUBJECT     = "Highlight Text Box Subject"
	NEW_HL_ANNOTATION_CONTENTS    = "Highlight annotation sample contents"

	NEW_SO_ANNOTATION_TEXT        = "NEW STRIKEOUT TEXT ANNOTATION"
	NEW_SO_ANNOTATION_DESCRIPTION = "This is a sample strikeout annotation"
	NEW_SO_ANNOTATION_SUBJECT     = "Strikeout Text Box Subject"
	NEW_SO_ANNOTATION_CONTENTS    = "Strikeout annotation sample contents"

	NEW_UL_ANNOTATION_TEXT        = "NEW UNDERLINE TEXT ANNOTATION"
	NEW_UL_ANNOTATION_DESCRIPTION = "This is a sample underline annotation"
	NEW_UL_ANNOTATION_SUBJECT     = "Underline Text Box Subject"
	NEW_UL_ANNOTATION_CONTENTS    = "Underline annotation sample contents"

	NEW_FT_ANNOTATION_TEXT        = "NEW FREE TEXT ANNOTATION"
	NEW_FT_ANNOTATION_DESCRIPTION = "This is a sample annotation"
	NEW_FT_ANNOTATION_SUBJECT     = "Free Text Box Subject"
	NEW_FT_ANNOTATION_CONTENTS    = "Free Text annotation sample contents"

	REPLACED_CONTENT = "This is a replaced sample annotation"
)

func InitPdfApi() *asposepdfcloud.PdfApiService {
	// Initialize Credentials and create Pdf.Cloud service object
	ClientId := "*********"     // Your Application Client ID
	ClientSecret := "*********" // Your Application Client Secret

	pdfApi := asposepdfcloud.NewPdfApiService(ClientId, ClientSecret, "")
	return pdfApi
}

func UploadFile(pdf_api *asposepdfcloud.PdfApiService, name string) {
	// Upload local file to the Pdf.Cloud folder
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	file, _ := os.Open(filepath.Join(LOCAL_FOLDER, name))
	_, httpResponse, err := pdf_api.UploadFile(filepath.Join(REMOTE_FOLDER, name), file, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println("Result file'" + name + "' successfully uploaded!")
	}
}

func SaveByteArrayToFile(local_folder string, file_name string, data []byte) {
	// Save byte array data to the local folder
	fileName := path.Join(local_folder, file_name)
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		size, err := f.Write(data)
		if err != nil || size == 0 {
			fmt.Println("Failures in downloading result!")
		} else {
			fmt.Println("Result file'" + fileName + "' successfully downloaded!")
		}
	}
}

func DownloadFile(pdf_api *asposepdfcloud.PdfApiService, name string, output_name string, prefix string) {
	// Download modified Pdf document to local folder from the Pdf.Cloud folder
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	result_data, _, _ := pdf_api.DownloadFile(path.Join(REMOTE_FOLDER, name), args)
	SaveByteArrayToFile(LOCAL_FOLDER, prefix+output_name, result_data)
}

func DeletePopupAnnotations(pdf_api *asposepdfcloud.PdfApiService, document_name string, parent_annotation string) {
	// Delete popup annotations for typed parent annotation in the page in the PDF document.
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	result, httpResponse, err := pdf_api.GetDocumentPopupAnnotations(document_name, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		for _, annotation := range result.Annotations.List {
			if annotation.Parent.Id == parent_annotation {
				_, response, err2 := pdf_api.DeleteAnnotation(document_name, annotation.Id, args)
				if err2 != nil {
					fmt.Println(err2)
				} else if response.StatusCode < 200 || response.StatusCode > 299 {
					fmt.Println("delete_popup_annotations(): Failed to delete popup annotation in the document.")
				} else {
					fmt.Println("delete_popup_annotations(): popup annotation id = '" + annotation.Id + "' for '" + annotation.Contents + "' deleted in the document '" + PDF_DOCUMENT_NAME + "'.")
				}
			}
		}
	}
}
