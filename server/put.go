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
	fn := r.MultipartForm.Value["filename"]
	if fn == nil || len(fn) != 1 {
		http.Redirect(w, r, "/error.html", http.StatusTemporaryRedirect)
		return
	}
	filename := fn[0]
	if len(filename) <= 0 {
		http.Redirect(w, r, "/error.html", http.StatusTemporaryRedirect)
		return
	}
	p := r.MultipartForm.Value["folder"]
	if p == nil || len(p) != 1 {
		http.Redirect(w, r, "/error.html", http.StatusTemporaryRedirect)
		return
	}
	folder := p[0]
	if len(folder) <= 0{
		http.Redirect(w, r, "/error.html", http.StatusTemporaryRedirect)
		return
	}
	f, err := r.MultipartForm.File["upload_file"][0].Open()
	if err != nil {
		http.Redirect(w, r, "/error.html", http.StatusTemporaryRedirect)
		return
	}
	err = pFile(f, folder + filename)
	if err != nil {
		http.Redirect(w, r, "/error.html", http.StatusTemporaryRedirect)
		return
	}
	http.Redirect(w, r, "/index.html", http.StatusTemporaryRedirect)
	return
}
