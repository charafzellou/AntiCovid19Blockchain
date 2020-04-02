package routes

import (
	"../contractManager"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetTreatments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func GetTreatmentsById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	treatmentId, err := strconv.ParseInt(params["_id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Expected treatment Id as a number"}`))
		return
	}

	cManager, err := contractManager.NewConractManager("xxx", "url")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errResponse := fmt.Sprintf(`{"error": "%v"}`, err)
		w.Write([]byte(errResponse))
		return
	}

	response, err := cManager.GetTreatment(treatmentId, "xxx")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errResponse := fmt.Sprintf(`{"error": "%v"}`, err)
		w.Write([]byte(errResponse))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func PostTreatment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}

	id, err := strconv.ParseInt(r.FormValue("idTreatment"),10,64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Expected patientId as a number"}`))
		return
	}
	activeAgent := r.FormValue("activeAgent")
	description := r.FormValue("description")
	steps := r.FormValue("steps")

	cManager, err := contractManager.NewConractManager("xxx", "url")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errResponse := fmt.Sprintf(`{"error": "%v"}`, err)
		w.Write([]byte(errResponse))
		return
	}

	response, err := cManager.CreateTreatment(
		id,
		activeAgent,
		description,
		steps,
		"xxx",
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errResponse := fmt.Sprintf(`{"error": "%v"}`, err)
		w.Write([]byte(errResponse))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(response))
}

