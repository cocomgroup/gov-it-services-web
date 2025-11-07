<script lang="ts">
  function scrollToContact() {
    const element = document.getElementById('contact');
    if (element) {
      element.scrollIntoView({ behavior: 'smooth' });
    }
  }
</script>

<section class="hero">
  <div class="container">
    <div class="hero-content">
      <h1>Trusted IT Solutions for Government & Defense</h1>
      <p class="subtitle">
        Delivering mission-critical technology services with FedRAMP authorization, 
        cleared workforce, and decades of experience supporting federal agencies and DoD.
      </p>
      <div class="hero-stats">
        <div class="stat">
          <div class="stat-value">15+</div>
          <div class="stat-label">Years Serving Government</div>
        </div>
        <div class="stat">
          <div class="stat-value">200+</div>
          <div class="stat-label">Cleared Professionals</div>
        </div>
        <div class="stat">
          <div class="stat-value">50+</div>
          <div class="stat-label">Active Contracts</div>
        </div>
      </div>
      <div class="hero-actions">
        <button class="btn-primary" on:click={scrollToContact}>Schedule Consultation</button>
        <button class="btn-secondary" on:click={() => document.getElementById('services')?.scrollIntoView({ behavior: 'smooth' })}>
          Our Services
        </button>
      </div>
    </div>
  </div>
</section>

<style>
  .hero {
    background: linear-gradient(135deg, #1e3a8a 0%, #3b82f6 100%);
    color: white;
    padding: 6rem 0;
    position: relative;
    overflow: hidden;
  }

  .hero::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: url('data:image/svg+xml,<svg width="100" height="100" xmlns="http://www.w3.org/2000/svg"><defs><pattern id="grid" width="40" height="40" patternUnits="userSpaceOnUse"><path d="M 40 0 L 0 0 0 40" fill="none" stroke="rgba(255,255,255,0.1)" stroke-width="1"/></pattern></defs><rect width="100" height="100" fill="url(%23grid)"/></svg>');
    opacity: 0.3;
  }

  .container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 2rem;
    position: relative;
    z-index: 1;
  }

  .hero-content {
    max-width: 900px;
    margin: 0 auto;
    text-align: center;
  }

  h1 {
    font-size: 3.5rem;
    margin-bottom: 1.5rem;
    line-height: 1.2;
  }

  .subtitle {
    font-size: 1.25rem;
    margin-bottom: 3rem;
    opacity: 0.95;
    line-height: 1.7;
  }

  .hero-stats {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: 2rem;
    margin-bottom: 3rem;
    padding: 2rem 0;
  }

  .stat {
    text-align: center;
  }

  .stat-value {
    font-size: 3rem;
    font-weight: 700;
    margin-bottom: 0.5rem;
  }

  .stat-label {
    font-size: 1rem;
    opacity: 0.9;
  }

  .hero-actions {
    display: flex;
    gap: 1rem;
    justify-content: center;
    flex-wrap: wrap;
  }

  .btn-primary, .btn-secondary {
    padding: 1rem 2rem;
    font-size: 1.1rem;
    font-weight: 600;
    border-radius: 8px;
    border: 2px solid white;
    cursor: pointer;
    transition: all 0.3s;
  }

  .btn-primary {
    background: white;
    color: var(--primary-color);
  }

  .btn-primary:hover {
    background: transparent;
    color: white;
    transform: translateY(-2px);
  }

  .btn-secondary {
    background: transparent;
    color: white;
  }

  .btn-secondary:hover {
    background: white;
    color: var(--primary-color);
    transform: translateY(-2px);
  }

  @media (max-width: 768px) {
    .hero {
      padding: 4rem 0;
    }

    h1 {
      font-size: 2.5rem;
    }

    .subtitle {
      font-size: 1.1rem;
    }

    .stat-value {
      font-size: 2.5rem;
    }

    .hero-actions {
      flex-direction: column;
    }

    .btn-primary, .btn-secondary {
      width: 100%;
    }
  }
</style>