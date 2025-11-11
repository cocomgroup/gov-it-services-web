<script lang="ts">
  const currentYear = new Date().getFullYear();

  function scrollToTop() {
    window.scrollTo({ top: 0, behavior: 'smooth' });
  }
</script>

<footer class="footer">
  <div class="container">
    <div class="footer-content">
      <div class="footer-section">
        <div class="logo">
          <svg width="40" height="40" viewBox="0 0 40 40" fill="none">
            <rect width="40" height="40" rx="8" fill="#1e3a8a"/>
            <path d="M12 20L18 26L28 14" stroke="white" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
          </svg>
          <span>SecureGov Solutions</span>
        </div>
        <p class="company-description">
          Delivering mission-critical IT solutions to government and defense agencies 
          with unmatched expertise, security, and compliance.
        </p>
        <div class="certifications-mini">
          <span>FedRAMP</span>
          <span>CMMC L3</span>
          <span>ISO 27001</span>
        </div>
      </div>

      <div class="footer-section">
        <h4>Services</h4>
        <ul>
          <li><a href="#services">Cloud Infrastructure</a></li>
          <li><a href="#services">Cybersecurity & ATO</a></li>
          <li><a href="#services">DevSecOps</a></li>
          <li><a href="#services">Data Analytics & AI</a></li>
          <li><a href="#services">Application Modernization</a></li>
        </ul>
      </div>

      <div class="footer-section">
        <h4>Company</h4>
        <ul>
          <li><a href="#capabilities">About Us</a></li>
          <li><a href="#capabilities">Capabilities</a></li>
          <li><a href="#contact">Contact</a></li>
          <li><a href="#careers">Careers</a></li>
          <li><a href="#partners">Partners</a></li>
        </ul>
      </div>

      <div class="footer-section">
        <h4>Resources</h4>
        <ul>
          <li><a href="#resources">Case Studies</a></li>
          <li><a href="#resources">White Papers</a></li>
          <li><a href="#resources">Security Compliance</a></li>
          <li><a href="#resources">Blog</a></li>
          <li><a href="#resources">Documentation</a></li>
        </ul>
      </div>
    </div>

    <div class="footer-bottom">
      <div class="footer-legal">
        <p>&copy; {currentYear} SecureGov Solutions. All rights reserved.</p>
        <div class="legal-links">
          <a href="#privacy">Privacy Policy</a>
          <span>|</span>
          <a href="#terms">Terms of Service</a>
          <span>|</span>
          <a href="#security">Security</a>
        </div>
      </div>
      <button class="back-to-top" onclick={scrollToTop} aria-label="Back to top">
        ↑ Back to Top
      </button>
    </div>
  </div>
</footer>

<style>
  .footer {
    background: #0f172a;
    color: white;
    padding: 4rem 0 2rem;
  }

  .container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 2rem;
  }

  .footer-content {
    display: grid;
    grid-template-columns: 2fr 1fr 1fr 1fr;
    gap: 3rem;
    margin-bottom: 3rem;
  }

  .footer-section h4 {
    color: white;
    margin-bottom: 1.5rem;
    font-size: 1.2rem;
  }

  .logo {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    font-weight: 600;
    font-size: 1.25rem;
    margin-bottom: 1rem;
  }

  .company-description {
    opacity: 0.8;
    line-height: 1.6;
    margin-bottom: 1.5rem;
  }

  .certifications-mini {
    display: flex;
    gap: 0.75rem;
    flex-wrap: wrap;
  }

  .certifications-mini span {
    background: rgba(59, 130, 246, 0.2);
    padding: 0.4rem 0.8rem;
    border-radius: 4px;
    font-size: 0.85rem;
    font-weight: 500;
  }

  .footer-section ul {
    list-style: none;
    padding: 0;
  }

  .footer-section ul li {
    margin-bottom: 0.75rem;
  }

  .footer-section a {
    color: rgba(255, 255, 255, 0.8);
    text-decoration: none;
    transition: color 0.3s;
  }

  .footer-section a:hover {
    color: #3b82f6;
  }

  .footer-bottom {
    border-top: 1px solid rgba(255, 255, 255, 0.1);
    padding-top: 2rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .footer-legal {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  .footer-legal p {
    opacity: 0.7;
    margin: 0;
  }

  .legal-links {
    display: flex;
    gap: 0.75rem;
    align-items: center;
    font-size: 0.9rem;
  }

  .legal-links span {
    opacity: 0.5;
  }

  .back-to-top {
    background: var(--secondary-color);
    color: white;
    border: none;
    padding: 0.75rem 1.5rem;
    border-radius: 6px;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.3s;
  }

  .back-to-top:hover {
    background: var(--primary-color);
    transform: translateY(-2px);
  }

  @media (max-width: 968px) {
    .footer-content {
      grid-template-columns: 1fr 1fr;
      gap: 2rem;
    }
  }

  @media (max-width: 640px) {
    .footer-content {
      grid-template-columns: 1fr;
    }

    .footer-bottom {
      flex-direction: column;
      gap: 1.5rem;
      text-align: center;
    }

    .footer-legal {
      align-items: center;
    }
  }
</style>