<script lang="ts">
  import { onMount } from 'svelte';

  interface Stat {
    value: string;
    label: string;
  }

  interface Capability {
    id: string;
    title: string;
    description: string;
    stats: Stat[];
  }

  let capabilities: Capability[] = [];
  let loading = true;
  let error = '';

  onMount(async () => {
    try {
      const response = await fetch('/api/capabilities');
      if (!response.ok) throw new Error('Failed to fetch capabilities');
      capabilities = await response.json();
    } catch (e) {
      error = 'Failed to load capabilities';
      console.error(e);
    } finally {
      loading = false;
    }
  });
</script>

<section id="capabilities" class="capabilities">
  <div class="container">
    <div class="section-header">
      <h2>Why Choose SecureGov Solutions</h2>
      <p>Trusted expertise and proven capabilities for mission-critical government IT services</p>
    </div>

    {#if loading}
      <div class="loading">Loading capabilities...</div>
    {:else if error}
      <div class="error">{error}</div>
    {:else}
      <div class="capabilities-grid">
        {#each capabilities as capability}
          <div class="capability-card">
            <h3>{capability.title}</h3>
            <p class="capability-description">{capability.description}</p>
            <div class="stats">
              {#each capability.stats as stat}
                <div class="stat-item">
                  <div class="stat-value">{stat.value}</div>
                  <div class="stat-label">{stat.label}</div>
                </div>
              {/each}
            </div>
          </div>
        {/each}
      </div>

      <div class="certifications">
        <h3>Certifications & Compliance</h3>
        <div class="cert-badges">
          <div class="badge">FedRAMP Authorized</div>
          <div class="badge">CMMC Level 3</div>
          <div class="badge">ISO 27001</div>
          <div class="badge">NIST 800-53</div>
          <div class="badge">DoD IL5</div>
          <div class="badge">SOC 2 Type II</div>
        </div>
      </div>
    {/if}
  </div>
</section>

<style>
  .capabilities {
    padding: 6rem 0;
    background: white;
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

  .capabilities-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
    gap: 2rem;
    margin-bottom: 4rem;
  }

  .capability-card {
    background: linear-gradient(135deg, #f9fafb 0%, #ffffff 100%);
    padding: 2.5rem;
    border-radius: 12px;
    border: 2px solid var(--border-color);
    transition: all 0.3s;
  }

  .capability-card:hover {
    border-color: var(--secondary-color);
    box-shadow: 0 8px 16px rgba(59, 130, 246, 0.1);
  }

  .capability-card h3 {
    color: var(--primary-color);
    margin-bottom: 1rem;
    font-size: 1.5rem;
  }

  .capability-description {
    color: var(--text-secondary);
    margin-bottom: 2rem;
    line-height: 1.6;
  }

  .stats {
    display: flex;
    gap: 2rem;
    justify-content: space-around;
  }

  .stat-item {
    text-align: center;
  }

  .stat-value {
    font-size: 2.5rem;
    font-weight: 700;
    color: var(--secondary-color);
    margin-bottom: 0.5rem;
  }

  .stat-label {
    font-size: 0.9rem;
    color: var(--text-secondary);
  }

  .certifications {
    text-align: center;
    padding: 3rem;
    background: var(--bg-light);
    border-radius: 12px;
  }

  .certifications h3 {
    color: var(--primary-color);
    margin-bottom: 2rem;
    font-size: 1.75rem;
  }

  .cert-badges {
    display: flex;
    flex-wrap: wrap;
    gap: 1rem;
    justify-content: center;
  }

  .badge {
    background: white;
    color: var(--primary-color);
    padding: 0.75rem 1.5rem;
    border-radius: 8px;
    font-weight: 600;
    border: 2px solid var(--secondary-color);
    transition: all 0.3s;
  }

  .badge:hover {
    background: var(--secondary-color);
    color: white;
    transform: translateY(-2px);
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
    .capabilities {
      padding: 4rem 0;
    }

    .capabilities-grid {
      grid-template-columns: 1fr;
    }

    .stats {
      flex-direction: column;
      gap: 1rem;
    }

    .cert-badges {
      gap: 0.75rem;
    }

    .badge {
      padding: 0.5rem 1rem;
      font-size: 0.9rem;
    }
  }
</style>