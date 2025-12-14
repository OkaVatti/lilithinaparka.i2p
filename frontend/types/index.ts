export * from "./api";

export interface ThemeOption {
  name: string;
  value: string;
  preview: {
    bg: string;
    fg: string;
    primary: string;
  };
}

export interface Pagination {
  page: number;
  limit: number;
  total: number;
  pages: number;
}

export interface MediaItem {
  id: number;
  title: string;
  description: string;
  file_name: string;
  file_size: number;
  mime_type: string;
  width: number;
  height: number;
  duration: number;
  thumbnail: string;
  category: string;
  tags: string;
  is_public: boolean;
  views: number;
  likes: number;
  artist: string;
  year: number;
  exif: string;
  created_at: string;
  updated_at: string;
}

export interface MediaCategory {
  id: number;
  name: string;
  description: string;
  slug: string;
  item_count: number;
  is_public: boolean;
  created_at: string;
  updated_at: string;
}

export interface AdminStats {
  blogPosts: number;
  mediaItems: number;
  games: number;
  bskyPosts: number;
  dbSize: string;
  lastBackup: string;
  uptime: number;
}

export interface SearchResult {
  type: "blog" | "media" | "game";
  id: number;
  title: string;
  description: string;
  slug: string;
  date?: string;
  score: number;
}

export interface User {
  id: number;
  username: string;
  email: string;
  role: "user" | "admin";
  profile: Record<string, any>;
  last_login: string;
  created_at: string;
}

export interface LoginResponse {
  token: string;
  expires_at: number;
  user: User;
}

export interface GameConfig {
  canvas_width: number;
  canvas_height: number;
  instructions: string;
  controls: {
    up: string;
    down: string;
    left: string;
    right: string;
    action: string;
  };
  scoring: {
    points_per_item: number;
    bonus_points: number;
    time_bonus: number;
  };
}

export interface GameMetadata {
  name: string;
  version: string;
  author: string;
  description: string;
  tags: string[];
  min_players: number;
  max_players: number;
  multiplayer: boolean;
  leaderboard: boolean;
  instructions: string;
  controls: Record<string, string>;
}

export interface Notification {
  id: number;
  user_id: number;
  type: "info" | "success" | "warning" | "error";
  title: string;
  message: string;
  read: boolean;
  action_url?: string;
  data?: string;
  expires_at?: string;
  created_at: string;
  updated_at: string;
}

export interface NotificationPreferences {
  id: number;
  user_id: number;
  email_enabled: boolean;
  push_enabled: boolean;
  in_app_enabled: boolean;
  notify_on_comment: boolean;
  notify_on_like: boolean;
  notify_on_follow: boolean;
  notify_on_post: boolean;
  quiet_hours_start?: string;
  quiet_hours_end?: string;
  created_at: string;
  updated_at: string;
}

export interface AnalyticsEvent {
  id: number;
  session_id: string;
  event_type: string;
  page: string;
  referrer: string;
  user_agent: string;
  ip_hash: string;
  data: string;
  created_at: string;
}

export interface DailyStats {
  id: number;
  date: string;
  page_views: number;
  visitors: number;
  sessions: number;
  blog_views: number;
  game_plays: number;
  media_views: number;
  data: string;
  created_at: string;
  updated_at: string;
}
