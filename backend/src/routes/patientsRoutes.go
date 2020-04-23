package routes

import (
	"anticovid/contractManager"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var ContractAddress *string
var RPCurl *string

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

	cManager, err := contractManager.NewContractManager(*ContractAddress, *RPCurl, "", "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errResponse := fmt.Sprintf(`{"error": "%v"}`, err)
		w.Write([]byte(errResponse))
		return
	}

	response, err := cManager.GetPatient(patientId)
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
	jsonKey := r.FormValue("jsonKey")
	passphrase := r.FormValue("passphrase")

	cManager, err := contractManager.NewContractManager(*ContractAddress, *RPCurl, jsonKey, passphrase)
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
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	patientId, err := strconv.ParseInt(params["_id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Expected patient Id as a number"}`))
		return
	}

	screening, err := strconv.ParseInt(params["screening"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Expected unix time"}`))
		return
	}
	jsonKey := r.FormValue("jsonKey")
	passphrase := r.FormValue("passphrase")

	cManager, err := contractManager.NewContractManager(*ContractAddress, *RPCurl, jsonKey, passphrase)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errResponse := fmt.Sprintf(`{"error": "%v"}`, err)
		w.Write([]byte(errResponse))
		return
	}

	response, err := cManager.SetScreeningDate(patientId, screening)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errResponse := fmt.Sprintf(`{"error": "%v"}`, err)
		w.Write([]byte(errResponse))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(response))
}

func PostPatientDeath(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	patientId, err := strconv.ParseInt(params["_id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Expected patient Id as a number"}`))
		return
	}

	death, err := strconv.ParseInt(params["death"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Expected unix time"}`))
		return
	}
	jsonKey := r.FormValue("jsonKey")
	passphrase := r.FormValue("passphrase")

	cManager, err := contractManager.NewContractManager(*ContractAddress, *RPCurl, jsonKey, passphrase)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errResponse := fmt.Sprintf(`{"error": "%v"}`, err)
		w.Write([]byte(errResponse))
		return
	}

	response, err := cManager.SetDeathDate(patientId, death)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errResponse := fmt.Sprintf(`{"error": "%v"}`, err)
		w.Write([]byte(errResponse))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(response))
}

func PostPatientRemission(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	patientId, err := strconv.ParseInt(params["_id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Expected patient Id as a number"}`))
		return
	}

	remission, err := strconv.ParseInt(params["remission"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Expected unix time"}`))
		return
	}
	jsonKey := r.FormValue("jsonKey")
	passphrase := r.FormValue("passphrase")

	cManager, err := contractManager.NewContractManager(*ContractAddress, *RPCurl, jsonKey, passphrase)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errResponse := fmt.Sprintf(`{"error": "%v"}`, err)
		w.Write([]byte(errResponse))
		return
	}

	response, err := cManager.SetRemissionDate(patientId, remission)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errResponse := fmt.Sprintf(`{"error": "%v"}`, err)
		w.Write([]byte(errResponse))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(response))
}

