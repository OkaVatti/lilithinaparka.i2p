<!-- app/layouts/bskyLayout.vue -->
<template>
  <div class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-50 dark:from-gray-900 dark:to-blue-900/20 text-text-primary">
    <!-- Navigation -->
    <header class="sticky top-0 z-50 backdrop-blur-lg bg-white/80 dark:bg-gray-900/80 border-b border-blue-200 dark:border-blue-800/50">
      <div class="container mx-auto px-4 py-4">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-3">
            <div class="w-8 h-8 rounded-full bg-gradient-to-r from-blue-500 to-sky-500"></div>
            <NuxtLink to="/" class="font-bold text-xl font-mono hover:text-blue-600 dark:hover:text-blue-400 transition-colors">
              Bluesky
            </NuxtLink>
          </div>
          <nav class="flex items-center gap-6">
            <NuxtLink to="/" class="text-sm hover:text-blue-600 dark:hover:text-blue-400 transition-colors">Home</NuxtLink>
            <NuxtLink to="/bsky" class="text-sm hover:text-blue-600 dark:hover:text-blue-400 transition-colors active">Bluesky</NuxtLink>
            <a href="https://bsky.app" target="_blank" class="text-sm hover:text-blue-600 dark:hover:text-blue-400 transition-colors">
              Open Bluesky
            </a>
          </nav>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="container mx-auto px-4 py-8">
      <!-- Bluesky header -->
      <div class="mb-12 text-center">
        <h1 class="text-4xl font-bold font-mono mb-4 bg-gradient-to-r from-blue-600 to-sky-500 bg-clip-text text-transparent">
          Bluesky Feed
        </h1>
        <p class="text-gray-600 dark:text-gray-300 max-w-2xl mx-auto">
          Latest posts from Bluesky. Follow me at 
          <a href="https://bsky.app/profile/yourhandle" target="_blank" 
             class="text-blue-600 dark:text-blue-400 hover:underline">
            @yourhandle
          </a>
        </p>
      </div>

      <!-- Bluesky feed container -->
      <div class="max-w-2xl mx-auto">
        <div class="space-y-4">
          <slot />
        </div>
        
        <!-- Bluesky stats -->
        <div class="mt-12 p-6 rounded-xl border border-blue-200 dark:border-blue-800/50 bg-white/50 dark:bg-gray-900/50 backdrop-blur-sm">
          <h3 class="text-lg font-bold font-mono mb-4">Bluesky Stats</h3>
          <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
            <div class="text-center p-4 rounded-lg bg-blue-50 dark:bg-blue-900/20">
              <div class="text-2xl font-bold text-blue-600 dark:text-blue-400">1.2k</div>
              <div class="text-sm text-gray-600 dark:text-gray-400">Followers</div>
            </div>
            <div class="text-center p-4 rounded-lg bg-blue-50 dark:bg-blue-900/20">
              <div class="text-2xl font-bold text-blue-600 dark:text-blue-400">245</div>
              <div class="text-sm text-gray-600 dark:text-gray-400">Following</div>
            </div>
            <div class="text-center p-4 rounded-lg bg-blue-50 dark:bg-blue-900/20">
              <div class="text-2xl font-bold text-blue-600 dark:text-blue-400">589</div>
              <div class="text-sm text-gray-600 dark:text-gray-400">Posts</div>
            </div>
            <div class="text-center p-4 rounded-lg bg-blue-50 dark:bg-blue-900/20">
              <div class="text-2xl font-bold text-blue-600 dark:text-blue-400">42</div>
              <div class="text-sm text-gray-600 dark:text-gray-400">Likes</div>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- Bluesky CTA -->
    <div class="my-12">
      <div class="container mx-auto px-4">
        <div class="max-w-2xl mx-auto p-8 rounded-2xl bg-gradient-to-r from-blue-600 to-sky-500 text-white text-center">
          <h3 class="text-2xl font-bold font-mono mb-4">Join Bluesky</h3>
          <p class="mb-6 opacity-90">The social network that puts you in control of your experience.</p>
          <a href="https://bsky.app" target="_blank" 
             class="inline-block px-8 py-3 bg-white text-blue-600 rounded-full font-bold hover:bg-gray-100 transition-colors">
            Sign up for Bluesky
          </a>
        </div>
      </div>
    </div>

    <!-- Footer -->
    <footer class="border-t border-blue-200 dark:border-blue-800/50 py-8">
      <div class="container mx-auto px-4">
        <div class="flex flex-col md:flex-row justify-between items-center">
          <div class="mb-4 md:mb-0">
            <p class="text-sm text-gray-600 dark:text-gray-400">Bluesky is an open social network</p>
          </div>
          <div class="flex gap-6">
            <a href="https://github.com/bluesky-social" target="_blank" 
               class="text-sm hover:text-blue-600 dark:hover:text-blue-400 transition-colors">GitHub</a>
            <a href="https://docs.bsky.app" target="_blank" 
               class="text-sm hover:text-blue-600 dark:hover:text-blue-400 transition-colors">API Docs</a>
            <a href="https://atproto.com" target="_blank" 
               class="text-sm hover:text-blue-600 dark:hover:text-blue-400 transition-colors">AT Protocol</a>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  layout: 'default'
})
</script>