<script lang="ts">
  import { onMount } from 'svelte';

  interface Service {
    id: string;
    title: string;
    description: string;
    icon: string;
  }

  let services: Service[] = [];

  onMount(async () => {
    const response = await fetch('/api/services');
    services = await response.json();
  });
</script>

<div class="dashboard">
  <!-- Header -->
  <header class="header">
    <div class="logo">
      <svg width="30" height="30" viewBox="0 0 40 40" fill="none">
        <rect width="40" height="40" rx="8" fill="#1e3a8a"/>
        <path d="M12 20L18 26L28 14" stroke="white" stroke-width="3"/>
      </svg>
      <span>CoCom SecureGov Solutions</span>
    </div>
    <nav>
      <a href="#services">Services</a>
      <a href="#contact">Contact</a>
      <a href="#about">About Us</a>
    </nav>
  </header>

  <!-- Main Content Grid -->
  <main class="content">
    <!-- Hero Section (Top) -->
    <section class="hero-compact">
      <h1>Trusted IT Solutions for Government & Defense</h1>
      <div class="stats">
        <div class="stat">
          <span class="value">15+</span>
          <span class="label">Years</span>
        </div>
        <div class="stat">
          <span class="value">200+</span>
          <span class="label">Cleared</span>
        </div>
        <div class="stat">
          <span class="value">50+</span>
          <span class="label">Contracts</span>
        </div>
      </div>
    </section>

    <!-- Services Grid (Middle) -->
    <section class="services-compact">
      <h2>Our Services</h2>
      <div class="service-grid">
        {#each services.slice(0, 6) as service}
          <div class="service-card-mini">
            <h3>{service.title}</h3>
            <p>{service.description}</p>
          </div>
        {/each}
      </div>
    </section>

    <!-- Contact (Bottom) -->
    <aside class="contact-mini">
      <h3>Contact Us</h3>
      <div class="contact-content">
        <p>📞 (202) 555-0100</p>
        <p>✉️ contracts@securegovsolutions.com</p>
        <button>Request Consultation</button>
      </div>
    </aside>
  </main>
</div>

<style>
  .dashboard {
    height: 100vh;
    display: flex;
    flex-direction: column;
    background: #f9fafb;
  }

  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem 2rem;
    background: white;
    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  }

  .logo {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    font-weight: 600;
    color: #1e3a8a;
  }

  .header nav {
    display: flex;
    gap: 2rem;
  }

  .header a {
    color: #1f2937;
    text-decoration: none;
  }

  .header a:hover {
    color: #3b82f6;
  }

  .content {
    flex: 1;
    display: grid;
    grid-template-rows: auto 1fr auto;
    gap: 1rem;
    padding: 1rem;
    overflow: hidden;
  }

  .hero-compact {
    background: linear-gradient(135deg, #1e3a8a, #3b82f6);
    color: white;
    padding: 1.5rem 2rem;
    border-radius: 12px;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .hero-compact h1 {
    font-size: 1.75rem;
    margin: 0;
    flex: 1;
  }

  .stats {
    display: flex;
    gap: 2rem;
  }

  .stat {
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  .stat .value {
    font-size: 2rem;
    font-weight: 700;
  }

  .stat .label {
    font-size: 0.75rem;
    opacity: 0.9;
  }

  .services-compact {
    background: white;
    padding: 1.5rem;
    border-radius: 12px;
    overflow-y: auto;
  }

  .services-compact h2 {
    margin-bottom: 1rem;
    color: #1e3a8a;
    font-size: 1.5rem;
  }

  .service-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 1rem;
  }

  .service-card-mini {
    padding: 1rem;
    background: #f9fafb;
    border-radius: 8px;
    border: 1px solid #e5e7eb;
    transition: all 0.3s;
  }

  .service-card-mini:hover {
    border-color: #3b82f6;
    box-shadow: 0 4px 8px rgba(59, 130, 246, 0.1);
  }

  .service-card-mini h3 {
    font-size: 1rem;
    margin-bottom: 0.5rem;
    color: #1e3a8a;
  }

  .service-card-mini p {
    font-size: 0.85rem;
    color: #6b7280;
    line-height: 1.4;
  }

  .contact-mini {
    background: white;
    padding: 1rem 2rem;
    border-radius: 12px;
  }

  .contact-content {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 2rem;
  }

  .contact-mini h3 {
    color: #1e3a8a;
    margin-bottom: 0.75rem;
  }

  .contact-mini p {
    font-size: 0.9rem;
    margin: 0;
  }

  .contact-mini button {
    background: #3b82f6;
    color: white;
    border: none;
    padding: 0.75rem 1.5rem;
    border-radius: 6px;
    cursor: pointer;
    font-weight: 600;
    white-space: nowrap;
  }

  .contact-mini button:hover {
    background: #1e3a8a;
  }

  @media (max-width: 1200px) {
    .service-grid {
      grid-template-columns: repeat(2, 1fr);
    }
  }

  @media (max-width: 768px) {
    .hero-compact {
      flex-direction: column;
      gap: 1rem;
      text-align: center;
    }

    .service-grid {
      grid-template-columns: 1fr;
    }

    .contact-content {
      flex-direction: column;
      align-items: flex-start;
    }
  }
</style>