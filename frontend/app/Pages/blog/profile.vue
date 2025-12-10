<template>
  <div class="blog-post-page">
    <div v-if="blogStore.loading" class="loading">Loading post</div>
    
    <div v-else-if="blogStore.error" class="error">
      <FeatherIcon name="alert-circle" size="20" />
      <span>{{ blogStore.error }}</span>
      <NuxtLink to="/blog" class="btn mt-2">
        <FeatherIcon name="arrow-left" size="18" />
        <span>Back to Blog</span>
      </NuxtLink>
    </div>
    
    <article v-else-if="post" class="blog-post">
      <header class="post-header">
        <NuxtLink to="/blog" class="back-link">
          <FeatherIcon name="arrow-left" size="18" />
          <span>Back to Blog</span>
        </NuxtLink>
        
        <h1 class="post-title">{{ post.title }}</h1>
        
        <div class="post-meta">
          <span class="meta-item">
            <FeatherIcon name="calendar" size="16" />
            {{ formatDate(post.date) }}
          </span>
          
          <span v-if="authors.length" class="meta-item">
            <FeatherIcon name="user" size="16" />
            {{ authors.join(', ') }}
          </span>
          
          <span v-if="categories.length" class="meta-item">
            <FeatherIcon name="folder" size="16" />
            {{ categories.join(', ') }}
          </span>
        </div>
        
        <div v-if="post.featured_image" class="featured-image">
          <img :src="post.featured_image" :alt="post.title" />
        </div>
        
        <p v-if="post.summary" class="post-summary">{{ post.summary }}</p>
      </header>
      
      <div class="post-content">
        <MarkdownRenderer :content="post.content" />
      </div>
      
      <footer class="post-footer">
        <div v-if="tags.length" class="post-tags">
          <h3>Tags:</h3>
          <div class="tags-list">
            <span v-for="tag in tags" :key="tag" class="tag">
              <FeatherIcon name="tag" size="14" />
              {{ tag }}
            </span>
          </div>
        </div>
        
        <div class="post-actions">
          <NuxtLink to="/blog" class="btn">
            <FeatherIcon name="arrow-left" size="18" />
            <span>Back to Blog</span>
          </NuxtLink>
        </div>
      </footer>
    </article>
  </div>
</template>

<script setup lang="ts">
import { useBlogStore } from '~~/stores/blog'

const route = useRoute()
const blogStore = useBlogStore()

const post = computed(() => blogStore.currentPost)

const categories = computed(() => {
  if (!post.value?.categories) return []
  try {
    return JSON.parse(post.value.categories)
  } catch {
    return []
  }
})

const tags = computed(() => {
  if (!post.value?.tags) return []
  try {
    return JSON.parse(post.value.tags)
  } catch {
    return []
  }
})

const authors = computed(() => {
  if (!post.value?.authors) return []
  try {
    return JSON.parse(post.value.authors)
  } catch {
    return []
  }
})

const formatDate = (dateStr: string) => {
  try {
    return new Date(dateStr).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'long',
      day: 'numeric'
    })
  } catch {
    return dateStr
  }
}

onMounted(() => {
  const slug = route.params.slug as string
  blogStore.fetchPostBySlug(slug)
})

useHead(() => ({
  title: post.value?.title || 'Blog Post',
  meta: [
    { name: 'description', content: post.value?.summary || '' }
  ]
}))
</script>

<style scoped>
.blog-post {
  max-width: 800px;
  margin: 0 auto;
}

.post-header {
  margin-bottom: 3rem;
}

.back-link {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 2rem;
  padding: 0.5rem 1rem;
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  color: var(--theme-accent);
  transition: all 0.2s;
}

.back-link:hover {
  border-color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.1);
}

.post-title {
  font-size: 3rem;
  margin-bottom: 1.5rem;
  color: var(--theme-primary);
  line-height: 1.2;
}

.post-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 1.5rem;
  margin-bottom: 2rem;
  color: var(--theme-fg);
  opacity: 0.7;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.featured-image {
  width: 100%;
  margin: 2rem 0;
  border-radius: 8px;
  overflow: hidden;
  border: 2px solid var(--theme-border);
}

.featured-image img {
  width: 100%;
  height: auto;
  display: block;
}

.post-summary {
  font-size: 1.3rem;
  line-height: 1.6;
  color: var(--theme-fg);
  opacity: 0.9;
  padding: 1.5rem;
  background: rgba(189, 147, 249, 0.1);
  border-left: 4px solid var(--theme-primary);
  border-radius: 4px;
}

.post-content {
  margin-bottom: 3rem;
  font-size: 1.1rem;
}

.post-footer {
  padding-top: 2rem;
  border-top: 2px solid var(--theme-border);
}

.post-tags {
  margin-bottom: 2rem;
}

.post-tags h3 {
  margin-bottom: 1rem;
  color: var(--theme-primary);
}

.tags-list {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
}

.tag {
  display: flex;
  align-items: center;
  gap: 0.25rem;
  padding: 0.5rem 1rem;
  background: rgba(189, 147, 249, 0.1);
  border: 1px solid var(--theme-primary);
  border-radius: 4px;
  color: var(--theme-primary);
}

.post-actions {
  text-align: center;
}

.post-actions .btn {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
}

@media (max-width: 768px) {
  .post-title {
    font-size: 2rem;
  }
  
  .post-summary {
    font-size: 1.1rem;
  }
  
  .post-content {
    font-size: 1rem;
  }
}
</style>