<template>
  <div class="notification-bell" v-click-outside="closeDropdown">
    <button @click="toggleDropdown" class="bell-button" :class="{ 'has-unread': unreadCount > 0 }">
      <FeatherIcon name="bell" size="20" />
      <span v-if="unreadCount > 0" class="badge">{{ unreadCount > 99 ? '99+' : unreadCount }}</span>
    </button>
    
    <div v-if="isOpen" class="notification-dropdown">
      <div class="dropdown-header">
        <h3>Notifications</h3>
        <div class="header-actions">
          <button v-if="unreadCount > 0" @click="markAllAsRead" class="btn-link">
            Mark all as read
          </button>
          <button @click="refresh" class="btn-link">
            <FeatherIcon name="refresh-cw" size="16" />
          </button>
        </div>
      </div>
      
      <div v-if="notifications.length === 0" class="no-notifications">
        <FeatherIcon name="bell-off" size="32" />
        <p>No notifications</p>
      </div>
      
      <div v-else class="notifications-list">
        <NotificationItem
          v-for="notification in notifications.slice(0, 10)"
          :key="notification.id"
          :notification="notification"
          @click="handleNotificationClick(notification)"
        />
      </div>
      
      <div class="dropdown-footer">
        <NuxtLink to="/notifications" @click="closeDropdown" class="view-all">
          View all notifications
        </NuxtLink>
        <button @click="openSettings" class="btn-link">
          <FeatherIcon name="settings" size="16" />
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useNotifications } from '../../composables/useNotifications'
import { useRouter } from 'vue-router'

const router = useRouter()
const { notifications, unreadCount, markAllAsRead, fetchNotifications } = useNotifications()

const isOpen = ref(false)

const toggleDropdown = () => {
  isOpen.value = !isOpen.value
}

const closeDropdown = () => {
  isOpen.value = false
}

const refresh = async () => {
  await fetchNotifications()
}

const handleNotificationClick = async (notification: any) => {
  // Mark as read
  if (!notification.read) {
    await markAsRead(notification.id)
  }
  
  // Navigate if there's an action URL
  if (notification.action_url) {
    router.push(notification.action_url)
  }
  
  closeDropdown()
}

const openSettings = () => {
  router.push('/settings/notifications')
  closeDropdown()
}

// Close dropdown on escape key
const handleKeydown = (e: KeyboardEvent) => {
  if (e.key === 'Escape' && isOpen.value) {
    closeDropdown()
  }
}

onMounted(() => {
  document.addEventListener('keydown', handleKeydown)
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleKeydown)
})
</script>

<style scoped>
.notification-bell {
  position: relative;
}

.bell-button {
  position: relative;
  background: none;
  border: none;
  color: var(--theme-fg);
  cursor: pointer;
  padding: 0.5rem;
  border-radius: 4px;
  transition: all 0.2s;
}

.bell-button:hover {
  background: rgba(255, 255, 255, 0.1);
  color: var(--theme-primary);
}

.bell-button.has-unread {
  color: var(--theme-primary);
}

.badge {
  position: absolute;
  top: 0;
  right: 0;
  background: var(--theme-error);
  color: white;
  font-size: 0.7rem;
  font-weight: bold;
  min-width: 18px;
  height: 18px;
  border-radius: 9px;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0 4px;
}

.notification-dropdown {
  position: absolute;
  top: calc(100% + 0.5rem);
  right: 0;
  width: 400px;
  max-height: 500px;
  background: var(--theme-bg);
  border: 2px solid var(--theme-border);
  border-radius: 8px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3);
  z-index: 1000;
  display: flex;
  flex-direction: column;
}

.dropdown-header {
  padding: 1rem;
  border-bottom: 1px solid var(--theme-border);
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.dropdown-header h3 {
  margin: 0;
  color: var(--theme-primary);
  font-size: 1.1rem;
}

.header-actions {
  display: flex;
  gap: 0.5rem;
}

.btn-link {
  background: none;
  border: none;
  color: var(--theme-accent);
  cursor: pointer;
  font-size: 0.9rem;
  padding: 0.25rem 0.5rem;
  border-radius: 4px;
}

.btn-link:hover {
  background: rgba(189, 147, 249, 0.1);
}

.no-notifications {
  padding: 2rem;
  text-align: center;
  color: var(--theme-fg);
  opacity: 0.6;
}

.no-notifications svg {
  margin-bottom: 1rem;
}

.notifications-list {
  flex: 1;
  overflow-y: auto;
  max-height: 400px;
}

.dropdown-footer {
  padding: 1rem;
  border-top: 1px solid var(--theme-border);
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.view-all {
  color: var(--theme-accent);
  text-decoration: none;
  font-size: 0.9rem;
}

.view-all:hover {
  text-decoration: underline;
}

@media (max-width: 768px) {
  .notification-dropdown {
    position: fixed;
    top: 60px;
    right: 1rem;
    left: 1rem;
    width: auto;
    max-height: calc(100vh - 80px);
  }
}
</style>