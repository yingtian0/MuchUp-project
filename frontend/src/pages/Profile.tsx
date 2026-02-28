import React from "react";
import { useNavigate } from "react-router-dom";

const Profile: React.FC = () => {
  const navigate = useNavigate();
  const username = localStorage.getItem("username") || "Unknown";
  const userId = localStorage.getItem("user_id") || "-";
  const token = localStorage.getItem("session_token") || "";

  const handleLogout = () => {
    localStorage.removeItem("session_token");
    localStorage.removeItem("user_id");
    localStorage.removeItem("username");
    navigate("/login");
  };

  return (
    <div className="px-6 py-10">
      <div className="mx-auto w-full max-w-4xl">
        <h1 className="text-3xl font-display">プロフィール</h1>
        <p className="mt-2 text-sm text-[#6a5f52]">
          ローカルに保存されているセッション情報です。
        </p>

        <div className="mt-8 grid gap-6 md:grid-cols-2">
          <div className="rounded-2xl border border-[#eadfce] bg-white p-6 shadow-sm">
            <div className="text-xs uppercase tracking-[0.3em] text-[#b06c4a]">
              User
            </div>
            <div className="mt-4 text-2xl font-display text-[#2b2620]">
              {username}
            </div>
            <div className="mt-2 text-sm text-[#6a5f52]">ID: {userId}</div>
          </div>
          <div className="rounded-2xl border border-[#eadfce] bg-white p-6 shadow-sm">
            <div className="text-xs uppercase tracking-[0.3em] text-[#b06c4a]">
              Token
            </div>
            <div className="mt-4 break-all font-mono text-xs text-[#2b2620]">
              {token || "未保存"}
            </div>
          </div>
        </div>

        <button
          type="button"
          onClick={handleLogout}
          className="mt-8 rounded-xl border border-[#d4c6b5] px-4 py-2 text-sm font-semibold text-[#6a5f52] transition hover:border-[#bda894] hover:text-[#2b2620]"
        >
          ログアウト
        </button>
      </div>
    </div>
  );
};

export default Profile;
