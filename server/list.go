package server

import (
	"encoding/json"
	"iot/s3fs"
	"net/http"
	"net/url"
)

func listFiles(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(400)
		return
	}
	prefix := r.URL.Path
	sa := r.Form.Get("start_after")
	if sa != "" {
		sa, err = url.QueryUnescape(sa)
	}
	if err != nil {
		w.WriteHeader(400)
		return
	}
	var fs *s3fs.FileList
	fs, err = ls(prefix, sa, 20)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	var b []byte
	b, err = json.Marshal(fs)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	_, _ =w.Write(b)
}
