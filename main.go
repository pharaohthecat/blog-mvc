package main

import (
	admin_models "github.com/pharaohthecat/blog-mvc/admin/models"
	"github.com/pharaohthecat/blog-mvc/config"
	"net/http"
)

func main(){
	admin_models.Post{}.Migrate()
	admin_models.User{}.Migrate()
	admin_models.Category{}.Migrate()
	http.ListenAndServe(":8080",config.Routes())
}
