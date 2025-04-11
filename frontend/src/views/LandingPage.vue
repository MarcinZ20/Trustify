<script setup lang="ts">
import SearchForm from '../components/SearchForm.vue'
import MainHeader from '../components/MainHeader.vue'
import Widget from '../components/Widget.vue'
import RecentSearches from '../components/searches/RecentSearches.vue'
import TrendingContainer from '../components/trending/TrendingContainer.vue'
</script>

<template>
  <div class="grid grid-cols-1 md:grid-rows-4 md:grid-cols-2">
    <div class="bg-white rounded-lg p-2 m-2 md:row-span-5">
      <Widget text="AI" class="w-[10%]" />
      <MainHeader class="md:mt-6" />
      <SearchForm />
    </div>
    <div class="bg-white rounded-lg row-span-3 p-2 m-2">
      <RecentSearches />
    </div>
    <div class="pb-2">
      <TrendingContainer />
    </div>
  </div>
</template>
