package main

import (
	"log"
	"net/http"

	"../routes"

	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()

	// Treatment
	treatments := router.PathPrefix("/v1/treatments").Subrouter()
	// GET
	//treatments.HandleFunc("", routes.GetTreatments).Methods(http.MethodGet)
	treatments.HandleFunc("/{id}", routes.GetTreatmentsById).Methods(http.MethodGet)
	// POST
	treatments.HandleFunc("", routes.PostTreatment).Methods(http.MethodPost)

	// Patient
	patients := router.PathPrefix("/v1/patients").Subrouter()
	// GET
	//patients.HandleFunc("", routes.GetPatient).Methods(http.MethodGet)
	patients.HandleFunc("/{_id}", routes.GetPatientById).Methods(http.MethodGet)
	// POST
	patients.HandleFunc("", routes.PostPatient).Methods(http.MethodPost)
	patients.HandleFunc("/{_id}/{screening}", routes.PostPatientScreening).Methods(http.MethodPost)
	patients.HandleFunc("/{_id}/{death}", routes.PostPatientDeath).Methods(http.MethodPost)
	patients.HandleFunc("/{_id}/{remission}", routes.PostPatientRemission).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":3001", router))
}
