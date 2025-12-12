<template>
  <header class="site-header">
    <div class="container">
      <div class="header-content">
        <NuxtLink to="/" class="site-title">
          <h1>{{ profileStore.profile?.name || 'Lilith in a Parka' }}</h1>
        </NuxtLink>
        
        <nav class="main-nav">
          <NuxtLink to="/" class="nav-link">
            <FeatherIcon name="home" size="18" />
            <span>Home</span>
          </NuxtLink>
          
          <NuxtLink to="/blog" class="nav-link">
            <FeatherIcon name="book-open" size="18" />
            <span>Blog</span>
          </NuxtLink>
          
          <NuxtLink to="/profile" class="nav-link">
            <FeatherIcon name="user" size="18" />
            <span>Profile</span>
          </NuxtLink>
          
          <NuxtLink to="/games" class="nav-link">
            <FeatherIcon name="play" size="18" />
            <span>Games</span>
          </NuxtLink>
        </nav>
        
        <ThemePicker />
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { useProfileStore } from '~/pages/profile/profile'

const profileStore = useProfileStore()
</script>

<style scoped>
.site-header {
  background: var(--theme-bg);
  border-bottom: 2px solid var(--theme-border);
  padding: 1rem 0;
  position: sticky;
  top: 0;
  z-index: 100;
}

.header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 2rem;
}

.site-title h1 {
  margin: 0;
  font-size: 1.5rem;
  color: var(--theme-primary);
}

.main-nav {
  display: flex;
  gap: 1rem;
  flex: 1;
  justify-content: center;
}

.nav-link {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 1rem;
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  transition: all 0.2s;
}

.nav-link:hover {
  border-color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.1);
  transform: translateY(-2px);
}

.nav-link.router-link-active {
  border-color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.2);
}

@media (max-width: 768px) {
  .header-content {
    flex-direction: column;
    gap: 1rem;
  }
  
  .main-nav {
    width: 100%;
    justify-content: space-around;
  }
  
  .nav-link span {
    display: none;
  }
}
</style>