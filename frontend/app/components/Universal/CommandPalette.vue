// PWD OF FILE
/home/lilith/code-shit/lilithinaparka.i2p/frontend/app/components/Universal
// FILE NAME
CommandPalette.vue

<template>
  <div class="command-palette-overlay" @click.self="$emit('close')">
    <div class="command-palette system-window">
      <div class="palette-header">
        <div class="palette-icon">
          <FeatherIcon name="command" size="20" />
        </div>
        <input
          ref="searchInput"
          v-model="searchQuery"
          class="palette-input"
          placeholder="Type a command or search..."
          @keydown.esc="$emit('close')"
          @keydown.up="navigateUp"
          @keydown.down="navigateDown"
          @keydown.enter="executeSelected"
        />
        <div class="palette-hint">
          <kbd>Esc</kbd> to close
        </div>
      </div>
      
      <div class="palette-results">
        <div 
          v-for="(result, index) in filteredResults"
          :key="result.id"
          class="result-item"
          :class="{ active: selectedIndex === index }"
          @click="executeCommand(result)"
          @mouseenter="selectedIndex = index"
        >
          <div class="result-icon">
            <FeatherIcon :name="result.icon" size="16" />
          </div>
          <div class="result-content">
            <div class="result-title">
              {{ result.title }}
              <span v-if="result.category" class="result-category">
                {{ result.category }}
              </span>
            </div>
            <div class="result-description">
              {{ result.description }}
            </div>
          </div>
          <div class="result-shortcut">
            <kbd v-for="key in result.shortcut" :key="key">{{ key }}</kbd>
          </div>
        </div>
        
        <div v-if="filteredResults.length === 0" class="no-results">
          <FeatherIcon name="search" size="24" />
          <p>No commands found for "{{ searchQuery }}"</p>
        </div>
      </div>
      
      <div class="palette-footer">
        <div class="footer-section">
          <div class="footer-item">
            <kbd>↑</kbd>
            <kbd>↓</kbd>
            <span>Navigate</span>
          </div>
          <div class="footer-item">
            <kbd>Enter</kbd>
            <span>Execute</span>
          </div>
        </div>
        <div class="footer-section">
          <div class="result-count">
            {{ filteredResults.length }} result{{ filteredResults.length !== 1 ? 's' : '' }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'

const emit = defineEmits(['close'])
const router = useRouter()
const searchInput = ref<HTMLInputElement | null>(null)
const searchQuery = ref('')
const selectedIndex = ref(0)

const commands = [
  // Navigation
  { 
    id: 'home', 
    title: 'Go to Home', 
    description: 'Navigate to dashboard',
    category: 'Navigation',
    icon: 'home',
    action: () => router.push('/'),
    shortcut: ['H']
  },
  { 
    id: 'blog', 
    title: 'Go to Blog', 
    description: 'Navigate to blog section',
    category: 'Navigation',
    icon: 'file-text',
    action: () => router.push('/blog'),
    shortcut: ['B']
  },
  { 
    id: 'games', 
    title: 'Go to Games', 
    description: 'Navigate to games section',
    category: 'Navigation',
    icon: 'play',
    action: () => router.push('/games'),
    shortcut: ['G']
  },
  { 
    id: 'art', 
    title: 'Go to Gallery', 
    description: 'Navigate to art gallery',
    category: 'Navigation',
    icon: 'image',
    action: () => router.push('/art'),
    shortcut: ['A']
  },
  { 
    id: 'profile', 
    title: 'Go to Profile', 
    description: 'Navigate to profile page',
    category: 'Navigation',
    icon: 'user',
    action: () => router.push('/profile'),
    shortcut: ['P']
  },
  
  // Theme
  { 
    id: 'theme-switcher', 
    title: 'Open Theme Switcher', 
    description: 'Change application theme',
    category: 'Appearance',
    icon: 'palette',
    action: () => {
      // Open theme switcher
      emit('close')
      // You would trigger your theme switcher modal here
    },
    shortcut: ['T']
  },
  { 
    id: 'toggle-dark', 
    title: 'Toggle Dark Mode', 
    description: 'Switch between dark and light mode',
    category: 'Appearance',
    icon: 'moon',
    action: () => {
      const isDark = document.documentElement.getAttribute('data-colorscheme') === 'dark'
      document.documentElement.setAttribute('data-colorscheme', isDark ? 'light' : 'dark')
    },
    shortcut: ['D']
  },
  
  // Actions
  { 
    id: 'refresh', 
    title: 'Refresh Page', 
    description: 'Reload current page',
    category: 'Actions',
    icon: 'refresh-cw',
    action: () => location.reload(),
    shortcut: ['R']
  },
  { 
    id: 'search', 
    title: 'Search Content', 
    description: 'Search across the site',
    category: 'Actions',
    icon: 'search',
    action: () => router.push('/search'),
    shortcut: ['/', 'S']
  },
  { 
    id: 'notifications', 
    title: 'View Notifications', 
    description: 'Open notification center',
    category: 'Actions',
    icon: 'bell',
    action: () => {
      emit('close')
      // Open notification center
    },
    shortcut: ['N']
  },
  
  // System
  { 
    id: 'settings', 
    title: 'Open Settings', 
    description: 'Open application settings',
    category: 'System',
    icon: 'settings',
    action: () => router.push('/settings'),
    shortcut: [',', 'S']
  },
  { 
    id: 'help', 
    title: 'Open Help', 
    description: 'View help documentation',
    category: 'System',
    icon: 'help-circle',
    action: () => router.push('/help'),
    shortcut: ['?']
  },
  { 
    id: 'about', 
    title: 'About', 
    description: 'View information about this site',
    category: 'System',
    icon: 'info',
    action: () => router.push('/about'),
    shortcut: []
  },
  
  // External
  { 
    id: 'github', 
    title: 'View Source Code', 
    description: 'Open GitHub repository',
    category: 'External',
    icon: 'github',
    action: () => window.open('https://github.com/lilithinaparka', '_blank'),
    shortcut: []
  },
  { 
    id: 'bluesky', 
    title: 'Open BlueSky', 
    description: 'Open BlueSky profile',
    category: 'External',
    icon: 'message-circle',
    action: () => window.open('https://bsky.app/profile/lilithinaparka', '_blank'),
    shortcut: []
  }
]

const filteredResults = computed(() => {
  if (!searchQuery.value.trim()) {
    return commands.slice(0, 10)
  }
  
  const query = searchQuery.value.toLowerCase()
  return commands.filter(cmd => 
    cmd.title.toLowerCase().includes(query) ||
    cmd.description.toLowerCase().includes(query) ||
    cmd.category?.toLowerCase().includes(query)
  ).slice(0, 10)
})

watch(searchQuery, () => {
  selectedIndex.value = 0
})

const navigateUp = () => {
  if (selectedIndex.value > 0) {
    selectedIndex.value--
  } else {
    selectedIndex.value = filteredResults.value.length - 1
  }
}

const navigateDown = () => {
  if (selectedIndex.value < filteredResults.value.length - 1) {
    selectedIndex.value++
  } else {
    selectedIndex.value = 0
  }
}

const executeSelected = () => {
  if (filteredResults.value[selectedIndex.value]) {
    executeCommand(filteredResults.value[selectedIndex.value])
  }
}

const executeCommand = (command: any) => {
  command.action()
  emit('close')
}

onMounted(() => {
  searchInput.value?.focus()
})
</script>

<style scoped>
@import '~/assets/css/design-system.css';

.command-palette-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.8);
  backdrop-filter: blur(4px);
  display: flex;
  align-items: flex-start;
  justify-content: center;
  padding-top: 10%;
  z-index: 2000;
  animation: fade-in var(--transition-fast);
}

.command-palette {
  width: 100%;
  max-width: 600px;
  max-height: 70vh;
  display: flex;
  flex-direction: column;
  animation: slide-down var(--transition-base);
  box-shadow: var(--shadow-xl);
}

@keyframes fade-in {
  from { opacity: 0; }
  to { opacity: 1; }
}

@keyframes slide-down {
  from { 
    opacity: 0;
    transform: translateY(-20px);
  }
  to { 
    opacity: 1;
    transform: translateY(0);
  }
}

.palette-header {
  padding: var(--space-md);
  border-bottom: 1px solid var(--theme-border);
  display: flex;
  align-items: center;
  gap: var(--space-md);
}

.palette-icon {
  color: var(--theme-primary);
}

.palette-input {
  flex: 1;
  background: transparent;
  border: none;
  color: var(--theme-fg);
  font-family: var(--font-ui);
  font-size: 1.125rem;
  outline: none;
  padding: 0;
}

.palette-input::placeholder {
  color: var(--theme-muted);
}

.palette-hint {
  display: flex;
  align-items: center;
  gap: var(--space-xs);
  font-size: 0.75rem;
  color: var(--theme-muted);
}

.palette-hint kbd {
  padding: 1px 4px;
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid var(--theme-border);
  border-radius: 3px;
  font-family: var(--font-mono);
  font-size: 0.7rem;
}

.palette-results {
  flex: 1;
  overflow-y: auto;
  max-height: 400px;
  padding: var(--space-sm);
}

.result-item {
  padding: var(--space-sm) var(--space-md);
  border-radius: var(--radius-md);
  display: flex;
  align-items: center;
  gap: var(--space-md);
  cursor: pointer;
  transition: all var(--transition-fast);
  position: relative;
  overflow: hidden;
}

.result-item:hover {
  background: rgba(255, 255, 255, 0.05);
}

.result-item.active {
  background: rgba(var(--theme-primary-rgb), 0.1);
}

.result-item::before {
  content: '';
  position: absolute;
  left: 0;
  top: 0;
  bottom: 0;
  width: 3px;
  background: var(--theme-primary);
  opacity: 0;
  transition: opacity var(--transition-fast);
}

.result-item.active::before {
  opacity: 1;
}

.result-icon {
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(255, 255, 255, 0.05);
  border-radius: var(--radius-sm);
  color: var(--theme-primary);
  flex-shrink: 0;
}

.result-content {
  flex: 1;
  min-width: 0;
}

.result-title {
  display: flex;
  align-items: center;
  gap: var(--space-sm);
  font-weight: 500;
  color: var(--theme-fg);
  margin-bottom: 2px;
}

.result-category {
  font-size: 0.75rem;
  color: var(--theme-muted);
  background: rgba(255, 255, 255, 0.05);
  padding: 1px 6px;
  border-radius: var(--radius-sm);
  font-weight: normal;
}

.result-description {
  font-size: 0.875rem;
  color: var(--theme-muted);
  line-height: 1.4;
}

.result-shortcut {
  display: flex;
  gap: 2px;
  flex-shrink: 0;
}

.result-shortcut kbd {
  padding: 2px 6px;
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid var(--theme-border);
  border-radius: 3px;
  font-family: var(--font-mono);
  font-size: 0.75rem;
  color: var(--theme-muted);
}

.no-results {
  padding: var(--space-xl) var(--space-md);
  text-align: center;
  color: var(--theme-muted);
}

.no-results p {
  margin: var(--space-md) 0 0 0;
}

.palette-footer {
  padding: var(--space-sm) var(--space-md);
  border-top: 1px solid var(--theme-border);
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 0.75rem;
  color: var(--theme-muted);
}

.footer-section {
  display: flex;
  align-items: center;
  gap: var(--space-lg);
}

.footer-item {
  display: flex;
  align-items: center;
  gap: 4px;
}

.footer-item kbd {
  padding: 1px 4px;
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid var(--theme-border);
  border-radius: 3px;
  font-family: var(--font-mono);
  font-size: 0.7rem;
}

.result-count {
  font-weight: 500;
  color: var(--theme-fg);
}

@media (max-width: 768px) {
  .command-palette-overlay {
    padding-top: 20%;
  }
  
  .command-palette {
    max-width: 90%;
    max-height: 80vh;
  }
  
  .result-shortcut {
    display: none;
  }
}
</style>

// END OF FILE