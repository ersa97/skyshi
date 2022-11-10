package controllers

import (
	"net/http"

	"gorm.io/gorm"
)

type SkyshiService struct {
	DB *gorm.DB
}

func (s *SkyshiService) GetAllActivityGroup(w http.ResponseWriter, r *http.Request) {

}

func (s *SkyshiService) GetOneActivityGroup(w http.ResponseWriter, r *http.Request) {

}

func (s *SkyshiService) CreateActivityGroup(w http.ResponseWriter, r *http.Request) {

}
func (s *SkyshiService) DeleteActivityGroup(w http.ResponseWriter, r *http.Request) {

}
func (s *SkyshiService) UpdateActivityGroup(w http.ResponseWriter, r *http.Request) {

}
