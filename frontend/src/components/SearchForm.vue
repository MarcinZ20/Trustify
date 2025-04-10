<script setup lang="ts">
import { reactive } from 'vue'
import { useRouter } from 'vue-router'
import type { SearchParams } from '@/types/searchParams'
import { searchContent } from '@/composables/useSearch'

const s = reactive<SearchParams>({
  headline: '',
  content: '',
  url: '',
})

const router = useRouter()

const sleep = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))
const onSubmit = async () => {
  // Push to the loading page
  router.push({name: 'LoadingPage'})

  try {
    const result = await searchContent({...s})
    await sleep(3000)
    // After success push to the result page
    router.push({
      name: 'ResultsPage',
      query: { headline: s.headline },
      state: { result }
    })
  } catch (error) {
    console.log(error)
    router.push({ name: 'ErrorPage' })
  }
}

</script>

<template>
  <div class="p-4 rounded-lg">
    <form @submit.prevent="onSubmit">
      <div class="grid grid-cols-1 gap-2 md:grid-cols-4">
        <div class="col-span-full md:col-span-3">
          <input
            v-model.lazy.trim="s.headline"
            type="text"
            id="searchbar"
            name="searchbar"
            required
            class="w-full bg-white px-3.5 py-2 rounded-md text-zinc-800 outline-1 -outline-offset-1 outline-zinc-300 placeholder:text-gray-400 focus:outline-zinc-500 sm:text-sm"
            placeholder="Violets are yellow ..."
          />
        </div>
        <button
          type="submit"
          id="submit-button"
          class="rounded-lg bg-zinc-800 text-white p-1.5 hover:bg-zinc-900"
        >
          Search
        </button>
      </div>
    </form>
  </div>
</template>

<style scoped></style>
