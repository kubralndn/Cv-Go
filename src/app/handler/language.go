package handler

import (
	"encoding/json"
	"limakcv/src/app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllLanguage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, id, w, r)
	if person == nil {
		return
	}

	language := []model.Language{}
	if err := db.Model(&person).Related(&language).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, language)
}

func CreateLanguage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, id, w, r)
	if person == nil {
		return
	}

	language := model.Language{PersonID: person.ID}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&language); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&language).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, language)
}

func GetLanguage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

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
	language := getLanguageOr404(db, id, w, r)
	if language == nil {
		return
	}
	respondJSON(w, http.StatusOK, language)
}

func UpdateLanguage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
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
	language := getLanguageOr404(db, id, w, r)
	if language == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&language); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&language).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, language)
}

func DeleteLanguage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
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
	skill := getLanguageOr404(db, id, w, r)
	if skill == nil {
		return
	}

	if err := db.Delete(&skill).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getLanguageOr404(db *gorm.DB, languageId int, w http.ResponseWriter, r *http.Request) *model.Language {
	language := model.Language{}
	if err := db.First(&language, languageId).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &language
}
