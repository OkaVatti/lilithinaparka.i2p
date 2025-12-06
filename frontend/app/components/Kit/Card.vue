<!-- app/components/Branding/Card.vue -->
<template>
  <div class="terminal-box card-component">
    <slot />
  </div>
</template>

<style scoped>
.card-component {
  transition: all 0.2s ease;
}

.card-component:hover {
  border-color: var(--accent);
  transform: translateY(-2px);
}
</style>