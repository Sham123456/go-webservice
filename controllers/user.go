package controllers

import (
	"net/http"
	"regexp"

	"github.com/Sham123456/go-webservice/models"
)

type userController struct {
	userIdPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.GetAll(w, r)
		}
	}
}

func (uc *userController) GetAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJson(models.GetUsers(), w)
}

func (uc *userController) Get(id int, w http.ResponseWriter) {
	u, err := models.GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJson(u, w)
}

func newUserController() *userController {
	return &userController{
		userIdPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
