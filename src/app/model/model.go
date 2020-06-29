package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Person struct {
	gorm.Model
	FirstName        string           `json:"FirstName"`
	LastName         string           `json:"LastName"`
	DateOfBirth      string           `json:"DateOfBirth"`
	EducationalLevel string           `json:"EducationalLevel"`
	Adress           string           `json:"Adress"`
	Tel              string           `json:"Tel"`
	Email            string           `json:"Email"`
	Summary          string           `json:"Summary"`
	WorkExperiences  []WorkExperience `gorm:"ForeignKey:PersonID" json:"WorkExperiences"`
	Skills           []Skill          `gorm:"ForeignKey:PersonID" json:"Skills"`
	Certifications   []Certification  `gorm:"ForeignKey:PersonID" json:"Certifications"`
	Language         []Language       `gorm:"ForeignKey:PersonID" json:"Languages"`
}

type Certification struct {
	gorm.Model
	CertificationName      string `json:"CertificationName"`
	CertificationAuthority string `json:"CertificationAuthority"`
	FromYear               string `json:"FromYear"`
	ToYear                 string `json:"ToYear"`
	PersonID               uint   `json:"person_id"`
}
type WorkExperience struct {
	gorm.Model
	Company     string `json:"WorkExperienceName"`
	Title       string `json:"Title"`
	FromYear    string `json:"FromYear"`
	ToYear      string `json:"ToYear"`
	Description string `json:"Description"`
	PersonID    uint   `json:"person_id"`
}
type Skill struct {
	gorm.Model
	SkillName string `json:"SkillName"`
	PersonID  uint   `json:"person_id"`
}

type Language struct {
	gorm.Model
	LanguageName  string `json:"LanguageName"`
	LanguageLevel string `json:"LanguageLevel"`
	PersonID      uint   `json:"person_id"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Person{}, &Certification{}, &WorkExperience{}, &Skill{}, &Language{})
	db.Model(&Language{}).AddForeignKey("person_id", "people(id)", "RESTRICT", "RESTRICT")
	db.Model(&Certification{}).AddForeignKey("person_id", "people(id)", "RESTRICT", "RESTRICT")
	db.Model(&WorkExperience{}).AddForeignKey("person_id", "people(id)", "RESTRICT", "RESTRICT")
	db.Model(&Skill{}).AddForeignKey("person_id", "people(id)", "RESTRICT", "RESTRICT")
	return db
}
