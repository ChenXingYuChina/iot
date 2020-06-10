package main

import (
	"iot/s3fs"
	"log"
)

func main() {
	s3fs.Start()
	_, err := s3fs.ListFiles("", "", 2)
	if err != nil {
		log.Println(err)
		return
	}
}
