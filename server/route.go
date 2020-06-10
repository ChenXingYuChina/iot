package server

import (
	"io"
	"iot/s3fs"
	"net/http"
)

var delFile func(filename string) error
var dlFile func(filename string) (io.ReadCloser, error)
var pFile func(f s3fs.Object, filename string) error
var ls func(prefix string, startAfter string, keyNumber int64) (*s3fs.FileList, error)

const (
	deleteApi = "/api/delete/"
	deleteMultiApi = "/api/deleteMulti"
	uploadApi = "/api/upload/"
	downloadApi = "/api/download/"
	lsApi = "/api/ls/"
)


func Setup(useS3 bool) {
	if useS3 {
		delFile = s3fs.DeleteFile
		dlFile = s3fs.DownloadFile
		pFile = s3fs.PutFile
		ls = s3fs.ListFiles
		s3fs.Start()
	} else {
		delFile = deleteFileLocal
		dlFile = downloadFileLocal
		pFile = putFileLocal
		ls = listFileLocal
	}

	http.HandleFunc(deleteApi, deleteFile)
	http.HandleFunc(lsApi, listFiles)
	http.HandleFunc(uploadApi, putFile)
	http.HandleFunc(downloadApi, downloadFile)
	http.HandleFunc(deleteMultiApi, deleteMulti)
	http.Handle("/", http.FileServer(http.Dir("./html")))
}

func Start() error {
	return http.ListenAndServe("0.0.0.0:80", nil)
}
