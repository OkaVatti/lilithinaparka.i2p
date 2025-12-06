// app/types/index.ts
export interface BlogPost {
  id: number
  created_at: string
  updated_at: string
  title: string
  date: string
  time: string
  authors: string
  tags: string
  categories: string
  draft: boolean
  share: boolean
  slug: string
  layout: string
  toc: boolean
  comments: boolean
  math: boolean
  featured_image: string
  featured_image_alt: string
  featured_video: string
  featured_video_alt: string
  summary: string
  content: string
  file_path: string
}

export interface BskyPost {
  id: number
  created_at: string
  updated_at: string
  uri: string
  cid: string
  author: string
  text: string
  posted_at: string
  reply_count: number
  repost_count: number
  like_count: number
  quote_count: number
  has_media: boolean
  media_urls: string
}

export interface Profile {
  id: number
  created_at: string
  updated_at: string
  pic: string
  name: string
  username: string
  cake_day: string
  bio: string
  interests: string
  location: string
  timezone: string
  website: string
  email: string
  github: string
  gitten: string
  bluesky: string
  bottletail: string
  rss_feed: string
  bitcoin: string
  ethereum: string
  solana: string
  monero: string
  rss: string
  bsky_display_name: string
  bsky_description: string
  bsky_avatar: string
  bsky_banner: string
  bsky_followers_count: number
  bsky_follows_count: number
  bsky_posts_count: number
}

export interface Artwork {
  id: number
  title: string
  image: string
  category: string
  tags: string[]
  description: string
  year: number
  created_at: string
  updated_at: string
}

export interface Project {
  id: number
  name: string
  description: string
  tags: string[]
  status: 'active' | 'development' | 'maintenance' | 'archived'
  stars: number
  language: string
  github: string
  created_at: string
  updated_at: string
}

export interface Game {
  id: number
  name: string
  description: string
  category: string
  tags: string[]
  players: string
  status: 'playable' | 'development' | 'active' | 'planned'
  features: string[]
  created_at: string
  updated_at: string
}

export interface SocialLink {
  platform: string
  username: string
  url: string
  icon: string
  color: string
  description: string
}