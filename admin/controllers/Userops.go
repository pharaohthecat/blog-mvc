package controllers

import (
	"crypto/sha256"
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pharaohthecat/blog-mvc/admin/helpers"
	"github.com/pharaohthecat/blog-mvc/admin/models"
)

type Userops struct{}

func (userops Userops) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("userops/login")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

// Login da página
func (userops Userops) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	username := r.FormValue("username")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(r.FormValue("password"))))

	user := models.User{}.Get("username = ? AND password = ?", username, password)
	if user.Username == username && user.Password == password {
		helpers.SetUser(w, r, username, password)
		helpers.SetAlert(w, r, "Bem-vindo!")
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	} else {
		helpers.SetAlert(w, r, "Nome de usuário ou senha incorreta")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}
}

// Logout do Usuário
func (userops Userops) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	helpers.RemoveUser(w, r)
	helpers.SetAlert(w, r, "Até a próxima!")
	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
}