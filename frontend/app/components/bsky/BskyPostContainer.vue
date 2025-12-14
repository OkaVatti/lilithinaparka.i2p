<template>
  <section class="bluesky-list">
    <header class="header-row">
      <h2>Bluesky Feed</h2>
      <p class="subtitle">A minimal preview of recent posts</p>
    </header>

    <div class="feed">
      <article v-for="post in posts" :key="post.id" class="bsky-card">
        <div class="bsky-head">
          <img class="avatar" :src="post.avatar" alt="avatar" v-if="post.avatar"/>
          <div>
            <div class="user">{{ post.user }}</div>
            <div class="meta"><small class="muted">{{ post.time }}</small></div>
          </div>
        </div>
        <div class="bsky-body">
          <p v-html="post.content"></p>
        </div>
      </article>
    </div>
  </section>
</template>

<script setup lang="ts">
const props = defineProps<{
  posts?: Array<{ id: string, user: string, avatar?: string, time?: string, content: string }>
}>()

const posts = props.posts ?? [
  { id: '1', user: 'alice', time: '2h', avatar: '', content: 'Experimenting with a new drawing workflow.' },
  { id: '2', user: 'bob', time: '6h', avatar: '', content: 'Privacy-first architectures rock.' }
]
</script>

<style scoped>
@import '~/assets/css/system.css';

.bluesky-list { display:flex; flex-direction:column; gap:1rem; }
.feed { display:flex; flex-direction:column; gap:.6rem; }

.bsky-card {
  border:1px solid var(--theme-border);
  padding:.7rem;
  border-radius:8px;
  background: linear-gradient(180deg, rgba(255,255,255,0.01), transparent);
}
.bsky-head { display:flex; gap:.6rem; align-items:center; margin-bottom:.4rem; }
.avatar { width:36px; height:36px; border-radius:50%; object-fit:cover; border:1px solid var(--theme-border); }
.user { font-weight:600; color:var(--theme-fg); }
.bsky-body p { margin:0; color:var(--theme-muted); }
</style>
