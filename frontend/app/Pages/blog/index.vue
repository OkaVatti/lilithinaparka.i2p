<template>
  <div class="blog-page">
    <header class="page-header">
      <h1>Blog</h1>
      <p>Thoughts, writings, and musings</p>
    </header>
    
    <div class="blog-filters">
      <div class="filter-group">
        <label>Filter by Category:</label>
        <select v-model="selectedCategory" @change="filterPosts">
          <option value="">All Categories</option>
          <option value="Casual">Casual</option>
          <option value="Interlude">Interlude</option>
          <option value="Serious">Serious</option>
        </select>
      </div>
    </div>
    
    <div v-if="blogStore.loading" class="loading">Loading posts</div>
    
    <div v-else-if="blogStore.error" class="error">
      <FeatherIcon name="alert-circle" size="20" />
      <span>{{ blogStore.error }}</span>
    </div>
    
    <div v-else-if="displayedPosts.length" class="grid grid-2">
      <BlogPostCard
        v-for="post in displayedPosts"
        :key="post.id"
        :post="post"
      />
    </div>
    
    <div v-else class="no-posts">
      <FeatherIcon name="file-text" size="48" />
      <p>No posts found.</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useBlogStore } from '~~/stores/blog'

const blogStore = useBlogStore()
const selectedCategory = ref('')

const displayedPosts = computed(() => blogStore.publishedPosts)

const filterPosts = async () => {
  if (selectedCategory.value) {
    await blogStore.fetchPostsByCategory(selectedCategory.value)
  } else {
    await blogStore.fetchPosts()
  }
}

onMounted(() => {
  blogStore.fetchPosts()
})
</script>

<style scoped>
.page-header {
  text-align: center;
  margin-bottom: 3rem;
  padding-bottom: 2rem;
  border-bottom: 2px solid var(--theme-border);
}

.page-header h1 {
  font-size: 3rem;
  margin-bottom: 0.5rem;
  color: var(--theme-primary);
}

.page-header p {
  font-size: 1.2rem;
  color: var(--theme-fg);
  opacity: 0.7;
}

.blog-filters {
  margin-bottom: 2rem;
  display: flex;
  justify-content: center;
}

.filter-group {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.filter-group label {
  color: var(--theme-fg);
}

.filter-group select {
  padding: 0.5rem 1rem;
  background: var(--theme-bg);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  color: var(--theme-fg);
  font-family: inherit;
  cursor: pointer;
}

.filter-group select:focus {
  outline: none;
  border-color: var(--theme-primary);
}

.no-posts {
  text-align: center;
  padding: 4rem 2rem;
  color: var(--theme-fg);
  opacity: 0.6;
}

.no-posts p {
  margin-top: 1rem;
  font-size: 1.2rem;
}
</style>