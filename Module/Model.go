package Module

import (
	"gorm.io/gorm"
	"time"
)

type Candidates struct {
	Id          int64          `json:"id" form:"id"`
	Name        string         `json:"name" form:"name"`
	PhoneNumber string         `json:"phone_number" form:"phone_number"`
	Mail        string         `json:"mail" form:"mail"`
	FaceBook    string         `json:"facebook" form:"facebook"`
	LinkedIn    string         `json:"linked_in" form:"linked_id"`
	CreatedAt   time.Time      `gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"type:TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"type:TIMESTAMP NULL DEFAULT NULL" json:"deleted_at"`
}

type CandidateDetail struct {
	Id          int64  `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Mail        string `json:"mail" form:"mail"`
	FaceBook    string `json:"facebook" form:"facebook"`
	LinkedIn    string `json:"linked_in" form:"linked_id"`
}

func (c Candidates) ConvertToDetail() CandidateDetail {
	return CandidateDetail{
		Id:          c.Id,
		Name:        c.Name,
		PhoneNumber: c.PhoneNumber,
		Mail:        c.Mail,
		FaceBook:    c.FaceBook,
		LinkedIn:    c.LinkedIn,
	}
}

func (Candidates) TableName() string { return "candidates" }
