import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { authApi } from "@/api/client";
import { useToast } from "@/components/Toast";

const Signup: React.FC = () => {
  const navigate = useNavigate();
  const toast = useToast();
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
      const res = await authApi.signup({ username, email, password });
      localStorage.setItem("session_token", res.token);
      localStorage.setItem("user_id", res.userId);
      localStorage.setItem("username", res.username);
      toast.push("登録が完了しました", "success");
      navigate("/chat");
    } catch (err: any) {
      toast.push(err.message || "登録に失敗しました", "error");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center px-6 py-12">
      <div className="w-full max-w-5xl grid gap-8 lg:grid-cols-[0.95fr_1.05fr]">
        <div className="rounded-3xl p-8 md:p-10 card-surface border border-[#eadfce]">
          <h2 className="text-3xl font-display">新規登録</h2>
          <p className="mt-2 text-sm text-[#6a5f52]">
            最初の一言が届く準備を整えます。
          </p>

          <form onSubmit={handleSubmit} className="mt-6 space-y-5">
            <label className="block text-sm font-medium text-[#5a5045]">
              ユーザー名
              <input
                type="text"
                required
                className="mt-2 w-full rounded-xl border border-[#e5d8c7] bg-white px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-[#d7815f]"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
              />
            </label>
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
              className="w-full rounded-xl bg-[#2b2620] py-3 text-sm font-semibold text-white transition hover:bg-[#1e1a16] disabled:opacity-60"
            >
              {loading ? "登録中..." : "登録する"}
            </button>
          </form>

          <div className="mt-6 flex items-center justify-between text-sm text-[#6a5f52]">
            <span>すでにアカウントをお持ちですか？</span>
            <Link
              to="/login"
              className="font-semibold text-[#b04b2f] hover:underline"
            >
              ログインへ
            </Link>
          </div>
        </div>

        <div className="hidden lg:flex flex-col justify-between rounded-3xl p-10 card-surface border border-[#eadfce]">
          <div>
            <p className="text-sm uppercase tracking-[0.4em] text-[#b06c4a] font-semibold">
              Match + Chat
            </p>
            <h1 className="mt-4 text-4xl font-display leading-tight">
              今日の相手と、
              <br />
              すぐにつながる。
            </h1>
            <p className="mt-4 text-base text-[#6a5f52]">
              マッチ後はルームIDを自動取得。
              すぐに会話をスタートできます。
            </p>
          </div>
          <div className="flex items-center gap-3 text-sm text-[#6a5f52]">
            <span className="h-2 w-2 rounded-full bg-[#2b2620]"></span>
            いつでもログインで再開
          </div>
        </div>
      </div>
    </div>
  );
};

export default Signup;
