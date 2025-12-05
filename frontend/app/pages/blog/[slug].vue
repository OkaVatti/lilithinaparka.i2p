<template>
  <div class="max-w-4xl mx-auto">
    <div v-if="loading" class="terminal-box">
      <p class="text-text-secondary">loading<span class="blink">_</span></p>
    </div>
    
    <div v-else-if="error" class="terminal-box">
      <p class="text-red-500">error: {{ error }}</p>
      <NuxtLink to="/blog" class="text-link hover:text-accent mt-4 inline-block">
        ← back to blog
      </NuxtLink>
    </div>
    
    <article v-else-if="post">
      <div class="mb-8">
        <NuxtLink to="/blog" class="text-sm text-link hover:text-accent mb-4 inline-block">
          ← back to blog
        </NuxtLink>
        
        <h1 class="text-3xl font-bold mb-2">
          {{ post.title }}
        </h1>
        
        <div class="flex items-center gap-4 text-sm text-text-secondary mb-4">
          <span>{{ post.date }}</span>
          <span v-if="post.time">{{ post.time }}</span>
          <span v-if="authors.length > 0">by {{ authors.join(', ') }}</span>
        </div>
        
        <div class="flex flex-wrap gap-2 mb-4">
          <span 
            v-for="category in categories" 
            :key="category"
            class="text-xs px-2 py-1 bg-bg-primary text-accent border border-accent"
          >
            {{ category }}
          </span>
        </div>
        
        <div class="flex flex-wrap gap-2">
          <span 
            v-for="tag in tags" 
            :key="tag"
            class="text-xs text-link"
          >
            #{{ tag }}
          </span>
        </div>
      </div>
      
      <div 
        v-if="post.featured_image" 
        class="mb-8 border border-border"
      >
        <img 
          :src="post.featured_image" 
          :alt="post.title"
          class="w-full"
        />
      </div>
      
      <div 
        class="prose prose-invert max-w-none"
        v-html="renderedContent"
      ></div>
      
      <div class="mt-12 pt-8 border-t border-border">
        <div class="flex justify-between items-center">
          <NuxtLink to="/blog" class="text-link hover:text-accent">
            ← back to blog
          </NuxtLink>
          
          <div v-if="post.share" class="text-sm text-text-secondary">
            share this post
          </div>
        </div>
      </div>
    </article>
  </div>
</template>

<script setup lang="ts">
import { useBlogStore } from '../../../stores/blog'

const route = useRoute()
const blogStore = useBlogStore()
const { renderMarkdown } = useMarkdown()

const slug = computed(() => route.params.slug as string)
const loading = computed(() => blogStore.loading)
const error = computed(() => blogStore.error)
const post = computed(() => blogStore.currentPost)

const authors = computed(() => {
  if (!post.value?.authors) return []
  try {
    return JSON.parse(post.value.authors)
  } catch {
    return []
  }
})

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

const renderedContent = computed(() => {
  if (!post.value?.content) return ''
  return renderMarkdown(post.value.content)
})

onMounted(async () => {
  await blogStore.fetchPostBySlug(slug.value)
})

watch(slug, async (newSlug) => {
  await blogStore.fetchPostBySlug(newSlug)
})

useHead(() => ({
  title: post.value ? `${post.value.title} - Lilith Parker` : 'Blog Post - Lilith Parker',
  meta: [
    { name: 'description', content: post.value?.summary || '' }
  ]
}))
</script>