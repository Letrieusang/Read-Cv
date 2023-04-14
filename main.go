package main

import (
	"awesomeProject4/Module"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	err := license.SetMeteredKey(`a9753e981af16db0abd1c4bf037227da1e8396e5a0f03f664c40b540255f92d4`)
	if err != nil {
		fmt.Printf("ERROR: Failed to set metered key: %v\n", err)
		fmt.Printf("Make sure to get a valid key from `https://cloud.unidoc.io`\n")
		panic(err)
	}
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
	surname := map[string]bool{"Lê": true, "Nguyễn": true, "Trần": true, "Phạm": true, "Hoàng": true, "Huỳnh": true, "Phan": true, "Vũ": true, "Võ": true, "Đặng": true, "Bùi": true, "Đỗ": true, "Hồ": true, "Ngô": true, "Dương": true, "Lý": true, "Nguyen": true, "NGUYỄN": true, "NGUYEN": true, "LUONG": true, "Tran": true, "NGUYÊN": true, "LÊ": true, "LE": true, "PHAM": true, "PHẠM": true, "TRAN": true, "TRẦN": true, "HOÀNG": true, "HOANG": true, "HUỲNH": true, "HUYNH": true}
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
			info["linked_in"] = lines[i]
		}
		if match, _ := regexp.MatchString("([\\+84|84|0]+(3|5|7|8|9|1[2|6|8|9]))+([0-9]{8})\\b", lines[i]); match == true {
			info["phone_number"] = lines[i]
		}
	}
	return info
}

func main() {
	files, _ := ioutil.ReadDir("CV")
	candidates := []Module.Candidates{}
	for _, file := range files {
		var candidate Module.Candidates
		filepath := fmt.Sprintf("CV/%s", file.Name())
		content, err := getFileReader(filepath)
		if err != nil {
			log.Fatal("Error: %v!", err)
		}
		/*
				thường thì các cv sẽ có phần họ tên, thông tin nằm ở gần đầu. cụ thể là từ dòng 10 đổ lại,
				nhưng với các cv chia thành hai cột trái phải thì hiện đang không rõ vị trí của chúng trong chuỗi kí tự,
				không có sự phân biệt rạch ròi giữa hai cột, các cột vẫn được đọc chung từ trên xuống dưới, trái sang phải như thường
			cần tìm cj đó để bám víu vào mà xác định mẫu, thứ gì đấy dễ xác định, bất biến, quy chuẩn chung
		*/
		candidate.Name = FindName(content, 0, 10)
		info := FindInfo(content, 0, 20)
		candidate.PhoneNumber = info["phone_number"]
		candidate.FaceBook = info["facebook"]
		candidate.Mail = info["mail"]
		candidate.LinkedIn = info["linked_in"]
		candidates = append(candidates, candidate)
		log.Println(content)
	}
	for _, candi := range candidates {
		log.Println(candi.Name)
		log.Println(candi.PhoneNumber)
		log.Println(candi.Id)
		log.Println(candi.Mail)
		log.Println(candi.FaceBook)
		log.Println(candi.LinkedIn)
		log.Println("/////////////////////////////////////////")
	}
	err := InsertDB(candidates)
	if err != nil {
		log.Println(err)
	}
}

func InsertDB(candidate []Module.Candidates) error {
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("Config/")
	viper.ReadInConfig()
	sqlPath := viper.GetString("sql_path")
	dsn := sqlPath
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	errCr := db.Table(Module.Candidates{}.TableName()).Create(&candidate).Error
	if errCr != nil {
		return errCr
	}
	return nil
}
