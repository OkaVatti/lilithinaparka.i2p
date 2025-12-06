<!-- app/components/Navigation/Footer.vue -->
<template>
  <footer class="border-t border-border mt-16 py-8 bg-bg-primary font-mono">
    <div class="container mx-auto px-4">
      <div class="flex flex-col md:flex-row justify-between items-center gap-4 text-sm">
        <div class="text-text-secondary">
          <span class="text-accent">©</span> {{ currentYear }} Lilith Parker
          <span class="mx-2 text-border">|</span>
          <span class="text-text-secondary">hosted on I2P</span>
        </div>
        
        <div class="flex gap-4">
          <a 
            href="https://github.com/lilithinaparka" 
            target="_blank" 
            rel="noopener noreferrer" 
            class="text-link hover:text-accent transition-colors"
          >
            <span class="text-accent">[</span>github<span class="text-accent">]</span>
          </a>
          <a 
            href="https://bsky.app/profile/lilithinaparka.bsky.social" 
            target="_blank" 
            rel="noopener noreferrer" 
            class="text-link hover:text-accent transition-colors"
          >
            <span class="text-accent">[</span>bluesky<span class="text-accent">]</span>
          </a>
          <a 
            href="/rss.xml" 
            class="text-link hover:text-accent transition-colors"
          >
            <span class="text-accent">[</span>rss<span class="text-accent">]</span>
          </a>
        </div>
      </div>
      
      <div class="text-center mt-4 text-xs text-text-secondary">
        <span class="inline-block animate-pulse">▂</span>
        <span class="mx-2">powered by nuxt + go + sqlite</span>
        <span class="inline-block animate-pulse">▂</span>
      </div>
    </div>
  </footer>
</template>

<script setup lang="ts">
const currentYear = new Date().getFullYear()
</script>