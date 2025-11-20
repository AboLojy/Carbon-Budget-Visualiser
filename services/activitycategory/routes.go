package activitycategory

import (
	"encoding/json"
	"net/http"
)

// Handler provides HTTP handlers for activity categories.
type Handler struct {
	store ActivityStore
}

// NewHandler creates a new Handler with the provided store.
func NewHandler(store ActivityStore) *Handler {
	return &Handler{store: store}
}

// RegisterRoutes registers activity category routes on the provided router.
// This registers GET /activities which returns all activity categories.
func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /activities", h.getActivities)
}

func (h *Handler) getActivities(w http.ResponseWriter, r *http.Request) {

	list, err := h.store.GetAll(r.Context())
	if err != nil {
		http.Error(w, "failed to fetch activities", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(list); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
