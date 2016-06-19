package main

import (
	"log"
	"os"
	"path/filepath"

	"../../net"
	"../../redlot"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Get work dir error: %v\n", err.Error())
	}
	dataPath := filepath.Join(pwd, "var/db")
	options := &redlot.Options{
		DataPath: dataPath,
	}

	net.Serve(":9999", options)
}
