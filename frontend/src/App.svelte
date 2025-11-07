<script lang="ts">
  import Header from './components/Header.svelte';
  import Hero from './components/Hero.svelte';
  import Services from './components/Services.svelte';
  import Capabilities from './components/Capabilities.svelte';
  import Contact from './components/Contact.svelte';
  import Footer from './components/Footer.svelte';
</script>

<main>
  <Header />
  <Hero />
  <Services />
  <Capabilities />
  <Contact />
  <Footer />
</main>

<style>
  main {
    width: 100%;
    overflow-x: hidden;
  }
</style>