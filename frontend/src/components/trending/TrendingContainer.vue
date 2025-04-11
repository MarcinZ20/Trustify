<script setup lang="ts">
import ImageOfTheDay from './ImageOfTheDay.vue'
import NewsOfTheDay from './NewsOfTheDay.vue'
</script>

<template>
  <div class="grid grid-cols-1 gap-2 md:grid-cols-2">
    <ImageOfTheDay class="col-span-1 p-2 rounded-lg mx-2 md:m-0 md:ml-2" />
    <div class="bg-zinc-800 text-white p-2 rounded-lg mx-2 md:m-0 md:mr-2">
      <NewsOfTheDay class="col-span-1" />
    </div>
  </div>
</template>
