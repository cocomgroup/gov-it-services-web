<script lang="ts">
  let mobileMenuOpen = $state(false);

  function toggleMenu() {
    mobileMenuOpen = !mobileMenuOpen;
  }

  function scrollToSection(sectionId: string) {
    const element = document.getElementById(sectionId);
    if (element) {
      element.scrollIntoView({ behavior: 'smooth' });
      mobileMenuOpen = false;
    }
  }
</script>

<header>
  <nav class="navbar">
    <div class="container">
      <div class="logo">
        <svg width="40" height="40" viewBox="0 0 40 40" fill="none">
          <rect width="40" height="40" rx="8" fill="#1e3a8a"/>
          <path d="M12 20L18 26L28 14" stroke="white" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
        </svg>
        <span>SecureGov Solutions</span>
      </div>

      <button class="mobile-menu-btn" onclick={toggleMenu} aria-label="Toggle menu">
        <svg width="24" height="24" fill="none" stroke="currentColor">
          {#if mobileMenuOpen}
            <path d="M6 18L18 6M6 6l12 12" stroke-width="2" stroke-linecap="round"/>
          {:else}
            <path d="M3 12h18M3 6h18M3 18h18" stroke-width="2" stroke-linecap="round"/>
          {/if}
        </svg>
      </button>

      <ul class="nav-menu" class:active={mobileMenuOpen}>
        <li><button onclick={() => scrollToSection('services')}>Services</button></li>
        <li><button onclick={() => scrollToSection('capabilities')}>Capabilities</button></li>
        <li><button onclick={() => scrollToSection('contact')}>Contact</button></li>
        <li><button class="cta-btn" onclick={() => scrollToSection('contact')}>Request Consultation</button></li>
      </ul>
    </div>
  </nav>
</header>

<style>
  header {
    background: white;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    position: sticky;
    top: 0;
    z-index: 100;
  }

  .navbar {
    padding: 1rem 0;
  }

  .container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 2rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .logo {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    font-weight: 600;
    font-size: 1.25rem;
    color: var(--primary-color);
  }

  .mobile-menu-btn {
    display: none;
    background: none;
    border: none;
    color: var(--primary-color);
    padding: 0.5rem;
  }

  .nav-menu {
    display: flex;
    list-style: none;
    gap: 2rem;
    align-items: center;
  }

  .nav-menu button {
    background: none;
    border: none;
    color: var(--text-primary);
    font-size: 1rem;
    padding: 0.5rem;
    transition: color 0.3s;
  }

  .nav-menu button:hover {
    color: var(--secondary-color);
  }

  .cta-btn {
    background: var(--secondary-color) !important;
    color: white !important;
    padding: 0.75rem 1.5rem !important;
    border-radius: 6px;
    font-weight: 500;
  }

  .cta-btn:hover {
    background: var(--primary-color) !important;
  }

  @media (max-width: 768px) {
    .mobile-menu-btn {
      display: block;
    }

    .nav-menu {
      position: fixed;
      left: -100%;
      top: 73px;
      flex-direction: column;
      background-color: white;
      width: 100%;
      text-align: center;
      transition: 0.3s;
      box-shadow: 0 10px 27px rgba(0, 0, 0, 0.05);
      padding: 2rem 0;
      gap: 1rem;
    }

    .nav-menu.active {
      left: 0;
    }
  }
</style>