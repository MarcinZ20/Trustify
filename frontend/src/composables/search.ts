import axios from 'axios'
import type SearchParams from '@/types/searchParams'

export async function useSearch(params: SearchParams) {
  try {
    const response = await axios.post('http://localhost:8080/api/v1/search', params)
    return { data: response.data, error: null }
  } catch (err) {
    return { data: null, error: err }
  }
}
