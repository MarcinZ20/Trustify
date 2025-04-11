<script setup lang="ts">
import RecentSearchEntry from './RecentSearchEntry.vue'

const recentSearches = [
  { query: 'Is this love?', url: 'https://example.com' },
  { query: 'Is Trum the president of US?', url: 'https://example.com' },
  { query: 'Cracow is the capitol of Moldova', url: 'https://example.com' },
  { query: 'Is the global warming real?', url: 'https://example.com' },
]
</script>

<template>
  <div class="grid grid-cols-1 grid-rows-5">
    <div class="p-3 row-span-1">
      <p class="text-lg">Recent Searches</p>
      <p class="text-sm text-gray-500">This is what's trending right now</p>
    </div>
    <ul class="row-span-3">
      <RecentSearchEntry
        v-for="(item, index) in recentSearches"
        :key="index"
        :query="item.query"
        :url="item.url"
        class="m-2"
      />
    </ul>
    <div class="row-span-1 flex items-center">
      <button class="flex flex-row gap-2 items-center bg-zinc-800 m-2 text-white rounded-full p-1.5 px-2">
        See more

        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="white"
          class="size-6"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M17.25 8.25 21 12m0 0-3.75 3.75M21 12H3"
          />
        </svg>
      </button>
    </div>
  </div>
</template>
