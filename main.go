package main

import (
	"bufio"
	"iot/s3fs"
	"iot/server"
	"os"
	"strings"
)

func main() {
	server.Setup(parseCredentials())
	server.Start()
}

func parseCredentials() bool {
	f, err := os.Open("cre/cre")
	if err != nil {
		return false
	}
	r := bufio.NewReader(f)
	_, _, err = r.ReadLine()
	if err != nil {
		return false
	}
	line, _, err := r.ReadLine()
	if err != nil {
		return false
	}
	i := strings.Index(string(line), "=")
	s3fs.Id = string(line)[i+1:]

	line, _, err = r.ReadLine()
	if err != nil {
		return false
	}
	i = strings.Index(string(line), "=")
	s3fs.Secret = string(line)[i+1:]

	line, _, err = r.ReadLine()
	if err != nil {
		return false
	}
	i = strings.Index(string(line), "=")
	s3fs.Token = string(line)[i+1:]
	return true
}
