<!-- app/components/Navigation/NavBar.vue -->
<template>
  <header class="nav-bar sticky top-0 z-50">
    <div class="nav-content">
      <!-- ASCII Logo -->
      <NuxtLink to="/" class="ascii-logo">
        <ASCIIHeader/>
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
  </header>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import ASCIIHeader from '../Branding/ASCIIHeader.vue'

const isMenuOpen = ref(false)

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value
}

const closeMenu = () => {
  isMenuOpen.value = false
}
</script>