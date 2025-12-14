// types/index.d.ts
export type BlogPost = {
  id?: number;
  slug: string;
  title: string;
  summary?: string;
  content?: string;
  featured_image?: string;
  date?: string; // RFC3339 string from backend
  authors?: string; // CSV in DB; UI may split into array
  tags?: string;
  categories?: string;
  draft?: boolean;
  share?: boolean;
};

export type Pagination<T = any> = {
  page: number;
  limit: number;
  total: number;
  pages: number;
  items?: T[];
};

export type MediaItem = {
  id: number;
  title: string;
  file_name: string;
  file_name_url?: string;
  file_size?: number;
  mime_type?: string;
  category?: string;
  tags?: string;
  description?: string;
  artist?: string;
  thumbnail?: string;
  width?: number;
  height?: number;
  exif?: string;
  is_public?: boolean;
  created_at?: string;
};

export type Game = {
  id?: number;
  slug: string;
  name: string;
  description?: string;
  config?: any; // GameConfig JSON stored by backend
  has_leaderboard?: boolean;
  created_at?: string;
};

export type GameScore = {
  id?: number;
  game_id?: number;
  alias?: string;
  score: number;
  metadata?: any;
  created_at?: string;
};
