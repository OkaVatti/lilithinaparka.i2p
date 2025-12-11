<template>
  <article class="media-item" @click="openLightbox">
    <div class="media-preview">
      <img 
        v-if="isImage"
        :src="thumbnailUrl" 
        :alt="item.title"
        loading="lazy"
      />
      <video 
        v-else-if="isVideo"
        :src="mediaUrl"
        :poster="thumbnailUrl"
        preload="metadata"
      ></video>
      
      <div class="media-overlay">
        <div class="media-type">
          <FeatherIcon :name="mediaIcon" size="20" />
        </div>
        <div v-if="item.duration" class="media-duration">
          {{ formatDuration(item.duration) }}
        </div>
      </div>
    </div>
    
    <div class="media-info">
      <h4 class="media-title">{{ item.title }}</h4>
      <p v-if="item.description" class="media-description">{{ item.description }}</p>
      
      <div class="media-meta">
        <span v-if="item.artist" class="meta-item">
          <FeatherIcon name="user" size="14" />
          {{ item.artist }}
        </span>
        <span v-if="item.year" class="meta-item">
          <FeatherIcon name="calendar" size="14" />
          {{ item.year }}
        </span>
        <span class="meta-item">
          <FeatherIcon name="eye" size="14" />
          {{ item.views }}
        </span>
        <span class="meta-item">
          <FeatherIcon name="heart" size="14" />
          {{ item.likes }}
        </span>
      </div>
      
      <div v-if="tags.length > 0" class="media-tags">
        <span v-for="tag in tags.slice(0, 3)" :key="tag" class="tag">
          {{ tag }}
        </span>
      </div>
      
      <div v-if="hasExif" class="media-exif">
        <button @click.stop="showExif = !showExif" class="exif-toggle">
          <FeatherIcon name="info" size="14" />
          <span>{{ showExif ? 'Hide' : 'Show' }} EXIF</span>
        </button>
        
        <div v-if="showExif" class="exif-data">
          <div v-for="(value, key) in exifData" :key="key" class="exif-item">
            <span class="exif-key">{{ formatExifKey(key.toString()) }}:</span>
            <span class="exif-value">{{ value }}</span>
          </div>
        </div>
      </div>
    </div>
  </article>
</template>

<script setup lang="ts">
import type { MediaItem } from '~~/types'

const props = defineProps<{
  item: MediaItem
}>()

const { apiFetch } = useApi()

const showExif = ref(false)

const isImage = computed(() => props.item.mime_type.startsWith('image/'))
const isVideo = computed(() => props.item.mime_type.startsWith('video/'))

const mediaIcon = computed(() => {
  if (isImage.value) return 'image'
  if (isVideo.value) return 'video'
  return 'file'
})

const thumbnailUrl = computed(() => {
  if (props.item.thumbnail) {
    return `/media${props.item.thumbnail}`
  }
  return `/media/uploads/${props.item.file_name}`
})

const mediaUrl = computed(() => {
  return `/media/uploads/${props.item.file_name}`
})

const tags = computed(() => {
  try {
    return JSON.parse(props.item.tags || '[]')
  } catch {
    return []
  }
})

const hasExif = computed(() => {
  return props.item.exif && props.item.exif !== '{}'
})

const exifData = computed(() => {
  if (!hasExif.value) return {}
  try {
    return JSON.parse(props.item.exif)
  } catch {
    return {}
  }
})

const formatDuration = (seconds: number) => {
  const mins = Math.floor(seconds / 60)
  const secs = seconds % 60
  return `${mins}:${secs.toString().padStart(2, '0')}`
}

const formatExifKey = (key: string) => {
  return key.split('_').map(word => 
    word.charAt(0).toUpperCase() + word.slice(1)
  ).join(' ')
}

const openLightbox = async () => {
  // Track view
  try {
    await apiFetch(`/media/${props.item.id}/view`, { method: 'POST' })
  } catch (error) {
    console.error('Failed to track view:', error)
  }
  
  // Emit event to parent to open lightbox
  // Or use a composable/store for lightbox management
}
</script>

<style scoped>
.media-item {
  break-inside: avoid;
  background: rgba(0, 0, 0, 0.2);
  border: 1px solid var(--theme-border);
  border-radius: 8px;
  overflow: hidden;
  cursor: pointer;
  transition: all 0.2s;
}

.media-item:hover {
  border-color: var(--theme-primary);
  transform: translateY(-2px);
  box-shadow: 0 4px 8px var(--theme-shadow);
}

.media-preview {
  position: relative;
  width: 100%;
  overflow: hidden;
  background: #000;
}

.media-preview img,
.media-preview video {
  width: 100%;
  height: auto;
  display: block;
}

.media-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(to bottom, rgba(0,0,0,0.5) 0%, transparent 50%, rgba(0,0,0,0.5) 100%);
  opacity: 0;
  transition: opacity 0.2s;
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  padding: 0.5rem;
}

.media-item:hover .media-overlay {
  opacity: 1;
}

.media-type {
  padding: 0.25rem 0.5rem;
  background: rgba(0, 0, 0, 0.8);
  border-radius: 4px;
  color: white;
}

.media-duration {
  padding: 0.25rem 0.5rem;
  background: rgba(0, 0, 0, 0.8);
  border-radius: 4px;
  color: white;
  font-size: 0.85rem;
}

.media-info {
  padding: 1rem;
}

.media-title {
  margin: 0 0 0.5rem 0;
  font-size: 1.1rem;
  color: var(--theme-primary);
}

.media-description {
  margin: 0 0 0.75rem 0;
  font-size: 0.9rem;
  color: var(--theme-fg);
  opacity: 0.8;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.media-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
  margin-bottom: 0.75rem;
  font-size: 0.85rem;
  color: var(--theme-fg);
  opacity: 0.7;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 0.25rem;
}

.media-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 0.25rem;
  margin-bottom: 0.75rem;
}

.tag {
  padding: 0.25rem 0.5rem;
  background: rgba(189, 147, 249, 0.1);
  border: 1px solid var(--theme-primary);
  border-radius: 4px;
  font-size: 0.75rem;
  color: var(--theme-primary);
}

.exif-toggle {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem;
  background: rgba(189, 147, 249, 0.1);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  color: var(--theme-fg);
  cursor: pointer;
  transition: all 0.2s;
  width: 100%;
  font-family: inherit;
}

.exif-toggle:hover {
  border-color: var(--theme-primary);
}

.exif-data {
  margin-top: 0.5rem;
  padding: 0.75rem;
  background: rgba(0, 0, 0, 0.3);
  border-radius: 4px;
  font-size: 0.85rem;
}

.exif-item {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 0.5rem;
  padding: 0.25rem 0;
  border-bottom: 1px solid var(--theme-border);
}

.exif-item:last-child {
  border-bottom: none;
}

.exif-key {
  color: var(--theme-fg);
  opacity: 0.7;
}

.exif-value {
  color: var(--theme-accent);
  text-align: right;
}
</style>