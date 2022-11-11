package controllers

import (
	"encoding/json"
	"net/http"
	"skyshi/models"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (s *SkyshiService) GetAllTodoItems(w http.ResponseWriter, r *http.Request) {

	var todo models.Todo
	result, err := todo.GetAllTodo(s.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Status:  "success",
		Message: "success",
		Data:    result,
	})
}

func (s *SkyshiService) GetOneTodoItems(w http.ResponseWriter, r *http.Request) {

	var todo models.Todo

	todo.Id, _ = strconv.Atoi(mux.Vars(r)["id"])

	result, err := todo.GetOneTodo(s.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Status:  "success",
		Message: "success",
		Data:    result,
	})
}

func (s *SkyshiService) CreateTodoItems(w http.ResponseWriter, r *http.Request) {

	var body models.Todo
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "Get Body Failed",
		})
		return
	}

	body.CreatedAt = time.Now()
	body.IsActive = 1

	result, err := body.CreateTodo(s.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Status:  "success",
		Message: "success",
		Data:    result,
	})
}
func (s *SkyshiService) DeleteTodoItems(w http.ResponseWriter, r *http.Request) {

	var todo models.Todo

	todo.Id, _ = strconv.Atoi(mux.Vars(r)["id"])

	err := todo.DeleteTodo(s.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Status:  "success",
		Message: "success",
	})

}
func (s *SkyshiService) UpdateTodoItems(w http.ResponseWriter, r *http.Request) {

	var body models.Todo
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "Get Body Failed",
		})
		return
	}

	body.Id, _ = strconv.Atoi(mux.Vars(r)["id"])

	body.UpdatedAt = time.Now()

	result, err := body.UpdateTodo(s.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Status:  "success",
		Message: "success",
		Data:    result,
	})

}
