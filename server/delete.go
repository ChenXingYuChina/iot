package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func deleteFile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len(deleteApi):]
	if len(path) <= 0 {
		w.WriteHeader(400)
		return
	}
	log.Println("call delete", path)
	err := delFile(path)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}

func deleteMulti(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(400)
		return
	}
	for k, v := range r.Form {
		log.Println(k, v)
	}
	//log.Println(r.Form)
	var data []string
	err = json.Unmarshal([]byte(r.Form["paths"][0]), &data)
	if err != nil {
		return
	}
	log.Println("call multi delete", data)
	for _, v := range data {
		err = delFile(v)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			continue
		}
	}
	w.WriteHeader(200)
}
