package main

import (
	"anticovid/routes"
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func init(){
	routes.ContractAddress = flag.String("contractAddr", "", "The address of the deployed HealthContract")
	routes.RPCurl = flag.String("rpcUrl", "http://localhost:8545", "RPC url")
	flag.Parse()
}

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
	patients.HandleFunc("/{_id}/screening/{screening}", routes.PostPatientScreening).Methods(http.MethodPost)
	patients.HandleFunc("/{_id}/death/{death}", routes.PostPatientDeath).Methods(http.MethodPost)
	patients.HandleFunc("/{_id}/remission/{remission}", routes.PostPatientRemission).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":3001", router))
}
