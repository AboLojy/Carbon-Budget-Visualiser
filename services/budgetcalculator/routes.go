package budgetcalculator

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/AboLojy/Carbon-Budget-Visualiser/types"
	"github.com/AboLojy/Carbon-Budget-Visualiser/utils"
)

// Handler provides HTTP handlers for budget calculations.
type Handler struct {
	store BudgetStore
}

// NewHandler creates a new Handler with the provided store.
func NewHandler(store BudgetStore) *Handler {
	return &Handler{store: store}
}

// RegisterRoutes registers budget calculator routes on the provided router.
// This registers GET /cities/{cityId}/budget-projection
func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /citybudget/{cityid}/{year}", h.getBudgetProjection)
	router.HandleFunc("POST /citybudget", h.createCityBudget)
}

func (h *Handler) getBudgetProjection(w http.ResponseWriter, r *http.Request) {

	cityIdstr := r.PathValue("cityid")
	year := r.PathValue("year")
	if cityIdstr == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing city ID"))
		return
	}
	if year == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing year"))
		return
	}
	startyear, err := strconv.Atoi(strings.TrimSpace(year))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid year: %v", err))
		return
	}
	cityId, err := strconv.Atoi(strings.TrimSpace(cityIdstr))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid city ID: %v", err))
		return
	}

	projection, err := h.store.CalculateBudgetConsumption(r.Context(), cityId, startyear)
	if err != nil {
		http.Error(w, "failed to calculate budget projection: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(projection)
}
func (h *Handler) createCityBudget(w http.ResponseWriter, r *http.Request) {
	var city types.City
	if err := json.NewDecoder(r.Body).Decode(&city); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	created, err := h.store.CreateCity(r.Context(), &city)
	if err != nil {
		http.Error(w, "failed to create city budget: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(created)
}
