package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"skyshi/models"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type SkyshiService struct {
	DB *gorm.DB
}

func (s *SkyshiService) GetAllActivityGroup(w http.ResponseWriter, r *http.Request) {

	var activity models.Activity
	result, err := activity.GetAllActivity(s.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "success",
		Data:    result,
	})
}

func (s *SkyshiService) GetOneActivityGroup(w http.ResponseWriter, r *http.Request) {
	var activity models.Activity

	activity.Id, _ = strconv.Atoi(mux.Vars(r)["id"])

	result, err := activity.GetOneActivity(s.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "success",
		Data:    result,
	})
}

func (s *SkyshiService) CreateActivityGroup(w http.ResponseWriter, r *http.Request) {

	var body models.Activity
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Status:  "failed",
			Message: "Get Body Failed",
		})
		return
	}

	body.CreatedAt = time.Now()

	fmt.Println(body)

	result, err := body.CreateActivity(s.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "success",
		Data:    result,
	})
}
func (s *SkyshiService) DeleteActivityGroup(w http.ResponseWriter, r *http.Request) {

	var activity models.Activity

	activity.Id, _ = strconv.Atoi(mux.Vars(r)["id"])

	err := activity.DeleteActivity(s.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "success",
	})

}
func (s *SkyshiService) UpdateActivityGroup(w http.ResponseWriter, r *http.Request) {

	var body models.Activity
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Status:  "failed",
			Message: "Get Body Failed",
		})
		return
	}

	body.Id, _ = strconv.Atoi(mux.Vars(r)["id"])

	body.UpdatedAt = time.Now()

	result, err := body.UpdateActivity(s.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "success",
		Data:    result,
	})

}
