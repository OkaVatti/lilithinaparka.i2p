<template>
  <footer class="site-footer">
    <div class="container">
      <div class="footer-content">
        <div class="footer-section">
          <h3>About</h3>
          <p>{{ profileStore.profile?.bio || 'A minimalist blog on I2P' }}</p>
        </div>
        
        <div class="footer-section">
          <h3>Links</h3>
          <div class="footer-links">
            <a v-if="profileStore.profile?.github" :href="profileStore.profile.github" target="_blank" rel="noopener">
              <FeatherIcon name="github" size="20" />
              <span>GitHub</span>
            </a>
            <a v-if="profileStore.profile?.bluesky" :href="profileStore.profile.bluesky" target="_blank" rel="noopener">
              <FeatherIcon name="cloud" size="20" />
              <span>BlueSky</span>
            </a>
            <a v-if="profileStore.profile?.rss_feed" :href="profileStore.profile.rss_feed">
              <FeatherIcon name="rss" size="20" />
              <span>RSS</span>
            </a>
          </div>
        </div>
        
        <div class="footer-section">
          <h3>Built With</h3>
          <p>Nuxt 4 • Vue 3 • Deno • System.css</p>
          <p class="copyright">© {{ new Date().getFullYear() }} {{ profileStore.profile?.name }}</p>
        </div>
      </div>
    </div>
  </footer>
</template>

<script setup lang="ts">
import { useProfileStore } from '~/pages/profile/profile'

const profileStore = useProfileStore()
</script>

<style scoped>
.site-footer {
  background: var(--theme-bg);
  border-top: 2px solid var(--theme-border);
  padding: 2rem 0;
  margin-top: 4rem;
}

.footer-content {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 2rem;
}

.footer-section h3 {
  margin-bottom: 1rem;
  font-size: 1.2rem;
  color: var(--theme-primary);
}

.footer-section p {
  margin: 0.5rem 0;
  color: var(--theme-fg);
  opacity: 0.8;
}

.footer-links {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.footer-links a {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem;
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  transition: all 0.2s;
}

.footer-links a:hover {
  border-color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.1);
}

.copyright {
  margin-top: 1rem;
  font-size: 0.9rem;
}

@media (max-width: 768px) {
  .footer-content {
    grid-template-columns: 1fr;
  }
}
</style>