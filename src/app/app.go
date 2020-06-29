package app

import (
	"log"
	"net/http"

	"limakcv/src/app/handler"
	"limakcv/src/app/model"
	"limakcv/src/config"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	// dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
	// 	config.DB.Username,
	// 	config.DB.Password,
	// 	config.DB.Host,
	// 	config.DB.Port,
	// 	config.DB.Name,
	// 	config.DB.Charset,
	// )

	//db, err := gorm.Open(config.DB.Dialect, dbURI)
	db, err := gorm.Open("postgres", "host=localhost user=postgres password=kubra dbname=kubDeneme sslmode=disable")
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Get("/persons", a.handleRequest(handler.GetAllPersons))
	a.Post("/persons", a.handleRequest(handler.CreatePerson))
	a.Get("/persons/{PersonID}", a.handleRequest(handler.GetPerson))
	a.Put("/persons/{PersonID}", a.handleRequest(handler.UpdatePerson))
	a.Delete("/persons/{PersonID}", a.handleRequest(handler.DeletePerson))

	a.Get("/persons/{PersonID}/certifications", a.handleRequest(handler.GetAllCertifications))
	a.Post("/persons/{PersonID}/certifications", a.handleRequest(handler.CreateCertification))
	a.Get("/persons/{PersonID}/certifications/{id:[0-9]+}", a.handleRequest(handler.GetCertification))
	a.Put("/persons/{PersonID}/certifications/{id:[0-9]+}", a.handleRequest(handler.UpdateCertification))
	a.Delete("/persons/{PersonID}/certifications/{id:[0-9]+}", a.handleRequest(handler.DeleteCertification))

	a.Get("/persons/{PersonID}/workexperiences", a.handleRequest(handler.GetAllWorkExperience))
	a.Post("/persons/{PersonID}/workexperiences", a.handleRequest(handler.CreateWorkExperience))
	a.Get("/persons/{PersonID}/workexperiences/{id:[0-9]+}", a.handleRequest(handler.GetWorkExperience))
	a.Put("/persons/{PersonID}/workexperiences/{id:[0-9]+}", a.handleRequest(handler.UpdateWorkExperience))
	a.Delete("/persons/{PersonID}/workexperiences/{id:[0-9]+}", a.handleRequest(handler.DeleteWorkExperience))

	a.Get("/persons/{PersonID}/skills", a.handleRequest(handler.GetAllSkill))
	a.Post("/persons/{PersonID}/skills", a.handleRequest(handler.CreateSkill))
	a.Get("/persons/{PersonID}/skills/{id:[0-9]+}", a.handleRequest(handler.GetSkill))
	a.Put("/persons/{PersonID}/skills/{id:[0-9]+}", a.handleRequest(handler.UpdateSkill))
	a.Delete("/persons/{PersonID}/skills/{id:[0-9]+}", a.handleRequest(handler.DeleteSkill))

	a.Get("/persons/{PersonID}/languages", a.handleRequest(handler.GetAllLanguage))
	a.Post("/persons/{PersonID}/languages", a.handleRequest(handler.CreateLanguage))
	a.Get("/persons/{PersonID}/languages/{id:[0-9]+}", a.handleRequest(handler.GetLanguage))
	a.Put("/persons/{PersonID}/languages/{id:[0-9]+}", a.handleRequest(handler.UpdateLanguage))
	a.Delete("/persons/{PersonID}/languages/{id:[0-9]+}", a.handleRequest(handler.DeleteLanguage))

}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}
