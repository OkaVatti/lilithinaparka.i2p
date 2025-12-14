// PWD OF FILE
/home/lilith/code-shit/lilithinaparka.i2p/frontend/app/layout
// FILE NAME
enhanced.vue

<template>
  <div class="enhanced-layout" :class="layoutClasses">
    <!-- Terminal-style status bar -->
    <div class="status-bar system-window">
      <div class="status-left">
        <div class="status-item">
          <FeatherIcon name="cpu" size="14" />
          <span class="status-label">SYSTEM</span>
          <span class="status-value pulse-animation">ONLINE</span>
        </div>
        <div class="status-item">
          <FeatherIcon name="shield" size="14" />
          <span class="status-label">SECURITY</span>
          <span class="status-value">I2P ENCRYPTED</span>
        </div>
        <div class="status-item">
          <FeatherIcon name="clock" size="14" />
          <span class="status-label">TIME</span>
          <span class="status-value">{{ currentTime }}</span>
        </div>
      </div>
      
      <div class="status-center">
        <div class="connection-status">
          <div class="connection-dot" :class="{ connected: isConnected }"></div>
          <span class="connection-text">{{ connectionText }}</span>
        </div>
      </div>
      
      <div class="status-right">
        <div class="status-item">
          <FeatherIcon name="thermometer" size="14" />
          <span class="status-label">THEME</span>
          <span class="status-value">{{ currentThemeName }}</span>
        </div>
        <div class="status-item">
          <FeatherIcon name="hard-drive" size="14" />
          <span class="status-label">MEM</span>
          <span class="status-value">{{ memoryUsage }}</span>
        </div>
        <ThemeSwitcherCompact />
      </div>
    </div>
    
    <!-- Main content with terminal split -->
    <div class="main-terminal">
      <!-- Left sidebar - Navigation -->
      <aside class="terminal-sidebar system-window">
        <div class="sidebar-header">
          <div class="sidebar-title">
            <FeatherIcon name="terminal" size="20" />
            <h3>NAVIGATION</h3>
          </div>
          <button 
            @click="toggleSidebar" 
            class="sidebar-toggle"
            :title="sidebarCollapsed ? 'Expand sidebar' : 'Collapse sidebar'"
          >
            <FeatherIcon :name="sidebarCollapsed ? 'chevron-right' : 'chevron-left'" size="16" />
          </button>
        </div>
        
        <transition name="slide">
          <nav v-if="!sidebarCollapsed" class="sidebar-nav">
            <div class="nav-section">
              <h4 class="nav-section-title">MAIN</h4>
              <NavLink to="/" icon="home" label="Dashboard" />
              <NavLink to="/blog" icon="file-text" label="Blog" />
              <NavLink to="/games" icon="play" label="Games" />
              <NavLink to="/art" icon="image" label="Gallery" />
              <NavLink to="/profile" icon="user" label="Profile" />
            </div>
            
            <div class="nav-section">
              <h4 class="nav-section-title">SOCIAL</h4>
              <NavLink to="/bluesky" icon="message-circle" label="BlueSky" />
              <NavLink to="/rss" icon="rss" label="RSS Feed" />
              <NavLink to="/contact" icon="mail" label="Contact" />
            </div>
            
            <div class="nav-section">
              <h4 class="nav-section-title">TOOLS</h4>
              <NavLink to="/search" icon="search" label="Search" />
              <NavLink to="/analytics" icon="bar-chart-2" label="Analytics" />
              <NavLink to="/settings" icon="settings" label="Settings" />
            </div>
            
            <div class="sidebar-footer">
              <div class="system-info">
                <div class="info-item">
                  <span class="info-label">UPTIME</span>
                  <span class="info-value">{{ uptime }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">VISITORS</span>
                  <span class="info-value">{{ visitorCount }}</span>
                </div>
              </div>
            </div>
          </nav>
        </transition>
      </aside>
      
      <!-- Main content area -->
      <main class="terminal-main">
        <!-- Terminal header -->
        <div class="terminal-header system-window">
          <div class="terminal-prompt">
            <span class="prompt-user">lilith@i2p:~</span>
            <span class="prompt-path">/{{ currentPath }}</span>
            <span class="prompt-symbol">$</span>
            <div class="terminal-input-line">
              <span class="input-cursor cursor-blink"></span>
            </div>
          </div>
          
          <div class="terminal-actions">
            <button @click="executeCommand('help')" class="action-btn" title="Help">
              <FeatherIcon name="help-circle" size="16" />
            </button>
            <button @click="clearTerminal" class="action-btn" title="Clear terminal">
              <FeatherIcon name="trash-2" size="16" />
            </button>
            <button @click="toggleFullscreen" class="action-btn" title="Fullscreen">
              <FeatherIcon :name="isFullscreen ? 'minimize' : 'maximize'" size="16" />
            </button>
          </div>
        </div>
        
        <!-- Terminal output/Content area -->
        <div class="terminal-output-area" ref="terminalOutput">
          <!-- Terminal welcome message -->
          <div v-if="showWelcome" class="terminal-welcome">
            <pre class="ascii-art">{{ asciiArt }}</pre>
            <div class="welcome-message">
              <p class="welcome-line">Welcome to lilithinaparka.i2p</p>
              <p class="welcome-line">Type 'help' for available commands</p>
              <p class="welcome-line">----------------------------------</p>
            </div>
          </div>
          
          <!-- Command history -->
          <div 
            v-for="(command, index) in commandHistory" 
            :key="index"
            class="command-history"
          >
            <div class="command-line">
              <span class="prompt-user">guest@i2p:~</span>
              <span class="prompt-symbol">$</span>
              <span class="command-text">{{ command.input }}</span>
            </div>
            <div class="command-output" v-html="command.output"></div>
          </div>
          
          <!-- Dynamic content slot -->
          <div class="dynamic-content">
            <slot />
          </div>
        </div>
        
        <!-- Terminal input -->
        <div class="terminal-input-area system-window">
          <div class="input-wrapper">
            <div class="input-prefix">
              <span class="prefix-user">guest@i2p:~</span>
              <span class="prefix-symbol">$</span>
            </div>
            <input
              v-model="currentCommand"
              @keyup.enter="executeCommand(currentCommand)"
              @keyup.up="navigateHistory(-1)"
              @keyup.down="navigateHistory(1)"
              class="terminal-input"
              placeholder="Type a command..."
              ref="commandInput"
            />
            <div class="input-suffix">
              <button @click="executeCommand(currentCommand)" class="execute-btn">
                <FeatherIcon name="arrow-right" size="16" />
              </button>
            </div>
          </div>
          
          <div class="input-hints">
            <span class="hint-label">Try:</span>
            <button 
              v-for="hint in commandHints" 
              :key="hint"
              @click="executeCommand(hint)"
              class="hint-btn"
            >
              {{ hint }}
            </button>
          </div>
        </div>
      </main>
      
      <!-- Right sidebar - Quick actions -->
      <aside class="quick-actions system-window">
        <div class="actions-header">
          <FeatherIcon name="zap" size="18" />
          <h4>QUICK ACTIONS</h4>
        </div>
        
        <div class="actions-grid">
          <QuickAction 
            icon="sun"
            label="Toggle Theme"
            :action="toggleTheme"
            color="var(--theme-primary)"
          />
          <QuickAction 
            icon="lock"
            label="Privacy Mode"
            :action="togglePrivacy"
            :active="privacyMode"
            color="var(--theme-accent)"
          />
          <QuickAction 
            icon="refresh-cw"
            label="Refresh Data"
            :action="refreshData"
            color="var(--theme-success)"
          />
          <QuickAction 
            icon="download"
            label="Backup"
            :action="createBackup"
            color="var(--theme-warning)"
          />
          <QuickAction 
            icon="code"
            label="View Source"
            :action="viewSource"
            color="var(--theme-info)"
          />
          <QuickAction 
            icon="moon"
            label="Night Mode"
            :action="toggleNightMode"
            :active="nightMode"
            color="var(--theme-muted)"
          />
        </div>
        
        <div class="quick-stats">
          <div class="stat-card">
            <FeatherIcon name="eye" size="14" />
            <div class="stat-content">
              <div class="stat-value">{{ pageViews }}</div>
              <div class="stat-label">Views Today</div>
            </div>
          </div>
          <div class="stat-card">
            <FeatherIcon name="users" size="14" />
            <div class="stat-content">
              <div class="stat-value">{{ activeUsers }}</div>
              <div class="stat-label">Active Now</div>
            </div>
          </div>
        </div>
      </aside>
    </div>
    
    <!-- Notification center -->
    <transition name="slide-up">
      <NotificationCenter v-if="showNotifications" />
    </transition>
    
    <!-- Command palette (Ctrl+K) -->
    <CommandPalette 
      v-if="showCommandPalette"
      @close="showCommandPalette = false"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useTheme } from '../composables/useTheme'
import ThemeSwitcherCompact from '../components/Universal/Theme/ThemeSwitcherCompact.vue'
import NavLink from '../components/Universal/NavLink.vue'
import QuickAction from '../components/Universal/QuickAction.vue'
import NotificationCenter from '../components/Universal/NotificationCenter.vue'
import CommandPalette from '../components/Universal/CommandPalette.vue'

const route = useRoute()
const router = useRouter()
const { currentTheme, themes, setTheme } = useTheme()

const currentTime = ref('')
const currentCommand = ref('')
const commandHistory = ref<any[]>([])
const historyIndex = ref(-1)
const sidebarCollapsed = ref(false)
const isFullscreen = ref(false)
const showWelcome = ref(true)
const privacyMode = ref(false)
const nightMode = ref(false)
const showNotifications = ref(false)
const showCommandPalette = ref(false)
const isConnected = ref(true)
const commandInput = ref<HTMLInputElement | null>(null)
const terminalOutput = ref<HTMLDivElement | null>(null)

const currentPath = computed(() => {
  return route.path.slice(1) || 'home'
})

const currentThemeName = computed(() => {
  const theme = themes.find(t => t.value === currentTheme.value)
  return theme?.name || 'Dracula'
})

const connectionText = computed(() => {
  return isConnected.value ? 'CONNECTED' : 'DISCONNECTED'
})

const memoryUsage = computed(() => {
  // Simulated memory usage
  const used = Math.floor(Math.random() * 4) + 2
  const total = 8
  return `${used}/${total}GB`
})

const uptime = computed(() => {
  const hours = Math.floor((Date.now() - startTime) / 3600000)
  const minutes = Math.floor(((Date.now() - startTime) % 3600000) / 60000)
  return `${hours}h ${minutes}m`
})

const visitorCount = computed(() => {
  // Simulated visitor count
  return Math.floor(Math.random() * 1000) + 42
})

const pageViews = computed(() => {
  return Math.floor(Math.random() * 500) + 100
})

const activeUsers = computed(() => {
  return Math.floor(Math.random() * 50) + 5
})

const layoutClasses = computed(() => ({
  'sidebar-collapsed': sidebarCollapsed.value,
  'fullscreen': isFullscreen.value,
  'privacy-mode': privacyMode.value,
  'night-mode': nightMode.value
}))

const commandHints = computed(() => {
  const base = ['help', 'clear', 'theme', 'blog', 'games']
  const path = route.path.slice(1)
  if (path && !base.includes(path)) {
    return ['help', path, 'clear', 'theme']
  }
  return base
})

const asciiArt = computed(() => {
  return `
╔══════════════════════════════════════╗
║                                      ║
║    ██╗     ██╗██╗██╗  ██╗██╗  ██╗   ║
║    ██║     ██║██║██║ ██╔╝██║  ██║   ║
║    ██║     ██║██║█████╔╝ ███████║   ║
║    ██║     ██║██║██╔═██╗ ██╔══██║   ║
║    ███████╗██║██║██║  ██╗██║  ██║   ║
║    ╚══════╝╚═╝╚═╝╚═╝  ╚═╝╚═╝  ╚═╝   ║
║                                      ║
║    in a parka.i2p                    ║
║                                      ║
╚══════════════════════════════════════╝
  `
})

const startTime = Date.now()

// Commands available in the terminal
const commands = {
  help: () => `
Available commands:
  help          - Show this help message
  clear         - Clear terminal
  theme [name]  - Change theme
  blog          - Go to blog
  games         - Go to games
  art           - Go to gallery
  profile       - Go to profile
  about         - Show about info
  date          - Show current date/time
  echo [text]   - Echo text
  ls            - List directory
  pwd           - Print working directory
  `,
  
  clear: () => {
    commandHistory.value = []
    showWelcome.value = false
    return 'Terminal cleared.'
  },
  
  theme: (args: string[]) => {
    if (args.length === 0) {
      return 'Usage: theme [theme_name]\nAvailable themes: ' + 
        themes.map(t => t.value).join(', ')
    }
    
    const themeName = args[0]
    const theme = themes.find(t => t.value === themeName)
    
    if (theme) {
      setTheme(theme.value)
      return `Theme changed to ${theme.name}`
    } else {
      return `Theme "${themeName}" not found`
    }
  },
  
  blog: () => {
    router.push('/blog')
    return 'Navigating to blog...'
  },
  
  games: () => {
    router.push('/games')
    return 'Navigating to games...'
  },
  
  art: () => {
    router.push('/art')
    return 'Navigating to gallery...'
  },
  
  profile: () => {
    router.push('/profile')
    return 'Navigating to profile...'
  },
  
  about: () => `
lilithinaparka.i2p - Privacy-first portal
Version: 2.0.0
Built with: Nuxt 3, Vue 3, Go, I2P
License: AGPL-3.0
Source: https://github.com/lilithinaparka
  `,
  
  date: () => {
    return new Date().toLocaleString()
  },
  
  echo: (args: string[]) => {
    return args.join(' ')
  },
  
  ls: () => {
    return `
./
../
blog/
games/
art/
profile/
about.md
config.ini
  `
  },
  
  pwd: () => {
    return route.path
  }
}

const updateTime = () => {
  const now = new Date()
  currentTime.value = now.toLocaleTimeString('en-US', {
    hour12: false,
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit'
  })
}

const executeCommand = (input: string) => {
  if (!input.trim()) return
  
  const args = input.trim().split(/\s+/)
  const commandName = args[0].toLowerCase()
  const commandArgs = args.slice(1)
  
  let output = ''
  
  if (commands[commandName as keyof typeof commands]) {
    const result = commands[commandName as keyof typeof commands](commandArgs)
    output = typeof result === 'string' ? result : result()
  } else {
    output = `Command not found: ${commandName}\nType 'help' for available commands`
  }
  
  commandHistory.value.push({
    input,
    output: output.replace(/\n/g, '<br>')
  })
  
  currentCommand.value = ''
  historyIndex.value = -1
  
  // Scroll to bottom of terminal
  setTimeout(() => {
    if (terminalOutput.value) {
      terminalOutput.value.scrollTop = terminalOutput.value.scrollHeight
    }
  }, 10)
}

const navigateHistory = (direction: number) => {
  if (commandHistory.value.length === 0) return
  
  historyIndex.value = Math.max(-1, 
    Math.min(commandHistory.value.length - 1, historyIndex.value + direction))
  
  if (historyIndex.value >= 0) {
    currentCommand.value = commandHistory.value[historyIndex.value].input
  } else {
    currentCommand.value = ''
  }
}

const clearTerminal = () => {
  executeCommand('clear')
}

const toggleSidebar = () => {
  sidebarCollapsed.value = !sidebarCollapsed.value
}

const toggleFullscreen = () => {
  if (!document.fullscreenElement) {
    document.documentElement.requestFullscreen()
    isFullscreen.value = true
  } else {
    document.exitFullscreen()
    isFullscreen.value = false
  }
}

const toggleTheme = () => {
  const currentIndex = themes.findIndex(t => t.value === currentTheme.value)
  const nextIndex = (currentIndex + 1) % themes.length
  setTheme(themes[nextIndex].value)
}

const togglePrivacy = () => {
  privacyMode.value = !privacyMode.value
  localStorage.setItem('privacyMode', privacyMode.value.toString())
}

const toggleNightMode = () => {
  nightMode.value = !nightMode.value
  document.documentElement.classList.toggle('night-mode', nightMode.value)
  localStorage.setItem('nightMode', nightMode.value.toString())
}

const refreshData = () => {
  // In a real app, this would refresh all data from the API
  location.reload()
}

const createBackup = () => {
  // Simulated backup creation
  const backupData = {
    timestamp: new Date().toISOString(),
    theme: currentTheme.value,
    settings: {
      privacyMode: privacyMode.value,
      nightMode: nightMode.value
    }
  }
  
  const blob = new Blob([JSON.stringify(backupData, null, 2)], {
    type: 'application/json'
  })
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = `backup-${new Date().toISOString().slice(0, 10)}.json`
  a.click()
  URL.revokeObjectURL(url)
}

const viewSource = () => {
  window.open('https://github.com/lilithinaparka', '_blank')
}

// Handle keyboard shortcuts
const handleKeydown = (e: KeyboardEvent) => {
  // Ctrl+K for command palette
  if ((e.ctrlKey || e.metaKey) && e.key === 'k') {
    e.preventDefault()
    showCommandPalette.value = !showCommandPalette.value
  }
  
  // Esc to close modals
  if (e.key === 'Escape') {
    if (showCommandPalette.value) {
      showCommandPalette.value = false
    }
    if (showNotifications.value) {
      showNotifications.value = false
    }
  }
  
  // Focus command input on /
  if (e.key === '/' && !e.ctrlKey && !e.metaKey) {
    e.preventDefault()
    commandInput.value?.focus()
  }
}

onMounted(() => {
  // Update time every second
  updateTime()
  const timeInterval = setInterval(updateTime, 1000)
  
  // Load saved settings
  privacyMode.value = localStorage.getItem('privacyMode') === 'true'
  nightMode.value = localStorage.getItem('nightMode') === 'true'
  
  // Add event listeners
  document.addEventListener('keydown', handleKeydown)
  document.addEventListener('fullscreenchange', () => {
    isFullscreen.value = !!document.fullscreenElement
  })
  
  // Simulate connection status changes
  const connectionInterval = setInterval(() => {
    if (Math.random() > 0.95) {
      isConnected.value = !isConnected.value
    }
  }, 10000)
  
  onUnmounted(() => {
    clearInterval(timeInterval)
    clearInterval(connectionInterval)
    document.removeEventListener('keydown', handleKeydown)
  })
})
</script>

<style scoped>
@import '~/assets/css/design-system.css';

.enhanced-layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background: var(--theme-bg);
  color: var(--theme-fg);
  font-family: var(--font-mono);
}

.status-bar {
  padding: var(--space-xs) var(--space-md);
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: var(--theme-surface);
  border-bottom: 2px solid var(--theme-border);
  font-size: 0.75rem;
}

.status-left,
.status-right {
  display: flex;
  gap: var(--space-lg);
  align-items: center;
}

.status-center {
  flex: 1;
  display: flex;
  justify-content: center;
}

.status-item {
  display: flex;
  align-items: center;
  gap: var(--space-xs);
  color: var(--theme-muted);
}

.status-label {
  font-weight: 600;
  letter-spacing: 0.05em;
  color: var(--theme-primary);
}

.status-value {
  font-weight: 500;
  color: var(--theme-fg);
}

.connection-status {
  display: flex;
  align-items: center;
  gap: var(--space-xs);
}

.connection-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: var(--theme-error);
}

.connection-dot.connected {
  background: var(--theme-success);
  animation: var(--animation-pulse);
}

.connection-text {
  font-weight: 500;
}

.main-terminal {
  flex: 1;
  display: flex;
  gap: var(--space-md);
  padding: var(--space-md);
  overflow: hidden;
}

.terminal-sidebar {
  width: var(--terminal-sidebar-width);
  min-width: var(--terminal-sidebar-width);
  display: flex;
  flex-direction: column;
  transition: all var(--transition-base);
}

.enhanced-layout.sidebar-collapsed .terminal-sidebar {
  width: 60px;
  min-width: 60px;
}

.sidebar-header {
  padding: var(--space-sm);
  border-bottom: 1px solid var(--theme-border);
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.sidebar-title {
  display: flex;
  align-items: center;
  gap: var(--space-xs);
}

.sidebar-title h3 {
  margin: 0;
  font-size: 0.875rem;
  color: var(--theme-primary);
  letter-spacing: 0.05em;
}

.sidebar-toggle {
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: transparent;
  border: 1px solid var(--theme-border);
  border-radius: var(--radius-sm);
  color: var(--theme-fg);
  cursor: pointer;
  transition: all var(--transition-fast);
}

.sidebar-toggle:hover {
  background: rgba(255, 255, 255, 0.1);
  border-color: var(--theme-primary);
}

.sidebar-nav {
  flex: 1;
  padding: var(--space-md);
  display: flex;
  flex-direction: column;
  gap: var(--space-lg);
  overflow-y: auto;
}

.nav-section {
  display: flex;
  flex-direction: column;
  gap: var(--space-xs);
}

.nav-section-title {
  margin: 0 0 var(--space-xs) 0;
  font-size: 0.75rem;
  font-weight: 600;
  color: var(--theme-muted);
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.sidebar-footer {
  margin-top: auto;
  padding-top: var(--space-md);
  border-top: 1px solid var(--theme-border);
}

.system-info {
  display: flex;
  flex-direction: column;
  gap: var(--space-xs);
}

.info-item {
  display: flex;
  justify-content: space-between;
  font-size: 0.75rem;
}

.info-label {
  color: var(--theme-muted);
}

.info-value {
  color: var(--theme-fg);
  font-weight: 500;
}

.terminal-main {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: var(--space-md);
  min-width: 0;
}

.terminal-header {
  padding: var(--space-sm) var(--space-md);
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: var(--theme-surface);
}

.terminal-prompt {
  display: flex;
  align-items: center;
  gap: var(--space-xs);
  font-family: var(--font-mono);
  font-size: 0.875rem;
}

.prompt-user {
  color: var(--theme-accent);
  font-weight: 600;
}

.prompt-path {
  color: var(--theme-primary);
}

.prompt-symbol {
  color: var(--theme-success);
  margin-right: var(--space-xs);
}

.terminal-input-line {
  flex: 1;
  min-width: 200px;
}

.input-cursor {
  display: inline-block;
  width: 8px;
  height: 16px;
  background: var(--theme-accent);
  vertical-align: middle;
}

.terminal-actions {
  display: flex;
  gap: var(--space-xs);
}

.action-btn {
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid var(--theme-border);
  border-radius: var(--radius-sm);
  color: var(--theme-fg);
  cursor: pointer;
  transition: all var(--transition-fast);
}

.action-btn:hover {
  background: rgba(255, 255, 255, 0.1);
  border-color: var(--theme-primary);
}

.terminal-output-area {
  flex: 1;
  overflow-y: auto;
  padding: var(--space-md);
  background: rgba(0, 0, 0, 0.3);
  border-radius: var(--radius-md);
  font-family: var(--font-mono);
  font-size: 0.875rem;
  line-height: 1.6;
}

.terminal-welcome {
  text-align: center;
  margin-bottom: var(--space-lg);
}

.ascii-art {
  color: var(--theme-primary);
  font-size: 0.75rem;
  line-height: 1.2;
  margin: 0 auto var(--space-md) auto;
  max-width: 400px;
}

.welcome-message {
  color: var(--theme-muted);
}

.welcome-line {
  margin: 0;
  font-size: 0.875rem;
}

.command-history {
  margin-bottom: var(--space-md);
}

.command-line {
  color: var(--theme-accent);
  margin-bottom: var(--space-xs);
}

.command-text {
  color: var(--theme-fg);
  font-weight: 500;
}

.command-output {
  color: var(--theme-muted);
  padding-left: var(--space-lg);
  white-space: pre-wrap;
  word-break: break-word;
}

.dynamic-content {
  margin-top: var(--space-lg);
}

.terminal-input-area {
  padding: var(--space-sm);
}

.input-wrapper {
  display: flex;
  align-items: center;
  gap: var(--space-xs);
  margin-bottom: var(--space-sm);
}

.input-prefix {
  display: flex;
  gap: var(--space-xs);
  color: var(--theme-accent);
  font-weight: 600;
  font-size: 0.875rem;
  white-space: nowrap;
}

.terminal-input {
  flex: 1;
  background: transparent;
  border: none;
  color: var(--theme-fg);
  font-family: var(--font-mono);
  font-size: 0.875rem;
  outline: none;
  padding: var(--space-xs) 0;
}

.terminal-input::placeholder {
  color: var(--theme-muted);
  opacity: 0.6;
}

.input-suffix {
  margin-left: auto;
}

.execute-btn {
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--theme-primary);
  border: none;
  border-radius: var(--radius-sm);
  color: white;
  cursor: pointer;
  transition: all var(--transition-fast);
}

.execute-btn:hover {
  background: color-mix(in srgb, var(--theme-primary) 90%, white);
  transform: translateX(2px);
}

.input-hints {
  display: flex;
  align-items: center;
  gap: var(--space-xs);
  flex-wrap: wrap;
}

.hint-label {
  font-size: 0.75rem;
  color: var(--theme-muted);
}

.hint-btn {
  padding: 2px 8px;
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid var(--theme-border);
  border-radius: var(--radius-sm);
  color: var(--theme-fg);
  font-size: 0.75rem;
  cursor: pointer;
  transition: all var(--transition-fast);
}

.hint-btn:hover {
  background: rgba(255, 255, 255, 0.1);
  border-color: var(--theme-primary);
}

.quick-actions {
  width: 200px;
  min-width: 200px;
  padding: var(--space-md);
  display: flex;
  flex-direction: column;
  gap: var(--space-md);
}

.actions-header {
  display: flex;
  align-items: center;
  gap: var(--space-xs);
  margin-bottom: var(--space-sm);
}

.actions-header h4 {
  margin: 0;
  font-size: 0.875rem;
  color: var(--theme-primary);
  letter-spacing: 0.05em;
}

.actions-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: var(--space-sm);
}

.quick-stats {
  display: flex;
  flex-direction: column;
  gap: var(--space-sm);
  margin-top: auto;
}

.stat-card {
  display: flex;
  align-items: center;
  gap: var(--space-sm);
  padding: var(--space-sm);
  background: rgba(255, 255, 255, 0.05);
  border-radius: var(--radius-sm);
}

.stat-content {
  display: flex;
  flex-direction: column;
}

.stat-value {
  font-size: 1.25rem;
  font-weight: 600;
  color: var(--theme-fg);
}

.stat-label {
  font-size: 0.75rem;
  color: var(--theme-muted);
}

/* Animations */
.slide-enter-active,
.slide-leave-active {
  transition: all var(--transition-base);
  overflow: hidden;
}

.slide-enter-from,
.slide-leave-to {
  max-width: 0;
  opacity: 0;
}

.slide-up-enter-active,
.slide-up-leave-active {
  transition: all var(--transition-base);
}

.slide-up-enter-from,
.slide-up-leave-to {
  transform: translateY(100%);
  opacity: 0;
}

/* Fullscreen mode */
.enhanced-layout.fullscreen {
  padding: 0;
}

.enhanced-layout.fullscreen .main-terminal {
  padding: 0;
}

.enhanced-layout.fullscreen .status-bar,
.enhanced-layout.fullscreen .terminal-sidebar,
.enhanced-layout.fullscreen .quick-actions {
  border-radius: 0;
  border: none;
}

/* Privacy mode */
.enhanced-layout.privacy-mode .status-bar,
.enhanced-layout.privacy-mode .sidebar-footer,
.enhanced-layout.privacy-mode .quick-stats {
  filter: blur(4px);
  opacity: 0.7;
}

/* Responsive */
@media (max-width: 1200px) {
  .quick-actions {
    display: none;
  }
}

@media (max-width: 992px) {
  .terminal-sidebar {
    position: fixed;
    left: 0;
    top: 0;
    bottom: 0;
    z-index: 1000;
    transform: translateX(-100%);
  }
  
  .enhanced-layout.sidebar-collapsed .terminal-sidebar {
    transform: translateX(0);
    width: var(--terminal-sidebar-width);
  }
}

@media (max-width: 768px) {
  .status-left,
  .status-right {
    display: none;
  }
  
  .status-center {
    justify-content: flex-start;
  }
  
  .terminal-prompt {
    font-size: 0.75rem;
  }
}
</style>

// END OF FILE