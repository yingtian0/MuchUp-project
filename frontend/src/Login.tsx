import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { authApi } from "@/api/client";

const Login: React.FC = () => {
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    setLoading(true);

    try {
      const res = await authApi.login({ email, password });
      // トークンを保存
      localStorage.setItem("session_token", res.token);
      localStorage.setItem("user_id", res.userId);
      // チャット画面（未実装）へ遷移
      // navigate('/chat');
      alert("ログイン成功！(チャット画面へ遷移します)");
    } catch (err: any) {
      setError(err.message || "ログインに失敗しました");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-50">
      <div className="w-full max-w-md p-8 space-y-6 bg-white rounded shadow-md">
        <h2 className="text-2xl font-bold text-center text-gray-800">
          MuchUp ログイン
        </h2>

        {error && (
          <div className="p-3 text-sm text-red-700 bg-red-100 rounded border border-red-200">
            {error}
          </div>
        )}

        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label className="block mb-1 text-sm font-medium text-gray-700">
              メールアドレス
            </label>
            <input
              type="email"
              required
              className="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
          </div>
          <div>
            <label className="block mb-1 text-sm font-medium text-gray-700">
              パスワード
            </label>
            <input
              type="password"
              required
              className="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>
          <button
            type="submit"
            disabled={loading}
            className="w-full py-2 font-bold text-white bg-blue-600 rounded hover:bg-blue-700 disabled:opacity-50"
          >
            {loading ? "ログイン中..." : "ログイン"}
          </button>
        </form>
        <p className="text-sm text-center text-gray-600">
          アカウントをお持ちでないですか？{" "}
          <Link to="/signup" className="text-blue-600 hover:underline">
            新規登録
          </Link>
        </p>
      </div>
    </div>
  );
};

export default Login;
