package routes

import (
	"BackEnd/handlers"
	"BackEnd/pkg/connection"
	"BackEnd/pkg/middleware"
	"BackEnd/repositories"

	"github.com/gorilla/mux"
)

func MusicRoutes(r *mux.Router) {
	musicRepository := repositories.RepositoryMusic(connection.DB)
	h := handlers.HandlerMusic(musicRepository)

	r.HandleFunc("/musics", h.FindMusics).Methods("GET")
	r.HandleFunc("/music/{id}", middleware.Auth(h.GetMusic)).Methods("GET")
	r.HandleFunc("/music", middleware.Auth(middleware.UploadFile(middleware.UploadMusic(h.CreateMusic)))).Methods("POST")
	r.HandleFunc("/music/{id}", middleware.Auth(middleware.UploadFile(h.UpdateMusic))).Methods("PATCH")
	r.HandleFunc("/music/{id}", middleware.Auth(h.DeleteMusic)).Methods("DELETE")
}
