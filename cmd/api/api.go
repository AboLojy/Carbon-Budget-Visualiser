package apiserver

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/AboLojy/Carbon-Budget-Visualiser/services/activitycategory"
	"github.com/AboLojy/Carbon-Budget-Visualiser/services/annualemission"
	"github.com/AboLojy/Carbon-Budget-Visualiser/services/budgetcalculator"
	"github.com/AboLojy/Carbon-Budget-Visualiser/services/emission"
)

type APIServer struct {
	address string
	db      *sql.DB
}

func NewAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{
		address: address,
		db:      db,
	}
}

func (s *APIServer) Start() error {
	router := http.NewServeMux()
	cityEmissionStore := emission.NewStore(s.db)
	activitycategoryStore := activitycategory.NewStore(s.db)
	annualemissionStor := annualemission.NewStore(s.db)
	budgetStor := budgetcalculator.NewStore(s.db)

	activitycategoryHandler := activitycategory.NewHandler(activitycategoryStore)
	emissionHandler := emission.NewHandler(cityEmissionStore)
	annualemissionHandler := annualemission.NewHandler(annualemissionStor)
	budgetcalculatorHandler := budgetcalculator.NewHandler(budgetStor)

	// Register routes with /api/v1 prefix
	apiRouter := http.NewServeMux()
	emissionHandler.RegisterRoutes(apiRouter)
	activitycategoryHandler.RegisterRoutes(apiRouter)
	annualemissionHandler.RegisterRoutes(apiRouter)
	budgetcalculatorHandler.RegisterRoutes(apiRouter)

	router.Handle("/api/v1/", http.StripPrefix("/api/v1", apiRouter))

	log.Println("Listening on ", s.address)
	return http.ListenAndServe(s.address, router)
}
