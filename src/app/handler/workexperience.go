package handler

import (
	"encoding/json"
	"limakcv/src/app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllWorkExperience(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, id, w, r)
	if person == nil {
		return
	}

	workexperience := []model.WorkExperience{}
	if err := db.Model(&person).Related(&workexperience).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, workexperience)
}

func CreateWorkExperience(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, id, w, r)
	if person == nil {
		return
	}

	workexperience := model.WorkExperience{PersonID: person.ID}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&workexperience); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&workexperience).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, workexperience)
}

func GetWorkExperience(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

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
	workexperience := getWorkExperienceOr404(db, id, w, r)
	if workexperience == nil {
		return
	}
	respondJSON(w, http.StatusOK, workexperience)
}

func UpdateWorkExperience(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
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
	workexperience := getWorkExperienceOr404(db, id, w, r)
	if workexperience == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&workexperience); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&workexperience).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, workexperience)
}

func DeleteWorkExperience(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
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
	workexperience := getWorkExperienceOr404(db, id, w, r)
	if workexperience == nil {
		return
	}

	if err := db.Delete(&workexperience).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getWorkExperienceOr404(db *gorm.DB, workexperienceId int, w http.ResponseWriter, r *http.Request) *model.WorkExperience {
	workexperience := model.WorkExperience{}
	if err := db.First(&workexperience, workexperienceId).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &workexperience
}
