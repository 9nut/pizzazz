package main

import (
	"log"
	"net/http"
	"os"
	"html/template"
)

const homePage = `
<html>
  <head>
    <title>Authentic with pizzazz!</title>
  </head>
  <body>
    <h1>Hello {{.}}!</h1>
    <p><i>You are authentic.</i></p>
  </body>
</html>
`

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

	tmpl := template.Must(template.New("homepage").Parse(homePage))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name, ok := r.Header["X-MS-CLIENT-PRINCIPAL-NAME"]
		if !ok {
			name = []string{"Unknown"}
		}
		tmpl.Execute(w, name[0])
	})
	log.Fatal(http.ListenAndServe(addr, nil))
}

