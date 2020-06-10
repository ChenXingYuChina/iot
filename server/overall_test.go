package server

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListFiles(t *testing.T) {
	t.Log(testHttpHelp(t, "/api/ls/", nil, 200, listFiles))
}

func TestLsAfter(t *testing.T) {
	t.Log(testHttpHelp(t, "/api/ls/?start_after=download.go", nil, 200, listFiles))
}

type multiHelp struct {
	*bytes.Buffer
}


func (*multiHelp) Close() error {
	return nil
}

func TestUpload(t *testing.T) {
	r := &http.Request{
		Header: map[string][]string{},
	}
	body := &bytes.Buffer{}
	w := multipart.NewWriter(body)
	field, err := w.CreateFormFile("upload_file", "filename")
	if err != nil {
		t.Error(err)
		return
	}
	_, err = field.Write([]byte("test files"))
	if err != nil {
		t.Error(err)
		return
	}
	err = w.Close()
	if err != nil {
		t.Error(err)
		return
	}
	r.Body = &multiHelp{body}
	r.Header.Add("Content-Type", w.FormDataContentType())
	t.Log(testHttpHelp(t, "/api/upload/folder/test_file2.txt", r, 307, putFile))
}

func TestDownloadFile(t *testing.T)  {
	t.Log(testHttpHelp(t, "/api/download/folder/test_file2.txt", nil, 200, downloadFile))
}

func TestDeleteFile(t *testing.T) {
	t.Log(testHttpHelp(t, "/api/delete/folder/test_file2.txt", nil, 200, deleteFile))
}

func TestDeleteFiles(t *testing.T) {
	body := url.Values{}
	g, err := json.Marshal([]string{"folder/test_file2.txt", "folder/test_file3.txt"})
	if err != nil {
		t.Error(err)
		return
	}
	body.Set("paths", string(g))
	r := &http.Request{
		Method:http.MethodPost,
		PostForm: body,
	}
	t.Log(testHttpHelp(t, "/api/deleteMulti", r, 200, deleteMulti))
}


func testHttpHelp(t *testing.T, u string, r *http.Request, expectCode int, f http.HandlerFunc) string {
	if r == nil {
		r = &http.Request{}
	}
	var err error
	if r.URL == nil {
		r.URL, err = url.Parse(u)
		if err != nil {
			t.Error(err)
			return ""
		}
	}
	rec := httptest.NewRecorder()
	f(rec, r)
	if rec.Code != expectCode {
		t.Error(rec.Code)
		return ""
	}
	if rec.Code == 307 {
		t.Log(rec.Header().Get("Location"))
	}
	return rec.Body.String()
}
