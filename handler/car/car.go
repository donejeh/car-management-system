package car

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/donejeh/car-management-system/service"
)

type CarHandler struct {
	service service.CarServiceInterface
}

func NewCarHandler(service service.CarServiceInterface) *CarHandler {
	return &CarHandler{
		service: service,
	}
}

func (h *CarHandler) GetCarById(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	vars := mux.Vars(r)
	id := vars["id"]

	res, err := h.service.GetCarById(ctx, id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error getting car by id: ", err)
		return
	}

	body, err := json.Marshal(res)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error marshalling car by id: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(body)

	if err != nil {
		log.Println("Error writing response: ", err)
	}

}

func (h *CarHandler) GetCarByBrand(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	brand := r.URL.Query().Get("brand")
	isEngine := r.URL.Query().Get("isEngine") == "true"

	res, err := h.service.GetCarByBrand(ctx, brand, isEngine)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error getting car by brand: ", err)
		return
	}

	body, err := json.Marshal(res)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error marshalling car by brand: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(body)

	if err != nil {
		log.Println("Error writing response: ", err)
	}

}
