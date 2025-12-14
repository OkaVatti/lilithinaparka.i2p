<template>
  <div class="theme-picker" v-click-outside="closeDropdown">
    <button @click="toggleDropdown" class="theme-button" :title="'Current theme: ' + currentThemeName">
      <FeatherIcon name="palette" size="20" />
      <span class="theme-label">{{ currentThemeName }}</span>
      <FeatherIcon name="chevron-down" size="16" :class="{ 'rotate-180': showDropdown }" />
    </button>
    
    <transition name="dropdown">
      <div v-if="showDropdown" class="theme-dropdown">
        <div class="dropdown-header">
          <FeatherIcon name="palette" size="18" />
          <span>Choose Theme</span>
        </div>
        
        <div class="theme-list">
          <button
            v-for="theme in themes"
            :key="theme.value"
            @click="selectTheme(theme.value)"
            class="theme-option"
            :class="{ active: currentTheme === theme.value }"
          >
            <div class="theme-preview">
              <span class="preview-bg" :style="{ backgroundColor: theme.preview.bg }"></span>
              <span class="preview-fg" :style="{ backgroundColor: theme.preview.fg }"></span>
              <span class="preview-primary" :style="{ backgroundColor: theme.preview.primary }"></span>
            </div>
            <span class="theme-name">{{ theme.name }}</span>
            <FeatherIcon v-if="currentTheme === theme.value" name="check" size="16" class="check-icon" />
          </button>
        </div>
      </div>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useTheme } from '~/composables/useTheme'

const { currentTheme, themes, setTheme, initTheme } = useTheme()
const showDropdown = ref(false)

const currentThemeName = computed(() => {
  const theme = themes.find(t => t.value === currentTheme.value)
  return theme?.name || 'Dracula'
})

const toggleDropdown = () => {
  showDropdown.value = !showDropdown.value
}

const closeDropdown = () => {
  showDropdown.value = false
}

const selectTheme = (themeValue: string) => {
  setTheme(themeValue)
  closeDropdown()
}

onMounted(() => {
  initTheme()
})
</script>

<style scoped>
.theme-picker {
  position: relative;
}

.theme-button {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 1rem;
  background: rgba(0, 0, 0, 0.2);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  color: var(--theme-fg);
  cursor: pointer;
  transition: all 0.2s;
  font-family: inherit;
}

.theme-button:hover {
  border-color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.1);
}

.theme-label {
  font-size: 0.9rem;
}

.theme-button svg:last-child {
  transition: transform 0.2s;
}

.rotate-180 {
  transform: rotate(180deg);
}

.theme-dropdown {
  position: absolute;
  top: calc(100% + 0.5rem);
  right: 0;
  width: 280px;
  background: var(--theme-bg);
  border: 2px solid var(--theme-border);
  border-radius: 8px;
  box-shadow: 0 4px 12px var(--theme-shadow);
  z-index: 1000;
  overflow: hidden;
}

.dropdown-header {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 1rem;
  background: rgba(189, 147, 249, 0.1);
  border-bottom: 1px solid var(--theme-border);
  color: var(--theme-primary);
  font-weight: bold;
}

.theme-list {
  max-height: 400px;
  overflow-y: auto;
}

.theme-option {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  width: 100%;
  padding: 0.75rem 1rem;
  background: transparent;
  border: none;
  border-bottom: 1px solid var(--theme-border);
  color: var(--theme-fg);
  cursor: pointer;
  transition: all 0.2s;
  text-align: left;
  font-family: inherit;
}

.theme-option:last-child {
  border-bottom: none;
}

.theme-option:hover {
  background: rgba(189, 147, 249, 0.1);
}

.theme-option.active {
  background: rgba(189, 147, 249, 0.2);
  color: var(--theme-primary);
}

.theme-preview {
  display: flex;
  gap: 2px;
  min-width: 60px;
}

.theme-preview span {
  width: 20px;
  height: 20px;
  border-radius: 3px;
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.theme-name {
  flex: 1;
  font-size: 0.9rem;
}

.check-icon {
  color: var(--theme-primary);
}

.dropdown-enter-active,
.dropdown-leave-active {
  transition: all 0.2s ease;
}

.dropdown-enter-from,
.dropdown-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

@media (max-width: 768px) {
  .theme-label {
    display: none;
  }
  
  .theme-dropdown {
    right: auto;
    left: 0;
  }
}
</style>