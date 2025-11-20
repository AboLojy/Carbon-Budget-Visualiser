package emission

import (
	"encoding/json"
	"net/http"

	"github.com/AboLojy/Carbon-Budget-Visualiser/types"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store    CityEmissionStore
	validate *validator.Validate
}

func NewHandler(store CityEmissionStore) *Handler {
	return &Handler{store: store, validate: validator.New()}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	// api router will be mounted at /api/v1/
	router.HandleFunc("POST /city", h.createCityEmission)
}

func (h *Handler) createCityEmission(w http.ResponseWriter, r *http.Request) {

	var c types.City
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(&c); err != nil {
		http.Error(w, "validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	created, err := h.store.CreateCityEmission(ctx, &c)
	if err != nil {
		http.Error(w, "failed to create city", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(created)
}
