package server

import (
	"io"
	"net/http"
)

func downloadFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	path := r.Form.Get("path")
	if len(path) <= 0 {
		w.WriteHeader(400)
		return
	}

	f, err := dlFile(path)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	_, err = io.Copy(w, f)
	if err != nil {
		w.WriteHeader(500)
	}
}
