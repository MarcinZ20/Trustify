<script setup lang="ts">
import LandingPageView from './views/LandingPageView.vue'
import AboutView from './views/AboutView.vue'
import ContactView from './views/ContactView.vue'
import Footer from './components/Footer.vue'
import Navbar from './components/Navbar.vue'
</script>

<template>
  <main class="bg-gray-100">
    <Navbar />
    <LandingPageView />
    <Footer />
  </main>
</template>
