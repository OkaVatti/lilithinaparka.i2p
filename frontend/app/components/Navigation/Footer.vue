<template>
  <footer class="border-t border-border mt-16 py-8">
    <div class="container mx-auto px-4">
      <div class="flex flex-col md:flex-row justify-between items-center gap-4">
        <div class="text-sm text-text-secondary">
          <span class="text-accent">©</span> {{ currentYear }} Lilith Parker
        </div>
        
        <div class="flex gap-4 text-sm">
          <a href="https://github.com/lilithinaparka" target="_blank" rel="noopener" class="hover:text-accent">
            github
          </a>
          <a href="https://bsky.app/profile/lilithinaparka.bsky.social" target="_blank" rel="noopener" class="hover:text-accent">
            bluesky
          </a>
          <a href="/rss.xml" class="hover:text-accent">
            rss
          </a>
        </div>
      </div>
      
      <div class="text-center mt-4 text-xs text-text-secondary">
        <span class="blink">_</span> hosted on I2P
      </div>
    </div>
  </footer>
</template>

<script setup lang="ts">
const currentYear = new Date().getFullYear()
</script>