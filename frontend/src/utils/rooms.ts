export interface RoomHistoryItem {
  roomId: string;
  ownerId: string;
  matchedAt: number;
}

const STORAGE_KEY = "room_history";

export const getRoomHistory = (): RoomHistoryItem[] => {
  const raw = localStorage.getItem(STORAGE_KEY);
  if (!raw) return [];
  try {
    const parsed = JSON.parse(raw) as RoomHistoryItem[];
    return Array.isArray(parsed) ? parsed : [];
  } catch {
    return [];
  }
};

export const saveRoomHistory = (history: RoomHistoryItem[]) => {
  localStorage.setItem(STORAGE_KEY, JSON.stringify(history));
};

export const addRoomHistory = (item: RoomHistoryItem) => {
  const history = getRoomHistory();
  const filtered = history.filter(entry => entry.roomId !== item.roomId);
  const next = [item, ...filtered].slice(0, 20);
  saveRoomHistory(next);
};

export const findRoomHistory = (roomId: string) => {
  return getRoomHistory().find(entry => entry.roomId === roomId) || null;
};

export const removeRoomHistory = (roomId: string) => {
  const history = getRoomHistory();
  const next = history.filter(entry => entry.roomId !== roomId);
  saveRoomHistory(next);
  return next;
};

export const clearRoomHistory = () => {
  saveRoomHistory([]);
};
