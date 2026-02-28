import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { authApi } from "@/api/client";
import { useToast } from "@/components/Toast";

const Login: React.FC = () => {
  const navigate = useNavigate();
  const toast = useToast();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
      const res = await authApi.login({ email, password });
      localStorage.setItem("session_token", res.token);
      localStorage.setItem("user_id", res.userId);
      localStorage.setItem("username", res.username);
      toast.push("ログインしました", "success");
      navigate("/chat");
    } catch (err: any) {
      toast.push(err.message || "ログインに失敗しました", "error");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center px-6 py-12">
      <div className="w-full max-w-5xl grid gap-8 lg:grid-cols-[1.1fr_0.9fr]">
        <div className="hidden lg:flex flex-col justify-between rounded-3xl p-10 card-surface border border-[#eadfce]">
          <div>
            <p className="text-sm uppercase tracking-[0.4em] text-[#b06c4a] font-semibold">
              MuchUp
            </p>
            <h1 className="mt-4 text-4xl font-display leading-tight">
              ふと話したくなる夜に、
              <br />
              ひらけるチャットルーム。
            </h1>
            <p className="mt-4 text-base text-[#6a5f52]">
              マッチからメッセージまで、最短で。
              今日はどんな会話が生まれる？
            </p>
          </div>
          <div className="flex items-center gap-3 text-sm text-[#6a5f52]">
            <span className="h-2 w-2 rounded-full bg-[#d7815f]"></span>
            安全なセッションで会話を継続
          </div>
        </div>

        <div className="rounded-3xl p-8 md:p-10 card-surface border border-[#eadfce]">
          <h2 className="text-3xl font-display">ログイン</h2>
          <p className="mt-2 text-sm text-[#6a5f52]">
            登録済みのメールアドレスで続行します。
          </p>

          <form onSubmit={handleSubmit} className="mt-6 space-y-5">
            <label className="block text-sm font-medium text-[#5a5045]">
              メールアドレス
              <input
                type="email"
                required
                className="mt-2 w-full rounded-xl border border-[#e5d8c7] bg-white px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-[#d7815f]"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
              />
            </label>
            <label className="block text-sm font-medium text-[#5a5045]">
              パスワード
              <input
                type="password"
                required
                className="mt-2 w-full rounded-xl border border-[#e5d8c7] bg-white px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-[#d7815f]"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </label>
            <button
              type="submit"
              disabled={loading}
              className="w-full rounded-xl bg-[#d7815f] py-3 text-sm font-semibold text-white transition hover:bg-[#b04b2f] disabled:opacity-60"
            >
              {loading ? "ログイン中..." : "ログインする"}
            </button>
          </form>

          <div className="mt-6 flex items-center justify-between text-sm text-[#6a5f52]">
            <span>アカウントがありませんか？</span>
            <Link
              to="/signup"
              className="font-semibold text-[#b04b2f] hover:underline"
            >
              新規登録へ
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Login;
