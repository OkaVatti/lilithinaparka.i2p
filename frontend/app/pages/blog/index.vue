<template>
  <div class="blog-index-page">
    <div class="container">
      <header class="page-header window">
        <div class="title-bar">
          <div class="title-bar-text">Blog</div>
        </div>
        <div class="window-body">
          <div class="header-content">
            <h1>Blog Posts</h1>
            <p class="subtitle">Thoughts on programming, privacy, and digital freedom</p>
            
            <div class="blog-filters">
              <div class="filter-group">
                <label>Category:</label>
                <select v-model="selectedCategory" @change="filterPosts">
                  <option value="">All</option>
                  <option value="Casual">Casual</option>
                  <option value="Interlude">Interlude</option>
                  <option value="Serious">Serious</option>
                </select>
              </div>
              
              <div class="search-group">
                <input 
                  v-model="searchQuery" 
                  type="text" 
                  placeholder="Search posts..."
                  @input="filterPosts"
                />
                <FeatherIcon name="search" size="18" />
              </div>
            </div>
          </div>
        </div>
      </header>

      <div v-if="blogStore.loading" class="loading">
        <FeatherIcon name="loader" size="24" class="spin" />
        <span>Loading posts...</span>
      </div>
      
      <div v-else-if="blogStore.error" class="error window">
        <div class="window-body">
          <FeatherIcon name="alert-circle" size="24" />
          <p>{{ blogStore.error }}</p>
          <button @click="retry" class="btn">
            <FeatherIcon name="refresh-cw" size="18" />
            <span>Retry</span>
          </button>
        </div>
      </div>
      
      <div v-else class="blog-grid">
        <BlogPostCard
          v-for="post in filteredPosts"
          :key="post.slug"
          :title="post.title"
          :excerpt="post.summary || extractExcerpt(post.content)"
          :date="post.date"
          :tags="parseTags(post.tags)"
          :to="`/blog/${post.slug}`"
          :image="post.featured_image"
        />
        
        <div v-if="filteredPosts.length === 0" class="no-posts window">
          <div class="window-body">
            <FeatherIcon name="file-text" size="48" />
            <p>No posts found</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useBlogStore } from '~~/stores/blog'
import BlogPostCard from '~/components/blog/BlogPostCard.vue'

const blogStore = useBlogStore()

const selectedCategory = ref('')
const searchQuery = ref('')

const filteredPosts = computed(() => {
  let posts = blogStore.publishedPosts
  
  if (selectedCategory.value) {
    posts = posts.filter(post => 
      post.categories.includes(selectedCategory.value)
    )
  }
  
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    posts = posts.filter(post =>
      post.title.toLowerCase().includes(query) ||
      post.summary?.toLowerCase().includes(query) ||
      post.content?.toLowerCase().includes(query)
    )
  }
  
  return posts
})

const parseTags = (tagsJson: string): string[] => {
  try {
    return JSON.parse(tagsJson)
  } catch {
    return []
  }
}

const extractExcerpt = (content: string): string => {
  if (!content) return ''
  const text = content.replace(/<[^>]+>/g, '')
  return text.slice(0, 160) + (text.length > 160 ? '...' : '')
}

const filterPosts = () => {
  // Filtering is handled by computed property
}

const retry = () => {
  blogStore.fetchPosts()
}

onMounted(() => {
  blogStore.fetchPosts()
})

useHead({
  title: 'Blog',
  meta: [
    { name: 'description', content: 'Privacy-focused blog on programming and digital freedom' }
  ]
})
</script>

<style scoped>
.blog-index-page {
  padding: 2rem 0;
}

.page-header {
  margin-bottom: 2rem;
}

.header-content {
  text-align: center;
  padding: 1rem;
}

.header-content h1 {
  margin: 0 0 0.5rem 0;
  font-size: 2.5rem;
  color: var(--theme-primary);
}

.subtitle {
  margin: 0 0 1.5rem 0;
  color: var(--theme-fg);
  opacity: 0.8;
}

.blog-filters {
  display: flex;
  gap: 1rem;
  justify-content: center;
  align-items: center;
  flex-wrap: wrap;
  max-width: 600px;
  margin: 0 auto;
}

.filter-group,
.search-group {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.filter-group label {
  font-size: 0.9rem;
  color: var(--theme-fg);
  font-weight: bold;
}

.filter-group select,
.search-group input {
  padding: 0.5rem 1rem;
  background: var(--theme-bg);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  color: var(--theme-fg);
  font-family: inherit;
}

.search-group {
  flex: 1;
  position: relative;
  min-width: 200px;
}

.search-group input {
  width: 100%;
  padding-right: 2.5rem;
}

.search-group svg {
  position: absolute;
  right: 1rem;
  top: 50%;
  transform: translateY(-50%);
  color: var(--theme-fg);
  opacity: 0.6;
  pointer-events: none;
}

.blog-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 2rem;
}

.loading,
.no-posts {
  grid-column: 1 / -1;
  text-align: center;
  padding: 3rem;
}

.loading {
  color: var(--theme-accent);
}

.loading svg.spin {
  animation: spin 1s linear infinite;
  margin-bottom: 0.5rem;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.error {
  grid-column: 1 / -1;
  text-align: center;
}

.error .window-body {
  padding: 2rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;
}

@media (max-width: 768px) {
  .blog-filters {
    flex-direction: column;
    align-items: stretch;
  }
  
  .blog-grid {
    grid-template-columns: 1fr;
  }
}
</style>