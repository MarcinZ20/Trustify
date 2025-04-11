<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router'
import { useSearchStore } from '@/stores/search'

const route = useRoute()
const router = useRouter()
const store = useSearchStore()

const goBack = () => {
  router.push({ name: 'LandingPage' })
}
</script>

<template>
  <div class="max-w-3xl mx-auto p-6">
    <template v-if="!store.responseData">
      <p class="text-center text-zinc-500 text-lg">No result found. Please search again.</p>
      <div class="text-center mt-4">
        <button @click="goBack" class="bg-zinc-800 text-white px-4 py-2 rounded hover:bg-zinc-900">
          ← Back to Search
        </button>
      </div>
    </template>

    <template v-else>
      <h1 class="text-3xl font-bold text-zinc-800 mb-4">
        Result for: "{{ route.query.headline }}"
      </h1>

      <div class="mb-6">
        <p class="text-xl text-zinc-700 mb-2">
          Verdict: <span class="font-semibold">{{ store.responseData.result }}</span>
        </p>
        <p class="text-lg text-zinc-600">
          Confidence Score: <span class="font-mono">{{ store.responseData.score }}%</span>
        </p>
      </div>

      <div class="mb-6">
        <h2 class="text-xl font-semibold text-zinc-800 mb-2">Sources:</h2>
        <ul class="list-disc list-inside space-y-1">
          <li v-for="(url, index) in store.responseData.sources" :key="index">
            <a :href="url" target="_blank" rel="noopener" class="text-blue-600 hover:underline">
              {{ url }}
            </a>
          </li>
        </ul>
      </div>

      <button
        @click="goBack"
        class="mt-6 px-4 py-2 rounded-md bg-zinc-800 text-white hover:bg-zinc-900 transition"
      >
        ← Go back
      </button>
    </template>
  </div>
</template>
