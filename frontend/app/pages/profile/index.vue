<template>
  <PageContainer>
    <ProfileContainer :user="user">
      <section class="posts">
        <h2>Recent posts</h2>

        <div class="post-grid">
          <BlogPostCard
            v-for="p in posts"
            :key="p.slug"
            :title="p.title"
            :excerpt="p.excerpt"
            :date="p.date"
            :tags="p.tags"
            :image="p.image"
            :to="`/blog/${p.slug}`"
          />
        </div>
      </section>
    </ProfileContainer>
  </PageContainer>
</template>

<script setup lang="ts">
import PageContainer from '~/components/layout/PageContainer.vue'
import ProfileContainer from '~/components/profile/ProfileContainer.vue'
import ProfileInfoCard from '~/components/profile/ProfileInfoCard.vue'
import BlogPostCard from '~/components/blog/BlogPostCard.vue'

const user = { name: 'Lilith in a Parka', avatar: '/images/avatar.jpg', bio: 'I make art and tinker with protocols.', stats: { posts: 26, followers: 900 } }

const posts = [
  { slug: 'making-a-moon', title: 'Making a Moon', excerpt: 'A short walk through my moon-painting process.', date: '2025-10-02', tags: ['art','process'], image: '/images/art/moon1.jpg' },
  { slug: 'privacy-first-css', title: 'Privacy-first CSS', excerpt: 'How I structure UX with minimal JS and strong CSP.', date: '2025-09-12', tags: ['web','privacy'], image: '/images/thumbs/csp.jpg' },
  { slug: 'glitch-experiments', title: 'Glitch Experiments', excerpt: 'Artifacts can be beautiful.', date: '2025-08-20', tags: ['art','glitch'], image: '/images/art/glitch.jpg' }
]
</script>

<style scoped>
@import '~/assets/css/system.css';

.posts { margin-top: .25rem; }
.post-grid { display:grid; grid-template-columns: repeat(auto-fill, minmax(280px,1fr)); gap:1rem; margin-top:.6rem; }
</style>
