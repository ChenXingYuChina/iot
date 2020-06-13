package server

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Prefix = "./"
	Setup(false)
	m.Run()
}

func TestList(t *testing.T) {
	t.Log(listFileLocal("", "", 20))
}

func TestList2(t *testing.T) {
	t.Log(listFileLocal("", "download.go", 2))
}

func TestList3(t *testing.T) {
	t.Log(listFileLocal("folder/", "", 1))
}

func TestDownload(t *testing.T) {
	r, err := downloadFileLocal("folder/test_file.txt")
	if err != nil {
		t.Error(err)
		return
	}
	d, err := ioutil.ReadAll(r)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(d))
}

func TestPut(t *testing.T) {
	f, err := os.Open("folder/test_file.txt")
	if err != nil {
		t.Error(err)
		return
	}
	err = putFileLocal(f, "folder/put_file.txt")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestDelete(t *testing.T) {
	err := deleteFileLocal("folder/put_file.txt")
	if err != nil {
		t.Error(err)
		return
	}
}
