import React, { useState } from "react";
import { Link } from "react-router-dom";
import {
  clearRoomHistory,
  getRoomHistory,
  removeRoomHistory,
  type RoomHistoryItem,
} from "@/utils/rooms";
import { useToast } from "@/components/Toast";

const Rooms: React.FC = () => {
  const toast = useToast();
  const [history, setHistory] = useState<RoomHistoryItem[]>(() =>
    getRoomHistory()
  );

  const handleClear = () => {
    clearRoomHistory();
    setHistory([]);
    toast.push("ルーム履歴をクリアしました", "success");
  };

  const handleRemove = (roomId: string) => {
    const next = removeRoomHistory(roomId);
    setHistory(next);
    toast.push("履歴から削除しました", "info");
  };

  return (
    <div className="px-6 py-10">
      <div className="mx-auto w-full max-w-5xl">
        <div className="flex flex-col gap-4 md:flex-row md:items-center md:justify-between">
          <div>
            <h1 className="text-3xl font-display">ルーム履歴</h1>
            <p className="mt-2 text-sm text-[#6a5f52]">
              過去にマッチしたルーム一覧です。
            </p>
          </div>
          <div className="flex items-center gap-3">
            <Link
              to="/chat"
              className="rounded-xl bg-[#2b2620] px-4 py-2 text-sm font-semibold text-white"
            >
              新しくマッチ
            </Link>
            <button
              type="button"
              onClick={handleClear}
              className="rounded-xl border border-[#d4c6b5] px-4 py-2 text-sm font-semibold text-[#6a5f52] transition hover:border-[#bda894] hover:text-[#2b2620]"
            >
              履歴をクリア
            </button>
          </div>
        </div>

        <div className="mt-8 space-y-4">
          {history.length === 0 && (
            <div className="rounded-2xl border border-dashed border-[#d8c7b4] bg-[#fffaf2] p-6 text-center text-sm text-[#6a5f52]">
              まだルーム履歴がありません。
            </div>
          )}
          {history.map(item => (
            <div
              key={item.roomId}
              className="flex flex-col gap-3 rounded-2xl border border-[#eadfce] bg-white px-6 py-4 shadow-sm md:flex-row md:items-center md:justify-between"
            >
              <div>
                <div className="text-xs uppercase tracking-[0.3em] text-[#b06c4a]">
                  Room
                </div>
                <div className="mt-2 font-mono text-sm text-[#2b2620]">
                  {item.roomId}
                </div>
                <div className="mt-1 text-xs text-[#6a5f52]">
                  Owner: {item.ownerId}
                </div>
              </div>
              <div className="flex items-center gap-3 text-xs text-[#6a5f52]">
                <span>
                  {new Date(item.matchedAt).toLocaleString("ja-JP", {
                    year: "numeric",
                    month: "2-digit",
                    day: "2-digit",
                    hour: "2-digit",
                    minute: "2-digit",
                  })}
                </span>
                <Link
                  to={`/chat?roomId=${encodeURIComponent(item.roomId)}`}
                  className="rounded-full border border-[#d4c6b5] px-4 py-1 text-xs font-semibold text-[#6a5f52] transition hover:border-[#bda894] hover:text-[#2b2620]"
                >
                  ルームを開く
                </Link>
                <button
                  type="button"
                  onClick={() => handleRemove(item.roomId)}
                  className="rounded-full border border-[#f0c4b6] px-4 py-1 text-xs font-semibold text-[#a14c35] transition hover:border-[#d99b86]"
                >
                  削除
                </button>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default Rooms;
