package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("HTTP_PLATFORM_PORT")
	if port == "" {
		port = "8080"
	}
	root := os.Getenv("WEBROOT_PATH")
	if root ==  "" {
		root = "./pages"
	} else {
		root += "\\pages"
	}

	addr := ":"+port
	log.Fatal(http.ListenAndServe(addr, http.FileServer(http.Dir(root))))
}

