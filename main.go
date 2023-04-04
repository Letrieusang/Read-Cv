package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

func init() {
	err := license.SetMeteredKey(`33792f026d20c3088ae83fd0f0ad67532761ce6cb5b5641440746ed449242917`)
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
	Mail        string `json:"mail" form:"mail"`
	FaceBook    string `json:"facebook" form:"facebook"`
	LinkedIn    string `json:"linked_in" form:"linked_id"`
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
func FindName(content string, firstline int, lastline int) string {
	surname := map[string]bool{"Lê": true, "Nguyễn": true, "Trần": true, "Phạm": true, "Hoàng": true, "Huỳnh": true, "Phan": true, "Vũ": true, "Võ": true, "Đặng": true, "Bùi": true, "Đỗ": true, "Hồ": true, "Ngô": true, "Dương": true, "Lý": true, "Nguyen": true, "NGUYỄN": true, "NGUYEN": true, "LUONG": true, "Tran": true, "NGUYÊN": true, "LÊ": true, "LE": true}
	lines := strings.Split(content, "\n")
	for i := firstline; i < lastline; i++ {
		word := strings.Split(lines[i], " ")
		if surname[word[0]] == true {
			return lines[i]
		}
	}
	return ""
}

func FindInfo(content string, firstline int, lastline int) map[string]string {
	info := map[string]string{}
	lines := strings.Split(content, "\n")
	for i := firstline; i < lastline; i++ {
		if strings.Contains(lines[i], "@gmail") == true {
			info["mail"] = lines[i]
		}
		if strings.Contains(lines[i], "facebook.com") == true {
			info["facebook"] = lines[i]
		}
		if strings.Contains(lines[i], "linkedin") == true {
			info["linkedin"] = lines[i]
		}
		if match, _ := regexp.MatchString("([\\+84|84|0]+(3|5|7|8|9|1[2|6|8|9]))+([0-9]{8})\\b", lines[i]); match == true {
			info["phone_number"] = lines[i]
		}
	}
	return info
}

func main() {
	files, _ := ioutil.ReadDir("CV")
	candidates := []Candidates{}
	for _, file := range files {
		var candidate Candidates
		filepath := fmt.Sprintf("CV/%s", file.Name())
		content, err := getFileReader(filepath)
		if err != nil {
			log.Fatal("Error: %v!", err)
		}
		if strings.Contains(content, "topcv") == true {
			candidate.Name = FindName(content, 0, 10)
			info := FindInfo(content, 0, 100)
			candidate.PhoneNumber = info["phone_number"]
			candidate.FaceBook = info["facebook"]
			candidate.Mail = info["mail"]
			candidate.LinkedIn = info["linked_in"]
		} else if strings.Contains(content, "viec365") == true {
			candidate.Name = FindName(content, 0, 10)
			info := FindInfo(content, 0, 100)
			candidate.PhoneNumber = info["phone_number"]
			candidate.FaceBook = info["facebook"]
			candidate.Mail = info["mail"]
			candidate.LinkedIn = info["linked_in"]
		} else {
			candidate.Name = FindName(content, 0, 10)
			info := FindInfo(content, 0, 100)
			candidate.PhoneNumber = info["phone_number"]
			candidate.FaceBook = info["facebook"]
			candidate.Mail = info["mail"]
			candidate.LinkedIn = info["linked_in"]
		}
		candidates = append(candidates, candidate)
		log.Println(content)
	}
	for _, candi := range candidates {
		log.Println(candi.Name)
		log.Println(candi.PhoneNumber)
		log.Println(candi.Id)
	}

}
