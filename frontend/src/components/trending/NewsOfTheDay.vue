<script setup lang="ts"></script>

<template>
  <div class="p-2 bg-zinc-800 rounded-lg">
    <p class="text-white text-lg">World is beautiful</p>
    <p class="text-xs text-gray-300 mb-2">
      Here is some context for the news. Click below to keep reading!
    </p>
    <div class="flex flex-row gap-1 items-center">
      <a href="https://example.com" class="text-xs text-underline">Read more</a>

      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="size-6"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M17.25 8.25 21 12m0 0-3.75 3.75M21 12H3"
        />
      </svg>
    </div>
  </div>
</template>
