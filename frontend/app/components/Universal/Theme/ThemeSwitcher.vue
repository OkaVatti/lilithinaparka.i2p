<template>
  <div class="theme-switcher">
    <div class="switcher-header glass-panel">
      <div class="header-content">
        <FeatherIcon name="palette" size="20" class="header-icon" />
        <h3 class="header-title">Visual Design</h3>
        <div class="header-actions">
          <button @click="toggleAdvanced" class="action-btn" title="Advanced settings">
            <FeatherIcon :name="showAdvanced ? 'chevron-up' : 'chevron-down'" size="16" />
          </button>
          <button @click="randomTheme" class="action-btn" title="Random theme">
            <FeatherIcon name="shuffle" size="16" />
          </button>
        </div>
      </div>
    </div>
    
    <transition name="slide-down">
      <div v-if="showAdvanced" class="advanced-settings glass-panel">
        <div class="settings-section">
          <h4 class="section-title">Color Scheme</h4>
          <div class="scheme-toggle">
            <button 
              @click="setColorScheme('light')"
              class="scheme-btn"
              :class="{ active: currentColorScheme === 'light' }"
              title="Light mode"
            >
              <FeatherIcon name="sun" size="18" />
              <span>Light</span>
            </button>
            <button 
              @click="setColorScheme('dark')"
              class="scheme-btn"
              :class="{ active: currentColorScheme === 'dark' }"
              title="Dark mode"
            >
              <FeatherIcon name="moon" size="18" />
              <span>Dark</span>
            </button>
            <button 
              @click="setColorScheme('auto')"
              class="scheme-btn"
              :class="{ active: currentColorScheme === 'auto' }"
              title="Auto (follow system)"
            >
              <FeatherIcon name="monitor" size="18" />
              <span>Auto</span>
            </button>
          </div>
        </div>
        
        <div class="settings-section">
          <h4 class="section-title">UI Density</h4>
          <div class="density-slider">
            <label class="density-label">Compact</label>
            <input 
              type="range" 
              min="0" 
              max="2" 
              v-model="density"
              class="density-range"
              @change="applyDensity"
            />
            <label class="density-label">Spacious</label>
          </div>
        </div>
        
        <div class="settings-section">
          <h4 class="section-title">Animation Level</h4>
          <div class="animation-toggle">
            <button 
              v-for="level in animationLevels"
              :key="level.value"
              @click="setAnimationLevel(level.value)"
              class="animation-btn"
              :class="{ active: animationLevel === level.value }"
              :title="level.description"
            >
              <FeatherIcon :name="level.icon" size="16" />
              <span>{{ level.label }}</span>
            </button>
          </div>
        </div>
      </div>
    </transition>
    
    <div class="theme-grid-container glass-panel">
      <div class="theme-categories">
        <button
          v-for="category in themeCategories"
          :key="category.id"
          @click="activeCategory = category.id"
          class="category-btn"
          :class="{ active: activeCategory === category.id }"
        >
          <FeatherIcon :name="category.icon" size="16" />
          <span>{{ category.label }}</span>
        </button>
      </div>
      
      <div class="theme-grid">
        <div
          v-for="theme in filteredThemes"
          :key="theme.value"
          @click="applyTheme(theme)"
          class="theme-card enhanced"
          :class="{ 
            active: currentTheme === theme.value,
            recommended: theme.recommended 
          }"
          :style="getThemeCardStyle(theme)"
        >
          <div class="theme-card-header">
            <div class="theme-preview">
              <div 
                v-for="(color, index) in getPreviewColors(theme)"
                :key="index"
                class="theme-preview-color"
                :style="{ backgroundColor: color }"
              />
            </div>
            <div class="theme-badges">
              <span v-if="theme.recommended" class="badge badge-recommended">
                <FeatherIcon name="star" size="12" />
              </span>
              <span v-if="theme.seasonal" class="badge badge-seasonal">
                <FeatherIcon :name="getSeasonIcon(theme.value)" size="12" />
              </span>
            </div>
          </div>
          
          <div class="theme-card-body">
            <h4 class="theme-name">{{ theme.name }}</h4>
            <p v-if="theme.description" class="theme-description">
              {{ theme.description }}
            </p>
            
            <div class="theme-stats">
              <div class="stat">
                <FeatherIcon name="contrast" size="12" />
                <span class="stat-label">Contrast: {{ theme.contrast || 'High' }}</span>
              </div>
              <div v-if="theme.accessibility" class="stat">
                <FeatherIcon name="eye" size="12" />
                <span class="stat-label">Accessible</span>
              </div>
            </div>
          </div>
          
          <div class="theme-card-footer">
            <button 
              v-if="currentTheme === theme.value"
              class="btn-active"
              disabled
            >
              <FeatherIcon name="check" size="16" />
              <span>Active</span>
            </button>
            <button 
              v-else
              class="btn-apply"
              @click.stop="applyTheme(theme)"
            >
              <FeatherIcon name="play" size="16" />
              <span>Apply</span>
            </button>
            
            <button 
              class="btn-preview"
              @click.stop="previewTheme(theme)"
              @mouseenter="previewTheme(theme)"
              @mouseleave="restoreTheme"
              title="Preview theme"
            >
              <FeatherIcon name="eye" size="14" />
            </button>
          </div>
          
          <div class="theme-card-glow"></div>
        </div>
      </div>
    </div>
    
    <div class="theme-presets glass-panel">
      <h4 class="presets-title">Quick Presets</h4>
      <div class="presets-grid">
        <button
          v-for="preset in quickPresets"
          :key="preset.id"
          @click="applyPreset(preset)"
          class="preset-btn"
          :style="{ 
            background: preset.gradient,
            color: preset.textColor 
          }"
        >
          <FeatherIcon :name="preset.icon" size="18" />
          <span class="preset-label">{{ preset.label }}</span>
        </button>
      </div>
    </div>
    
    <transition name="fade">
      <div v-if="showThemePreview" class="theme-preview-overlay">
        <div class="preview-content">
          <div class="preview-header">
            <h3>Preview Mode</h3>
            <p>Hover over themes to preview</p>
          </div>
          <button @click="restoreTheme" class="btn-restore">
            <FeatherIcon name="x" size="16" />
            <span>Restore Original</span>
          </button>
        </div>
      </div>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useTheme } from '../../../composables/useTheme'

const { currentTheme, currentColorScheme, themes, seasonalThemes, setTheme, setColorScheme, initTheme } = useTheme()

const showAdvanced = ref(false)
const activeCategory = ref('all')
const density = ref(1)
const animationLevel = ref('reduced')
const showThemePreview = ref(false)
let previewTimeout: NodeJS.Timeout | null = null

const themeCategories = [
  { id: 'all', label: 'All Themes', icon: 'grid' },
  { id: 'modern', label: 'Modern', icon: 'star' },
  { id: 'classic', label: 'Classic', icon: 'book' },
  { id: 'seasonal', label: 'Seasonal', icon: 'cloud' },
  { id: 'accessible', label: 'Accessible', icon: 'eye' },
  { id: 'experimental', label: 'Experimental', icon: 'zap' }
]

const animationLevels = [
  { value: 'none', label: 'Off', icon: 'minus', description: 'No animations' },
  { value: 'reduced', label: 'Reduced', icon: 'activity', description: 'Minimal animations' },
  { value: 'normal', label: 'Normal', icon: 'play', description: 'Standard animations' },
  { value: 'enhanced', label: 'Enhanced', icon: 'zap', description: 'All animations' }
]

const quickPresets = [
  { id: 'terminal', label: 'Terminal', icon: 'terminal', 
    theme: 'terminal', scheme: 'dark', gradient: 'linear-gradient(135deg, #0a0a0a 0%, #1a1a1a 100%)', textColor: '#00ff00' },
  { id: 'dracula', label: 'Dracula', icon: 'moon', 
    theme: 'dracula', scheme: 'dark', gradient: 'linear-gradient(135deg, #282a36 0%, #44475a 100%)', textColor: '#f8f8f2' },
  { id: 'nord', label: 'Nord', icon: 'snowflake', 
    theme: 'nord', scheme: 'dark', gradient: 'linear-gradient(135deg, #2e3440 0%, #4c566a 100%)', textColor: '#eceff4' },
  { id: 'light', label: 'Light', icon: 'sun', 
    theme: 'solarized', scheme: 'light', gradient: 'linear-gradient(135deg, #fdf6e3 0%, #eee8d5 100%)', textColor: '#657b83' },
  { id: 'retro', label: 'Retro', icon: 'monitor', 
    theme: 'gruvbox', scheme: 'dark', gradient: 'linear-gradient(135deg, #282828 0%, #3c3836 100%)', textColor: '#ebdbb2' }
]

// Enhanced themes with more metadata
const enhancedThemes = computed(() => {
  return themes.map(theme => ({
    ...theme,
    description: getThemeDescription(theme.value),
    contrast: getThemeContrast(theme.value),
    accessibility: isAccessibleTheme(theme.value),
    recommended: ['dracula', 'nord', 'gruvbox'].includes(theme.value),
    seasonal: ['winter', 'spring', 'summer', 'autumn'].includes(theme.value),
    category: getThemeCategory(theme.value)
  }))
})

const filteredThemes = computed(() => {
  if (activeCategory.value === 'all') return enhancedThemes.value
  
  return enhancedThemes.value.filter(theme => {
    if (activeCategory.value === 'seasonal') return theme.seasonal
    if (activeCategory.value === 'accessible') return theme.accessibility
    if (activeCategory.value === 'modern') return theme.recommended
    if (activeCategory.value === 'classic') return !theme.seasonal && !theme.recommended
    if (activeCategory.value === 'experimental') return theme.value.includes('experimental')
    return true
  })
})

const getThemeDescription = (themeValue: string) => {
  const descriptions: Record<string, string> = {
    dracula: 'Dark theme with high contrast and vibrant colors',
    nord: 'Arctic, north-bluish color palette',
    gruvbox: 'Retro groove color scheme',
    tokyo: 'Clean, dark theme with vivid colors',
    solarized: 'Carefully designed for solarized fans',
    terminal: 'Classic green-on-black terminal look',
    winter: 'Cool blue tones for winter',
    spring: 'Fresh greens for spring',
    summer: 'Warm golden tones for summer',
    autumn: 'Rich oranges and browns for autumn'
  }
  return descriptions[themeValue] || 'A beautiful color theme'
}

const getThemeContrast = (themeValue: string) => {
  const highContrast = ['dracula', 'terminal', 'monokai']
  const mediumContrast = ['nord', 'gruvbox', 'tokyo']
  return highContrast.includes(themeValue) ? 'High' : mediumContrast.includes(themeValue) ? 'Medium' : 'Standard'
}

const isAccessibleTheme = (themeValue: string) => {
  const accessibleThemes = ['dracula', 'nord', 'solarized']
  return accessibleThemes.includes(themeValue)
}

const getThemeCategory = (themeValue: string) => {
  if (['winter', 'spring', 'summer', 'autumn'].includes(themeValue)) return 'seasonal'
  if (['dracula', 'nord', 'gruvbox'].includes(themeValue)) return 'modern'
  if (['solarized', 'monokai', 'terminal'].includes(themeValue)) return 'classic'
  return 'modern'
}

const getSeasonIcon = (themeValue: string) => {
  const icons: Record<string, string> = {
    winter: 'cloud-snow',
    spring: 'cloud-rain',
    summer: 'sun',
    autumn: 'leaf'
  }
  return icons[themeValue] || 'circle'
}

const getThemeCardStyle = (theme: any) => {
  return {
    '--card-bg': theme.preview.bg,
    '--card-fg': theme.preview.fg,
    '--card-primary': theme.preview.primary
  }
}

const getPreviewColors = (theme: any) => {
  return [
    theme.preview.bg,
    theme.preview.fg,
    theme.preview.primary,
    adjustColor(theme.preview.bg, 20),
    adjustColor(theme.preview.primary, -20)
  ]
}

const adjustColor = (hex: string, percent: number) => {
  const num = parseInt(hex.replace('#', ''), 16)
  const amt = Math.round(2.55 * percent)
  const R = Math.min(255, Math.max(0, (num >> 16) + amt))
  const G = Math.min(255, Math.max(0, (num >> 8 & 0x00FF) + amt))
  const B = Math.min(255, Math.max(0, (num & 0x0000FF) + amt))
  return '#' + (0x1000000 + R * 0x10000 + G * 0x100 + B).toString(16).slice(1)
}

const applyTheme = (theme: any) => {
  setTheme(theme.value)
  localStorage.setItem('lastUsedTheme', theme.value)
  
  // Apply density and animation settings
  applyDensity()
  setAnimationLevel(animationLevel.value)
}

const previewTheme = (theme: any) => {
  if (previewTimeout) clearTimeout(previewTimeout)
  setTheme(theme.value)
  showThemePreview.value = true
  
  previewTimeout = setTimeout(() => {
    showThemePreview.value = false
  }, 3000)
}

const restoreTheme = () => {
  const savedTheme = localStorage.getItem('theme') || 'dracula'
  setTheme(savedTheme)
  showThemePreview.value = false
  if (previewTimeout) clearTimeout(previewTimeout)
}

const toggleAdvanced = () => {
  showAdvanced.value = !showAdvanced.value
}

const randomTheme = () => {
  const randomIndex = Math.floor(Math.random() * enhancedThemes.value.length)
  applyTheme(enhancedThemes.value[randomIndex])
}

const applyDensity = () => {
  const densities = ['compact', 'normal', 'spacious']
  document.documentElement.setAttribute('data-density', densities[density.value])
  localStorage.setItem('uiDensity', density.value.toString())
}

const setAnimationLevel = (level: string) => {
  animationLevel.value = level
  document.documentElement.setAttribute('data-animations', level)
  localStorage.setItem('animationLevel', level)
}

const applyPreset = (preset: any) => {
  setTheme(preset.theme)
  setColorScheme(preset.scheme)
}

onMounted(() => {
  initTheme()
  
  // Load saved settings
  const savedDensity = localStorage.getItem('uiDensity')
  if (savedDensity) {
    density.value = parseInt(savedDensity)
    applyDensity()
  }
  
  const savedAnimationLevel = localStorage.getItem('animationLevel')
  if (savedAnimationLevel) {
    animationLevel.value = savedAnimationLevel
    setAnimationLevel(savedAnimationLevel)
  }
  
  // Listen for system color scheme changes
  const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)')
  const handleColorSchemeChange = (e: MediaQueryListEvent) => {
    if (currentColorScheme.value === 'auto') {
      setColorScheme(e.matches ? 'dark' : 'light')
    }
  }
  
  mediaQuery.addEventListener('change', handleColorSchemeChange)
  
  onUnmounted(() => {
    mediaQuery.removeEventListener('change', handleColorSchemeChange)
  })
})
</script>

<style scoped>
@import '~/assets/css/design-system.css';

.theme-switcher {
  display: flex;
  flex-direction: column;
  gap: var(--space-md);
  max-width: 800px;
  margin: 0 auto;
}

.switcher-header {
  padding: var(--space-md);
}

.header-content {
  display: flex;
  align-items: center;
  gap: var(--space-sm);
}

.header-icon {
  color: var(--theme-primary);
}

.header-title {
  margin: 0;
  font-size: 1.25rem;
  font-weight: 600;
  color: var(--theme-fg);
  flex: 1;
}

.header-actions {
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
  transform: translateY(-1px);
}

.advanced-settings {
  padding: var(--space-md);
  display: flex;
  flex-direction: column;
  gap: var(--space-lg);
}

.settings-section {
  display: flex;
  flex-direction: column;
  gap: var(--space-sm);
}

.section-title {
  margin: 0;
  font-size: 0.875rem;
  font-weight: 600;
  color: var(--theme-muted);
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.scheme-toggle,
.animation-toggle {
  display: flex;
  gap: var(--space-xs);
}

.scheme-btn,
.animation-btn {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: var(--space-xs);
  padding: var(--space-sm) var(--space-md);
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid var(--theme-border);
  border-radius: var(--radius-sm);
  color: var(--theme-fg);
  font-size: 0.875rem;
  cursor: pointer;
  transition: all var(--transition-fast);
}

.scheme-btn:hover,
.animation-btn:hover {
  background: rgba(255, 255, 255, 0.1);
  border-color: var(--theme-border);
}

.scheme-btn.active,
.animation-btn.active {
  background: var(--theme-primary);
  border-color: var(--theme-primary);
  color: white;
  font-weight: 500;
}

.density-slider {
  display: flex;
  align-items: center;
  gap: var(--space-md);
}

.density-label {
  font-size: 0.875rem;
  color: var(--theme-muted);
  min-width: 60px;
}

.density-range {
  flex: 1;
  height: 4px;
  background: rgba(255, 255, 255, 0.1);
  border-radius: var(--radius-pill);
  outline: none;
  appearance: none;
}

.density-range::-webkit-slider-thumb {
  -webkit-appearance: none;
  width: 16px;
  height: 16px;
  background: var(--theme-primary);
  border-radius: 50%;
  cursor: pointer;
  transition: all var(--transition-fast);
}

.density-range::-webkit-slider-thumb:hover {
  transform: scale(1.2);
}

.theme-grid-container {
  padding: var(--space-md);
}

.theme-categories {
  display: flex;
  gap: var(--space-xs);
  margin-bottom: var(--space-lg);
  padding-bottom: var(--space-md);
  border-bottom: 1px solid var(--theme-border);
  overflow-x: auto;
}

.category-btn {
  flex-shrink: 0;
  display: flex;
  align-items: center;
  gap: var(--space-xs);
  padding: var(--space-sm) var(--space-md);
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid var(--theme-border);
  border-radius: var(--radius-pill);
  color: var(--theme-fg);
  font-size: 0.875rem;
  cursor: pointer;
  transition: all var(--transition-fast);
  white-space: nowrap;
}

.category-btn:hover {
  background: rgba(255, 255, 255, 0.1);
  border-color: var(--theme-border);
}

.category-btn.active {
  background: var(--theme-primary);
  border-color: var(--theme-primary);
  color: white;
}

.theme-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
  gap: var(--space-md);
}

.theme-card.enhanced {
  position: relative;
  padding: var(--space-md);
  background: var(--card-bg);
  border: 2px solid var(--card-fg);
  border-radius: var(--radius-md);
  color: var(--card-fg);
  cursor: pointer;
  transition: all var(--transition-base);
  overflow: hidden;
}

.theme-card.enhanced:hover {
  transform: translateY(-4px) scale(1.02);
  border-color: var(--card-primary);
  box-shadow: var(--shadow-lg);
}

.theme-card.enhanced.active {
  border-color: var(--card-primary);
  box-shadow: 
    0 0 0 2px var(--card-primary),
    0 8px 32px rgba(0, 0, 0, 0.3);
}

.theme-card.enhanced.recommended::before {
  content: '★ Recommended';
  position: absolute;
  top: 8px;
  right: -30px;
  background: var(--card-primary);
  color: white;
  padding: 2px 30px;
  font-size: 0.75rem;
  font-weight: 600;
  transform: rotate(45deg);
  transform-origin: center;
}

.theme-card-header {
  margin-bottom: var(--space-md);
}

.theme-preview {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 2px;
  margin-bottom: var(--space-xs);
}

.theme-preview-color {
  height: 16px;
  border-radius: var(--radius-sm);
}

.theme-badges {
  display: flex;
  gap: var(--space-xs);
  justify-content: flex-end;
}

.badge {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  font-size: 0.75rem;
}

.badge-recommended {
  background: gold;
  color: black;
}

.badge-seasonal {
  background: var(--card-primary);
  color: white;
}

.theme-card-body {
  margin-bottom: var(--space-md);
}

.theme-name {
  margin: 0 0 var(--space-xs) 0;
  font-size: 1rem;
  font-weight: 600;
  color: var(--card-fg);
}

.theme-description {
  margin: 0 0 var(--space-sm) 0;
  font-size: 0.75rem;
  color: color-mix(in srgb, var(--card-fg) 80%, transparent);
  line-height: 1.4;
}

.theme-stats {
  display: flex;
  gap: var(--space-sm);
  font-size: 0.75rem;
}

.stat {
  display: flex;
  align-items: center;
  gap: 2px;
  color: color-mix(in srgb, var(--card-fg) 70%, transparent);
}

.stat-label {
  font-size: 0.7rem;
}

.theme-card-footer {
  display: flex;
  gap: var(--space-xs);
}

.btn-active,
.btn-apply,
.btn-preview {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: var(--space-xs);
  padding: var(--space-xs) var(--space-sm);
  border: none;
  border-radius: var(--radius-sm);
  font-size: 0.75rem;
  font-weight: 500;
  cursor: pointer;
  transition: all var(--transition-fast);
}

.btn-active {
  background: var(--card-primary);
  color: white;
  cursor: default;
}

.btn-apply {
  background: rgba(255, 255, 255, 0.1);
  color: var(--card-fg);
}

.btn-apply:hover {
  background: var(--card-primary);
  color: white;
}

.btn-preview {
  width: 32px;
  background: rgba(255, 255, 255, 0.05);
  color: var(--card-fg);
}

.btn-preview:hover {
  background: rgba(255, 255, 255, 0.1);
}

.theme-card-glow {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: radial-gradient(
    circle at var(--mouse-x, 50%) var(--mouse-y, 50%),
    rgba(255, 255, 255, 0.1) 0%,
    transparent 50%
  );
  opacity: 0;
  pointer-events: none;
  transition: opacity var(--transition-fast);
}

.theme-card.enhanced:hover .theme-card-glow {
  opacity: 1;
}

.theme-presets {
  padding: var(--space-md);
}

.presets-title {
  margin: 0 0 var(--space-md) 0;
  font-size: 0.875rem;
  font-weight: 600;
  color: var(--theme-muted);
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.presets-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(140px, 1fr));
  gap: var(--space-sm);
}

.preset-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: var(--space-xs);
  padding: var(--space-sm);
  border: none;
  border-radius: var(--radius-md);
  font-size: 0.875rem;
  font-weight: 500;
  cursor: pointer;
  transition: all var(--transition-base);
}

.preset-btn:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-md);
}

.preset-label {
  font-weight: 600;
}

.theme-preview-overlay {
  position: fixed;
  bottom: var(--space-lg);
  left: 50%;
  transform: translateX(-50%);
  background: var(--glass-bg);
  backdrop-filter: var(--backdrop-blur);
  border: 1px solid var(--glass-border);
  border-radius: var(--radius-lg);
  padding: var(--space-md);
  box-shadow: var(--glass-shadow);
  z-index: 1000;
  animation: float 2s ease-in-out infinite;
}

.preview-header {
  margin-bottom: var(--space-sm);
  text-align: center;
}

.preview-header h3 {
  margin: 0 0 var(--space-xs) 0;
  color: var(--theme-primary);
}

.preview-header p {
  margin: 0;
  font-size: 0.875rem;
  color: var(--theme-muted);
}

.btn-restore {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: var(--space-xs);
  width: 100%;
  padding: var(--space-sm) var(--space-md);
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid var(--theme-border);
  border-radius: var(--radius-md);
  color: var(--theme-fg);
  font-size: 0.875rem;
  font-weight: 500;
  cursor: pointer;
  transition: all var(--transition-fast);
}

.btn-restore:hover {
  background: var(--theme-primary);
  border-color: var(--theme-primary);
  color: white;
}

/* Animations */
.slide-down-enter-active,
.slide-down-leave-active {
  transition: all var(--transition-base);
  max-height: 500px;
  overflow: hidden;
}

.slide-down-enter-from,
.slide-down-leave-to {
  max-height: 0;
  opacity: 0;
  transform: translateY(-10px);
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity var(--transition-fast);
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

/* Responsive */
@media (max-width: 768px) {
  .theme-grid {
    grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
  }
  
  .presets-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 480px) {
  .theme-grid {
    grid-template-columns: 1fr;
  }
  
  .theme-categories {
    flex-wrap: wrap;
  }
  
  .category-btn {
    flex: 1;
    min-width: 100px;
  }
}
</style>