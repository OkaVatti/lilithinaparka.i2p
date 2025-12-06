<!-- app/layouts/ArtworkLayout.vue -->
<template>
  <div class="min-h-screen bg-bg-primary text-text-primary">
    <!-- Navigation -->
    <header class="sticky top-0 z-50 backdrop-blur-lg bg-bg-primary/80 border-b border-border">
      <div class="container mx-auto px-4 py-4">
        <div class="flex items-center justify-between">
          <NuxtLink to="/" class="font-bold text-xl font-mono hover:text-accent transition-colors">
            Gallery
          </NuxtLink>
          <nav class="flex items-center gap-6">
            <NuxtLink to="/" class="text-sm hover:text-accent transition-colors">Home</NuxtLink>
            <NuxtLink to="/artwork" class="text-sm hover:text-accent transition-colors">Artwork</NuxtLink>
            <NuxtLink to="/blog" class="text-sm hover:text-accent transition-colors">Blog</NuxtLink>
            <NuxtLink to="/profile" class="text-sm hover:text-accent transition-colors">Profile</NuxtLink>
          </nav>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="container mx-auto px-4 py-8">
      <div class="mb-8">
        <h1 class="text-4xl font-bold font-mono mb-2">Artwork Gallery</h1>
        <p class="text-text-secondary">Digital art, illustrations, and creative projects</p>
      </div>
      
      <!-- Artwork-specific content -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <slot />
      </div>
    </main>

    <!-- Footer -->
    <footer class="border-t border-border mt-12 py-6">
      <div class="container mx-auto px-4">
        <div class="flex flex-col md:flex-row justify-between items-center">
          <div class="mb-4 md:mb-0">
            <p class="text-sm text-text-secondary">&copy; {{ new Date().getFullYear() }} All artwork rights reserved</p>
          </div>
          <div class="flex gap-4">
            <NuxtLink to="/artwork/archive" class="text-sm hover:text-accent transition-colors">Archive</NuxtLink>
            <NuxtLink to="/artwork/commissions" class="text-sm hover:text-accent transition-colors">Commissions</NuxtLink>
            <NuxtLink to="/artwork/process" class="text-sm hover:text-accent transition-colors">Process</NuxtLink>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
// Optional: Add artwork-specific logic here
</script>

<style scoped>
/* Artwork-specific styles */
.grid {
  perspective: 1000px;
}
</style>