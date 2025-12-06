<!-- app/components/Posts/SocialLink.vue -->
<template>
  <a :href="link.url" target="_blank" class="social-link-card block">
    <div class="p-4 rounded-lg border border-border bg-bg-secondary/30 hover:border-accent/30 transition-all duration-300">
      <div class="flex items-start justify-between mb-3">
        <div class="flex items-center gap-3">
          <div class="w-10 h-10 rounded-full flex items-center justify-center" :class="link.color">
            <span class="text-xl">{{ link.icon }}</span>
          </div>
          <div>
            <h3 class="text-lg font-bold" :class="link.color">{{ link.platform }}</h3>
            <p class="text-text-secondary text-sm">{{ link.username }}</p>
          </div>
        </div>
      </div>
      <p class="text-text-secondary text-sm mb-4">{{ link.description }}</p>
      <div class="text-xs text-link hover:text-accent transition-colors font-mono">
        visit →
      </div>
    </div>
  </a>
</template>

<script setup lang="ts">
defineProps<{
  link: {
    platform: string
    username: string
    url: string
    icon: string
    color: string
    description: string
  }
}>()
</script>