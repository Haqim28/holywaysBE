package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	AuthRoutes(r)
	// ProductRoutes(r)
	DonationRoutes(r)
	FundRoutes(r)
	TransactionRoutes(r)
}
