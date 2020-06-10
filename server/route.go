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

func Setup(useS3 bool) {
	if useS3 {
		delFile = s3fs.DeleteFile
		dlFile = s3fs.DownloadFile
		pFile = s3fs.PutFile
		ls = s3fs.ListFiles
	} else {
		delFile = deleteFileLocal
		dlFile = downloadFileLocal
		pFile = putFileLocal
		ls = listFileLocal
	}

	http.HandleFunc("/api/delete/*", deleteFile)
	http.HandleFunc("/api/ls/*", listFiles)
	http.HandleFunc("/api/upload/*", putFile)
	http.HandleFunc("/api/download/*", downloadFile)
	http.HandleFunc("/api/deleteMulti/*", deleteMulti)
}
