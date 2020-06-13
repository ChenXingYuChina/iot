package server

import (
	"log"
	"net/http"
)

func putFile(w http.ResponseWriter, r *http.Request) {
	//body, err2 := ioutil.ReadAll(r.Body)
	//if err2 != nil {
	//	log.Println(err2)
	//	return
	//}
	//log.Println(body)
	err := r.ParseMultipartForm(1024 * 1024)
	if err != nil {
		log.Println("parse fail", err)
		http.Redirect(w, r, "/error.html", http.StatusTemporaryRedirect)
		return
	}
	log.Println(r.MultipartForm.Value)
	fn := r.MultipartForm.Value["filename"]
	if fn == nil || len(fn) != 1 {
		log.Println("filename fail")
		http.Redirect(w, r, "/error.html", http.StatusTemporaryRedirect)
		return
	}
	filename := fn[0]
	if len(filename) <= 0 {
		log.Println("filename fail 2")
		http.Redirect(w, r, "/error.html", http.StatusTemporaryRedirect)
		return
	}
	p := r.MultipartForm.Value["folder"]
	if p == nil || len(p) != 1 {
		log.Println("folder fail")
		http.Redirect(w, r, "/error.html", http.StatusTemporaryRedirect)
		return
	}
	folder := p[0]
	if len(folder) <= 0{
		log.Println("folder fail 2")
		http.Redirect(w, r, "/error.html", http.StatusTemporaryRedirect)
		return
	}
	log.Println(r.MultipartForm.File)
	f := r.MultipartForm.File["file_upload"]
	if f == nil || len(f) == 0 {
		log.Println("file fail")
		http.Redirect(w, r, "/error.html", http.StatusTemporaryRedirect)
		return
	}
	file, err := f[0].Open()
	if err != nil {
		log.Println("file fail")
		http.Redirect(w, r, "/error.html", http.StatusTemporaryRedirect)
		return
	}
	err = pFile(file, folder + filename)
	if err != nil {
		log.Println("file fail 2")
		http.Redirect(w, r, "/error.html", http.StatusTemporaryRedirect)
		return
	}
	http.Redirect(w, r, "/index.html", http.StatusTemporaryRedirect)
	return
}
