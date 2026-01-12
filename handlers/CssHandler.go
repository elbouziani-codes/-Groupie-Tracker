package handlers

import "net/http"

func CssHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "static/style.css")
}