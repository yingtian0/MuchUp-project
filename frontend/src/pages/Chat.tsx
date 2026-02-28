import React, { useCallback, useEffect, useMemo, useRef, useState } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";
import { chatApi, type ChatMessage } from "@/api/client";
import { addRoomHistory, findRoomHistory } from "@/utils/rooms";
import { useToast } from "@/components/Toast";

const DEFAULT_POLL_INTERVAL = Number(
  import.meta.env.VITE_POLL_INTERVAL_MS || 5000
);
const POLL_OPTIONS = [3000, 5000, 10000];

const Chat: React.FC = () => {
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const toast = useToast();
  const [roomId, setRoomId] = useState<string | null>(null);
  const [ownerId, setOwnerId] = useState<string | null>(null);
  const [messages, setMessages] = useState<ChatMessage[]>([]);
  const [text, setText] = useState("");
  const [loading, setLoading] = useState(false);
  const [sending, setSending] = useState(false);
  const [unreadCount, setUnreadCount] = useState(0);
  const [pollInterval, setPollInterval] = useState(() => {
    const stored = localStorage.getItem("chat_poll_interval");
    if (stored) {
      const parsed = Number(stored);
      if (!Number.isNaN(parsed)) return parsed;
    }
    return DEFAULT_POLL_INTERVAL;
  });

  const inFlightRef = useRef(false);
  const listRef = useRef<HTMLDivElement | null>(null);
  const bottomRef = useRef<HTMLDivElement | null>(null);
  const lastCountRef = useRef(0);

  const userId = useMemo(() => localStorage.getItem("user_id"), []);
  const username = useMemo(() => localStorage.getItem("username"), []);

  useEffect(() => {
    if (!userId) {
      navigate("/login");
    }
  }, [navigate, userId]);

  const isNearBottom = () => {
    const el = listRef.current;
    if (!el) return true;
    const threshold = 80;
    return el.scrollHeight - el.scrollTop - el.clientHeight < threshold;
  };

  const scrollToBottom = () => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth", block: "end" });
  };

  const loadMessages = useCallback(
    async (targetRoomId: string) => {
      if (inFlightRef.current) return;
      inFlightRef.current = true;
      setLoading(true);
      try {
        const res = await chatApi.getMessages(targetRoomId);
        setMessages(res);
      } catch (err: any) {
        toast.push(err.message || "メッセージ取得に失敗しました", "error");
      } finally {
        setLoading(false);
        inFlightRef.current = false;
      }
    },
    [toast]
  );

  const handleMatch = async () => {
    if (!userId) return;
    setLoading(true);
    try {
      const res = await chatApi.matchRoom(userId);
      setRoomId(res.roomId);
      setOwnerId(res.ownerId);
      addRoomHistory({
        roomId: res.roomId,
        ownerId: res.ownerId,
        matchedAt: Date.now(),
      });
      await loadMessages(res.roomId);
      toast.push("ルームを作成しました", "success");
    } catch (err: any) {
      toast.push(err.message || "マッチに失敗しました", "error");
    } finally {
      setLoading(false);
    }
  };

  const handleRefresh = async () => {
    if (!roomId) return;
    await loadMessages(roomId);
  };

  const handleSend = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!userId || !roomId || !text.trim()) return;

    setSending(true);
    try {
      const trimmed = text.trim();
      const res = await chatApi.sendMessage({
        userId,
        roomId,
        text: trimmed,
      });

      setMessages(prev => [
        ...prev,
        {
          messageId: res.messageId,
          senderId: userId,
          text: trimmed,
          createdAt: res.createdAt,
        },
      ]);
      setText("");
      setUnreadCount(0);
      requestAnimationFrame(scrollToBottom);
    } catch (err: any) {
      toast.push(err.message || "送信に失敗しました", "error");
    } finally {
      setSending(false);
    }
  };

  useEffect(() => {
    const paramRoomId = searchParams.get("roomId");
    if (paramRoomId) {
      setRoomId(paramRoomId);
      const entry = findRoomHistory(paramRoomId);
      setOwnerId(entry?.ownerId || null);
      loadMessages(paramRoomId);
    }
  }, [loadMessages, searchParams]);

  useEffect(() => {
    if (!roomId) return;
    const interval = window.setInterval(() => {
      loadMessages(roomId);
    }, pollInterval);
    return () => window.clearInterval(interval);
  }, [loadMessages, pollInterval, roomId]);

  useEffect(() => {
    if (messages.length === 0) {
      lastCountRef.current = 0;
      setUnreadCount(0);
      return;
    }

    const isBottom = isNearBottom();
    if (messages.length > lastCountRef.current) {
      const delta = messages.length - lastCountRef.current;
      if (!isBottom) {
        setUnreadCount(prev => prev + delta);
      } else {
        requestAnimationFrame(scrollToBottom);
        setUnreadCount(0);
      }
      lastCountRef.current = messages.length;
    }
  }, [messages]);

  const handleScroll = () => {
    if (isNearBottom()) {
      setUnreadCount(0);
    }
  };

  const formatTime = (createdAt: number) => {
    const ts = createdAt < 1_000_000_000_000 ? createdAt * 1000 : createdAt;
    const date = new Date(ts);
    return date.toLocaleString("ja-JP", {
      hour: "2-digit",
      minute: "2-digit",
    });
  };

  const handlePollChange = (value: number) => {
    setPollInterval(value);
    localStorage.setItem("chat_poll_interval", String(value));
  };

  return (
    <div className="min-h-screen px-6 py-10">
      <div className="mx-auto w-full max-w-6xl grid gap-6 lg:grid-cols-[0.38fr_0.62fr]">
        <section className="rounded-3xl border border-[#eadfce] card-surface p-6 md:p-8">
          <div className="flex items-start justify-between gap-4">
            <div>
              <p className="text-xs uppercase tracking-[0.4em] text-[#b06c4a] font-semibold">
                Session
              </p>
              <h2 className="mt-3 text-2xl font-display">ルームを開く</h2>
              <p className="mt-2 text-sm text-[#6a5f52]">
                マッチするとルームIDが発行されます。
              </p>
            </div>
            <div className="text-right text-xs text-[#6a5f52]">
              <div>ユーザー</div>
              <div className="font-semibold text-[#2b2620]">
                {username || "Unknown"}
              </div>
            </div>
          </div>

          <div className="mt-6 space-y-4">
            <button
              type="button"
              onClick={handleMatch}
              disabled={loading}
              className="w-full rounded-xl bg-[#2b2620] py-3 text-sm font-semibold text-white transition hover:bg-[#1e1a16] disabled:opacity-60"
            >
              {loading ? "マッチ中..." : "マッチしてルームを作成"}
            </button>
            <div className="rounded-2xl border border-[#efe3d2] bg-[#fff6e8] p-4 text-sm text-[#5a5045]">
              <div className="flex items-center justify-between">
                <span>Room ID</span>
                <span className="font-mono text-xs">{roomId || "未発行"}</span>
              </div>
              <div className="mt-2 flex items-center justify-between">
                <span>Owner</span>
                <span className="font-mono text-xs">{ownerId || "-"}</span>
              </div>
            </div>
            <button
              type="button"
              onClick={handleRefresh}
              disabled={!roomId || loading}
              className="w-full rounded-xl border border-[#d7815f] py-2 text-sm font-semibold text-[#b04b2f] transition hover:bg-[#fdece6] disabled:opacity-50"
            >
              受信を更新
            </button>
            <div className="rounded-2xl border border-[#eadfce] bg-white px-4 py-3 text-xs text-[#6a5f52]">
              <div className="mb-2 font-semibold text-[#2b2620]">
                更新間隔
              </div>
              <div className="flex items-center gap-2">
                {POLL_OPTIONS.map(option => (
                  <button
                    key={option}
                    type="button"
                    onClick={() => handlePollChange(option)}
                    className={`rounded-full border px-3 py-1 text-xs font-semibold transition ${
                      pollInterval === option
                        ? "border-[#2b2620] text-[#2b2620]"
                        : "border-[#d4c6b5] text-[#6a5f52] hover:border-[#bda894]"
                    }`}
                  >
                    {option / 1000}s
                  </button>
                ))}
              </div>
            </div>
          </div>
        </section>

        <section className="rounded-3xl border border-[#eadfce] card-surface p-6 md:p-8 flex flex-col">
          <div className="flex items-center justify-between">
            <div>
              <h2 className="text-2xl font-display">チャット</h2>
              <p className="text-sm text-[#6a5f52]">
                {roomId ? "ルームに参加中" : "ルームを作成して開始"}
              </p>
            </div>
            <div className="text-xs text-[#6a5f52]">
              {loading ? "取得中..." : `${messages.length} messages`}
            </div>
          </div>

          <div
            ref={listRef}
            onScroll={handleScroll}
            className="mt-6 flex-1 space-y-4 overflow-y-auto pr-2"
          >
            {!roomId && (
              <div className="rounded-2xl border border-dashed border-[#d8c7b4] bg-[#fffaf2] p-6 text-center text-sm text-[#6a5f52]">
                まずはルームを作成して会話を始めましょう。
              </div>
            )}
            {roomId && messages.length === 0 && !loading && (
              <div className="rounded-2xl border border-dashed border-[#d8c7b4] bg-[#fffaf2] p-6 text-center text-sm text-[#6a5f52]">
                まだメッセージがありません。
              </div>
            )}
            {messages.map(msg => {
              const isOwn = msg.senderId === userId;
              return (
                <div
                  key={msg.messageId}
                  className={`flex ${isOwn ? "justify-end" : "justify-start"}`}
                >
                  <div
                    className={`max-w-[75%] rounded-2xl px-4 py-3 text-sm shadow-sm ${
                      isOwn
                        ? "bg-[#d7815f] text-white"
                        : "bg-white border border-[#ede0cf] text-[#2b2620]"
                    }`}
                  >
                    <p>{msg.text}</p>
                    <div
                      className={`mt-2 text-xs ${
                        isOwn ? "text-white/80" : "text-[#8a7c6c]"
                      }`}
                    >
                      {formatTime(msg.createdAt)}
                    </div>
                  </div>
                </div>
              );
            })}
            <div ref={bottomRef} />
          </div>

          {unreadCount > 0 && (
            <button
              type="button"
              onClick={scrollToBottom}
              className="mt-3 self-center rounded-full border border-[#d7815f] px-4 py-1 text-xs font-semibold text-[#b04b2f]"
            >
              新着 {unreadCount} 件を表示
            </button>
          )}

          <form onSubmit={handleSend} className="mt-6 flex gap-3">
            <input
              type="text"
              placeholder={
                roomId ? "メッセージを入力" : "ルーム作成後に入力できます"
              }
              className="flex-1 rounded-xl border border-[#e5d8c7] bg-white px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-[#d7815f]"
              value={text}
              onChange={(e) => setText(e.target.value)}
              disabled={!roomId || sending}
            />
            <button
              type="submit"
              disabled={!roomId || sending || !text.trim()}
              className="rounded-xl bg-[#2b2620] px-5 text-sm font-semibold text-white transition hover:bg-[#1e1a16] disabled:opacity-40"
            >
              {sending ? "送信中" : "送信"}
            </button>
          </form>
        </section>
      </div>
    </div>
  );
};

export default Chat;
