# SecureGov Solutions - Government & DoD IT Services Website

A professional IT services website built with TypeScript/Svelte frontend and Go backend, designed specifically for government and Department of Defense clients.

## 🏗️ Architecture

### Frontend
- **Framework**: Svelte 5 with TypeScript
- **Build Tool**: Vite
- **Styling**: Custom CSS with CSS variables
- **Features**:
  - Responsive design
  - Smooth scrolling navigation
  - Dynamic content loading from API
  - Contact form with validation
  - Mobile-friendly interface

### Backend
- **Language**: Go 1.21+
- **Router**: Gorilla Mux
- **Features**:
  - RESTful API endpoints
  - CORS enabled
  - Service catalog management
  - Contact form handling
  - Health check endpoint

## 📋 Prerequisites

- **Node.js** 18+ and npm
- **Go** 1.21+
- **Git**

## 🚀 Getting Started

### Backend Setup

1. Navigate to the backend directory:
```bash
cd backend
```

2. Install Go dependencies:
```bash
go mod download
```

3. Run the Go server:
```bash
go run main.go
```

The backend API will start on `http://localhost:8080`

### Frontend Setup

1. Navigate to the frontend directory:
```bash
cd frontend
```

2. Install npm dependencies:
```bash
npm install
```

3. Start the development server:
```bash
npm run dev
```

The frontend will start on `http://localhost:5173`

## 🔌 API Endpoints

### GET /api/services
Returns list of IT services offered
```json
[
  {
    "id": "cloud-infrastructure",
    "title": "Secure Cloud Infrastructure",
    "description": "FedRAMP authorized cloud solutions...",
    "icon": "cloud",
    "features": ["AWS GovCloud", "FedRAMP High", ...]
  }
]
```

### GET /api/capabilities
Returns company capabilities and statistics
```json
[
  {
    "id": "security-clearance",
    "title": "Security Cleared Workforce",
    "description": "Our team maintains active security clearances...",
    "stats": [
      {"value": "200+", "label": "Cleared Professionals"}
    ]
  }
]
```

### POST /api/contact
Submit contact form
```json
{
  "name": "John Doe",
  "email": "john@agency.gov",
  "organization": "Department of Defense",
  "clearance": "secret",
  "message": "We need cloud migration services..."
}
```

### GET /health
Health check endpoint
```json
{
  "status": "healthy"
}
```

## 🎨 Features

### Hero Section
- Eye-catching gradient background
- Key statistics display
- Clear call-to-action buttons
- Responsive design

### Services Section
- 6 core service offerings:
  - Secure Cloud Infrastructure
  - Cybersecurity & ATO Services
  - DevSecOps & CI/CD
  - Data Analytics & AI/ML
  - Application Modernization
  - Managed IT Services
- Detailed feature lists
- Hover animations

### Capabilities Section
- Company strengths and differentiators
- Statistical proof points
- Certification badges (FedRAMP, CMMC, ISO, etc.)
- Compliance framework display

### Contact Section
- Professional contact form
- Security clearance level selection
- Form validation
- Success/error messaging
- Company contact information

### Footer
- Quick navigation links
- Company information
- Legal links
- Back to top functionality

## 🔒 Security Features

The website emphasizes government-specific security requirements:
- FedRAMP authorization
- CMMC Level 3 compliance
- NIST 800-53 adherence
- DoD Impact Level 5 support
- ISO 27001 certification
- SOC 2 Type II compliance

## 🎯 Government Focus

Tailored specifically for:
- Federal agencies
- Department of Defense
- State and local government
- Cleared programs
- Classified environments

## 📱 Responsive Design

- Mobile-first approach
- Breakpoints for tablets and desktop
- Touch-friendly interface
- Optimized performance

## 🛠️ Development

### Build for Production

Frontend:
```bash
cd frontend
npm run build
```

Backend:
```bash
cd backend
go build -o server main.go
```

### Code Quality

Frontend:
```bash
npm run check  # Type checking
```

Backend:
```bash
go fmt ./...   # Format code
go vet ./...   # Vet code
```

## 📦 Project Structure

```
gov-it-services/
├── backend/
│   ├── main.go           # Main server file
│   └── go.mod           # Go dependencies
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   │   ├── Header.svelte
│   │   │   ├── Hero.svelte
│   │   │   ├── Services.svelte
│   │   │   ├── Capabilities.svelte
│   │   │   ├── Contact.svelte
│   │   │   └── Footer.svelte
│   │   ├── App.svelte
│   │   ├── main.ts
│   │   └── app.css
│   ├── index.html
│   ├── package.json
│   ├── tsconfig.json
│   ├── vite.config.ts
│   └── svelte.config.js
└── README.md
```

## 🌟 Key Differentiators

Unlike typical IT services websites, this site specifically targets:
- Government compliance requirements
- Security clearance workforce
- Mission-critical operations
- Federal acquisition processes
- DoD-specific certifications

## 📝 Customization

To customize for your organization:

1. Update branding in `Header.svelte` and `Footer.svelte`
2. Modify services in `backend/main.go`
3. Adjust color scheme in `app.css` (CSS variables)
4. Update contact information in `Contact.svelte`
5. Add your certifications in `Capabilities.svelte`

## 🤝 Contributing

This is a template project. Feel free to:
- Add new services
- Enhance the API
- Improve styling
- Add authentication
- Integrate with CMS

## 📄 License

This project is provided as a template for government IT services websites.

## 🔗 Inspired By

Design patterns inspired by TekSystems and other leading government IT services providers, with a focus on:
- Professional appearance
- Clear value proposition
- Trust indicators
- Easy navigation
- Government-specific messaging

## 📞 Support

For questions or issues, please refer to the documentation or contact your development team.