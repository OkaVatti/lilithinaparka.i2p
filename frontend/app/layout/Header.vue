<template>
  <header class="site-header">
    <div class="container">
      <div class="window">
        <div class="title-bar">
          <div class="title-bar-text">
            <NuxtLink to="/" style="color: inherit; text-decoration: none;">
              {{ profileStore.profile?.name || 'Lilith in a Parka' }}
            </NuxtLink>
          </div>
          <div class="title-bar-controls">
            <ThemeButton />
          </div>
        </div>
        <div class="window-body">
          <nav class="main-nav">
            <menu role="tablist">
              <NuxtLink to="/" class="nav-link">
                <button :aria-selected="route.path === '/'">
                  Home
                </button>
              </NuxtLink>
              
              <NuxtLink to="/blog" class="nav-link">
                <button :aria-selected="route.path.startsWith('/blog')">
                  Blog
                </button>
              </NuxtLink>
              
              <NuxtLink to="/profile" class="nav-link">
                <button :aria-selected="route.path.startsWith('/profile')">
                  Profile
                </button>
              </NuxtLink>
              
              <NuxtLink to="/games" class="nav-link">
                <button :aria-selected="route.path.startsWith('/games')">
                  Games
                </button>
              </NuxtLink>

              <NuxtLink v-if="auth.isAdmin.value" to="/admin" class="nav-link">
                <button :aria-selected="route.path.startsWith('/admin')">
                  Admin
                </button>
              </NuxtLink>
            </menu>
          </nav>
        </div>
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { useProfileStore } from '~~/stores/profile'
import { useAuth } from '~/composables/useAuth'
import { useRoute } from '#vue-router'
import ThemeButton from '~/components/Universal/Theme/ThemeButton.vue'

const profileStore = useProfileStore()
const auth = useAuth()
const route = useRoute()

onMounted(() => {
  profileStore.fetchProfile()
})
</script>

<style scoped>
.site-header {
  padding: 1rem 0;
  position: sticky;
  top: 0;
  z-index: 100;
  background: var(--theme-bg);
  border-bottom: 2px solid var(--theme-border);
}

.title-bar-controls {
  display: flex;
  gap: 0.5rem;
  align-items: center;
}

.main-nav menu {
  display: flex;
  gap: 0.5rem;
  padding: 0;
  margin: 0;
  list-style: none;
  flex-wrap: wrap;
}

.nav-link {
  text-decoration: none;
}

.nav-link button {
  margin: 0;
}

.nav-link button[aria-selected="true"] {
  background: var(--theme-primary);
  color: var(--theme-bg);
}

@media (max-width: 768px) {
  .main-nav menu {
    justify-content: space-around;
  }
  
  .nav-link button {
    padding: 0.5rem;
    font-size: 0.8rem;
  }
}
</style>