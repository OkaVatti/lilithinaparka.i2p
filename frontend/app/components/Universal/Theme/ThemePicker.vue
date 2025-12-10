<template>
  <div class="theme-picker">
    <button class="theme-button" @click="showPicker = !showPicker">
      <FeatherIcon name="palette" size="18" />
      <span>Theme</span>
    </button>
    
    <div v-if="showPicker" class="theme-dropdown">
      <div class="theme-options">
        <button
          v-for="theme in themes"
          :key="theme.value"
          class="theme-option"
          :class="{ active: currentTheme === theme.value }"
          @click="selectTheme(theme.value)"
        >
          <div class="theme-preview">
            <span class="preview-color" :style="{ backgroundColor: theme.preview.bg }"></span>
            <span class="preview-color" :style="{ backgroundColor: theme.preview.fg }"></span>
            <span class="preview-color" :style="{ backgroundColor: theme.preview.primary }"></span>
          </div>
          <span class="theme-name">{{ theme.name }}</span>
          <FeatherIcon v-if="currentTheme === theme.value" name="check" size="16" />
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue'
import { useTheme } from '~/composables/useTheme'

const { currentTheme, themes, setTheme, initTheme } = useTheme()
const showPicker = ref(false)

const selectTheme = (theme: string) => {
  setTheme(theme)
  showPicker.value = false
}

const handleClickOutside = (event: MouseEvent) => {
  const target = event.target as HTMLElement
  if (!target.closest('.theme-picker')) {
    showPicker.value = false
  }
}

onMounted(() => {
  initTheme()
  document.addEventListener('click', handleClickOutside)
})

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside)
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
  background: var(--theme-bg);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  color: var(--theme-fg);
  cursor: pointer;
  transition: all 0.2s;
}

.theme-button:hover {
  border-color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.1);
}

.theme-dropdown {
  position: absolute;
  top: calc(100% + 0.5rem);
  right: 0;
  background: var(--theme-bg);
  border: 2px solid var(--theme-border);
  border-radius: 4px;
  padding: 0.5rem;
  min-width: 200px;
  box-shadow: 0 4px 8px var(--theme-shadow);
  z-index: 1000;
}

.theme-options {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.theme-option {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.5rem;
  background: transparent;
  border: 1px solid transparent;
  border-radius: 4px;
  color: var(--theme-fg);
  cursor: pointer;
  transition: all 0.2s;
  text-align: left;
}

.theme-option:hover {
  border-color: var(--theme-border);
  background: rgba(255, 255, 255, 0.05);
}

.theme-option.active {
  border-color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.1);
}

.theme-preview {
  display: flex;
  gap: 2px;
}

.preview-color {
  width: 16px;
  height: 16px;
  border-radius: 2px;
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.theme-name {
  flex: 1;
}

@media (max-width: 768px) {
  .theme-button span {
    display: none;
  }
  
  .theme-dropdown {
    right: auto;
    left: 0;
  }
}
</style>