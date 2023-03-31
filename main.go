package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

func init() {
	err := license.SetMeteredKey(`b0de89db72db36b3d8d4521ff734b62b254e36b9953fb14da3f98aa014300c17`)
	if err != nil {
		fmt.Printf("ERROR: Failed to set metered key: %v\n", err)
		fmt.Printf("Make sure to get a valid key from https://cloud.unidoc.io\n")
		panic(err)
	}
}

type Candidates struct {
	Id          int64  `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Email       string `json:"email" form:"email"`
}

func getFileReader(inputPath string) (string, error) {
	var content string
	file, err := os.Open(inputPath)

	if err != nil {
		log.Fatal("Error: file was not opened!")
		return "", err
	}

	defer file.Close()

	pdfReader, error := model.NewPdfReader(file)
	if err != nil {
		log.Fatal("Error: file was not opened!")
		return "", error
	}
	numPages, error := pdfReader.GetNumPages()
	if error != nil {
		fmt.Printf("Did not get the number of pages!")
	}

	fmt.Printf("This PDF has %v number of page(s)\n", numPages)

	for i := 0; i < numPages; i++ {
		pageNum := i + 1
		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			log.Fatal("Error %v", err)
		}
		textExtractor, err := extractor.New(page)
		if err != nil {
			log.Fatal("Error %v", err)
		}

		text, err := textExtractor.ExtractText()
		if err != nil {
			log.Fatal("Error %v", err)
		}
		content = fmt.Sprintf("%s\n%s", content, text)
	}
	return content, nil

}
func main() {

	files, _ := ioutil.ReadDir("CV")
	for _, file := range files {
		filepath := fmt.Sprintf("CV/%s", file.Name())
		content, err := getFileReader(filepath)
		if err != nil {
			log.Fatal("Error: %v!", err)
		}
		lines := strings.Split(content, "\n")
		for i, line := range lines {
			fmt.Print(i, " ")
			fmt.Println(line)
		}
		fmt.Println("//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////")
	}
}
