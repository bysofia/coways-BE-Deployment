package handlers

import (
	musicdto "BackEnd/dto/music"
	dto "BackEnd/dto/result"
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

type handlerMusic struct {
	MusicRepository repositories.MusicRepository
}

func HandlerMusic(MusicRepository repositories.MusicRepository) *handlerMusic {
	return &handlerMusic{MusicRepository}
}

func (h *handlerMusic) FindMusics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	musics, err := h.MusicRepository.FindMusics()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: musics}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerMusic) GetMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var music models.Music
	music, err := h.MusicRepository.GetMusic(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: music}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerMusic) CreateMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
	dataMusicContext := r.Context().Value("dataMusic")

	filepath := dataContex.(string)
	musicFile := dataMusicContext.(string)

	year, _ := strconv.Atoi(r.FormValue("year"))
	singer_id, _ := strconv.Atoi(r.FormValue("singer_id"))
	request := musicdto.MusicRequest{
		Title:     r.FormValue("title"),
		Year:      year,
		Thumbnail: filepath,
		SingerID:  singer_id,
		MusicFile: musicFile,
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

	respMusic, err := cld.Upload.Upload(ctx, musicFile, uploader.UploadParams{Folder: "coways"})

	if err != nil {
		fmt.Println(err.Error())
	}

	music := models.Music{
		Title:     request.Title,
		Year:      request.Year,
		Thumbnail: resp.SecureURL,
		SingerID:  request.SingerID,
		MusicFile: respMusic.SecureURL,
	}

	dataMusic, err := h.MusicRepository.CreateMusic(music)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	dataMusic, _ = h.MusicRepository.GetMusic(dataMusic.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: dataMusic}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerMusic) UpdateMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
	dataMusicContext := r.Context().Value("dataMusic")

	filepath := dataContex.(string)
	musicFile := dataMusicContext.(string)

	year, _ := strconv.Atoi(r.FormValue("year"))
	singer_id, _ := strconv.Atoi(r.FormValue("singer_id"))

	request := musicdto.MusicRequest{
		Title:     r.FormValue("title"),
		Year:      year,
		SingerID:  singer_id,
		Thumbnail: filepath,
		MusicFile: musicFile,
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

	respMusic, err := cld.Upload.Upload(ctx, musicFile, uploader.UploadParams{Folder: "coways"})

	if err != nil {
		fmt.Println(err.Error())
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	music, _ := h.MusicRepository.GetMusic(id)
	music.SingerID = singer_id
	music.Title = request.Title
	music.Year = request.Year
	music.Thumbnail = respMusic.SecureURL

	if filepath != "false" && musicFile != "false" {
		music.Thumbnail = resp.SecureURL
		music.MusicFile = respMusic.SecureURL
	}

	music, err = h.MusicRepository.UpdateMusic(music)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: music}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerMusic) DeleteMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	music, err := h.MusicRepository.GetMusic(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.MusicRepository.DeleteMusic(music)
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
