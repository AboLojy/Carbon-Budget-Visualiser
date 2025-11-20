package annualemission

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/AboLojy/Carbon-Budget-Visualiser/types"
	"github.com/go-playground/validator/v10"
)

// Handler provides HTTP handlers for annual emissions.
type Handler struct {
	store    AnnualEmissionStore
	validate *validator.Validate
}

// NewHandler creates a new Handler with the provided store.
func NewHandler(store AnnualEmissionStore) *Handler {
	return &Handler{
		store:    store,
		validate: validator.New(),
	}
}

// RegisterRoutes registers annual emission routes on the provided router.
func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/annualemissions", h.handleAnnualEmissions)
}

func (h *Handler) handleAnnualEmissions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createAnnualEmission(w, r)
	case http.MethodGet:
		h.getAnnualEmissions(w, r)
	case http.MethodDelete:
		h.deleteAnnualEmission(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) createAnnualEmission(w http.ResponseWriter, r *http.Request) {
	var ae types.AnnualEmission
	if err := json.NewDecoder(r.Body).Decode(&ae); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(&ae); err != nil {
		http.Error(w, "validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	created, err := h.store.CreateAnnualEmission(r.Context(), &ae)
	if err != nil {
		http.Error(w, "failed to create annual emission", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(created)
}

func (h *Handler) getAnnualEmissions(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.GetAll(r.Context())
	if err != nil {
		http.Error(w, "failed to fetch annual emissions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(list)
}

func (h *Handler) deleteAnnualEmission(w http.ResponseWriter, r *http.Request) {
	// Extract emission ID from query string: DELETE /annualemissions?id=123
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		// Try path-based: DELETE /annualemissions/123 (if router supports it)
		parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/annualemissions"), "/")
		if len(parts) > 1 && parts[1] != "" {
			idStr = parts[1]
		}
	}

	if idStr == "" {
		http.Error(w, "emission_id required (use ?id=123 or /annualemissions/123)", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid emission_id", http.StatusBadRequest)
		return
	}

	if err := h.store.DeleteByID(r.Context(), id); err != nil {
		http.Error(w, "failed to delete annual emission", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
