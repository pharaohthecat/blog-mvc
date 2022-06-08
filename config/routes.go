package config

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	admin "github.com/pharaohthecat/blog-mvc/admin/controllers"
	site "github.com/pharaohthecat/blog-mvc/site/controllers"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	//ADMIN
	//Blog Posts
	r.GET("/admin",admin.Dashboard{}.Index)
	r.GET("/admin/new",admin.Dashboard{}.NewItem)
	r.POST("/admin/add",admin.Dashboard{}.Add)
	r.GET("/admin/delete/:id",admin.Dashboard{}.Delete)
	r.GET("/admin/edit/:id",admin.Dashboard{}.Edit)
	r.POST("/admin/update/:id",admin.Dashboard{}.Update)

	//Categories
	r.GET("/admin/categories",admin.Categories{}.Index)
	r.POST("/admin/categories/add",admin.Categories{}.Add)
	r.GET("/admin/categories/delete/:id",admin.Categories{}.Delete)

	//Userops
	r.GET("/admin/login",admin.Userops{}.Index)
	r.POST("/admin/do_login",admin.Userops{}.Login)
	r.GET("/admin/logout",admin.Userops{}.Logout)

	//SITE
	//Homepage
	r.GET("/",site.Homepage{}.Index)
	r.GET("/post/:slug",site.Homepage{}.Detail)

	// SERVER FILES
	r.ServeFiles("/admin/assets/*filepath",http.Dir("admin/assets"))
	r.ServeFiles("/assets/*filepath",http.Dir("site/assets"))
	r.ServeFiles("/uploads/*filepath",http.Dir("uploads"))
	return r
}