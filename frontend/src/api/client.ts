/**
 * API Client Definition
 * スキーマ定義に基づき手動で実装しています。
 */

const BASE_PATH = import.meta.env.VITE_API_BASE_URL || "http://localhost:8080";

export interface LoginRequest {
  email: string;
  password?: string;
}

export interface SignupRequest {
  username: string;
  email: string;
  password?: string;
}

export interface AuthResponse {
  token: string;
  userId: string;
  username: string;
}

export interface ErrorResponse {
  code: number;
  message: string;
}

export interface ChatMessage {
  messageId: string;
  senderId: string;
  text: string;
  createdAt: number;
}

export interface MatchRoomResponse {
  ownerId: string;
  roomId: string;
}

class BaseApi {
  protected basePath: string;

  constructor(basePath: string) {
    this.basePath = basePath;
  }

  protected async request<T>(
    path: string,
    method: string,
    body?: any
  ): Promise<T> {
    const headers: HeadersInit = {
      "Content-Type": "application/json",
    };

    const token = localStorage.getItem("session_token");
    if (token) {
      headers["Authorization"] = `Bearer ${token}`;
    }

    const response = await fetch(`${this.basePath}${path}`, {
      method,
      headers,
      body: body ? JSON.stringify(body) : undefined,
    });

    if (!response.ok) {
      if (response.status === 401 || response.status === 403) {
        throw new Error("認証エラーが発生しました。");
      }
      const errorData = await response.json().catch(() => ({}));
      throw new Error(errorData.message || `API Error: ${response.status}`);
    }

    return response.json();
  }
}

class AuthApi extends BaseApi {
  /**
   * ユーザーログイン
   */
  async login(req: LoginRequest): Promise<AuthResponse> {
    return this.request<AuthResponse>("/v1/auth/login", "POST", req);
  }

  /**
   * ユーザー登録
   */
  async signup(req: SignupRequest): Promise<AuthResponse> {
    return this.request<AuthResponse>("/v1/auth/signup", "POST", req);
  }
}

class ChatApi extends BaseApi {
  async matchRoom(): Promise<MatchRoomResponse> {
    const res = await this.request<{ owner_id: string; room_id: string }>(
      "/v1/chat/match",
      "POST",
      {}
    );

    return {
      ownerId: res.owner_id,
      roomId: res.room_id,
    };
  }

  async getMessages(roomId: string): Promise<ChatMessage[]> {
    const res = await this.request<{ message: any[] }>(
      `/v1/chat/rooms/${roomId}/messages`,
      "GET"
    );

    return (res.message || []).map(item => ({
      messageId: item.message_id,
      senderId: item.sender_id,
      text: item.text,
      createdAt: item.created_at,
    }));
  }

  // Legacy fallback until the browser-side WebSocket auth/connect flow is wired.
  async sendMessage(params: {
    userId: string;
    roomId: string;
    text: string;
  }): Promise<{ messageId: string; roomId: string; createdAt: number }> {
    const res = await this.request<{
      message_id: string;
      room_id: string;
      created_at: number;
    }>(`/v1/chat/rooms/${params.roomId}/messages`, "POST", {
      user_id: params.userId,
      room_id: params.roomId,
      text: params.text,
    });

    return {
      messageId: res.message_id,
      roomId: res.room_id,
      createdAt: res.created_at,
    };
  }
}

export const authApi = new AuthApi(BASE_PATH);
export const chatApi = new ChatApi(BASE_PATH);
