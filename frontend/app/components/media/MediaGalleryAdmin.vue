<template>
  <div class="media-gallery-admin">
    <div class="gallery-filters">
      <div class="filter-group">
        <label>Category:</label>
        <select v-model="selectedCategory" @change="filterMedia">
          <option value="">All Categories</option>
          <option value="art">Art</option>
          <option value="photography">Photography</option>
          <option value="video">Video</option>
          <option value="other">Other</option>
        </select>
      </div>
      
      <div class="search-group">
        <input 
          v-model="searchQuery" 
          type="text" 
          placeholder="Search media..." 
          @input="filterMedia"
        />
        <FeatherIcon name="search" size="18" />
      </div>
      
      <div class="view-toggle">
        <button 
          @click="viewMode = 'grid'" 
          :class="{ active: viewMode === 'grid' }"
          title="Grid view"
        >
          <FeatherIcon name="grid" size="18" />
        </button>
        <button 
          @click="viewMode = 'list'" 
          :class="{ active: viewMode === 'list' }"
          title="List view"
        >
          <FeatherIcon name="list" size="18" />
        </button>
      </div>
    </div>
    
    <div v-if="loading" class="loading">
      <FeatherIcon name="loader" size="24" class="spin" />
      <span>Loading media...</span>
    </div>
    
    <div v-else-if="error" class="error">
      <FeatherIcon name="alert-circle" size="20" />
      <span>{{ error }}</span>
    </div>
    
    <div v-else-if="filteredMedia.length === 0" class="no-media">
      <FeatherIcon name="image" size="48" />
      <p>No media found</p>
    </div>
    
    <div v-else :class="['media-container', viewMode]">
      <div v-for="item in filteredMedia" :key="item.id" class="media-item">
        <div class="media-preview">
          <img 
            v-if="item.mime_type.startsWith('image/')"
            :src="getMediaUrl(item)" 
            :alt="item.title"
            @click="viewMedia(item)"
          />
          <div v-else-if="item.mime_type.startsWith('video/')" class="video-preview">
            <FeatherIcon name="play-circle" size="48" />
            <span>{{ item.title }}</span>
          </div>
          <div v-else class="file-preview">
            <FeatherIcon name="file" size="48" />
            <span>{{ item.file_name }}</span>
          </div>
          
          <div class="media-actions">
            <button @click="editMedia(item)" class="action-btn" title="Edit">
              <FeatherIcon name="edit" size="16" />
            </button>
            <button @click="deleteMedia(item)" class="action-btn danger" title="Delete">
              <FeatherIcon name="trash" size="16" />
            </button>
          </div>
        </div>
        
        <div class="media-info">
          <h4>{{ item.title || item.file_name }}</h4>
          <div class="media-meta">
            <span class="file-size">{{ formatFileSize(item.file_size) }}</span>
            <span class="category">{{ item.category }}</span>
            <span v-if="!item.is_public" class="private-badge">Private</span>
          </div>
        </div>
      </div>
    </div>
    
    <div v-if="showPagination" class="pagination">
      <button 
        @click="previousPage" 
        :disabled="currentPage === 1"
        class="btn btn-secondary"
      >
        <FeatherIcon name="chevron-left" size="18" />
        Previous
      </button>
      
      <span class="page-info">Page {{ currentPage }} of {{ totalPages }}</span>
      
      <button 
        @click="nextPage" 
        :disabled="currentPage === totalPages"
        class="btn btn-secondary"
      >
        Next
        <FeatherIcon name="chevron-right" size="18" />
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useApi } from '~/composables/useApi'
import type { MediaItem } from '~~/types'

const { apiFetch } = useApi()

const media = ref<MediaItem[]>([])
const loading = ref(true)
const error = ref<string | null>(null)
const selectedCategory = ref('')
const searchQuery = ref('')
const viewMode = ref<'grid' | 'list'>('grid')
const currentPage = ref(1)
const itemsPerPage = 24

const filteredMedia = computed(() => {
  let result = media.value
  
  if (selectedCategory.value) {
    result = result.filter(item => item.category === selectedCategory.value)
  }
  
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    result = result.filter(item => 
      item.title?.toLowerCase().includes(query) ||
      item.description?.toLowerCase().includes(query) ||
      item.tags?.toLowerCase().includes(query)
    )
  }
  
  const start = (currentPage.value - 1) * itemsPerPage
  const end = start + itemsPerPage
  return result.slice(start, end)
})

const totalPages = computed(() => {
  let result = media.value
  
  if (selectedCategory.value) {
    result = result.filter(item => item.category === selectedCategory.value)
  }
  
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    result = result.filter(item => 
      item.title?.toLowerCase().includes(query) ||
      item.description?.toLowerCase().includes(query) ||
      item.tags?.toLowerCase().includes(query)
    )
  }
  
  return Math.ceil(result.length / itemsPerPage)
})

const showPagination = computed(() => totalPages.value > 1)

const fetchMedia = async () => {
  loading.value = true
  error.value = null
  
  try {
    const response = await apiFetch<{ items: MediaItem[] }>('/media')
    media.value = response.items || []
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'Failed to load media'
  } finally {
    loading.value = false
  }
}

const filterMedia = () => {
  currentPage.value = 1
}

const getMediaUrl = (item: MediaItem) => {
  return item.thumbnail || `/media/${item.file_name}`
}

const formatFileSize = (bytes: number) => {
  if (!bytes) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return Math.round(bytes / Math.pow(k, i) * 100) / 100 + ' ' + sizes[i]
}

const viewMedia = (item: MediaItem) => {
  window.open(getMediaUrl(item), '_blank')
}

const editMedia = (item: MediaItem) => {
  console.log('Edit media:', item)
}

const deleteMedia = async (item: MediaItem) => {
  if (!confirm(`Delete "${item.title || item.file_name}"?`)) return
  
  try {
    await apiFetch(`/media/${item.id}`, { method: 'DELETE' })
    await fetchMedia()
  } catch (err) {
    console.error('Failed to delete media:', err)
    alert('Failed to delete media')
  }
}

const previousPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
  }
}

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
  }
}

onMounted(() => {
  fetchMedia()
})
</script>

<style scoped>
.gallery-filters {
  display: flex;
  gap: 1rem;
  margin-bottom: 1.5rem;
  align-items: center;
}

.filter-group {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.filter-group label {
  font-size: 0.9rem;
  color: var(--theme-fg);
  font-weight: bold;
}

.filter-group select {
  padding: 0.5rem;
  background: var(--theme-bg);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  color: var(--theme-fg);
  font-family: inherit;
}

.search-group {
  flex: 1;
  position: relative;
}

.search-group input {
  width: 100%;
  padding: 0.5rem 2.5rem 0.5rem 1rem;
  background: var(--theme-bg);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  color: var(--theme-fg);
  font-family: inherit;
}

.search-group svg {
  position: absolute;
  right: 0.75rem;
  top: 50%;
  transform: translateY(-50%);
  color: var(--theme-fg);
  opacity: 0.6;
}

.view-toggle {
  display: flex;
  gap: 0.25rem;
  background: rgba(0, 0, 0, 0.2);
  border-radius: 4px;
  padding: 0.25rem;
}

.view-toggle button {
  padding: 0.5rem;
  background: transparent;
  border: none;
  color: var(--theme-fg);
  cursor: pointer;
  border-radius: 4px;
  transition: all 0.2s;
}

.view-toggle button:hover,
.view-toggle button.active {
  background: var(--theme-primary);
  color: var(--theme-bg);
}

.media-container.grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 1rem;
}

.media-container.list {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.media-item {
  background: rgba(0, 0, 0, 0.2);
  border: 1px solid var(--theme-border);
  border-radius: 8px;
  overflow: hidden;
  transition: all 0.2s;
}

.media-item:hover {
  border-color: var(--theme-primary);
  transform: translateY(-2px);
}

.media-preview {
  position: relative;
  width: 100%;
  padding-bottom: 100%;
  overflow: hidden;
  background: rgba(0, 0, 0, 0.3);
}

.list .media-preview {
  padding-bottom: 0;
  height: 100px;
  width: 100px;
}

.media-preview img,
.video-preview,
.file-preview {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  cursor: pointer;
}

.video-preview,
.file-preview {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  color: var(--theme-fg);
  opacity: 0.6;
}

.media-actions {
  position: absolute;
  top: 0.5rem;
  right: 0.5rem;
  display: flex;
  gap: 0.25rem;
  opacity: 0;
  transition: opacity 0.2s;
}

.media-item:hover .media-actions {
  opacity: 1;
}

.action-btn {
  padding: 0.5rem;
  background: rgba(0, 0, 0, 0.8);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  color: var(--theme-fg);
  cursor: pointer;
  transition: all 0.2s;
}

.action-btn:hover {
  background: var(--theme-primary);
  color: var(--theme-bg);
}

.action-btn.danger:hover {
  background: var(--theme-error);
}

.media-info {
  padding: 0.75rem;
}

.media-info h4 {
  margin: 0 0 0.5rem 0;
  font-size: 0.9rem;
  color: var(--theme-fg);
}

.media-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  font-size: 0.75rem;
}

.media-meta span {
  padding: 0.25rem 0.5rem;
  background: rgba(0, 0, 0, 0.3);
  border-radius: 3px;
  color: var(--theme-fg);
  opacity: 0.8;
}

.private-badge {
  color: var(--theme-warning) !important;
}

.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 1rem;
  margin-top: 2rem;
}

.page-info {
  color: var(--theme-fg);
  font-size: 0.9rem;
}

.loading,
.error,
.no-media {
  text-align: center;
  padding: 3rem;
  color: var(--theme-fg);
}

.spin {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}
</style>