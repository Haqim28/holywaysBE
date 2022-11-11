package routes

import (
	"holyways/handlers"
	"holyways/pkg/mysql"
	"holyways/repositories"

	"github.com/gorilla/mux"
)

func DonationRoutes(r *mux.Router) {
	donationRepository := repositories.RepositoryDonation(mysql.DB)
	h := handlers.HandlerDonation(donationRepository)

	r.HandleFunc("/donation", h.CreateDonation).Methods("POST")
	r.HandleFunc("/donations/{id}", h.GetDonationByFund).Methods("GET")
	r.HandleFunc("/donation/{id}", h.GetDonationByUser).Methods("GET")
	r.HandleFunc("/donation/{id}", h.DeleteDonation).Methods("DELETE")
	// r.HandleFunc("/donation/{id}", h.UpdateDonation).Methods("PATCH")

}
