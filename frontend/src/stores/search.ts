import { defineStore } from 'pinia'
import type { ResponseParams } from '@/types/responseParams'
import type { SearchParams } from '@/types/searchParams'
import { useSearch } from '@/composables/search'

export const useSearchStore = defineStore('search', {
  state: () => ({
    searchQuery: null as SearchParams | null,
    responseData: null as ResponseParams | null,
    loading: false,
    error: null,
  }),
  getters: {},
  actions: {
    async submitSearchQuery(query: SearchParams, router: any) {
      this.searchQuery = query
      this.loading = true
      this.error = null

      const { data, error } = await useSearch(query)

      if (error !== null) {
        this.error = error
        this.loading = false
        router.push({ name: 'ErrorPage' })
      }

      this.responseData = data
      this.loading = false

      router.push({
        name: 'ResultsPage',
        query: { headline: query.headline },
      })
    },
  },
})
