package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// TemplateData structure for data passed to templates
type TemplateData struct {
	PageTitle string
	Projects  []Project  // For future project display
	BlogPosts []BlogPost // For future blog post display
}

// Dummy structures for demonstration (will be replaced with actual data)
type Project struct {
	Name        string
	Description string
	DemoURL     string
	RepoURL     string
}

type BlogPost struct {
	Title   string
	Content template.HTML
	Slug    string
}

// Parse templates here; ensure the path is correct
// Make sure all HTML files exist in the 'templates' directory
var templates = template.Must(template.ParseFiles(
	filepath.Join("templates", "index.html"),
	filepath.Join("templates", "about.html"),
	filepath.Join("templates", "projects.html"),
	filepath.Join("templates", "blog.html"),
	filepath.Join("templates", "demo_pii_discovery.html"), // Example for GDPR projects
	filepath.Join("templates", "demo_synthetic_data.html"),
	filepath.Join("templates", "demo_consent_manager.html"),
	filepath.Join("templates", "demo_risk_assessment.html"),
))

func renderTemplate(w http.ResponseWriter, tmpl string, data TemplateData) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		log.Printf("Error rendering template %s: %v", tmpl, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{PageTitle: "Andre Kuzma - GoLang Portfolio"}
	renderTemplate(w, "index.html", data)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{PageTitle: "About Me"}
	renderTemplate(w, "about.html", data)
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	// Project data would be loaded from a future kuzma-dev-blog-content repository or database
	projects := []Project{
		{Name: "GDPR Data Discovery", Description: "Intelligent tool for discovering and cataloging PII/SPI data.", DemoURL: "/demo/data-discovery", RepoURL: "https://github.com/Kuzma-Dev/go-gdpr-data-discovery-service"},
		{Name: "Synthetic Data Generator", Description: "Tool for generating synthetic data for testing.", DemoURL: "/demo/synthetic-data", RepoURL: "https://github.com/Kuzma-Dev/go-gdpr-synthetic-data-generator-service"},
		{Name: "GDPR Consent Manager", Description: "Manage user consents for GDPR compliance.", DemoURL: "/demo/consent-manager", RepoURL: "https://github.com/Kuzma-Dev/go-gdpr-consent-manager-service"},
		{Name: "GDPR Risk Assessment", Description: "Tool for assessing data re-identification risk.", DemoURL: "/demo/risk-assessment", RepoURL: "https://github.com/Kuzma-Dev/go-gdpr-risk-assessment-service"},
		// More projects here
	}
	data := TemplateData{PageTitle: "My Projects", Projects: projects}
	renderTemplate(w, "projects.html", data)
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	// Blog posts would be loaded from a future kuzma-dev-blog-content repository
	blogPosts := []BlogPost{
		{Title: "Welcome to my GoLang blog!", Content: template.HTML("First post about my GoLang portfolio."), Slug: "welcome-to-golang-blog"},
	}
	data := TemplateData{PageTitle: "My Blog", BlogPosts: blogPosts}
	renderTemplate(w, "blog.html", data)
}

// Demo Handlers for GDPR projects
func demoPIIDiscoveryHandler(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{PageTitle: "Demo: PII Discovery"}
	renderTemplate(w, "demo_pii_discovery.html", data)
}

func demoSyntheticDataHandler(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{PageTitle: "Demo: Synthetic Data Generator"}
	renderTemplate(w, "demo_synthetic_data.html", data)
}

func demoConsentManagerHandler(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{PageTitle: "Demo: Consent Manager"}
	renderTemplate(w, "demo_consent_manager.html", data)
}

func demoRiskAssessmentHandler(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{PageTitle: "Demo: Risk Assessment"}
	renderTemplate(w, "demo_risk_assessment.html", data)
}

func main() {
	// Set up file paths
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Portfolio routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/projects", projectsHandler)
	http.HandleFunc("/blog", blogHandler)

	// GDPR demo routes
	http.HandleFunc("/demo/data-discovery", demoPIIDiscoveryHandler)
	http.HandleFunc("/demo/synthetic-data", demoSyntheticDataHandler)
	http.HandleFunc("/demo/consent-manager", demoConsentManagerHandler)
	http.HandleFunc("/demo/risk-assessment", demoRiskAssessmentHandler)

	log.Println("GoLang Portfolio running on http://0.0.0.0:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
