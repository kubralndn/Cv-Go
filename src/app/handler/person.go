package handler

import (
	"encoding/json"
	"limakcv/src/app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllPersons(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	persons := []model.Person{}
	db.Find(&persons)
	for i, _ := range persons {
		db.Model(persons[i]).Related(&persons[i].WorkExperiences)
		db.Model(persons[i]).Related(&persons[i].Language)
		db.Model(persons[i]).Related(&persons[i].Skills)
		db.Model(persons[i]).Related(&persons[i].Certifications)
	}
	respondJSON(w, http.StatusOK, persons)
}

func CreatePerson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	person := model.Person{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&person); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&person).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, person)
}

func GetPerson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, id, w, r)
	if person == nil {
		return
	}
	db.Model(person).Related(&person.Certifications)
	db.Model(person).Related(&person.Skills)
	db.Model(person).Related(&person.Language)
	db.Model(person).Related(&person.WorkExperiences)
	respondJSON(w, http.StatusOK, person)
}

func UpdatePerson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, id, w, r)
	if person == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&person); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&person).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, person)
}

func DeletePerson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, id, w, r)
	if person == nil {
		return
	}
	if err := db.Delete(&person).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getPersonOr404(db *gorm.DB, personID int, w http.ResponseWriter, r *http.Request) *model.Person {
	person := model.Person{}
	if err := db.First(&person, model.Person{Model: gorm.Model{ID: uint(personID)}}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &person
}
