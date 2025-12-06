<!-- app/layouts/defaultLayout.vue -->
<template>
  <div class="min-h-screen bg-bg-primary text-text-primary font-mono">
    <!-- Navigation -->
    <NavBar />
    
    <!-- Main Content -->
    <main class="container mx-auto px-4 py-8">
      <slot />
    </main>
    
    <!-- Footer -->
    <footer class="border-t border-border mt-12 py-6 bg-bg-secondary">
      <div class="container mx-auto px-4">
        <div class="flex flex-col md:flex-row justify-between items-center gap-4">
          <div class="text-text-secondary text-sm">
            <p>© 2024 lilthinaparka.i2p</p>
            <p class="text-xs mt-1">hosted on <span class="text-accent">i2p</span> · <span class="text-accent">tor</span> · <span class="text-accent">loki</span></p>
          </div>
          
          <div class="flex gap-4 text-xs">
            <NuxtLink to="/privacy" class="text-link hover:text-accent transition-colors">privacy</NuxtLink>
            <span class="text-border">|</span>
            <NuxtLink to="/terms" class="text-link hover:text-accent transition-colors">terms</NuxtLink>
            <span class="text-border">|</span>
            <a href="/rss.xml" class="text-link hover:text-accent transition-colors">rss</a>
            <span class="text-border">|</span>
            <NuxtLink to="/contact" class="text-link hover:text-accent transition-colors">contact</NuxtLink>
          </div>
        </div>
        
        <div class="mt-6 text-center text-xs text-text-secondary">
          <p>this site uses: deno · nuxt 4 · tailwindcss · typescript</p>
          <p class="mt-1">backend: go · gorm · echo · sqlite</p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import NavBar from '../components/Navigation/Navbar.vue'
</script>