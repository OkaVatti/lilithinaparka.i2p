<template>
  <div class="theme-button-wrapper" v-click-outside="closeStyleCard">
    <button @click="toggleStyleCard" class="theme-toggle-btn" title="Change Theme">
      <FeatherIcon name="palette" size="20" />
      <span class="theme-name">{{ currentThemeName }}</span>
      <FeatherIcon 
        name="chevron-down" 
        size="16" 
        :class="{ 'rotate-180': showStyleCard }" 
      />
    </button>
    
    <transition name="slide-fade">
      <StyleCard v-if="showStyleCard" @close="closeStyleCard" />
    </transition>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useTheme } from '~/composables/useTheme'
import StyleCard from './StyleCard.vue'

const { currentTheme, themes, initTheme } = useTheme()
const showStyleCard = ref(false)

const currentThemeName = computed(() => {
  const theme = themes.find(t => t.value === currentTheme.value)
  return theme?.name || 'Dracula'
})

const toggleStyleCard = () => {
  showStyleCard.value = !showStyleCard.value
}

const closeStyleCard = () => {
  showStyleCard.value = false
}

onMounted(() => {
  initTheme()
})
</script>

<style scoped>
.theme-button-wrapper {
  position: relative;
}

.theme-toggle-btn {
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

.theme-toggle-btn:hover {
  border-color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.1);
}

.theme-name {
  font-size: 0.9rem;
}

.theme-toggle-btn svg:last-child {
  transition: transform 0.2s;
}

.rotate-180 {
  transform: rotate(180deg);
}

.slide-fade-enter-active,
.slide-fade-leave-active {
  transition: all 0.2s ease;
}

.slide-fade-enter-from,
.slide-fade-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

@media (max-width: 768px) {
  .theme-name {
    display: none;
  }
}
</style>