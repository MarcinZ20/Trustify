<script setup lang="ts">
defineProps({
  text: String,
})
</script>

<template>
  <div class="bg-zinc-800 rounded-lg text-center">
    <p class="text-white text-sm p-1 font-bold">{{ text }}</p>
  </div>
</template>
