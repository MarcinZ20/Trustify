<script setup lang="ts">
import Footer from './components/Footer.vue'
import Navbar from './components/Navbar.vue'
</script>

<template>
  <main class="flex flex-col justify-between h-screen bg-gray-100">
    <Navbar />
    <router-view ></router-view>
    <Footer />
  </main>
</template>
