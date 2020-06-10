package server

import (
	"io"
	"io/ioutil"
	"iot/s3fs"
	"net/url"
	"os"
	"strings"
)

var prefix string

func deleteFileLocal(filename string) error {
	return os.Remove(prefix + filename)
}

func downloadFileLocal(filename string) (io.ReadCloser, error) {
	return os.Open(prefix + filename)
}

func putFileLocal(f s3fs.Object, filename string) error {
	filename = prefix + filename
	i := strings.LastIndex(filename, "/")
	_ = os.MkdirAll(filename[i:], os.ModePerm)
	file, err := os.Create(prefix + filename)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, f)
	return err
}

func listFileLocal(folder string, startAfter string, keyNumber int64) (*s3fs.FileList, error) {
	if len(startAfter) == 0 {
		startAfter = folder
	}
	baseName := startAfter[len(folder):]
	visualFolder := folder
	folder = prefix + folder
	dir, err := ioutil.ReadDir(folder[:len(folder)-1])
	if err != nil {
		return nil, err
	}
	goal := &s3fs.FileList{}
	i := int64(0)
	after := false
	for loc, v := range dir {
		if after {
			if v.IsDir() {
				goal.Files = append(goal.Files, visualFolder+v.Name()+"/")
			} else {
				goal.Files = append(goal.Files, visualFolder+v.Name())
			}
			i++
			if i >= keyNumber {
				if loc != len(dir) - 1 {
					goal.Next = "/api/ls/"+visualFolder +"?start_after=" + url.QueryEscape(goal.Files[len(goal.Files) -1])
				}
				break
			}
			continue
		}
		after = v.Name() > baseName
		if after {
			if v.IsDir() {
				goal.Files = append(goal.Files, visualFolder+v.Name()+"/")
			} else {
				goal.Files = append(goal.Files, visualFolder+v.Name())
			}
			i++
			if i >= keyNumber {
				if loc != len(dir) - 1 {
					goal.Next = "/api/ls/"+visualFolder +"?start_after=" + url.QueryEscape(goal.Files[len(goal.Files) -1])
				}
				break
			}
		}
	}
	return goal, nil
}
