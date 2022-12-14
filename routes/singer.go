package routes

import (
	"BackEnd/handlers"
	"BackEnd/pkg/connection"
	"BackEnd/pkg/middleware"
	"BackEnd/repositories"

	"github.com/gorilla/mux"
)

func SingerRoutes(r *mux.Router) {
	singerRepository := repositories.RepositorySinger(connection.DB)
	h := handlers.HandlerSinger(singerRepository)

	r.HandleFunc("/singers", h.FindSingers).Methods("GET")
	r.HandleFunc("/singer/{id}", middleware.Auth(h.GetSinger)).Methods("GET")
	r.HandleFunc("/singer", middleware.Auth(middleware.UploadFile(h.CreateSinger))).Methods("POST")
	r.HandleFunc("/singer/{id}", middleware.Auth(middleware.UploadFile(h.UpdateSinger))).Methods("PATCH")
	r.HandleFunc("/singer/{id}", middleware.Auth(h.DeleteSinger)).Methods("DELETE")
}
