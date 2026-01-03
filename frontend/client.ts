/**
 * API Client Definition
 * 本来は openapi-generator 等で生成されるコードを想定していますが、
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

class AuthApi {
  private basePath: string;

  constructor(basePath: string) {
    this.basePath = basePath;
  }

  private async request<T>(
    path: string,
    method: string,
    body?: any
  ): Promise<T> {
    const headers: HeadersInit = {
      "Content-Type": "application/json",
    };

    // セッショントークンが保存されていれば付与 (Envoyでの検証用)
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
      // エラーハンドリング
      if (response.status === 401 || response.status === 403) {
        throw new Error("認証エラーが発生しました。");
      }
      const errorData = await response.json().catch(() => ({}));
      throw new Error(errorData.message || `API Error: ${response.status}`);
    }

    return response.json();
  }

  /**
   * ユーザーログイン
   */
  async login(req: LoginRequest): Promise<AuthResponse> {
    return this.request<AuthResponse>("/auth/login", "POST", req);
  }

  /**
   * ユーザー登録
   */
  async signup(req: SignupRequest): Promise<AuthResponse> {
    return this.request<AuthResponse>("/auth/signup", "POST", req);
  }
}

export const authApi = new AuthApi(BASE_PATH);
