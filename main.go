package main

import "iot/server"

func main() {
	//s3fs.Start()
	server.Setup(false)
	server.Start()
}
