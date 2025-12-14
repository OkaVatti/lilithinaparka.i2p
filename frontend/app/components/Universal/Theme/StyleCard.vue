<template>
  <div class="style-card">
    <div class="style-card-header">
      <FeatherIcon name="palette" size="18" />
      <span>Appearance</span>
      <button @click="$emit('close')" class="close-btn">
        <FeatherIcon name="x" size="18" />
      </button>
    </div>
    
    <div class="style-card-body">
      <div class="section">
        <h4>Color Schemes</h4>
        <div class="theme-grid">
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
            <span class="theme-label">{{ theme.name }}</span>
            <FeatherIcon v-if="currentTheme === theme.value" name="check" size="14" class="check-icon" />
          </button>
        </div>
      </div>
      
      <div class="section">
        <h4>Seasonal Themes</h4>
        <div class="seasonal-grid">
          <button
            v-for="season in seasonalThemes"
            :key="season.value"
            @click="selectTheme(season.value)"
            class="seasonal-option"
            :class="{ active: currentTheme === season.value }"
          >
            <FeatherIcon :name="getSeasonIcon(season.value)" size="20" />
            <span>{{ season.name }}</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useTheme } from '~/composables/useTheme'

const emit = defineEmits(['close'])

const { currentTheme, themes, seasonalThemes, setTheme } = useTheme()

const selectTheme = (themeValue: string) => {
  setTheme(themeValue)
  setTimeout(() => emit('close'), 200)
}

const getSeasonIcon = (season: string) => {
  const icons: Record<string, string> = {
    winter: 'cloud-snow',
    spring: 'cloud-rain',
    summer: 'sun',
    autumn: 'wind'
  }
  return icons[season] || 'circle'
}
</script>

<style scoped>
.style-card {
  position: absolute;
  top: calc(100% + 0.5rem);
  right: 0;
  width: 320px;
  max-height: 500px;
  overflow-y: auto;
  background: var(--theme-bg);
  border: 2px solid var(--theme-border);
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.5);
  z-index: 1000;
}

.style-card-header {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 1rem;
  background: rgba(189, 147, 249, 0.1);
  border-bottom: 1px solid var(--theme-border);
  color: var(--theme-primary);
  font-weight: bold;
}

.close-btn {
  margin-left: auto;
  padding: 0.25rem;
  background: transparent;
  border: none;
  color: var(--theme-fg);
  cursor: pointer;
  border-radius: 4px;
  transition: all 0.2s;
}

.close-btn:hover {
  background: rgba(255, 255, 255, 0.1);
}

.style-card-body {
  padding: 1rem;
}

.section {
  margin-bottom: 1.5rem;
}

.section:last-child {
  margin-bottom: 0;
}

.section h4 {
  margin: 0 0 0.75rem 0;
  font-size: 0.85rem;
  color: var(--theme-fg);
  opacity: 0.7;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.theme-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 0.5rem;
}

.theme-option {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  width: 100%;
  padding: 0.75rem;
  background: transparent;
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  color: var(--theme-fg);
  cursor: pointer;
  transition: all 0.2s;
  text-align: left;
  font-family: inherit;
}

.theme-option:hover {
  background: rgba(189, 147, 249, 0.1);
  border-color: var(--theme-primary);
}

.theme-option.active {
  background: rgba(189, 147, 249, 0.2);
  border-color: var(--theme-primary);
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

.theme-label {
  flex: 1;
  font-size: 0.9rem;
}

.check-icon {
  color: var(--theme-primary);
}

.seasonal-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 0.5rem;
}

.seasonal-option {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.5rem;
  padding: 1rem;
  background: transparent;
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  color: var(--theme-fg);
  cursor: pointer;
  transition: all 0.2s;
  font-family: inherit;
}

.seasonal-option:hover {
  background: rgba(189, 147, 249, 0.1);
  border-color: var(--theme-primary);
}

.seasonal-option.active {
  background: rgba(189, 147, 249, 0.2);
  border-color: var(--theme-primary);
  color: var(--theme-primary);
}

@media (max-width: 768px) {
  .style-card {
    width: 280px;
    right: auto;
    left: 0;
  }
}
</style>