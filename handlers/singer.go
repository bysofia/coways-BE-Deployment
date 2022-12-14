package handlers

import (
	dto "BackEnd/dto/result"
	singerdto "BackEnd/dto/singer"
	"BackEnd/models"
	"BackEnd/repositories"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerSinger struct {
	SingerRepository repositories.SingerRepository
}

func HandlerSinger(SingerRepository repositories.SingerRepository) *handlerSinger {
	return &handlerSinger{SingerRepository}
}

func (h *handlerSinger) FindSingers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	singers, err := h.SingerRepository.FindSingers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: singers}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerSinger) GetSinger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var singer models.Singer
	singer, err := h.SingerRepository.GetSinger(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: singer}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerSinger) CreateSinger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
	filepath := dataContex.(string)

	startCareer, _ := strconv.Atoi(r.FormValue("start_career"))
	old, _ := strconv.Atoi(r.FormValue("old"))
	request := singerdto.SingerRequest{
		Title:       r.FormValue("name"),
		Old:         old,
		Thumbnail:   filepath,
		Category:    r.FormValue("category"),
		StartCareer: startCareer,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "coways"})

	if err != nil {
		fmt.Println(err.Error())
	}

	singer := models.Singer{
		Name:        request.Title,
		Old:         request.Old,
		Category:    request.Category,
		StartCareer: request.StartCareer,
		Thumbnail:   resp.SecureURL,
	}

	dataSinger, err := h.SingerRepository.CreateSinger(singer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	dataSinger, _ = h.SingerRepository.GetSinger(dataSinger.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: dataSinger}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerSinger) UpdateSinger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
	filepath := dataContex.(string)

	startCareer, _ := strconv.Atoi(r.FormValue("start_career"))
	old, _ := strconv.Atoi(r.FormValue("old"))

	request := singerdto.SingerRequest{
		Title:       r.FormValue("name"),
		Old:         old,
		Thumbnail:   filepath,
		Category:    r.FormValue("category"),
		StartCareer: startCareer,
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "coways"})

	if err != nil {
		fmt.Println(err.Error())
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	singer, _ := h.SingerRepository.GetSinger(id)
	singer.Name = request.Title
	singer.Old = request.Old
	singer.Category = request.Category
	singer.Thumbnail = resp.SecureURL
	singer.StartCareer = request.StartCareer

	if filepath != "false" {
		singer.Thumbnail = resp.SecureURL
	}

	singer, err = h.SingerRepository.UpdateSinger(singer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: singer}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerSinger) DeleteSinger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	singer, err := h.SingerRepository.GetSinger(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.SingerRepository.DeleteSinger(singer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}
