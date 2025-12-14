<template>
  <PageContainer>
    <div v-if="art">
      <ArtPostContainer
        :title="art.title"
        :image="art.image"
        :description="art.description"
        :author="art.author"
        :date="art.date"
        :tags="art.tags"
      />
    </div>

    <div v-else class="not-found">
      <h2>Artwork not found</h2>
      <NuxtLink to="/art/gallery" class="btn">Back to gallery</NuxtLink>
    </div>
  </PageContainer>
</template>

<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router'
import PageContainer from '~/components/layout/PageContainer.vue'
import ArtPostContainer from '~/components/art/ArtPostContainer.vue'

const route = useRoute()
const router = useRouter()
const slug = route.params.slug as string

// NOTE: replace this with an actual fetch to your backend or content source.
const sample = [
  { slug: 'moon-study', title: 'Moon Study', image: '/images/art/moon1.jpg', author: 'Lilith', date: '2025-10-02', description: '<p>A nocturne study in cool tones.</p>', tags: ['digital','study'] },
  { slug: 'forest-glow', title: 'Forest Glow', image: '/images/art/forest.jpg', author: 'Lilith', date: '2025-09-12', description: '<p>Light through leaves. Mixed media.</p>', tags: ['mixed-media'] },
  { slug: 'glitch-portrait', title: 'Glitch Portrait', image: '/images/art/glitch.jpg', author: 'Lilith', date: '2025-08-20', description: '<p>Portrait with intentional artifacts.</p>', tags: ['experimental'] }
]

const art = sample.find(a => a.slug === slug) ?? null
</script>

<style scoped>
@import '~/assets/css/system.css';

.not-found { text-align:center; padding:2rem; color:var(--theme-muted); }
</style>
