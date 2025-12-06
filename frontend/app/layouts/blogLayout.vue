<!-- app/layouts/blogLayout.vue -->
<template>
  <div class="min-h-screen bg-bg-primary text-text-primary">
    <!-- Navigation -->
    <header class="sticky top-0 z-50 backdrop-blur-lg bg-bg-primary/80 border-b border-border">
      <div class="container mx-auto px-4 py-4">
        <div class="flex items-center justify-between">
          <NuxtLink to="/" class="font-bold text-xl font-mono hover:text-accent transition-colors">
            Blog
          </NuxtLink>
          <nav class="flex items-center gap-6">
            <NuxtLink to="/" class="text-sm hover:text-accent transition-colors">Home</NuxtLink>
            <NuxtLink to="/artwork" class="text-sm hover:text-accent transition-colors">Artwork</NuxtLink>
            <NuxtLink to="/blog" class="text-sm hover:text-accent transition-colors active">Blog</NuxtLink>
            <NuxtLink to="/profile" class="text-sm hover:text-accent transition-colors">Profile</NuxtLink>
          </nav>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="container mx-auto px-4 py-8">
      <!-- Blog header -->
      <div class="mb-12 text-center">
        <h1 class="text-5xl font-bold font-mono mb-4">Blog</h1>
        <p class="text-text-secondary text-lg max-w-2xl mx-auto">
          Thoughts, tutorials, and updates from my creative journey
        </p>
        <div class="flex flex-wrap justify-center gap-2 mt-6">
          <button class="px-4 py-2 rounded-full border border-border hover:border-accent/50 transition-colors text-sm">
            All Posts
          </button>
          <button v-for="category in categories" :key="category" 
                  class="px-4 py-2 rounded-full border border-border hover:border-accent/50 transition-colors text-sm">
            {{ category }}
          </button>
        </div>
      </div>

      <!-- Blog content -->
      <div class="max-w-4xl mx-auto">
        <slot />
      </div>

      <!-- Pagination/Archive -->
      <div class="mt-16 pt-8 border-t border-border">
        <div class="flex justify-between items-center">
          <NuxtLink to="/blog/archive" class="text-sm hover:text-accent transition-colors">
            View Archive →
          </NuxtLink>
          <div class="text-text-secondary text-sm">
            {{ totalPosts }} posts total
          </div>
        </div>
      </div>
    </main>

    <!-- Newsletter Signup -->
    <div class="border-y border-border my-12">
      <div class="container mx-auto px-4 py-12">
        <div class="max-w-md mx-auto text-center">
          <h3 class="text-2xl font-bold font-mono mb-4">Subscribe to Updates</h3>
          <p class="text-text-secondary mb-6">Get notified about new posts and artwork</p>
          <form class="flex gap-2">
            <input type="email" placeholder="your@email.com" 
                   class="flex-1 px-4 py-2 rounded-lg border border-border bg-bg-secondary focus:outline-none focus:border-accent">
            <button type="submit" class="px-6 py-2 bg-accent text-white rounded-lg hover:bg-accent/80 transition-colors">
              Subscribe
            </button>
          </form>
        </div>
      </div>
    </div>

    <!-- Footer -->
    <footer class="border-t border-border py-8">
      <div class="container mx-auto px-4">
        <div class="flex flex-col md:flex-row justify-between items-center">
          <div class="mb-4 md:mb-0">
            <p class="text-sm text-text-secondary">Written with ☕ and ❤️</p>
          </div>
          <div class="flex gap-6">
            <NuxtLink to="/blog/rss" class="text-sm hover:text-accent transition-colors">RSS</NuxtLink>
            <NuxtLink to="/blog/atom" class="text-sm hover:text-accent transition-colors">Atom</NuxtLink>
            <NuxtLink to="/blog/json" class="text-sm hover:text-accent transition-colors">JSON Feed</NuxtLink>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
// Example blog categories
const categories = ['Tutorials', 'Thoughts', 'Updates', 'Technical', 'Creative']
const totalPosts = 42 // This would come from your data
</script>