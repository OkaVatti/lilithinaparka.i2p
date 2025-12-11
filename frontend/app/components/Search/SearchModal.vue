<template>
  <div class="search-modal" :class="{ active: isOpen }" @click.self="close">
    <div class="search-container">
      <div class="search-header">
        <div class="search-input-container">
          <FeatherIcon name="search" size="20" class="search-icon" />
          <input
            ref="searchInput"
            v-model="query"
            type="text"
            placeholder="Search blog posts, games, media..."
            @input="onInput"
            @keydown.enter="performSearch"
            @keydown.esc="close"
          />
          <button v-if="query" @click="clearSearch" class="clear-btn">
            <FeatherIcon name="x" size="18" />
          </button>
        </div>
        <button @click="close" class="close-btn">
          <FeatherIcon name="x" size="24" />
        </button>
      </div>

      <div v-if="isLoading" class="loading">
        <FeatherIcon name="loader" size="24" class="spin" />
        <span>Searching...</span>
      </div>

      <div v-else-if="error" class="error">
        <FeatherIcon name="alert-circle" size="24" />
        <span>{{ error }}</span>
      </div>

      <div v-else-if="query && results.length === 0" class="no-results">
        <FeatherIcon name="search" size="48" />
        <h3>No results found</h3>
        <p>Try different keywords or check your spelling.</p>
        <div v-if="suggestions.length > 0" class="suggestions">
          <p>Suggestions:</p>
          <div class="suggestion-tags">
            <button
              v-for="suggestion in suggestions"
              :key="suggestion"
              @click="applySuggestion(suggestion)"
              class="suggestion-tag"
            >
              {{ suggestion }}
            </button>
          </div>
        </div>
      </div>

      <div v-else-if="results.length > 0" class="search-results">
        <div class="results-header">
          <h3>{{ totalResults }} results for "{{ query }}"</h3>
          <div class="result-filters">
            <select v-model="selectedType" @change="filterResults">
              <option value="">All Types</option>
              <option value="blog">Blog Posts</option>
              <option value="game">Games</option>
              <option value="media">Media</option>
            </select>
            <select v-model="sortBy" @change="filterResults">
              <option value="relevance">Relevance</option>
              <option value="date">Date</option>
              <option value="title">Title</option>
            </select>
          </div>
        </div>

        <div class="results-list">
          <SearchResult
            v-for="result in filteredResults"
            :key="`${result.type}-${result.item_id}`"
            :result="result"
            @click="handleResultClick(result)"
          />
        </div>

        <div v-if="totalPages > 1" class="pagination">
          <button
            :disabled="currentPage === 1"
            @click="goToPage(currentPage - 1)"
            class="page-btn"
          >
            <FeatherIcon name="chevron-left" size="16" />
            Previous
          </button>
          
          <div class="page-numbers">
            <button
              v-for="page in visiblePages"
              :key="page"
              class="page-number"
              :class="{ active: page === currentPage }"
              @click="goToPage(page)"
            >
              {{ page }}
            </button>
          </div>
          
          <button
            :disabled="currentPage === totalPages"
            @click="goToPage(currentPage + 1)"
            class="page-btn"
          >
            Next
            <FeatherIcon name="chevron-right" size="16" />
          </button>
        </div>
      </div>

      <div v-else class="search-tips">
        <div class="tip-section">
          <h4>Search Tips</h4>
          <ul>
            <li>Use quotes for exact matches: "hello world"</li>
            <li>Use AND/OR for boolean searches: programming AND games</li>
            <li>Filter by type: type:blog programming</li>
            <li>Search within specific tags: tag:retro games</li>
          </ul>
        </div>
        
        <div class="recent-searches" v-if="recentSearches.length > 0">
          <h4>Recent Searches</h4>
          <div class="recent-tags">
            <button
              v-for="recent in recentSearches"
              :key="recent"
              @click="setQuery(recent)"
              class="recent-tag"
            >
              {{ recent }}
            </button>
          </div>
        </div>
        
        <div class="popular-searches">
          <h4>Popular Searches</h4>
          <div class="popular-tags">
            <button
              v-for="popular in popularSearches"
              :key="popular"
              @click="setQuery(popular)"
              class="popular-tag"
            >
              {{ popular }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted, onUnmounted } from 'vue'
import { useApi } from '~/composables/useApi'
import { useRouter } from 'vue-router'

const props = defineProps<{
  isOpen: boolean
}>()

const emit = defineEmits<{
  close: []
}>()

const router = useRouter()
const { apiFetch } = useApi()

const query = ref('')
const searchInput = ref<HTMLInputElement | null>(null)
const isLoading = ref(false)
const error = ref<string | null>(null)
const results = ref<any[]>([])
const suggestions = ref<string[]>([])
const selectedType = ref('')
const sortBy = ref('relevance')
const currentPage = ref(1)
const totalResults = ref(0)
const totalPages = ref(0)
const recentSearches = ref<string[]>([])
const popularSearches = ref<string[]>([
  'programming', 'games', 'blog', 'retro', 'i2p',
  'privacy', 'vue', 'go', 'tutorial'
])

const filteredResults = computed(() => {
  if (!selectedType.value) return results.value
  
  return results.value.filter(result => 
    result.type === selectedType.value
  )
})

const visiblePages = computed(() => {
  const pages: number[] = []
  const maxVisible = 5
  
  let start = Math.max(1, currentPage.value - 2)
  let end = Math.min(totalPages.value, start + maxVisible - 1)
  
  // Adjust start if we're near the end
  if (end - start + 1 < maxVisible) {
    start = Math.max(1, end - maxVisible + 1)
  }
  
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  
  return pages
})

let searchTimeout: NodeJS.Timeout | null = null

const onInput = () => {
  if (searchTimeout) {
    clearTimeout(searchTimeout)
  }
  
  if (query.value.length >= 2) {
    searchTimeout = setTimeout(() => {
      performSearch()
    }, 300)
  } else {
    results.value = []
    suggestions.value = []
  }
}

const performSearch = async (page = 1) => {
  if (!query.value.trim()) {
    results.value = []
    return
  }
  
  isLoading.value = true
  error.value = null
  
  try {
    const response = await apiFetch<any>('/search', {
      method: 'POST',
      body: JSON.stringify({
        query: query.value,
        types: selectedType.value ? [selectedType.value] : [],
        limit: 10,
        page,
        sort_by: sortBy.value,
        sort_dir: 'desc'
      })
    })
    
    results.value = response.results
    suggestions.value = response.suggestions || []
    totalResults.value = response.total
    totalPages.value = response.total_pages
    currentPage.value = page
    
    // Save to recent searches
    saveToRecentSearches(query.value)
    
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'Search failed'
    console.error('Search error:', err)
  } finally {
    isLoading.value = false
  }
}

const filterResults = () => {
  currentPage.value = 1
  performSearch(1)
}

const goToPage = (page: number) => {
  if (page < 1 || page > totalPages.value) return
  currentPage.value = page
  performSearch(page)
}

const clearSearch = () => {
  query.value = ''
  results.value = []
  suggestions.value = []
  if (searchInput.value) {
    searchInput.value.focus()
  }
}

const setQuery = (q: string) => {
  query.value = q
  performSearch()
}

const applySuggestion = (suggestion: string) => {
  query.value = suggestion
  performSearch()
}

const handleResultClick = (result: any) => {
  let route = ''
  
  switch (result.type) {
    case 'blog':
      route = `/blog/${result.slug}`
      break
    case 'game':
      route = `/games/${result.slug}`
      break
    case 'media':
      route = `/media/${result.item_id}`
      break
  }
  
  if (route) {
    router.push(route)
    close()
  }
}

const close = () => {
  emit('close')
}

const saveToRecentSearches = (search: string) => {
  if (!search.trim()) return
  
  const recent = recentSearches.value.filter(s => s !== search)
  recent.unshift(search)
  
  if (recent.length > 10) {
    recent.pop()
  }
  
  recentSearches.value = recent
  
  if (process.client) {
    localStorage.setItem('recent_searches', JSON.stringify(recent))
  }
}

const loadRecentSearches = () => {
  if (process.client) {
    const stored = localStorage.getItem('recent_searches')
    if (stored) {
      try {
        recentSearches.value = JSON.parse(stored)
      } catch {
        recentSearches.value = []
      }
    }
  }
}

const handleKeydown = (e: KeyboardEvent) => {
  if (e.key === '/' && e.target === document.body) {
    e.preventDefault()
    if (!props.isOpen) {
      // Open search modal - would need parent to handle
    }
  }
}

watch(() => props.isOpen, (isOpen) => {
  if (isOpen && searchInput.value) {
    searchInput.value.focus()
  }
})

onMounted(() => {
  loadRecentSearches()
  document.addEventListener('keydown', handleKeydown)
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleKeydown)
  if (searchTimeout) {
    clearTimeout(searchTimeout)
  }
})
</script>

<style scoped>
.search-modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.8);
  display: flex;
  align-items: flex-start;
  justify-content: center;
  z-index: 9999;
  opacity: 0;
  visibility: hidden;
  transition: all 0.3s;
  padding-top: 4rem;
}

.search-modal.active {
  opacity: 1;
  visibility: visible;
}

.search-container {
  width: 90%;
  max-width: 800px;
  max-height: 80vh;
  background: var(--theme-bg);
  border-radius: 12px;
  border: 2px solid var(--theme-primary);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.5);
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.search-header {
  display: flex;
  align-items: center;
  padding: 1.5rem;
  border-bottom: 2px solid var(--theme-border);
  background: rgba(0, 0, 0, 0.2);
}

.search-input-container {
  flex: 1;
  position: relative;
  display: flex;
  align-items: center;
}

.search-icon {
  position: absolute;
  left: 1rem;
  color: var(--theme-fg);
  opacity: 0.6;
}

.search-input-container input {
  width: 100%;
  padding: 1rem 3rem 1rem 3rem;
  font-size: 1.2rem;
  background: var(--theme-bg);
  border: 2px solid var(--theme-border);
  border-radius: 8px;
  color: var(--theme-fg);
  font-family: 'Courier New', monospace;
}

.search-input-container input:focus {
  outline: none;
  border-color: var(--theme-primary);
}

.clear-btn {
  position: absolute;
  right: 1rem;
  background: none;
  border: none;
  color: var(--theme-fg);
  opacity: 0.6;
  cursor: pointer;
  padding: 0.25rem;
}

.clear-btn:hover {
  opacity: 1;
}

.close-btn {
  margin-left: 1rem;
  background: none;
  border: none;
  color: var(--theme-fg);
  cursor: pointer;
  padding: 0.5rem;
  border-radius: 4px;
}

.close-btn:hover {
  background: rgba(255, 255, 255, 0.1);
}

.loading,
.error,
.no-results,
.search-tips,
.search-results {
  flex: 1;
  overflow-y: auto;
  padding: 2rem;
}

.loading,
.error,
.no-results {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  text-align: center;
  min-height: 300px;
}

.loading svg.spin {
  animation: spin 1s linear infinite;
  margin-bottom: 1rem;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.error svg {
  color: var(--theme-error);
  margin-bottom: 1rem;
}

.no-results svg {
  margin-bottom: 1rem;
  opacity: 0.6;
}

.no-results h3 {
  margin-bottom: 0.5rem;
  color: var(--theme-primary);
}

.no-results p {
  margin-bottom: 1.5rem;
  color: var(--theme-fg);
  opacity: 0.8;
}

.suggestions {
  margin-top: 1.5rem;
}

.suggestions p {
  margin-bottom: 0.5rem;
  font-weight: bold;
}

.suggestion-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  justify-content: center;
}

.suggestion-tag {
  padding: 0.5rem 1rem;
  background: rgba(189, 147, 249, 0.1);
  border: 1px solid var(--theme-primary);
  border-radius: 4px;
  color: var(--theme-primary);
  cursor: pointer;
  transition: all 0.2s;
  font-size: 0.9rem;
}

.suggestion-tag:hover {
  background: rgba(189, 147, 249, 0.2);
  transform: translateY(-2px);
}

.results-header {
  margin-bottom: 1.5rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.results-header h3 {
  color: var(--theme-primary);
  margin: 0;
}

.result-filters {
  display: flex;
  gap: 0.5rem;
}

.result-filters select {
  padding: 0.5rem;
  background: var(--theme-bg);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  color: var(--theme-fg);
  font-family: inherit;
  cursor: pointer;
}

.result-filters select:focus {
  outline: none;
  border-color: var(--theme-primary);
}

.results-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 1rem;
  margin-top: 2rem;
  padding-top: 1.5rem;
  border-top: 1px solid var(--theme-border);
}

.page-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 1rem;
  background: var(--theme-bg);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  color: var(--theme-fg);
  cursor: pointer;
  transition: all 0.2s;
}

.page-btn:hover:not(:disabled) {
  border-color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.1);
}

.page-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.page-numbers {
  display: flex;
  gap: 0.25rem;
}

.page-number {
  width: 2.5rem;
  height: 2.5rem;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--theme-bg);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  color: var(--theme-fg);
  cursor: pointer;
  transition: all 0.2s;
}

.page-number:hover {
  border-color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.1);
}

.page-number.active {
  border-color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.2);
  color: var(--theme-primary);
}

.search-tips {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
}

.tip-section,
.recent-searches,
.popular-searches {
  background: rgba(0, 0, 0, 0.2);
  border-radius: 8px;
  padding: 1.5rem;
  border: 1px solid var(--theme-border);
}

.tip-section h4,
.recent-searches h4,
.popular-searches h4 {
  margin-top: 0;
  margin-bottom: 1rem;
  color: var(--theme-primary);
}

.tip-section ul {
  margin: 0;
  padding-left: 1.5rem;
  color: var(--theme-fg);
  opacity: 0.8;
  line-height: 1.6;
}

.tip-section li {
  margin-bottom: 0.5rem;
}

.recent-tags,
.popular-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
}

.recent-tag,
.popular-tag {
  padding: 0.5rem 1rem;
  background: rgba(189, 147, 249, 0.1);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  color: var(--theme-fg);
  cursor: pointer;
  transition: all 0.2s;
  font-size: 0.9rem;
}

.recent-tag:hover,
.popular-tag:hover {
  border-color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.2);
}

@media (max-width: 768px) {
  .search-modal {
    padding-top: 1rem;
    align-items: stretch;
  }
  
  .search-container {
    width: 100%;
    max-height: 100vh;
    border-radius: 0;
    border: none;
  }
  
  .search-header {
    flex-direction: column;
    gap: 1rem;
  }
  
  .search-input-container {
    width: 100%;
  }
  
  .close-btn {
    align-self: flex-end;
  }
  
  .results-header {
    flex-direction: column;
    align-items: stretch;
    gap: 1rem;
  }
  
  .search-tips {
    grid-template-columns: 1fr;
  }
}
</style>