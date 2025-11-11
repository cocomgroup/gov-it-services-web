<script lang="ts">
  import { onMount } from 'svelte';

  interface Service {
    id: string;
    title: string;
    description: string;
    icon: string;
    features: string[];
  }

  let services: Service[] = $state([]);
  let loading = $state(true);
  let error = $state('');

  const iconMap: Record<string, string> = {
    cloud: '☁️',
    shield: '🛡️',
    code: '💻',
    chart: '📊',
    refresh: '🔄',
    settings: '⚙️'
  };

  onMount(async () => {
    try {
      const response = await fetch('/api/services');
      if (!response.ok) throw new Error('Failed to fetch services');
      services = await response.json();
    } catch (e) {
      error = 'Failed to load services';
      console.error(e);
    } finally {
      loading = false;
    }
  });
</script>

<section id="services" class="services">
  <div class="container">
    <div class="section-header">
      <h2>Mission-Critical IT Services</h2>
      <p>Comprehensive technology solutions designed specifically for government and defense requirements</p>
    </div>

    {#if loading}
      <div class="loading">Loading services...</div>
    {:else if error}
      <div class="error">{error}</div>
    {:else}
      <div class="services-grid">
        {#each services as service}
          <div class="service-card">
            <div class="service-icon">{iconMap[service.icon] || '📋'}</div>
            <h3>{service.title}</h3>
            <p class="service-description">{service.description}</p>
            <ul class="features-list">
              {#each service.features as feature}
                <li>{feature}</li>
              {/each}
            </ul>
          </div>
        {/each}
      </div>
    {/if}
  </div>
</section>

<style>
  .services {
    padding: 6rem 0;
    background: var(--bg-light);
  }

  .container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 2rem;
  }

  .section-header {
    text-align: center;
    margin-bottom: 4rem;
  }

  .section-header h2 {
    color: var(--primary-color);
    margin-bottom: 1rem;
  }

  .section-header p {
    font-size: 1.2rem;
    color: var(--text-secondary);
    max-width: 700px;
    margin: 0 auto;
  }

  .services-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
    gap: 2rem;
  }

  .service-card {
    background: white;
    padding: 2rem;
    border-radius: 12px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    transition: transform 0.3s, box-shadow 0.3s;
  }

  .service-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 8px 16px rgba(0, 0, 0, 0.15);
  }

  .service-icon {
    font-size: 3rem;
    margin-bottom: 1rem;
  }

  .service-card h3 {
    color: var(--primary-color);
    margin-bottom: 1rem;
    font-size: 1.5rem;
  }

  .service-description {
    color: var(--text-secondary);
    margin-bottom: 1.5rem;
    line-height: 1.6;
  }

  .features-list {
    list-style: none;
    padding: 0;
  }

  .features-list li {
    padding: 0.5rem 0;
    padding-left: 1.5rem;
    position: relative;
    color: var(--text-primary);
  }

  .features-list li::before {
    content: '✓';
    position: absolute;
    left: 0;
    color: var(--secondary-color);
    font-weight: bold;
  }

  .loading, .error {
    text-align: center;
    padding: 2rem;
    font-size: 1.2rem;
  }

  .error {
    color: var(--accent-color);
  }

  @media (max-width: 768px) {
    .services {
      padding: 4rem 0;
    }

    .services-grid {
      grid-template-columns: 1fr;
    }
  }
</style>