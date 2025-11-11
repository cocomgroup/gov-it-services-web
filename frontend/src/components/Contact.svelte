<script lang="ts">
  let formData = $state({
    name: '',
    email: '',
    organization: '',
    clearance: '',
    message: ''
  });

  let submitting = $state(false);
  let submitStatus: 'idle' | 'success' | 'error' = $state('idle');
  let statusMessage = $state('');

  async function handleSubmit(event: Event) {
    event.preventDefault();
    submitting = true;
    submitStatus = 'idle';

    try {
      const response = await fetch('/api/contact', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData)
      });

      if (!response.ok) throw new Error('Failed to submit form');

      const data = await response.json();
      submitStatus = 'success';
      statusMessage = data.message;

      // Reset form
      formData = {
        name: '',
        email: '',
        organization: '',
        clearance: '',
        message: ''
      };
    } catch (error) {
      submitStatus = 'error';
      statusMessage = 'Failed to submit form. Please try again.';
      console.error(error);
    } finally {
      submitting = false;
    }
  }
</script>

<section id="contact" class="contact">
  <div class="container">
    <div class="contact-wrapper">
      <div class="contact-info">
        <h2>Ready to Secure Your Mission</h2>
        <p class="intro">
          Contact our team to discuss your government IT requirements. 
          We're ready to support your mission with secure, compliant solutions.
        </p>

        <div class="info-items">
          <div class="info-item">
            <div class="icon">📞</div>
            <div>
              <h4>Phone</h4>
              <p>(202) 555-0100</p>
            </div>
          </div>

          <div class="info-item">
            <div class="icon">✉️</div>
            <div>
              <h4>Email</h4>
              <p>contracts@securegovsolutions.com</p>
            </div>
          </div>

          <div class="info-item">
            <div class="icon">📍</div>
            <div>
              <h4>Location</h4>
              <p>Washington, DC Metro Area</p>
            </div>
          </div>

          <div class="info-item">
            <div class="icon">🔒</div>
            <div>
              <h4>Facility Clearance</h4>
              <p>Secret & Top Secret</p>
            </div>
          </div>
        </div>
      </div>

      <div class="contact-form">
        <form onsubmit={handleSubmit}>
          <div class="form-group">
            <label for="name">Full Name *</label>
            <input
              type="text"
              id="name"
              bind:value={formData.name}
              required
              placeholder="John Doe"
            />
          </div>

          <div class="form-group">
            <label for="email">Email Address *</label>
            <input
              type="email"
              id="email"
              bind:value={formData.email}
              required
              placeholder="john.doe@agency.gov"
            />
          </div>

          <div class="form-group">
            <label for="organization">Organization *</label>
            <input
              type="text"
              id="organization"
              bind:value={formData.organization}
              required
              placeholder="Department/Agency Name"
            />
          </div>

          <div class="form-group">
            <label for="clearance">Security Clearance Level</label>
            <select id="clearance" bind:value={formData.clearance}>
              <option value="">Select clearance level</option>
              <option value="none">None</option>
              <option value="confidential">Confidential</option>
              <option value="secret">Secret</option>
              <option value="ts">Top Secret</option>
              <option value="ts-sci">TS/SCI</option>
            </select>
          </div>

          <div class="form-group">
            <label for="message">Project Requirements *</label>
            <textarea
              id="message"
              bind:value={formData.message}
              required
              rows="5"
              placeholder="Please describe your project requirements and timeline..."
            ></textarea>
          </div>

          {#if submitStatus === 'success'}
            <div class="alert alert-success">
              {statusMessage}
            </div>
          {/if}

          {#if submitStatus === 'error'}
            <div class="alert alert-error">
              {statusMessage}
            </div>
          {/if}

          <button type="submit" class="submit-btn" disabled={submitting}>
            {submitting ? 'Submitting...' : 'Request Consultation'}
          </button>
        </form>
      </div>
    </div>
  </div>
</section>

<style>
  .contact {
    padding: 6rem 0;
    background: linear-gradient(135deg, #1e3a8a 0%, #3b82f6 100%);
    color: white;
  }

  .container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 2rem;
  }

  .contact-wrapper {
    display: grid;
    grid-template-columns: 1fr 1.2fr;
    gap: 4rem;
    align-items: start;
  }

  .contact-info h2 {
    font-size: 2.5rem;
    margin-bottom: 1.5rem;
  }

  .intro {
    font-size: 1.1rem;
    line-height: 1.7;
    margin-bottom: 3rem;
    opacity: 0.95;
  }

  .info-items {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
  }

  .info-item {
    display: flex;
    gap: 1rem;
    align-items: start;
  }

  .icon {
    font-size: 2rem;
    min-width: 50px;
  }

  .info-item h4 {
    margin-bottom: 0.25rem;
    font-size: 1.1rem;
  }

  .info-item p {
    opacity: 0.9;
    margin: 0;
  }

  .contact-form {
    background: white;
    padding: 2.5rem;
    border-radius: 12px;
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.2);
  }

  .form-group {
    margin-bottom: 1.5rem;
  }

  label {
    display: block;
    margin-bottom: 0.5rem;
    color: var(--text-primary);
    font-weight: 500;
  }

  input, select, textarea {
    width: 100%;
    padding: 0.75rem;
    border: 2px solid var(--border-color);
    border-radius: 6px;
    font-size: 1rem;
    font-family: inherit;
    transition: border-color 0.3s;
  }

  input:focus, select:focus, textarea:focus {
    outline: none;
    border-color: var(--secondary-color);
  }

  textarea {
    resize: vertical;
    min-height: 120px;
  }

  .submit-btn {
    width: 100%;
    padding: 1rem;
    background: var(--secondary-color);
    color: white;
    border: none;
    border-radius: 6px;
    font-size: 1.1rem;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.3s;
  }

  .submit-btn:hover:not(:disabled) {
    background: var(--primary-color);
    transform: translateY(-2px);
  }

  .submit-btn:disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }

  .alert {
    padding: 1rem;
    border-radius: 6px;
    margin-bottom: 1rem;
  }

  .alert-success {
    background: #d1fae5;
    color: #065f46;
    border: 1px solid #6ee7b7;
  }

  .alert-error {
    background: #fee2e2;
    color: #991b1b;
    border: 1px solid #fca5a5;
  }

  @media (max-width: 968px) {
    .contact {
      padding: 4rem 0;
    }

    .contact-wrapper {
      grid-template-columns: 1fr;
      gap: 2rem;
    }

    .contact-info h2 {
      font-size: 2rem;
    }
  }
</style>