<!-- app/pages/profile.vue -->
<template>
  <Profile :profile="profileData" />
</template>

<script setup lang="ts">
import Profile from '~/components/Profile/Profile.vue'
import type { Profile as ProfileType } from '../../types'

definePageMeta({
  layout: 'profile'
})

const profileData: ProfileType = {
  id: 1,
  created_at: '2023-01-01T00:00:00Z',
  updated_at: '2024-01-01T00:00:00Z',
  pic: "/images/profile-pic-1.png",
  name: "Lilith Parker",
  username: "@lilithinaparka",
  cake_day: "2004-01-01",
  bio: "chronically online cat girl who likes programming, music, and burritos",
  interests: "programming, music, burritos, art, cryptography, privacy, open source",
  location: "Earth, Sol System, Milky Way Galaxy",
  timezone: "UTC-5",
  website: "https://lilithinaparka.i2p",
  email: "lilithinaparka@mail.i2p",
  github: "https://github.com/okavatti",
  gitten: "lilithinaparka.gitten.i2p/profile",
  bluesky: "@lilithinaparka.bsky.social",
  bottletail: "#lilithinaparka.bttl.dev",
  rss_feed: "https://lilithinaparka.i2p/rss.xml",
  bitcoin: "bc1q89rsfakvfm5l9px3l9gwcs7tvra6ctxwpeerd4",
  ethereum: "0x060cc45f815f7342880d9c267b216f0d9cd7b110",
  solana: "9ALUmLHVcTQU7zQ4oKYB5Yib7cpccbXheAKjja48LkBF",
  monero: "45qivL5DtkdaSWSU1kfFhhjL9G4LhWXr4Zxw5a4VtnoQAi3LaKbkD6Dc6xy92Zc4aLDTikA3rhCW9Gf1J5jhMEXUPpsa2vu",
  rss: "https://lilithinaparka.i2p/rss.xml",
  bsky_display_name: "Lilith Parker",
  bsky_description: "Cat girl programmer and artist",
  bsky_avatar: "/images/profile-pic-1.png",
  bsky_banner: "/images/banner.jpg",
  bsky_followers_count: 1500,
  bsky_follows_count: 450,
  bsky_posts_count: 235
}
</script>