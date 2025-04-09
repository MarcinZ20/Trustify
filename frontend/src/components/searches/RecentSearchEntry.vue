<script setup lang="ts">
defineProps({
  query: String,
  url: String,
})
</script>

<template>
  <div class="flex flex-row place-content-between p-2 bg-gray-200 rounded-xl">
    <p class="text-zinc-800">{{ query }}</p>
    <a :href="url" class="bg-zinc-800 rounded-lg text-white p-1 text-xs">Search</a>
  </div>
</template>
