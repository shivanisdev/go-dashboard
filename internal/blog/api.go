package blog

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterHandlers add routes of blogs
func RegisterHandlers(routers *mux.Router) {
	routers.HandleFunc("/blogs/{title}", getBlogs).Methods("GET")
}

func getBlogs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Blogs: %v\n", vars["title"])
}
