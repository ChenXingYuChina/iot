package s3fs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Start()
	m.Run()
}

func TestPutFile(t *testing.T) {
	f, err := os.Open("test_file.txt")
	if err != nil {
		t.Error(err)
		return
	}
	err = PutFile(f, "folder/testFile2.txt")
	if err != nil {
		t.Error(err)
		return
	}
	f.Close()
}

func TestListFiles(t *testing.T) {
	f, err := ListFiles("folder/", "folder/testFile.txt", 1)
	if err != nil {
		t.Error(err)
		return
	}
	g, err := json.Marshal(f)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(g))
}

func TestListFilesStartAfter(t *testing.T) {
	f, err := ListFiles("", "folder2/", 10)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(f)
}

func TestDownloadFile(t *testing.T) {
	f, err := DownloadFile("testFile.txt")
	if err != nil {
		t.Error(err)
		return
	}
	all, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(all))
}

func TestDeleteFile(t *testing.T) {
	err := DeleteFile("testFile.txt")
	if err != nil {
		t.Error(err)
		return
	}
	TestPutFile(t)
}
