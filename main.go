package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var tmpl *template.Template

func init() {
	var err error
	tmpl, err = template.New("index").Parse(indexTemplate)
	if err != nil {
		log.Fatalf("Failed to parse index template: %v", err)
	}
	tmpl, err = tmpl.New("navbar").Parse(navbarTemplate)
	if err != nil {
		log.Fatalf("Failed to parse navbar template: %v", err)
	}
	tmpl, err = tmpl.New("home").Parse(homeTemplate)
	if err != nil {
		log.Fatalf("Failed to parse home template: %v", err)
	}
	tmpl, err = tmpl.New("modals").Parse(modalsTemplate)
	if err != nil {
		log.Fatalf("Failed to parse modals template: %v", err)
	}
	tmpl, err = tmpl.New("swot").Parse(swotTemplate)
	if err != nil {
		log.Fatalf("Failed to parse swot template: %v", err)
	}
	tmpl, err = tmpl.New("bmc").Parse(bmcTemplate)
	if err != nil {
		log.Fatalf("Failed to parse bmc template: %v", err)
	}
	tmpl, err = tmpl.New("sr").Parse(srTemplate)
	if err != nil {
		log.Fatalf("Failed to parse sr template: %v", err)
	}
	tmpl, err = tmpl.New("contact").Parse(contactTemplate)
	if err != nil {
		log.Fatalf("Failed to parse contact template: %v", err)
	}
	tmpl, err = tmpl.New("footer").Parse(footerTemplate)
	if err != nil {
		log.Fatalf("Failed to parse footer template: %v", err)
	}
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Year int
	}{
		Year: time.Now().Year(),
	}

	err := tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		log.Printf("Error rendering template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	// Serve static files (CSS, JS, images, videos)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle root route
	http.HandleFunc("/", renderTemplate)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}