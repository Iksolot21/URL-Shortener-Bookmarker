// interfaces/interfaces.ts

export interface User {
    id: number;
    username: string;
    email: string;
    createdAt: string;
}

export interface Bookmark {
    id: number;
    userId: number;
    url: string;
    description: string;
    tags: string[];
    createdAt: string;
}
 export interface AuthResponse {
   token: string
}

export interface ShortenResponse {
  short_url: string
}