export interface BlogPost {
  id: number;
  created_at: string;
  updated_at: string;
  title: string;
  date: string;
  time: string;
  authors: string;
  tags: string;
  categories: string;
  draft: boolean;
  share: boolean;
  slug: string;
  layout: string;
  toc: boolean;
  comments: boolean;
  math: boolean;
  featured_image: string;
  featured_image_alt: string;
  featured_video: string;
  featured_video_alt: string;
  summary: string;
  content: string;
  file_path: string;
}

export interface BskyPost {
  id: number;
  created_at: string;
  updated_at: string;
  uri: string;
  cid: string;
  author: string;
  text: string;
  posted_at: string;
  reply_count: number;
  repost_count: number;
  like_count: number;
  quote_count: number;
  has_media: boolean;
  media_urls: string;
}

export interface Profile {
  id: number;
  created_at: string;
  updated_at: string;
  pic: string;
  name: string;
  cake_day: string;
  bio: string;
  interests: string;
  location: string;
  timezone: string;
  website: string;
  email: string;
  github: string;
  bluesky: string;
  rss_feed: string;
  bitcoin_donation_addr: string;
  ethereum_donation_addr: string;
  solana_donation_addr: string;
  monero_donation_addr: string;
  bsky_display_name: string;
  bsky_description: string;
  bsky_avatar: string;
  bsky_banner: string;
  bsky_followers_count: number;
  bsky_follows_count: number;
  bsky_posts_count: number;
}

export interface Game {
  id: number;
  created_at: string;
  updated_at: string;
  slug: string;
  name: string;
  description: string;
  category: string;
  tags: string;
  min_players: number;
  max_players: number;
  multiplayer_supported: boolean;
  has_leaderboard: boolean;
  version: string;
  config: string;
}

export interface GameScore {
  id: number;
  created_at: string;
  updated_at: string;
  game_id: number;
  alias: string;
  score: number;
  level: number;
  data: string;
  ip_hash: string;
}
