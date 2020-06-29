package handler

import (
	"encoding/json"
	"limakcv/src/app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllSkill(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, id, w, r)
	if person == nil {
		return
	}

	skill := []model.Skill{}
	if err := db.Model(&person).Related(&skill).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, skill)
}

func CreateSkill(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, id, w, r)
	if person == nil {
		return
	}

	skill := model.Skill{PersonID: person.ID}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&skill); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&skill).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, skill)
}

func GetSkill(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

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
	skill := getSkillOr404(db, id, w, r)
	if skill == nil {
		return
	}
	respondJSON(w, http.StatusOK, skill)
}

func UpdateSkill(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
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
	skill := getSkillOr404(db, id, w, r)
	if skill == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&skill); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&skill).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, skill)
}

func DeleteSkill(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
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
	skill := getSkillOr404(db, id, w, r)
	if skill == nil {
		return
	}

	if err := db.Delete(&skill).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getSkillOr404(db *gorm.DB, skillId int, w http.ResponseWriter, r *http.Request) *model.Skill {
	skill := model.Skill{}
	if err := db.First(&skill, skillId).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &skill
}
