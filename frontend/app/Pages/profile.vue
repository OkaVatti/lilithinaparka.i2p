<template>
  <div class="profile-page">
    <div v-if="profileStore.loading" class="loading">
      <FeatherIcon name="loader" size="24" class="spin" />
      <span>Loading profile...</span>
    </div>
    
    <div v-else-if="profileStore.error" class="error">
      <FeatherIcon name="alert-circle" size="20" />
      <span>{{ profileStore.error }}</span>
      <button @click="retry" class="btn mt-2">
        <FeatherIcon name="refresh-cw" size="18" />
        <span>Retry</span>
      </button>
    </div>
    
    <div v-else-if="profileStore.profile">
      <ProfileCard :profile="profileStore.profile" />
      
      <div v-if="hasBskyData" class="bsky-section">
        <h2>BlueSky Profile</h2>
        <div class="bsky-stats">
          <div class="stat-item">
            <FeatherIcon name="users" size="20" />
            <div>
              <span class="stat-value">{{ formatNumber(profileStore.profile.bsky_followers_count) }}</span>
              <span class="stat-label">Followers</span>
            </div>
          </div>
          
          <div class="stat-item">
            <FeatherIcon name="user-plus" size="20" />
            <div>
              <span class="stat-value">{{ formatNumber(profileStore.profile.bsky_follows_count) }}</span>
              <span class="stat-label">Following</span>
            </div>
          </div>
          
          <div class="stat-item">
            <FeatherIcon name="edit" size="20" />
            <div>
              <span class="stat-value">{{ formatNumber(profileStore.profile.bsky_posts_count) }}</span>
              <span class="stat-label">Posts</span>
            </div>
          </div>
        </div>
        
        <div v-if="profileStore.profile.bsky_banner" class="bsky-banner">
          <img :src="profileStore.profile.bsky_banner" alt="Banner" />
        </div>
        
        <p v-if="profileStore.profile.bsky_description" class="bsky-description">
          {{ profileStore.profile.bsky_description }}
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useProfileStore } from '~~/stores/profile'

const profileStore = useProfileStore()

const hasBskyData = computed(() => {
  return profileStore.profile?.bsky_display_name || 
         profileStore.profile?.bsky_followers_count
})

const formatNumber = (num: number) => {
  if (!num) return '0'
  if (num >= 1000000) return (num / 1000000).toFixed(1) + 'M'
  if (num >= 1000) return (num / 1000).toFixed(1) + 'K'
  return num.toString()
}

const retry = () => {
  profileStore.fetchProfile()
}

onMounted(() => {
  if (!profileStore.profile) {
    profileStore.fetchProfile()
  }
})

useHead({
  title: `${profileStore.profile?.name || 'Profile'} - Lilith in a Parka`,
  meta: [
    { name: 'description', content: profileStore.profile?.bio || '' }
  ]
})
</script>

<style scoped>
.profile-page {
  max-width: 900px;
  margin: 0 auto;
}

.bsky-section {
  margin-top: 3rem;
  padding: 2rem;
  background: rgba(0, 0, 0, 0.2);
  border: 2px solid var(--theme-border);
  border-radius: 8px;
}

.bsky-section h2 {
  margin-top: 0;
  margin-bottom: 1.5rem;
  color: var(--theme-accent);
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.bsky-stats {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 1.5rem;
  margin-bottom: 2rem;
}

.stat-item {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1rem;
  background: rgba(189, 147, 249, 0.1);
  border: 1px solid var(--theme-primary);
  border-radius: 4px;
}

.stat-item svg {
  color: var(--theme-accent);
}

.stat-value {
  display: block;
  font-size: 1.5rem;
  font-weight: bold;
  color: var(--theme-primary);
}

.stat-label {
  display: block;
  font-size: 0.85rem;
  color: var(--theme-fg);
  opacity: 0.7;
}

.bsky-banner {
  width: 100%;
  height: 200px;
  margin-bottom: 1.5rem;
  border-radius: 8px;
  overflow: hidden;
}

.bsky-banner img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.bsky-description {
  line-height: 1.6;
  color: var(--theme-fg);
  opacity: 0.9;
}

@media (max-width: 768px) {
  .bsky-stats {
    grid-template-columns: 1fr;
  }
}
</style>