package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/omarsabri666/url_shorter/helpers"
	"github.com/omarsabri666/url_shorter/model/url"
	service "github.com/omarsabri666/url_shorter/service/url"
	"github.com/omarsabri666/url_shorter/validators"
)

type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(service *service.URLService) *URLHandler {
	return &URLHandler{service: service}
}

func (h *URLHandler) GetURLHttp(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.PathValue("url")
	if shortUrl == "" {
		helpers.WriteJson(w, 400, url.GetUrlResponse{Success: false, Message: "url not provided"})
		// http.Error(w, "short url is required", http.StatusBadRequest)

	}
	// 2️⃣ Call service with request context
	u, err := h.service.GetURL(shortUrl, r.Context())
	if err != nil {
		log.Println(err)

		helpers.WriteJson(w, 500, url.GetUrlResponse{Success: false, Message: err.Error()})
		return
	}

	// 3️⃣ Redirect to long URL
	http.Redirect(w, r, u.LongUrl, http.StatusFound)
}

func (h *URLHandler) CreateURLHttp(w http.ResponseWriter, r *http.Request) {
	var req url.CreateURLRequest
	var res url.CreateUrlResponse

	defer r.Body.Close()
	if r.ContentLength == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		res.Message = "empty body"
		res.Success = false
		json.NewEncoder(w).Encode(res)

		// http.Error(w, "empty body", http.StatusBadRequest)
		return
	}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusBadRequest)
		// res.Message = "invalid json body"
		// res.Success = false
		// json.NewEncoder(w).Encode(res)
		helpers.WriteJson(w, 400, url.CreateUrlResponse{Message: "invalid json body", Success: false})

		// http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}

	// if err := validato.Validate.Struct(req); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return

	// }
	if err := validators.Validate.Struct(req); err != nil {
		helpers.WriteJson(w, 400, url.CreateUrlResponse{Message: err.Error(), Success: false})
		// http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	shortUrl, err := h.service.CreateURL(req, r.Context())
	if err != nil {
		log.Println(err)

		http.Error(w, "failed to create url", http.StatusInternalServerError)
		return
	}
	helpers.WriteJson(w, 201, url.CreateUrlResponse{Data: map[string]string{"url": shortUrl}, Success: true, Message: "Url created successfully"})

}
