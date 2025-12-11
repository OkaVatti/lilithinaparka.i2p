<template>
  <div 
    class="notification-item" 
    :class="[notification.type, { unread: !notification.read }]"
    @click="$emit('click', notification)"
  >
    <div class="notification-icon">
      <FeatherIcon :name="notificationIcon" size="18" />
    </div>
    
    <div class="notification-content">
      <h4 v-if="notification.title" class="notification-title">{{ notification.title }}</h4>
      <p class="notification-message">{{ notification.message }}</p>
      <span class="notification-time">{{ formatTimeAgo(notification.created_at) }}</span>
    </div>
    
    <div class="notification-actions">
      <button 
        v-if="!notification.read" 
        @click.stop="markAsRead"
        class="action-btn"
        title="Mark as read"
      >
        <FeatherIcon name="check" size="16" />
      </button>
      <button 
        @click.stop="deleteNotification"
        class="action-btn delete"
        title="Delete"
      >
        <FeatherIcon name="x" size="16" />
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Notification } from '~~/types'

const props = defineProps<{
  notification: Notification
}>()

const emit = defineEmits<{
  click: [notification: Notification]
}>()

const { apiFetch } = useApi()

const notificationIcon = computed(() => {
  switch (props.notification.type) {
    case 'success': return 'check-circle'
    case 'warning': return 'alert-triangle'
    case 'error': return 'alert-circle'
    case 'info':
    default: return 'info'
  }
})

const markAsRead = async () => {
  try {
    await apiFetch(`/notifications/${props.notification.id}/read`, {
      method: 'PUT'
    })
    // The parent component should handle refreshing
  } catch (error) {
    console.error('Failed to mark as read:', error)
  }
}

const deleteNotification = async () => {
  try {
    await apiFetch(`/notifications/${props.notification.id}`, {
      method: 'DELETE'
    })
    // The parent component should handle refreshing
  } catch (error) {
    console.error('Failed to delete notification:', error)
  }
}

const formatTimeAgo = (dateStr: string) => {
  const date = new Date(dateStr)
  const now = new Date()
  const diffMs = now.getTime() - date.getTime()
  const diffMins = Math.floor(diffMs / 60000)
  
  if (diffMins < 1) return 'Just now'
  if (diffMins < 60) return `${diffMins}m ago`
  if (diffMins < 1440) return `${Math.floor(diffMins / 60)}h ago`
  if (diffMins < 10080) return `${Math.floor(diffMins / 1440)}d ago`
  
  return date.toLocaleDateString('en-US', {
    month: 'short',
    day: 'numeric',
    year: date.getFullYear() !== now.getFullYear() ? 'numeric' : undefined
  })
}
</script>

<style scoped>
.notification-item {
  display: flex;
  gap: 1rem;
  padding: 1rem;
  border-bottom: 1px solid var(--theme-border);
  cursor: pointer;
  transition: all 0.2s;
}

.notification-item:hover {
  background: rgba(189, 147, 249, 0.05);
}

.notification-item:last-child {
  border-bottom: none;
}

.notification-item.unread {
  background: rgba(189, 147, 249, 0.1);
}

.notification-icon {
  flex-shrink: 0;
  width: 36px;
  height: 36px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
}

.notification-item.info .notification-icon {
  background: rgba(139, 147, 233, 0.2);
  color: #8b8be9;
}

.notification-item.success .notification-icon {
  background: rgba(80, 250, 123, 0.2);
  color: #50fa7b;
}

.notification-item.warning .notification-icon {
  background: rgba(255, 184, 108, 0.2);
  color: #ffb86c;
}

.notification-item.error .notification-icon {
  background: rgba(255, 85, 85, 0.2);
  color: #ff5555;
}

.notification-content {
  flex: 1;
  min-width: 0;
}

.notification-title {
  margin: 0 0 0.25rem 0;
  font-size: 0.95rem;
  color: var(--theme-primary);
}

.notification-message {
  margin: 0 0 0.5rem 0;
  font-size: 0.9rem;
  color: var(--theme-fg);
  line-height: 1.4;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.notification-time {
  font-size: 0.8rem;
  color: var(--theme-fg);
  opacity: 0.6;
}

.notification-actions {
  display: flex;
  gap: 0.25rem;
  align-items: flex-start;
}

.action-btn {
  padding: 0.25rem;
  background: none;
  border: none;
  color: var(--theme-fg);
  opacity: 0.6;
  cursor: pointer;
  border-radius: 4px;
  transition: all 0.2s;
}

.action-btn:hover {
  opacity: 1;
  background: rgba(255, 255, 255, 0.1);
}

.action-btn.delete:hover {
  color: var(--theme-error);
}
</style>