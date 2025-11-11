package main

import (
	"encoding/json"
	"os"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"database/sql"
)

type Service struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Icon        string   `json:"icon"`
	Features    []string `json:"features"`
}

type ContactForm struct {
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Organization string    `json:"organization"`
	Clearance    string    `json:"clearance"`
	Message      string    `json:"message"`
	Timestamp    time.Time `json:"timestamp"`
}

type Capability struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Stats       []Stat `json:"stats"`
}

type Stat struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

var db *sql.DB

var services = []Service{
	{
		ID:          "cloud-infrastructure",
		Title:       "Secure Cloud Infrastructure",
		Description: "FedRAMP authorized cloud solutions with IL4/IL5 compliance for sensitive government data.",
		Icon:        "cloud",
		Features: []string{
			"AWS GovCloud & Azure Government",
			"FedRAMP High & DoD Impact Level 5",
			"Continuous monitoring & compliance",
			"Zero-trust architecture implementation",
		},
	},
	{
		ID:          "cybersecurity",
		Title:       "Cybersecurity & ATO Services",
		Description: "Comprehensive security solutions ensuring NIST 800-53 compliance and accelerated ATO processes.",
		Icon:        "shield",
		Features: []string{
			"NIST 800-53 & CMMC compliance",
			"Security assessment & authorization",
			"Penetration testing & vulnerability management",
			"24/7 security operations center",
		},
	},
	{
		ID:          "devops",
		Title:       "DevSecOps & CI/CD",
		Description: "Automated, secure software delivery pipelines that meet government security requirements.",
		Icon:        "code",
		Features: []string{
			"Automated security scanning & compliance checks",
			"Container orchestration (Kubernetes)",
			"Infrastructure as Code (Terraform, Ansible)",
			"Continuous integration & deployment",
		},
	},
	{
		ID:          "data-analytics",
		Title:       "Data Analytics & AI/ML",
		Description: "Advanced analytics and machine learning solutions for mission-critical decision making.",
		Icon:        "chart",
		Features: []string{
			"Big data processing & analytics",
			"Machine learning model development",
			"Predictive analytics for operations",
			"Data visualization & dashboards",
		},
	},
	{
		ID:          "application-modernization",
		Title:       "Application Modernization",
		Description: "Transform legacy systems into modern, cloud-native applications with enhanced security.",
		Icon:        "refresh",
		Features: []string{
			"Legacy system assessment & migration",
			"Microservices architecture",
			"API development & integration",
			"Mobile application development",
		},
	},
	{
		ID:          "managed-services",
		Title:       "Managed IT Services",
		Description: "Comprehensive IT operations support with cleared personnel and 24/7 availability.",
		Icon:        "settings",
		Features: []string{
			"24/7 managed operations",
			"Cleared personnel (Secret/TS/SCI)",
			"IT service management (ITIL)",
			"Help desk & user support",
		},
	},
}

var capabilities = []Capability{
	{
		ID:          "security-clearance",
		Title:       "Security Cleared Workforce",
		Description: "Our team maintains active security clearances to support classified programs.",
		Stats: []Stat{
			{Value: "200+", Label: "Cleared Professionals"},
			{Value: "98%", Label: "Clearance Success Rate"},
		},
	},
	{
		ID:          "compliance",
		Title:       "Compliance & Certifications",
		Description: "Certified and compliant with all major government security frameworks.",
		Stats: []Stat{
			{Value: "FedRAMP", Label: "Authorized"},
			{Value: "CMMC L3", Label: "Certified"},
		},
	},
	{
		ID:          "experience",
		Title:       "Mission-Focused Experience",
		Description: "Decades of experience supporting critical government missions.",
		Stats: []Stat{
			{Value: "15+ Years", Label: "Government Focus"},
			{Value: "50+", Label: "Active Contracts"},
		},
	},
}
// Authentication middleware
//func authMiddleware(next http.Handler) http.Handler { ... }

// Database operations
//func getServicesFromDB() ([]Service, error) { ... }
//func saveContactToDB(contact ContactForm) error { ... }

// Environment configuration
var (
    port = os.Getenv("PORT")
    dbURL = os.Getenv("DATABASE_URL")
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getServices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func getCapabilities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(capabilities)
}

func submitContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var contact ContactForm
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contact.Timestamp = time.Now()
	log.Printf("Contact form submitted: %+v", contact)

	response := map[string]string{
		"status":  "success",
		"message": "Thank you for your inquiry. Our team will contact you within 24 hours.",
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/services", getServices).Methods("GET")
	api.HandleFunc("/capabilities", getCapabilities).Methods("GET")
	api.HandleFunc("/contact", submitContact).Methods("POST")

	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
	}).Methods("GET")

	handler := enableCORS(r)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}