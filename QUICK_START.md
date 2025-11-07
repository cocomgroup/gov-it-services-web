# Quick Setup Guide

## 🎯 What You Have

A complete, production-ready IT services website built for Government & DoD clients with:
- **Frontend**: TypeScript + Svelte (modern, reactive UI)
- **Backend**: Go REST API (fast, efficient server)
- **Design**: Inspired by TekSystems, tailored for government sector

## 📁 Project Structure

```
gov-it-services/
├── backend/          # Go API server
│   ├── main.go       # Main server with REST endpoints
│   └── go.mod        # Go dependencies
├── frontend/         # Svelte TypeScript app
│   ├── src/
│   │   ├── components/  # All UI components
│   │   ├── App.svelte   # Main app
│   │   ├── main.ts      # Entry point
│   │   └── app.css      # Global styles
│   ├── index.html
│   └── package.json
├── PREVIEW.html      # Standalone HTML preview (open this first!)
├── start-dev.sh      # Quick start script
└── README.md         # Full documentation
```

## ⚡ Quick Start (3 Steps)

### Option 1: Preview (No Installation)
1. Open `PREVIEW.html` in your browser to see the design

### Option 2: Full Development Setup

**Prerequisites**: Install Node.js 18+ and Go 1.21+

**Step 1 - Start Backend:**
```bash
cd backend
go mod download
go run main.go
# Server runs on http://localhost:8080
```

**Step 2 - Start Frontend (new terminal):**
```bash
cd frontend
npm install
npm run dev
# App runs on http://localhost:5173
```

**Step 3 - Open Browser:**
Visit http://localhost:5173

### Option 3: Use Quick Start Script
```bash
chmod +x start-dev.sh
./start-dev.sh
```

## 🎨 Key Features

### Services Section
- 6 government-focused IT services
- Cloud Infrastructure (FedRAMP, AWS GovCloud)
- Cybersecurity & ATO
- DevSecOps
- Data Analytics & AI
- Application Modernization
- Managed Services

### Professional Design
- Responsive (mobile, tablet, desktop)
- Government color scheme (navy blue, professional)
- Trust indicators (certifications, clearances)
- Smooth animations and interactions

### API Endpoints
- `GET /api/services` - List all services
- `GET /api/capabilities` - Company capabilities
- `POST /api/contact` - Submit contact form
- `GET /health` - Health check

## 🔧 Customization

### Change Company Name/Branding
Edit these files:
- `frontend/src/components/Header.svelte` (logo & name)
- `frontend/src/components/Footer.svelte` (footer info)
- `frontend/index.html` (page title)

### Modify Services
Edit: `backend/main.go` (services array)

### Change Colors
Edit: `frontend/src/app.css` (CSS variables at top)

### Update Contact Info
Edit: `frontend/src/components/Contact.svelte`

## 📊 What Makes This Government-Focused

Unlike general IT websites, this includes:
- ✅ Security clearance levels in contact form
- ✅ FedRAMP, CMMC, NIST certifications prominent
- ✅ DoD Impact Level references
- ✅ Government-specific service descriptions
- ✅ Professional, trust-building design
- ✅ Compliance-focused messaging

## 🚀 Production Deployment

### Build Frontend
```bash
cd frontend
npm run build
# Output in frontend/dist/
```

### Build Backend
```bash
cd backend
go build -o server main.go
# Run with: ./server
```

### Deploy Options
- Static hosting: Netlify, Vercel (frontend)
- Cloud platforms: AWS, Google Cloud, Azure (full stack)
- Government clouds: AWS GovCloud, Azure Government

## 📞 Need Help?

- Check `README.md` for full documentation
- Review component files for implementation details
- Inspect `PREVIEW.html` for visual reference

## 🎯 Next Steps

1. ✅ Preview the design (PREVIEW.html)
2. ✅ Run the development environment
3. ✅ Customize branding and content
4. ✅ Add your real contact information
5. ✅ Update certifications to match your company
6. ✅ Test the contact form
7. ✅ Deploy to production

---

**Built with modern tech for government IT services** 🇺🇸
TypeScript • Svelte • Go • REST API