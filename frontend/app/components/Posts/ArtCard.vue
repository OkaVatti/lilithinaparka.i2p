<!-- app/components/Posts/ArtCard.vue -->
<template>
  <div class="art-card" :class="{ 'masonry-item': masonry }">
    <div class="relative overflow-hidden rounded-lg border border-border bg-bg-secondary">
      <img 
        :src="art.image" 
        :alt="art.title"
        class="w-full h-48 object-cover hover:scale-110 transition-transform duration-500"
        loading="lazy"
      />
      <div class="absolute inset-0 bg-gradient-to-t from-bg-primary/80 to-transparent opacity-0 hover:opacity-100 transition-opacity duration-300">
        <div class="absolute bottom-0 left-0 right-0 p-4">
          <h3 class="text-lg font-bold text-text-primary mb-1">{{ art.title }}</h3>
          <p class="text-text-secondary text-sm mb-2">{{ art.description }}</p>
          <div class="flex flex-wrap gap-1">
            <span v-for="tag in art.tags.slice(0, 2)" :key="tag" class="art-tag">
              {{ tag }}
            </span>
            <span v-if="art.tags.length > 2" class="art-tag">+{{ art.tags.length - 2 }}</span>
          </div>
          <div v-if="art.year" class="mt-2 text-xs text-text-secondary font-mono">
            {{ art.year }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  art: {
    id: number
    title: string
    image: string
    category: string
    tags: string[]
    description?: string
    year?: number
  }
  masonry?: boolean
}>()
</script>

<style scoped>
.art-card.masonry-item {
  break-inside: avoid;
  margin-bottom: 1rem;
}

.art-tag {
  @apply px-2 py-0.5 rounded-full bg-bg-primary/50 text-xs text-text-secondary 
         border border-border;
}
</style>