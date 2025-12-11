<template>
  <div class="media-gallery">
    <div class="gallery-filters">
      <div class="categories">
        <button
          v-for="category in categories"
          :key="category.id"
          class="category-btn"
          :class="{ active: selectedCategory === category.slug }"
          @click="selectCategory(category.slug)"
        >
          {{ category.name }}
          <span class="count">{{ category.item_count }}</span>
        </button>
      </div>
      
      <div class="filter-controls">
        <select v-model="sortBy" @change="fetchMedia">
          <option value="newest">Newest First</option>
          <option value="oldest">Oldest First</option>
          <option value="views">Most Views</option>
          <option value="likes">Most Likes</option>
        </select>
        
        <select v-model="itemsPerPage" @change="fetchMedia">
          <option value="12">12 per page</option>
          <option value="24">24 per page</option>
          <option value="48">48 per page</option>
        </select>
      </div>
    </div>
    
    <div v-if="loading" class="loading">
      <FeatherIcon name="loader" size="24" class="spin" />
      <span>Loading media...</span>
    </div>
    
    <div v-else-if="error" class="error">
      <FeatherIcon name="alert-circle" size="24" />
      <span>{{ error }}</span>
      <button @click="fetchMedia" class="btn mt-2">
        Retry
      </button>
    </div>
    
    <template v-else>
      <div class="masonry-grid">
        <MediaItem
          v-for="item in mediaItems"
          :key="item.id"
          :item="item"
        />
      </div>
      
      <div v-if="mediaItems.length === 0" class="no-media">
        <FeatherIcon name="image" size="48" />
        <p>No media found in this category.</p>
      </div>
    </template>
    
    <div v-if="pagination && pagination.pages > 1" class="pagination">
      <button 
        :disabled="pagination.page === 1"
        @click="goToPage(pagination.page - 1)"
        class="page-btn"
      >
        <FeatherIcon name="chevron-left" size="18" />
        Previous
      </button>
      
      <div class="page-numbers">
        <button
          v-for="page in visiblePages"
          :key="page"
          class="page-number"
          :class="{ active: page === pagination.page }"
          @click="goToPage(page)"
        >
          {{ page }}
        </button>
        
        <span v-if="hasMorePages" class="page-ellipsis">...</span>
      </div>
      
      <button 
        :disabled="pagination.page === pagination.pages"
        @click="goToPage(pagination.page + 1)"
        class="page-btn"
      >
        Next
        <FeatherIcon name="chevron-right" size="18" />
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useApi } from '~~/composables/useApi'

interface MediaItem {
  id: number
  title: string
  description: string
  file_name: string
  thumbnail: string
  mime_type: string
  width: number
  height: number
  views: number
  likes: number
  category: string
  artist: string
  created_at: string
}

interface Category {
  id: number
  name: string
  slug: string
  item_count: number
}

interface Pagination {
  page: number
  limit: number
  total: number
  pages: number
}

const { apiFetch } = useApi()

const mediaItems = ref<MediaItem[]>([])
const categories = ref<Category[]>([])
const loading = ref(false)
const error = ref<string | null>(null)
const selectedCategory = ref('')
const sortBy = ref('newest')
const itemsPerPage = ref(12)
const pagination = ref<Pagination | null>(null)

const visiblePages = computed(() => {
  if (!pagination.value) return []
  
  const pages = pagination.value.pages
  const current = pagination.value.page
  const visible: number[] = []
  
  // Always show first page
  visible.push(1)
  
  // Show pages around current
  for (let i = Math.max(2, current - 2); i <= Math.min(pages - 1, current + 2); i++) {
    visible.push(i)
  }
  
  // Always show last page
  if (pages > 1) {
    visible.push(pages)
  }
  
  return [...new Set(visible)].sort((a, b) => a - b)
})

const hasMorePages = computed(() => {
  if (!pagination.value) return false
  return pagination.value.pages > 5 && 
         pagination.value.page < pagination.value.pages - 3
})

const fetchCategories = async () => {
  try {
    const data = await apiFetch<Category[]>('/media/categories')
    categories.value = data
  } catch (err) {
    console.error('Failed to fetch categories:', err)
  }
}

const fetchMedia = async (page = 1) => {
  loading.value = true
  error.value = null
  
  try {
    const params = new URLSearchParams({
      page: page.toString(),
      limit: itemsPerPage.value.toString(),
      sort: sortBy.value
    })
    
    if (selectedCategory.value) {
      params.append('category', selectedCategory.value)
    }
    
    const data = await apiFetch<{
      items: MediaItem[]
      pagination: Pagination
    }>(`/media?${params}`)
    
    mediaItems.value = data.items
    pagination.value = data.pagination
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'Failed to load media'
    console.error('Failed to fetch media:', err)
  } finally {
    loading.value = false
  }
}

const selectCategory = (categorySlug: string) => {
  selectedCategory.value = selectedCategory.value === categorySlug ? '' : categorySlug
  fetchMedia(1)
}

const goToPage = (page: number) => {
  if (page < 1 || (pagination.value && page > pagination.value.pages)) return
  fetchMedia(page)
}

onMounted(async () => {
  await Promise.all([
    fetchCategories(),
    fetchMedia()
  ])
})
</script>

<style scoped>
.media-gallery {
  padding: 2rem 0;
}

.gallery-filters {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  margin-bottom: 2rem;
}

.categories {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
}

.category-btn {
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
  font-family: inherit;
}

.category-btn:hover {
  border-color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.1);
}

.category-btn.active {
  border-color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.2);
  color: var(--theme-primary);
}

.count {
  font-size: 0.8rem;
  padding: 0.1rem 0.4rem;
  background: var(--theme-border);
  border-radius: 2px;
}

.filter-controls {
  display: flex;
  gap: 1rem;
}

.filter-controls select {
  padding: 0.5rem 1rem;
  background: var(--theme-bg);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  color: var(--theme-fg);
  font-family: inherit;
  cursor: pointer;
}

.filter-controls select:focus {
  outline: none;
  border-color: var(--theme-primary);
}

.masonry-grid {
  column-count: 3;
  column-gap: 1rem;
  margin-bottom: 2rem;
}

.masonry-grid > * {
  break-inside: avoid;
  margin-bottom: 1rem;
}

.loading,
.error,
.no-media {
  grid-column: 1 / -1;
  text-align: center;
  padding: 4rem 2rem;
  color: var(--theme-fg);
}

.loading svg.spin {
  animation: spin 1s linear infinite;
  margin-bottom: 1rem;
}

.no-media {
  opacity: 0.6;
}

.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 1rem;
  margin-top: 2rem;
  padding-top: 2rem;
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
  font-family: inherit;
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

.page-ellipsis {
  display: flex;
  align-items: center;
  padding: 0 0.5rem;
  color: var(--theme-fg);
  opacity: 0.6;
}

@media (max-width: 1200px) {
  .masonry-grid {
    column-count: 2;
  }
}

@media (max-width: 768px) {
  .masonry-grid {
    column-count: 1;
  }
  
  .filter-controls {
    flex-direction: column;
  }
  
  .pagination {
    flex-direction: column;
    gap: 0.5rem;
  }
  
  .page-numbers {
    order: -1;
  }
}
</style>