import React from "react";
import { Link, NavLink, Outlet, useNavigate } from "react-router-dom";

const AppShell: React.FC = () => {
  const navigate = useNavigate();
  const username = localStorage.getItem("username");

  const handleLogout = () => {
    localStorage.removeItem("session_token");
    localStorage.removeItem("user_id");
    localStorage.removeItem("username");
    navigate("/login");
  };

  return (
    <div className="min-h-screen">
      <header className="sticky top-0 z-40 border-b border-[#eadfce] bg-[#f6f1e7]/80 backdrop-blur">
        <div className="mx-auto flex w-full max-w-6xl items-center justify-between px-6 py-4">
          <Link to="/chat" className="font-display text-xl">
            MuchUp
          </Link>
          <nav className="flex items-center gap-4 text-sm font-medium text-[#6a5f52]">
            <NavLink
              to="/chat"
              className={({ isActive }) =>
                isActive ? "text-[#2b2620]" : "hover:text-[#2b2620]"
              }
            >
              チャット
            </NavLink>
            <NavLink
              to="/rooms"
              className={({ isActive }) =>
                isActive ? "text-[#2b2620]" : "hover:text-[#2b2620]"
              }
            >
              履歴
            </NavLink>
            <NavLink
              to="/profile"
              className={({ isActive }) =>
                isActive ? "text-[#2b2620]" : "hover:text-[#2b2620]"
              }
            >
              プロフィール
            </NavLink>
          </nav>
          <div className="flex items-center gap-3 text-sm">
            <span className="text-[#6a5f52]">{username || "Guest"}</span>
            <button
              type="button"
              onClick={handleLogout}
              className="rounded-full border border-[#d4c6b5] px-4 py-1 text-xs font-semibold text-[#6a5f52] transition hover:border-[#bda894] hover:text-[#2b2620]"
            >
              ログアウト
            </button>
          </div>
        </div>
      </header>
      <main>
        <Outlet />
      </main>
    </div>
  );
};

export default AppShell;
