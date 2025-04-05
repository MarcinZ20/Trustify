<script setup lang="ts">
import { ref } from 'vue'

const urls = ref<string[]>(['http://example.de', 'http://dummywebsite.com', 'http://foo.pl'])
</script>
<template>
  <div>
    <p class="text-2xl text-red-500">Your news is fake</p>
    <p class="text-xl text-sky-700">Credibility: 70%</p>
    <ul>
      <li v-for="(url, index) in urls" :key="index">
        {{ url }}
      </li>
    </ul>
  </div>
</template>
<style scoped></style>
