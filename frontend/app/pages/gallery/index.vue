<!-- app/pages/gallery/index.vue -->
<template>
  <div class="gallery-page">
    <!-- Gallery Header -->
    <div class="gallery-header mb-8">
      <h1 class="text-4xl font-bold text-accent mb-4 font-mono">$ gallery</h1>
      <p class="text-text-secondary font-mono">
        > my artwork, digital creations, and visual experiments
      </p>
      
      <!-- Gallery Stats -->
      <div class="flex flex-wrap gap-4 mt-6">
        <div class="gallery-stat">
          <div class="stat-number">{{ totalArtworks }}</div>
          <div class="stat-label">Artworks</div>
        </div>
        <div class="gallery-stat">
          <div class="stat-number">{{ categories.length }}</div>
          <div class="stat-label">Categories</div>
        </div>
        <div class="gallery-stat">
          <div class="stat-number">2018</div>
          <div class="stat-label">Since</div>
        </div>
      </div>
    </div>
    
    <!-- Category Filter -->
    <div class="flex flex-wrap gap-2 mb-8">
      <button
        v-for="category in categories"
        :key="category"
        @click="setActiveCategory(category)"
        class="gallery-filter"
        :class="{ 'active': activeCategory === category }"
      >
        {{ category }}
      </button>
    </div>
    
    <!-- Masonry Grid -->
    <div class="masonry-grid">
      <ArtCard
        v-for="art in filteredArtworks"
        :key="art.id"
        :art="art"
        :masonry="true"
        @click="viewArtwork(art)"
      />
    </div>
    
    <!-- No Artworks Message -->
    <div v-if="filteredArtworks.length === 0" class="text-center py-12">
      <p class="text-text-secondary font-mono">> no artworks found in this category</p>
    </div>
    
    <!-- Gallery Information -->
    <div class="mt-12 p-6 rounded-xl bg-bg-secondary/50 border border-accent/10">
      <h3 class="text-xl font-bold mb-4 font-mono text-accent">$ about the gallery</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <p class="text-text-primary mb-4">
            This gallery serves as my personal digital exhibition space. Since I can't upload my artwork to mainstream platforms due to content policies and censorship concerns, this is where I showcase my creations.
          </p>
          <p class="text-text-secondary text-sm">
            All artworks are original creations unless otherwise noted. I work with various mediums including digital painting, pixel art, 3D modeling, and generative art.
          </p>
        </div>
        <div>
          <h4 class="text-lg font-bold mb-3 font-mono text-accent">Art Categories</h4>
          <ul class="space-y-2">
            <li v-for="category in categories" :key="category" class="flex items-center gap-2">
              <span class="text-accent">→</span>
              <span class="text-text-primary">{{ category }}</span>
              <span class="text-text-secondary text-sm">
                ({{ artworks.filter(a => a.category === category).length }})
              </span>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import ArtCard from '../../components/Posts/ArtCard.vue'
import { ref, computed } from 'vue'
import { useRouter } from '#imports'

definePageMeta({
  layout: 'artwork'
})

const router = useRouter()

// Sample artwork data
const artworks = [
  { id: 1, title: 'Digital Cat', image: '/images/art/cat-digital.jpg', category: 'Digital', tags: ['cat', 'anime', 'cute'], description: 'A digital painting of my cat in anime style', year: 2024 },
  { id: 2, title: 'Cyberpunk City', image: '/images/art/cyberpunk-city.jpg', category: 'Digital', tags: ['cyberpunk', 'cityscape', 'neon'], description: 'Futuristic cityscape with neon lights', year: 2023 },
  { id: 3, title: 'Forest Spirit', image: '/images/art/forest-spirit.jpg', category: 'Fantasy', tags: ['fantasy', 'nature', 'spirit'], description: 'Mystical forest spirit in an enchanted woods', year: 2023 },
  { id: 4, title: 'Binary Dreams', image: '/images/art/binary-dreams.jpg', category: 'Abstract', tags: ['abstract', 'code', 'digital'], description: 'Abstract representation of digital consciousness', year: 2024 },
  { id: 5, title: 'Pixel Space', image: '/images/art/pixel-space.jpg', category: 'Pixel', tags: ['pixel', 'space', 'retro'], description: 'Retro pixel art space scene', year: 2022 },
  { id: 6, title: 'Chaos Theory', image: '/images/art/chaos-theory.jpg', category: 'Abstract', tags: ['chaos', 'fractal', 'mathematical'], description: 'Visualization of chaotic mathematical systems', year: 2023 },
  { id: 7, title: 'Neon Samurai', image: '/images/art/neon-samurai.jpg', category: 'Digital', tags: ['samurai', 'neon', 'cyberpunk'], description: 'Cyberpunk samurai in neon-lit alley', year: 2024 },
  { id: 8, title: 'Crystal Cavern', image: '/images/art/crystal-cavern.jpg', category: 'Fantasy', tags: ['crystal', 'cavern', 'magical'], description: 'Magical crystal cavern with glowing mushrooms', year: 2023 },
  { id: 9, title: 'Data Flow', image: '/images/art/data-flow.jpg', category: 'Abstract', tags: ['data', 'network', 'flow'], description: 'Visualization of data flowing through networks', year: 2024 },
  { id: 10, title: 'Retro Computer', image: '/images/art/retro-computer.jpg', category: 'Pixel', tags: ['computer', 'retro', '80s'], description: 'Pixel art of an 80s computer setup', year: 2022 },
  { id: 11, title: 'Ocean Depths', image: '/images/art/ocean-depths.jpg', category: 'Digital', tags: ['ocean', 'deep-sea', 'creatures'], description: 'Deep sea scene with bioluminescent creatures', year: 2023 },
  { id: 12, title: 'Circuit Garden', image: '/images/art/circuit-garden.jpg', category: 'Abstract', tags: ['circuit', 'garden', 'nature-tech'], description: 'Hybrid of organic nature and electronic circuits', year: 2024 }
]

const categories = ['All', 'Digital', 'Fantasy', 'Abstract', 'Pixel']
const activeCategory = ref('All')

const filteredArtworks = computed(() => {
  if (activeCategory.value === 'All') return artworks
  return artworks.filter(art => art.category === activeCategory.value)
})

const totalArtworks = computed(() => artworks.length)

const setActiveCategory = (category: string) => {
  activeCategory.value = category
}

const viewArtwork = (art: any) => {
  router.push(`/gallery/${art.id}`)
}
</script>

<style scoped>
.gallery-page {
  @apply max-w-7xl mx-auto;
}

.gallery-header {
  @apply pb-8 border-b border-accent/20;
}

.gallery-stat {
  @apply px-6 py-3 rounded-lg bg-bg-secondary/30 border border-accent/10;
}

.gallery-stat .stat-number {
  @apply text-2xl font-bold text-accent font-mono;
}

.gallery-stat .stat-label {
  @apply text-text-secondary text-sm mt-1;
}

.gallery-filter {
  @apply px-4 py-2 rounded-full border border-accent/20 text-sm font-mono
         text-text-secondary hover:text-accent hover:border-accent/40
         transition-all duration-200;
}

.gallery-filter.active {
  @apply bg-accent/10 text-accent border-accent/40;
}

.masonry-grid {
  column-count: 1;
  column-gap: 1rem;
}

@media (min-width: 640px) {
  .masonry-grid {
    column-count: 2;
  }
}

@media (min-width: 1024px) {
  .masonry-grid {
    column-count: 3;
  }
}

@media (min-width: 1280px) {
  .masonry-grid {
    column-count: 4;
  }
}

.masonry-grid > * {
  break-inside: avoid;
  margin-bottom: 1rem;
}
</style>