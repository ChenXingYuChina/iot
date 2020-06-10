package server

import (
	"log"
	"net/http"
)

func putFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024*1024)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}
	path := r.URL.Path[len(uploadApi):]
	if len(path) <= 0 {
		w.WriteHeader(400)
		return
	}
	f, err := r.MultipartForm.File["upload_file"][0].Open()
	if err != nil {
		w.WriteHeader(400)
		return
	}
	err = pFile(f, path)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	http.Redirect(w, r, "../index.html", http.StatusTemporaryRedirect)
	return
}
