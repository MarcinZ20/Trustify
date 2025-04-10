<script setup lang="ts">
import axios from 'axios'
</script>

<template>
  <div class="p-4 rounded-lg">
    <form action="" method="post">
      <div class="grid grid-cols-1 gap-2 md:grid-cols-4">
        <div class="col-span-full md:col-span-3">
          <input
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
