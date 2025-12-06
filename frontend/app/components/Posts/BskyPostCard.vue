<!-- app/components/Posts/BskyPostCard.vue -->
<template>
  <div class="bsky-post-card">
    <div class="p-4 rounded-lg border border-border bg-bg-secondary/30">
      <!-- Post Header -->
      <div class="flex items-start justify-between mb-3">
        <div class="flex items-center gap-2">
          <div class="w-8 h-8 rounded-full bg-accent/10 flex items-center justify-center">
            <span class="text-accent text-sm">𝕏</span>
          </div>
          <div>
            <div class="bsky-author">{{ post.author }}</div>
            <div class="bsky-timestamp">{{ post.timestamp }}</div>
          </div>
        </div>
        <a v-if="post.author" :href="`https://bsky.app/profile/${post.author.replace('@', '')}`" 
           target="_blank" class="text-xs text-link hover:text-accent transition-colors">
          view
        </a>
      </div>
      
      <!-- Post Content -->
      <div class="bsky-content mb-4">
        <p class="text-text-primary whitespace-pre-wrap">{{ post.text }}</p>
      </div>
      
      <!-- Media -->
      <div v-if="post.hasMedia && post.media" class="bsky-media mb-4">
        <img :src="post.media[0]" alt="Bluesky media" class="rounded-lg max-h-64 object-cover w-full">
      </div>
      
      <!-- Post Stats -->
      <div class="bsky-stats">
        <div class="flex items-center gap-6 text-text-secondary text-sm">
          <span class="flex items-center gap-1">
            <span>♥</span>
            <span>{{ post.likes }}</span>
          </span>
          <span class="flex items-center gap-1">
            <span>🔄</span>
            <span>{{ post.reposts }}</span>
          </span>
          <span class="flex items-center gap-1">
            <span>💬</span>
            <span>{{ post.replies }}</span>
          </span>
        </div>
      </div>
      
      <!-- Hashtags -->
      <div v-if="extractHashtags(post.text).length > 0" class="mt-3 flex flex-wrap gap-1">
        <span v-for="tag in extractHashtags(post.text)" :key="tag" class="bsky-tag">
          {{ tag }}
        </span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  post: {
    id: number
    text: string
    author: string
    likes: number
    reposts: number
    replies: number
    timestamp: string
    hasMedia?: boolean
    media?: string[]
  }
}>()

const extractHashtags = (text: string): string[] => {
  const hashtags = text.match(/#[a-zA-Z0-9_]+/g)
  return hashtags ? hashtags.slice(0, 3) : []
}
</script>

<style scoped>
.bsky-author {
  @apply text-sm font-bold text-accent;
}

.bsky-timestamp {
  @apply text-xs text-text-secondary;
}

.bsky-content {
  @apply text-sm leading-relaxed;
}

.bsky-stats {
  @apply border-t border-border pt-3;
}

.bsky-tag {
  @apply px-2 py-0.5 rounded-full bg-bg-primary/50 text-xs text-text-secondary;
}
</style>