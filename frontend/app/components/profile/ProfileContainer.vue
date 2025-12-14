<template>
  <div class="profile-shell">
    <aside class="left">
      <ProfileInfoCard :name="user.name" :avatar="user.avatar" :bio="user.bio" :stats="user.stats" />
    </aside>

    <section class="right">
      <slot />
    </section>
  </div>
</template>

<script setup lang="ts">
import ProfileInfoCard from './ProfileInfoCard.vue'
const props = defineProps<{
  user?: { name:string, avatar?:string, bio?:string, stats?: Record<string,number> }
}>()

const user = props.user ?? { name: 'Lilith', avatar: '', bio: 'A privacy-first creative. Art, code, & bad puns.', stats: { posts: 12, followers: 420 } }
</script>

<style scoped>
@import '~/assets/css/system.css';

.profile-shell {
  display:grid;
  grid-template-columns: 280px 1fr;
  gap:1.2rem;
}
.left { position: relative; }
.right { min-height: 200px; }

/* responsive */
@media (max-width: 900px) {
  .profile-shell { grid-template-columns: 1fr; }
  .left { order: 2; }
  .right { order: 1; }
}
</style>
