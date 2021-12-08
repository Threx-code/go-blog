package routes

import (
	"net/http"
	"regexp"

	"github.com/Threx-code/go-blog/package/controllers"
	"github.com/gorilla/mux"
)

var validPath = regexp.MustCompile("^/blogs/([0-9]?)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}

var BlogRoutes = func(router *mux.Router) {
	router.HandleFunc("/blogs/", makeHandler(controllers.Index)).Methods("GET")
	router.HandleFunc("/blogs/{id}", makeHandler(controllers.Read)).Methods("GET")
	router.HandleFunc("/blogs/{id}", makeHandler(controllers.Update)).Methods("PUT")
	router.HandleFunc("/blogs/{id}", makeHandler(controllers.Destroy)).Methods("DELETE")
	router.HandleFunc("/blogs/", makeHandler(controllers.Store)).Methods("POST")
}
