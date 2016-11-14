package main

import (
	"flag"
	"fmt"
	"html/template"
	_ "image/gif"
	_ "image/jpeg"
	"log"
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/ledyba/IdPhotoMaker/photo"
)

var port = flag.Int("port", 8080, "")

func assumeFile(name string) string {
	bytes, _ := Asset(name)
	return string(bytes)
}

func render(templateName string, dat interface{}, w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.New("").Parse(assumeFile("templates/main.html")))
	templates, _ = templates.Parse(assumeFile(fmt.Sprintf("templates/%s.html", templateName)))

	err := templates.ExecuteTemplate(w, "base", dat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	flag.Parse() // Scan the arguments list
	photo.InitCache()
	http.HandleFunc("/", mainHandler)
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(
			&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "static"})))
	http.HandleFunc("/upload", uploadHandler)

	log.Printf("Start at http://localhost:%d/", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
