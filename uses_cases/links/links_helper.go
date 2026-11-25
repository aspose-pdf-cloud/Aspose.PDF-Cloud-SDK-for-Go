package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26"
)

const (
	REMOTE_FOLDER      = "Your_Temp_Pdf_Cloud"
	LOCAL_FOLDER       = "test_data"
	PDF_DOCUMENT       = "PdfWithLinks.pdf"
	PDF_OUTPUT         = "output_sample_add_link.pdf"
	PDF_OUTPUT_REMOVE  = "output_sample_rem_link.pdf"
	PDF_OUTPUT_REPLACE = "output_sample_rep_link.pdf"
	NEW_LINK_ACTION    = "https://reference.aspose.cloud/pdf/"
	PAGE_NUMBER        = 1
	LINK_ID            = "GE5UO32UN5AWG5DJN5XDWOBYFQ3TGMJMGEZTMLBXGQ3A"
)

var (
	LINK_RECT = asposepdfcloud.Rectangle{LLX: 70, LLY: 731.229, URX: 90, URY: 745.185}
)

func initPdfApi() *asposepdfcloud.PdfApiService {
	ClientId := "******"
	ClientSecret := "******"

	pdfApi := asposepdfcloud.NewPdfApiService(ClientId, ClientSecret, "")
	return pdfApi
}

// Upload local file to the remote folder with check errors
func uploadFile(pdf_api *asposepdfcloud.PdfApiService, name string) {
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	file, err := os.Open(filepath.Join(LOCAL_FOLDER, name))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		result, httpResponse, err := pdf_api.UploadFile(path.Join(REMOTE_FOLDER, name), file, args)
		if err != nil {
			fmt.Println(err.Error())
		} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
			fmt.Println("Unexpected error!")
		} else {
			fmt.Println(result)
		}
	}
}

// Download file from remote folder and save it locally with check errors
func downloadFile(pdf_api *asposepdfcloud.PdfApiService, name string, output_name string) {
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	result_data, httpResponse, err := pdf_api.DownloadFile(path.Join(REMOTE_FOLDER, name), args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fileName := path.Join(LOCAL_FOLDER, output_name)
		f, _ := os.Create(fileName)
		_, _ = f.Write(result_data)
		fmt.Println("File '" + fileName + "'successfully downloaded.")
	}
}

func showLinks(links *[]asposepdfcloud.LinkAnnotation) {
	for i := 0; i < len(*links); i++ {
		fmt.Print("Link #")
		fmt.Println(i)
		fmt.Print("\tId == '")
		fmt.Print((*links)[i].Id)
		fmt.Println("'")
		fmt.Print("\tActionType == '")
		fmt.Print((*links)[i].ActionType)
		fmt.Println("'")
		fmt.Print("\tAction == '")
		fmt.Print((*links)[i].Action)
		fmt.Println("'")
	}
}
