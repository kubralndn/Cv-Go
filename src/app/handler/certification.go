package handler

import (
	"encoding/json"
	"limakcv/src/app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllCertifications(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, id, w, r)
	if person == nil {
		return
	}

	certifications := []model.Certification{}
	if err := db.Model(&person).Related(&certifications).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, certifications)
}

func CreateCertification(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, id, w, r)
	if person == nil {
		return
	}

	certification := model.Certification{PersonID: person.ID}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&certification); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&certification).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, certification)
}

func GetCertification(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	personid, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, personid, w, r)
	if person == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	certification := getCertificateOr404(db, id, w, r)
	if certification == nil {
		return
	}
	respondJSON(w, http.StatusOK, certification)
}

func UpdateCertification(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	personid, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, personid, w, r)
	if person == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	certification := getCertificateOr404(db, id, w, r)
	if certification == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&certification); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&certification).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, certification)
}

func DeleteCertification(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	personid, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, personid, w, r)
	if person == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	certification := getCertificateOr404(db, id, w, r)
	if certification == nil {
		return
	}

	if err := db.Delete(&certification).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getCertificateOr404(db *gorm.DB, certificationId int, w http.ResponseWriter, r *http.Request) *model.Certification {
	certification := model.Certification{}
	if err := db.First(&certification, certificationId).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &certification
}
