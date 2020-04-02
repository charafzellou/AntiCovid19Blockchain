package routes

import (
	"../contractManager"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetPatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func GetPatientById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	patientId, err := strconv.ParseInt(params["_id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Expected patient Id as a number"}`))
		return
	}

	cManager, err := contractManager.NewConractManager("xxx", "url")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errResponse := fmt.Sprintf(`{"error": "%v"}`, err)
		w.Write([]byte(errResponse))
		return
	}

	response, err := cManager.GetPatient(patientId, "xxx")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errResponse := fmt.Sprintf(`{"error": "%v"}`, err)
		w.Write([]byte(errResponse))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func PostPatient(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}

	id, err := strconv.ParseInt(r.FormValue("idPatient"),10,64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Expected patientId as a number"}`))
		return
	}
	treatmentId, err := strconv.ParseInt(r.FormValue("treatmentId"),10,64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Expected treatmentId as a number"}`))
		return
	}
	screeningDate, _ := strconv.ParseInt(r.FormValue("screeningDate"),10,64)
	birthYear, _ := strconv.ParseInt(r.FormValue("birthYear"),10,64)
	sex, _ := strconv.ParseBool(r.FormValue("sex"))
	postCode := r.FormValue("postCode")
	country := r.FormValue("country")
	refPhysician := r.FormValue("refPhysician")
	medicalHistory := r.FormValue("medicalHistory")

	cManager, err := contractManager.NewConractManager("xxx", "url")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errResponse := fmt.Sprintf(`{"error": "%v"}`, err)
		w.Write([]byte(errResponse))
		return
	}

	response, err := cManager.CreatePatient(
		id,
		treatmentId,
		screeningDate,
		birthYear,
		sex,
		postCode,
		country,
		refPhysician,
		medicalHistory,
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

func PostPatientScreening(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func PostPatientDeath(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func PostPatientRemission(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

