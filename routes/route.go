package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"skyshi/controllers"
	"skyshi/models"

	"github.com/gorilla/mux"
)

func Mux(skyshiService controllers.SkyshiService) {
	r := mux.NewRouter()
	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(rw, r)
		})
	})

	//activity
	r.HandleFunc("/activity-groups", skyshiService.GetAllActivityGroup).Methods("GET")
	r.HandleFunc("/activity-groups/{id}", skyshiService.GetOneActivityGroup).Methods("GET")
	r.HandleFunc("/activity-groups", skyshiService.CreateActivityGroup).Methods("POST")
	r.HandleFunc("/activity-groups/{id}", skyshiService.DeleteActivityGroup).Methods("DELETE")
	r.HandleFunc("/activity-groups/{id}", skyshiService.UpdateActivityGroup).Methods("PATCH")

	//TODO
	r.HandleFunc("/todo-items", skyshiService.GetAllTodoItems).Methods("GET")
	r.HandleFunc("/todo-items/{id}", skyshiService.GetOneTodoItems).Methods("GET")
	r.HandleFunc("/todo-items", skyshiService.CreateTodoItems).Methods("POST")
	r.HandleFunc("/todo-items/{id}", skyshiService.DeleteTodoItems).Methods("DELETE")
	r.HandleFunc("/todo-items/{id}", skyshiService.UpdateTodoItems).Methods("PATCH")

	r.Use(mux.CORSMethodMiddleware(r))

	r.NotFoundHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		json.NewEncoder(response).Encode(models.Response{
			Message: "route not found",
			Data:    nil,
		})
	})

	r.MethodNotAllowedHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		json.NewEncoder(response).Encode(models.Response{
			Message: "method not allowed",
			Data:    nil,
		})

	})

	appPort := os.Getenv("APPLICATION_PORT")

	log.Println("Running at " + os.Getenv("APP_URL") + ":" + appPort + "/")

	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", appPort), r)

}
