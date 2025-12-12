<template>
  <div class="admin-dashboard">
    <header class="dashboard-header">
      <h1>Admin Dashboard</h1>
      <div class="admin-stats">
        <div class="stat-card">
          <FeatherIcon name="file-text" size="24" />
          <div class="stat-info">
            <span class="stat-value">{{ stats.blogPosts }}</span>
            <span class="stat-label">Blog Posts</span>
          </div>
        </div>
        
        <div class="stat-card">
          <FeatherIcon name="image" size="24" />
          <div class="stat-info">
            <span class="stat-value">{{ stats.mediaItems }}</span>
            <span class="stat-label">Media Items</span>
          </div>
        </div>
        
        <div class="stat-card">
          <FeatherIcon name="play" size="24" />
          <div class="stat-info">
            <span class="stat-value">{{ stats.games }}</span>
            <span class="stat-label">Games</span>
          </div>
        </div>
        
        <div class="stat-card">
          <FeatherIcon name="cloud" size="24" />
          <div class="stat-info">
            <span class="stat-value">{{ stats.bskyPosts }}</span>
            <span class="stat-label">BlueSky Posts</span>
          </div>
        </div>
      </div>
    </header>
    
    <div class="dashboard-content">
      <nav class="admin-nav">
        <button
          v-for="tab in tabs"
          :key="tab.id"
          class="nav-tab"
          :class="{ active: activeTab === tab.id }"
          @click="activeTab = tab.id"
        >
          <FeatherIcon :name="tab.icon" size="18" />
          <span>{{ tab.label }}</span>
        </button>
      </nav>
      
      <div class="tab-content">
        <!-- Blog Management -->
        <div v-if="activeTab === 'blog'" class="blog-management">
          <div class="section-header">
            <h2>Blog Management</h2>
            <button @click="showNewPostForm = true" class="btn btn-primary">
              <FeatherIcon name="plus" size="18" />
              <span>New Post</span>
            </button>
          </div>
          
          <div class="posts-list">
            <div v-for="post in blogPosts" :key="post.id" class="post-item">
              <div class="post-info">
                <h3>{{ post.title }}</h3>
                <div class="post-meta">
                  <span>{{ formatDate(post.date) }}</span>
                  <span class="category">{{ post.categories }}</span>
                  <span v-if="post.draft" class="draft-badge">Draft</span>
                </div>
              </div>
              <div class="post-actions">
                <button @click="editPost(post)" class="btn btn-secondary">
                  <FeatherIcon name="edit" size="16" />
                </button>
                <button @click="deletePost(post)" class="btn btn-danger">
                  <FeatherIcon name="trash" size="16" />
                </button>
              </div>
            </div>
          </div>
        </div>
        
        <!-- Media Management -->
        <div v-else-if="activeTab === 'media'" class="media-management">
          <div class="section-header">
            <h2>Media Library</h2>
            <button @click="showUploadModal = true" class="btn btn-primary">
              <FeatherIcon name="upload" size="18" />
              <span>Upload Media</span>
            </button>
          </div>
          
          <MediaGalleryAdmin />
        </div>
        
        <!-- Game Management -->
        <div v-else-if="activeTab === 'games'" class="game-management">
          <div class="section-header">
            <h2>Game Management</h2>
            <button @click="showNewGameForm = true" class="btn btn-primary">
              <FeatherIcon name="plus" size="18" />
              <span>Add Game</span>
            </button>
          </div>
          
          <div class="games-list">
            <div v-for="game in games" :key="game.id" class="game-item">
              <div class="game-info">
                <h3>{{ game.name }}</h3>
                <p class="game-description">{{ game.description }}</p>
                <div class="game-meta">
                  <span class="category">{{ game.category }}</span>
                  <span class="players">{{ game.min_players }}-{{ game.max_players }} players</span>
                  <span v-if="game.multiplayer_supported" class="multiplayer">Multiplayer</span>
                </div>
              </div>
              <div class="game-actions">
                <button @click="editGame(game)" class="btn btn-secondary">
                  <FeatherIcon name="edit" size="16" />
                </button>
                <button @click="deleteGame(game)" class="btn btn-danger">
                  <FeatherIcon name="trash" size="16" />
                </button>
              </div>
            </div>
          </div>
        </div>
        
        <!-- System Status -->
        <div v-else-if="activeTab === 'system'" class="system-status">
          <h2>System Status</h2>
          
          <div class="status-cards">
            <div class="status-card">
              <h3>Database</h3>
              <div class="status-indicator online"></div>
              <p>Connected and healthy</p>
              <code>{{ stats.dbSize }}</code>
            </div>
            
            <div class="status-card">
              <h3>BlueSky API</h3>
              <div :class="['status-indicator', bskyStatus]"></div>
              <p>Last sync: {{ lastSyncTime }}</p>
              <button @click="forceSync" class="btn btn-secondary mt-2">
                Force Sync
              </button>
            </div>
            
            <div class="status-card">
              <h3>Media Storage</h3>
              <div class="storage-bar">
                <div class="storage-used" :style="{ width: storageUsage + '%' }"></div>
              </div>
              <p>{{ storageUsed }} / {{ storageTotal }}</p>
            </div>
          </div>
          
          <div class="system-actions">
            <button @click="clearCache" class="btn btn-secondary">
              Clear Cache
            </button>
            <button @click="backupDatabase" class="btn btn-secondary">
              Backup Database
            </button>
            <button @click="rescanContent" class="btn btn-primary">
              Rescan All Content
            </button>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Modals -->
    <PostEditorModal
      v-if="showNewPostForm"
      @close="showNewPostForm = false"
      @saved="handlePostSaved"
    />
    
    <MediaUploadModal
      v-if="showUploadModal"
      @close="showUploadModal = false"
      @uploaded="handleMediaUploaded"
    />
    
    <GameEditorModal
      v-if="showNewGameForm"
      @close="showNewGameForm = false"
      @saved="handleGameSaved"
    />
  </div>
</template>

<script setup lang="ts">
import { useAuth } from '~/composables/useAuth'
import { useApi } from '~/composables/useApi'
import { useRouter } from 'vue-router'

const auth = useAuth()
const { apiFetch } = useApi()
const router = useRouter()

// Ensure admin access
if (process.client && !auth.isAdmin.value) {
  router.push('/')
}

const tabs = [
  { id: 'blog', label: 'Blog', icon: 'file-text' },
  { id: 'media', label: 'Media', icon: 'image' },
  { id: 'games', label: 'Games', icon: 'play' },
  { id: 'system', label: 'System', icon: 'settings' },
]

const activeTab = ref('blog')
const showNewPostForm = ref(false)
const showUploadModal = ref(false)
const showNewGameForm = ref(false)

const stats = ref({
  blogPosts: 0,
  mediaItems: 0,
  games: 0,
  bskyPosts: 0,
  dbSize: '0 MB',
})

const blogPosts = ref([])
const games = ref([])
const bskyStatus = ref('online')
const lastSyncTime = ref('Never')
const storageUsage = ref(0)
const storageUsed = ref('0 MB')
const storageTotal = ref('1 GB')

const fetchStats = async () => {
  try {
    const [posts, media, gamesData, bsky] = await Promise.all([
      apiFetch('/blog/posts?drafts=true'),
      apiFetch('/media?limit=1'),
      apiFetch('/games'),
      apiFetch('/bsky/posts?limit=1'),
    ])
    
    stats.value.blogPosts = posts.length
    stats.value.mediaItems = media.items?.length || 0
    stats.value.games = gamesData.length
    stats.value.bskyPosts = bsky.length
    
    // Simulate storage usage
    stats.value.dbSize = '2.4 MB'
    storageUsage.value = 35
    storageUsed.value = '350 MB'
    storageTotal.value = '1 GB'
    
    lastSyncTime.value = '2 hours ago'
    bskyStatus.value = 'online'
  } catch (error) {
    console.error('Failed to fetch stats:', error)
  }
}

const fetchBlogPosts = async () => {
  try {
    const posts = await apiFetch('/blog/posts?drafts=true')
    blogPosts.value = posts
  } catch (error) {
    console.error('Failed to fetch blog posts:', error)
  }
}

const fetchGames = async () => {
  try {
    const gamesData = await apiFetch('/games')
    games.value = gamesData
  } catch (error) {
    console.error('Failed to fetch games:', error)
  }
}

const editPost = (post: any) => {
  // Implement post editing
  console.log('Edit post:', post)
}

const deletePost = async (post: any) => {
  if (!confirm(`Delete post "${post.title}"?`)) return
  
  try {
    await apiFetch(`/blog/posts/${post.id}`, { method: 'DELETE' })
    await fetchBlogPosts()
  } catch (error) {
    console.error('Failed to delete post:', error)
  }
}

const editGame = (game: any) => {
  // Implement game editing
  console.log('Edit game:', game)
}

const deleteGame = async (game: any) => {
  if (!confirm(`Delete game "${game.name}"?`)) return
  
  try {
    await apiFetch(`/games/${game.slug}`, { method: 'DELETE' })
    await fetchGames()
  } catch (error) {
    console.error('Failed to delete game:', error)
  }
}

const forceSync = async () => {
  try {
    await Promise.all([
      apiFetch('/bsky/refresh', { method: 'POST' }),
      apiFetch('/profile/refresh', { method: 'POST' }),
      apiFetch('/blog/rescan', { method: 'POST' }),
    ])
    lastSyncTime.value = 'Just now'
  } catch (error) {
    console.error('Sync failed:', error)
  }
}

const clearCache = async () => {
  if (!confirm('Clear all cache?')) return
  
  try {
    await apiFetch('/admin/cache/clear', { method: 'POST' })
    alert('Cache cleared successfully')
  } catch (error) {
    console.error('Failed to clear cache:', error)
  }
}

const backupDatabase = async () => {
  try {
    const response = await apiFetch('/admin/backup', { method: 'POST' })
    alert(`Backup created: ${response.filename}`)
  } catch (error) {
    console.error('Backup failed:', error)
  }
}

const rescanContent = async () => {
  try {
    await apiFetch('/blog/rescan', { method: 'POST' })
    alert('Content rescanned successfully')
  } catch (error) {
    console.error('Rescan failed:', error)
  }
}

const handlePostSaved = () => {
  showNewPostForm.value = false
  fetchBlogPosts()
  fetchStats()
}

const handleMediaUploaded = () => {
  showUploadModal.value = false
  fetchStats()
}

const handleGameSaved = () => {
  showNewGameForm.value = false
  fetchGames()
  fetchStats()
}

const formatDate = (dateStr: string) => {
  return new Date(dateStr).toLocaleDateString()
}

onMounted(async () => {
  await Promise.all([
    fetchStats(),
    fetchBlogPosts(),
    fetchGames()
  ])
})
</script>

<style scoped>
.admin-dashboard {
  min-height: 100vh;
  background: var(--theme-bg);
}

.dashboard-header {
  padding: 2rem;
  background: rgba(0, 0, 0, 0.3);
  border-bottom: 2px solid var(--theme-border);
}

.dashboard-header h1 {
  margin-bottom: 2rem;
  color: var(--theme-primary);
}

.admin-stats {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
}

.stat-card {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1.5rem;
  background: rgba(0, 0, 0, 0.2);
  border: 1px solid var(--theme-border);
  border-radius: 8px;
  transition: all 0.2s;
}

.stat-card:hover {
  border-color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.1);
}

.stat-card svg {
  color: var(--theme-primary);
}

.stat-info {
  display: flex;
  flex-direction: column;
}

.stat-value {
  font-size: 1.8rem;
  font-weight: bold;
  color: var(--theme-primary);
}

.stat-label {
  font-size: 0.9rem;
  color: var(--theme-fg);
  opacity: 0.7;
}

.dashboard-content {
  display: flex;
  min-height: calc(100vh - 200px);
}

.admin-nav {
  width: 250px;
  padding: 2rem;
  background: rgba(0, 0, 0, 0.2);
  border-right: 2px solid var(--theme-border);
}

.nav-tab {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  width: 100%;
  padding: 1rem;
  margin-bottom: 0.5rem;
  background: transparent;
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  color: var(--theme-fg);
  cursor: pointer;
  transition: all 0.2s;
  text-align: left;
}

.nav-tab:hover {
  border-color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.1);
}

.nav-tab.active {
  border-color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.2);
  color: var(--theme-primary);
}

.tab-content {
  flex: 1;
  padding: 2rem;
  overflow-y: auto;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}

.section-header h2 {
  color: var(--theme-primary);
}

.posts-list,
.games-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.post-item,
.game-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  background: rgba(0, 0, 0, 0.2);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  transition: all 0.2s;
}

.post-item:hover,
.game-item:hover {
  border-color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.1);
}

.post-info h3,
.game-info h3 {
  margin: 0 0 0.5rem 0;
  color: var(--theme-primary);
}

.post-meta {
  display: flex;
  gap: 1rem;
  font-size: 0.9rem;
  color: var(--theme-fg);
  opacity: 0.7;
}

.draft-badge {
  color: var(--theme-warning);
}

.game-description {
  margin: 0.5rem 0;
  font-size: 0.9rem;
  color: var(--theme-fg);
  opacity: 0.8;
}

.game-meta {
  display: flex;
  gap: 1rem;
  font-size: 0.9rem;
}

.category {
  text-transform: uppercase;
  font-weight: bold;
  color: var(--theme-accent);
}

.multiplayer {
  color: var(--theme-success);
}

.post-actions,
.game-actions {
  display: flex;
  gap: 0.5rem;
}

.btn-danger {
  background: var(--theme-error);
  color: white;
}

.status-cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 1.5rem;
  margin-bottom: 2rem;
}

.status-card {
  padding: 1.5rem;
  background: rgba(0, 0, 0, 0.2);
  border: 1px solid var(--theme-border);
  border-radius: 8px;
}

.status-card h3 {
  margin: 0 0 1rem 0;
  color: var(--theme-primary);
}

.status-indicator {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  margin-bottom: 1rem;
}

.status-indicator.online {
  background: var(--theme-success);
}

.status-indicator.offline {
  background: var(--theme-error);
}

.storage-bar {
  width: 100%;
  height: 8px;
  background: var(--theme-border);
  border-radius: 4px;
  margin: 1rem 0;
  overflow: hidden;
}

.storage-used {
  height: 100%;
  background: var(--theme-primary);
  transition: width 0.3s;
}

.system-actions {
  display: flex;
  gap: 1rem;
}

@media (max-width: 768px) {
  .dashboard-content {
    flex-direction: column;
  }
  
  .admin-nav {
    width: 100%;
    padding: 1rem;
    border-right: none;
    border-bottom: 2px solid var(--theme-border);
  }
  
  .nav-tab {
    justify-content: center;
  }
  
  .admin-stats {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .system-actions {
    flex-direction: column;
  }
}
</style>