<script setup lang="ts">
import LandingPageView from './views/LandingPageView.vue'
import AboutView from './views/AboutView.vue'
import ContactView from './views/ContactView.vue'
import Footer from './components/Footer.vue'
import Navbar from './components/Navbar.vue'
</script>

<template>
  <main class="flex flex-col justify-between h-screen bg-gray-100">
    <Navbar />
    <router-view class=""></router-view>
    <Footer />
  </main>
</template>
