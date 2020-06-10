package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func deleteFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(400)
		return
	}
	err = delFile(r.Form["path"][0])
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
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
	var data []string
	err = json.Unmarshal([]byte(r.Form["paths"][0]), &data)
	if err != nil {
		return
	}
	for _, v := range data {
		err = delFile(v)
		if err != nil {
			log.Println(err)
			continue
		}
		w.WriteHeader(500)
	}
	w.WriteHeader(200)
}
