package server

import "net/http"

func putFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024*1024)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	f, err := r.MultipartForm.File["upload_file"][0].Open()
	if err != nil {
		w.WriteHeader(400)
		return
	}
	paths := r.MultipartForm.Value["ab_path"]
	if paths == nil || len(paths) != 1 {
		w.WriteHeader(400)
		return
	}
	err = pFile(f, paths[0])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	return
}
