<!-- frontend/app/components/Navigation/Navbar.vue -->
<template>
  <header class="nav-bar sticky top-0 z-50" :class="{ 'scrolled': isScrolled }">
    <div class="nav-content">
      <!-- ASCII Logo -->
      <NuxtLink to="/" class="ascii-logo">
        <pre class="text-xs sm:text-sm leading-tight tracking-tight">
          <BrandingASCIIHeader/>
        </pre>
      </NuxtLink>
      
      <!-- Desktop Navigation -->
      <nav class="nav-menu hidden md:flex">
        <NuxtLink to="/" class="nav-link" :class="{ 'active': $route.path === '/' }">
          home
        </NuxtLink>
        <NuxtLink to="/blog" class="nav-link" :class="{ 'active': $route.path.startsWith('/blog') }">
          blog
        </NuxtLink>
        <NuxtLink to="/gallery" class="nav-link" :class="{ 'active': $route.path.startsWith('/gallery') }">
          gallery
        </NuxtLink>
        <NuxtLink to="/projects" class="nav-link" :class="{ 'active': $route.path.startsWith('/projects') }">
          projects
        </NuxtLink>
        <NuxtLink to="/social" class="nav-link" :class="{ 'active': $route.path.startsWith('/social') }">
          social
        </NuxtLink>
        <NuxtLink to="/profile" class="nav-link" :class="{ 'active': $route.path.startsWith('/profile') }">
          profile
        </NuxtLink>
      </nav>
      
      <!-- Status Indicator -->
      <div class="hidden lg:flex items-center gap-2 text-xs font-mono">
        <div class="w-2 h-2 rounded-full" :class="connectionStatus.class"></div>
        <span class="text-text-secondary">{{ connectionStatus.text }}</span>
      </div>
      
      <!-- Mobile Navigation Toggle -->
      <button @click="toggleMenu" class="nav-toggle md:hidden">
        <svg class="nav-toggle-icon w-6 h-6" :class="{ 'open': isMenuOpen }" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path v-if="!isMenuOpen" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
          <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>
    </div>
    
    <!-- Mobile Navigation Menu -->
    <div class="nav-mobile-menu" :class="{ 'open': isMenuOpen }">
      <div class="nav-mobile-links">
        <NuxtLink to="/" class="nav-mobile-link" @click="closeMenu" :class="{ 'active': $route.path === '/' }">
          home
        </NuxtLink>
        <NuxtLink to="/blog" class="nav-mobile-link" @click="closeMenu" :class="{ 'active': $route.path.startsWith('/blog') }">
          blog
        </NuxtLink>
        <NuxtLink to="/gallery" class="nav-mobile-link" @click="closeMenu" :class="{ 'active': $route.path.startsWith('/gallery') }">
          gallery
        </NuxtLink>
        <NuxtLink to="/projects" class="nav-mobile-link" @click="closeMenu" :class="{ 'active': $route.path.startsWith('/projects') }">
          projects
        </NuxtLink>
        <NuxtLink to="/social" class="nav-mobile-link" @click="closeMenu" :class="{ 'active': $route.path.startsWith('/social') }">
          social
        </NuxtLink>
        <NuxtLink to="/profile" class="nav-mobile-link" @click="closeMenu" :class="{ 'active': $route.path.startsWith('/profile') }">
          profile
        </NuxtLink>
      </div>
    </div>
    
    <!-- Progress Bar for Loading States -->
    <div v-if="isLoading" class="nav-progress loading"></div>
  </header>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'

const isMenuOpen = ref(false)
const isScrolled = ref(false)
const isLoading = ref(false)
const isConnected = ref(true)
let eventSource: EventSource | null = null

const connectionStatus = computed(() => {
  if (isConnected.value) {
    return { class: 'bg-success animate-pulse', text: 'connected' }
  }
  return { class: 'bg-error', text: 'disconnected' }
})

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value
}

const closeMenu = () => {
  isMenuOpen.value = false
}

const handleScroll = () => {
  isScrolled.value = window.scrollY > 20
}

// Connect to SSE for real-time updates
const connectSSE = () => {
  const config = useRuntimeConfig()
  eventSource = new EventSource(`${config.public.apiBase}/events`)
  
  eventSource.onopen = () => {
    isConnected.value = true
    console.log('SSE connected')
  }
  
  eventSource.onerror = () => {
    isConnected.value = false
    console.error('SSE connection error')
  }
  
  eventSource.onmessage = (event) => {
    try {
      const data = JSON.parse(event.data)
      if (data.type === 'new_post') {
        console.log('New post detected:', data.path)
        // Trigger a notification or refresh
      }
    } catch (e) {
      console.error('Failed to parse SSE message:', e)
    }
  }
}

onMounted(() => {
  window.addEventListener('scroll', handleScroll)
  connectSSE()
})

onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll)
  if (eventSource) {
    eventSource.close()
  }
})
</script>