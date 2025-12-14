<template>
  <div class="page">
    <PageContainer>
      <header class="gallery-header">
        <h1>Art Gallery</h1>
        <p class="subtitle">A curated masonry of recent art posts.</p>
      </header>

      <div class="masonry" aria-live="polite">
        <ArtGalleryCard
          v-for="art in arts"
          :key="art.slug"
          :title="art.title"
          :image="art.image"
          :author="art.author"
          :date="art.date"
          :to="`/art/${art.slug}`"
        />
      </div>
    </PageContainer>
  </div>
</template>

<script setup lang="ts">
import PageContainer from '~/components/layout/PageContainer.vue'
import ArtGalleryCard from '~/components/art/ArtGalleryCard.vue'

// example local data — replace with useAsyncData/useFetch to load from API
const arts = [
  { slug: 'moon-study', title: 'Moon Study', image: '/images/art/moon1.jpg', author: 'Lilith', date: '2025-10-02' },
  { slug: 'forest-glow', title: 'Forest Glow', image: '/images/art/forest.jpg', author: 'Lilith', date: '2025-09-12' },
  { slug: 'glitch-portrait', title: 'Glitch Portrait', image: '/images/art/glitch.jpg', author: 'Lilith', date: '2025-08-20' },
  { slug: 'ocean-echo', title: 'Ocean Echo', image: '/images/art/ocean.jpg', author: 'Lilith', date: '2025-01-11' },
  { slug: 'city-night', title: 'City Night', image: '/images/art/city.jpg', author: 'Lilith', date: '2025-03-06' }
]
</script>

<style scoped>
@import '~/assets/css/system.css';

.gallery-header h1 { margin:0; color:var(--theme-fg); }
.subtitle { color:var(--theme-muted); margin-top:.2rem; }

.masonry {
  column-count: 3;
  column-gap: 1rem;
  margin-top: 1rem;
}

/* responsive column counts */
@media (max-width: 1100px) { .masonry { column-count: 2; } }
@media (max-width: 700px) { .masonry { column-count: 1; } }

/* ensure gallery cards break inside column nicely */
.masonry > * { display: inline-block; width: 100%; margin-bottom: 12px; break-inside: avoid; }
</style>
