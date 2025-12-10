<template>
  <article class="bsky-post-card card">
    <div class="post-header">
      <div class="author-info">
        <FeatherIcon name="cloud" size="18" />
        <span class="author-handle">@{{ post.author }}</span>
      </div>
      <span class="post-date">
        {{ formatDate(post.posted_at) }}
      </span>
    </div>
    
    <div class="post-content">
      <p>{{ post.text }}</p>
    </div>
    
    <div v-if="post.has_media && mediaUrls.length" class="post-media">
      <img
        v-for="(url, index) in mediaUrls"
        :key="index"
        :src="url"
        :alt="`Media ${index + 1}`"
        loading="lazy"
      />
    </div>
    
    <div class="post-stats">
      <span class="stat">
        <FeatherIcon name="message-circle" size="14" />
        {{ post.reply_count }}
      </span>
      <span class="stat">
        <FeatherIcon name="repeat" size="14" />
        {{ post.repost_count }}
      </span>
      <span class="stat">
        <FeatherIcon name="heart" size="14" />
        {{ post.like_count }}
      </span>
      <span v-if="post.quote_count > 0" class="stat">
        <FeatherIcon name="message-square" size="14" />
        {{ post.quote_count }}
      </span>
    </div>
    
    <a :href="getPostUrl(post.uri)" target="_blank" rel="noopener" class="view-on-bsky">
      View on BlueSky <FeatherIcon name="external-link" size="14" />
    </a>
  </article>
</template>

<script setup lang="ts">
import type { BskyPost } from '~~/types'

const props = defineProps<{
  post: BskyPost
}>()

const mediaUrls = computed(() => {
  try {
    return JSON.parse(props.post.media_urls || '[]')
  } catch {
    return []
  }
})

const formatDate = (dateStr: string) => {
  try {
    const date = new Date(dateStr)
    const now = new Date()
    const diffMs = now.getTime() - date.getTime()
    const diffMins = Math.floor(diffMs / 60000)
    
    if (diffMins < 60) return `${diffMins}m ago`
    if (diffMins < 1440) return `${Math.floor(diffMins / 60)}h ago`
    if (diffMins < 10080) return `${Math.floor(diffMins / 1440)}d ago`
    
    return date.toLocaleDateString('en-US', {
      month: 'short',
      day: 'numeric',
      year: 'numeric'
    })
  } catch {
    return dateStr
  }
}

const getPostUrl = (uri: string) => {
  const parts = uri.split('/')
  const did = parts[2]
  const postId = parts[parts.length - 1]
  return `https://bsky.app/profile/${did}/post/${postId}`
}
</script>

<style scoped>
.bsky-post-card {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.post-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-bottom: 0.5rem;
  border-bottom: 1px solid var(--theme-border);
}

.author-info {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: var(--theme-accent);
  font-weight: bold;
}

.post-date {
  font-size: 0.85rem;
  color: var(--theme-fg);
  opacity: 0.6;
}

.post-content p {
  margin: 0;
  line-height: 1.6;
  white-space: pre-wrap;
  word-wrap: break-word;
}

.post-media {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 0.5rem;
}

.post-media img {
  width: 100%;
  height: 200px;
  object-fit: cover;
  border-radius: 4px;
  border: 1px solid var(--theme-border);
}

.post-stats {
  display: flex;
  gap: 1.5rem;
  padding-top: 0.75rem;
  border-top: 1px solid var(--theme-border);
}

.stat {
  display: flex;
  align-items: center;
  gap: 0.25rem;
  color: var(--theme-fg);
  opacity: 0.7;
  font-size: 0.9rem;
}

.view-on-bsky {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  color: var(--theme-accent);
  font-size: 0.9rem;
  margin-top: 0.5rem;
}

.view-on-bsky:hover {
  color: var(--theme-secondary);
}
</style>