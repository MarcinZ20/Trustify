import axios from 'axios'
import type SearchParams from '@/types/searchParams'

export async function searchContent(params: SearchParams) {
  const response = await axios.post('http://localhost:8080/api/v1/search', params)
  return response.data
}
