<template>
  <article class="terminal-box mb-4">
    <div class="flex flex-col gap-3">
      <div class="flex items-center gap-2">
        <span class="text-sm font-bold text-accent">@{{ post.author }}</span>
        <span class="text-xs text-text-secondary">
          {{ formattedDate }}
        </span>
      </div>
      
      <p class="text-sm whitespace-pre-wrap">{{ post.text }}</p>
      
      <div v-if="hasMedia && mediaUrls.length > 0" class="grid grid-cols-2 gap-2">
        <img 
          v-for="(url, index) in mediaUrls" 
          :key="index"
          :src="url"
          :alt="`Media ${index + 1}`"
          class="w-full border border-border"
        />
      </div>
      
      <div class="flex gap-6 text-xs text-text-secondary">
        <span>↩ {{ post.reply_count }}</span>
        <span>↻ {{ post.repost_count }}</span>
        <span>♥ {{ post.like_count }}</span>
        <span v-if="post.quote_count > 0">💬 {{ post.quote_count }}</span>
      </div>
    </div>
  </article>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { BskyPost } from '../../../types'

const props = defineProps<{
  post: BskyPost
}>()

const formattedDate = computed(() => {
  const date = new Date(props.post.posted_at)
  return date.toLocaleDateString('en-US', { 
    year: 'numeric', 
    month: 'short', 
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
})

const hasMedia = computed(() => props.post.has_media)

const mediaUrls = computed(() => {
  try {
    return JSON.parse(props.post.media_urls)
  } catch {
    return []
  }
})
</script>