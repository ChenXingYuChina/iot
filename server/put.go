package server

import (
	"net/http"
)

func putFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024 * 1024)
	if err != nil {
		http.Redirect(w, r, "/error.html", http.StatusTemporaryRedirect)
		return
	}
	path := r.URL.Path[len(uploadApi):]
	if len(path) <= 0 {
		http.Redirect(w, r, "/error.html", http.StatusTemporaryRedirect)
		return
	}
	f, err := r.MultipartForm.File["upload_file"][0].Open()
	if err != nil {
		http.Redirect(w, r, "/error.html", http.StatusTemporaryRedirect)
		return
	}
	err = pFile(f, path)
	if err != nil {
		http.Redirect(w, r, "/error.html", http.StatusTemporaryRedirect)
		return
	}
	http.Redirect(w, r, "/index.html", http.StatusTemporaryRedirect)
	return
}
