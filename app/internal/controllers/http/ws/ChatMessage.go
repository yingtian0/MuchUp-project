package ws

import (
	"MuchUp/app/internal/usecase/dto"
	"MuchUp/app/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"MuchUp/app/internal/controllers/usecase"
	"MuchUp/app/pkg/logger"
	"MuchUp/app/pkg/middleware"

	"github.com/gorilla/websocket"
)

// WebSocket 接続への HTTP Upgrade 設定。
// CheckOrigin が常に true なので、どの Origin からの接続も許可する。
// 開発中は便利だが、本番ではセキュリティ上かなり緩い設定。
var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Hub は WebSocket クライアント全体を管理する中心。
// 役割:
// - 接続中クライアントの管理
// - クライアントの登録/解除
// - メッセージのブロードキャスト
type Hub struct {
	Clients    map[*Client]bool // 接続中クライアント一覧
	Broadcast  chan []byte      // 全体送信用チャネル
	Register   chan *Client     // 新規接続登録用チャネル
	Unregister chan *Client     // 切断通知用チャネル
	Mutex      sync.RWMutex     // Clients への同時アクセス保護
}

// Client は 1 本の WebSocket 接続を表す。
type Client struct {
	Hub    *Hub            // このクライアントが所属する Hub
	Conn   *websocket.Conn // 実際の WebSocket 接続
	Send   chan []byte     // クライアントへ送信するメッセージキュー
	UserID string          // 接続ユーザーID
	RoomID string          // 現在参加中のグループID
}

// ChatHandler は WebSocket メッセージ処理とユースケース呼び出しを担当する。
type ChatHandler struct {
	Hub            *Hub
	MessageUsecase usecase.MessageUsecase
	UserUsecase    usecase.UserUsecase
	Logger         logger.Logger
}

// WebSocket 上でやり取りする共通メッセージ形式。
// Type でイベント種別を判定し、Data に実データを積む。
type WebSocketMessage struct {
	Type      string      `json:"type"`
	Data      interface{} `json:"data"`
	UserID    string      `json:"user_id,omitempty"`
	RoomID    string      `json:"room_id,omitempty"`
	Timestamp int64       `json:"timestamp,omitempty"`
}

// チャット本文をフロントへ返すときの整形済みメッセージ。
type ChatMessage struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	UserID    string `json:"user_id"`
	RoomID    string `json:"room_id"`
	Username  string `json:"username"`
	Timestamp int64  `json:"timestamp"`
}

// Hub のイベントループ。
// Register / Unregister / Broadcast の各チャネルを監視し、
// クライアント管理やメッセージ配信を行う。
func (h *Hub) Run() {
	for {
		select {
		// 新規クライアント登録
		case client := <-h.Register:
			h.Mutex.Lock()
			h.Clients[client] = true
			h.Mutex.Unlock()

			// ユーザー接続通知を同じグループへ送る
			message := WebSocketMessage{
				Type: "user_connected",
				Data: map[string]string{
					"user_id": client.UserID,
					"room_id": client.RoomID,
				},
			}
			if data, err := json.Marshal(message); err == nil {
				h.BroadcastToGroup(data, client.RoomID)
			}

		// クライアント切断
		case client := <-h.Unregister:
			h.Mutex.Lock()
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
			h.Mutex.Unlock()

			// ユーザー切断通知を同じグループへ送る
			message := WebSocketMessage{
				Type: "user_disconnected",
				Data: map[string]string{
					"user_id": client.UserID,
					"room_id": client.RoomID,
				},
			}
			if data, err := json.Marshal(message); err == nil {
				h.BroadcastToGroup(data, client.RoomID)
			}

		// 全体ブロードキャスト
		case message := <-h.Broadcast:
			h.Mutex.RLock()
			for client := range h.Clients {
				select {
				case client.Send <- message:
					// 送信キューへ積めたら何もしない
				default:
					// 送信キューが詰まっているクライアントは切断扱い
					// ※ ここでは RLock 中に delete しているため設計上は危険
					close(client.Send)
					delete(h.Clients, client)
				}
			}
			h.Mutex.RUnlock()
		}
	}
}

// 指定グループに属するクライアントへメッセージを送る。
func (h *Hub) BroadcastToGroup(message []byte, roomID string) {
	h.Mutex.RLock()
	defer h.Mutex.RUnlock()

	for client := range h.Clients {
		if client.RoomID == roomID {
			select {
			case client.Send <- message:
				// 送信キューへ積めた
			default:
				// キュー詰まり時は切断扱い
				// ※ ここでも RLock 中に delete しているため危険
				close(client.Send)
				delete(h.Clients, client)
			}
		}
	}
}

// HandleWebSocket handles WebSocket connections for chat.
// @Summary WebSocket Chat
// @Description Connects via WebSocket for direct and group chat messaging
// @Tags chat
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Success 101 {string} string "Switching Protocols"
// @Router /ws/chat [get]
func (ch *ChatHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// middleware で設定された userID を context から取得
	userID, ok := r.Context().Value(middleware.UserIDContextKey).(string)
	if !ok || userID == "" {
		http.Error(w, "user ID not found in context (middleware missing or invalid)", http.StatusInternalServerError)
		return
	}

	// groupID は無くても許容している
	RoomID, ok := r.Context().Value(middleware.GroupIDContextKey).(string)
	if !ok {
		RoomID = ""
	}

	// HTTP 接続を WebSocket にアップグレード
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	// 接続ごとの Client オブジェクトを作成
	client := &Client{
		Hub:    ch.Hub,
		Conn:   conn,
		Send:   make(chan []byte, 256), // 送信用バッファ
		UserID: userID,
		RoomID: RoomID,
	}

	// Hub へ登録し、送信/受信ループを goroutine で開始
	client.Hub.Register <- client
	go client.WritePump()
	go client.ReadPump(ch)
}

// ReadPump はクライアントから届くメッセージを受信するループ。
// 受信した JSON を WebSocketMessage に変換し、Type ごとに処理を分岐する。
func (c *Client) ReadPump(handler *ChatHandler) {
	defer func() {
		// 接続終了時は Hub に解除通知し、コネクションを閉じる
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, messageBytes, err := c.Conn.ReadMessage()
		if err != nil {
			// 異常切断の場合のみログ出力
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}

		// JSON -> WebSocketMessage へ変換
		var wsMessage WebSocketMessage
		if err := json.Unmarshal(messageBytes, &wsMessage); err != nil {
			log.Printf("Message unmarshal error: %v", err)
			continue
		}

		// メッセージ種別ごとのハンドリング
		switch wsMessage.Type {
		case "chat_message":
			handler.HandleChatMessage(c, wsMessage)
		case "typing":
			handler.HandleTyping(c, wsMessage)
		case "join_group":
			handler.HandleJoinGroup(c, wsMessage)
		}
	}
}

// WritePump はサーバーからクライアントへの送信専用ループ。
// c.Send チャネルに積まれたメッセージを WebSocket 経由で送る。
func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()

	for message := range c.Send {
		if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("WebSocket write error: %v", err)
			return
		}
	}

	// Send チャネルが閉じられたら CloseMessage を送って終了
	c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
}

// チャットメッセージ受信時の処理。
// 主な流れ:
// 1. 受信データを取り出す
// 2. domain entity を作る
// 3. DB 保存
// 4. ユーザー名を取得
// 5. フロント向けレスポンスを作り配信
func (ch *ChatHandler) HandleChatMessage(client *Client, wsMessage WebSocketMessage) {
	// Data は interface{} なので map として解釈する
	data, ok := wsMessage.Data.(map[string]interface{})
	if !ok {
		return
	}

	// フロントから渡された値を取り出す
	content, _ := data["content"].(string)
	recipientId, _ := data["recipientId"].(string)
	roomID, _ := data["room_id"].(string)
	if roomID == "" {
		roomID, _ = data["roomId"].(string)
	}
	if roomID == "" {
		roomID, _ = data["groupId"].(string)
	}
	if roomID == "" {
		roomID = client.RoomID
	}

	// ここでは wsMessage.Type を messageType に入れている
	// ただし呼び出し元 switch は "chat_message" しか来ないため、
	// 後続の "direct"/"group" 判定とは整合していない
	messageType := wsMessage.Type

	input := dto.SendChatMessageInput{
		EventType: dto.MessageEvent,
		SenderID:  client.UserID,
		RoomID:    roomID,
		MessageID: utils.GenerateUUID(),
		Content:   content,
		CreatedAt: time.Now(),
	}

	if err := ch.MessageUsecase.SendChatMessage(context.Background(), input); err != nil {
		log.Printf("Failed to publish message: %v", err)
		return
	}

	// 送信者情報を取得してニックネームを付与
	user, err := ch.UserUsecase.GetUserByID(client.UserID)
	if err != nil {
		logger.NewLogger().WithError(err)
		return
	}

	// フロントへ返す整形済みメッセージ
	chatMessage := ChatMessage{
		ID:        input.MessageID,
		Content:   input.Content,
		UserID:    input.SenderID,
		RoomID:    input.RoomID,
		Username:  user.NickName,
		Timestamp: input.CreatedAt.Unix(),
	}

	response := WebSocketMessage{
		Type: "new_message",
		Data: chatMessage,
	}

	// JSON 化して送信
	if responseData, err := json.Marshal(response); err == nil {
		// ここは messageType の値次第で送信先を変える意図だが、
		// 実際には messageType == "chat_message" のため、
		// この条件には入らない可能性が高い
		if messageType == "direct" && recipientId != "" {
			ch.Hub.SendToUser(responseData, recipientId)
		} else if roomID != "" {
			ch.Hub.BroadcastToGroup(responseData, roomID)
		}
	}
}

// 指定 userID を持つクライアントへ 1:1 送信する。
func (h *Hub) SendToUser(message []byte, userID string) {
	h.Mutex.RLock()
	defer h.Mutex.RUnlock()

	for client := range h.Clients {
		if client.UserID == userID {
			select {
			case client.Send <- message:
				// 送信キューへ積めた
			default:
				// ※ ここも RLock 中の delete なので危険
				close(client.Send)
				delete(h.Clients, client)
			}
		}
	}
}

// タイピング通知処理。
// 誰が typing 中かを同一グループへブロードキャストする。
func (ch *ChatHandler) HandleTyping(client *Client, wsMessage WebSocketMessage) {
	response := WebSocketMessage{
		Type:   "typing",
		UserID: client.UserID,
		Data: map[string]interface{}{
			"user_id": client.UserID,
			"room_id": client.RoomID,
			"typing":  wsMessage.Data,
		},
	}

	if responseData, err := json.Marshal(response); err == nil {
		ch.Hub.BroadcastToGroup(responseData, client.RoomID)
	}
}

// グループ参加処理。
// 既存グループがあれば退出通知を出し、新しいグループへ所属を切り替える。
func (ch *ChatHandler) HandleJoinGroup(client *Client, wsMessage WebSocketMessage) {
	data, ok := wsMessage.Data.(map[string]interface{})
	if !ok {
		return
	}

	newGroupID, ok := data["group_id"].(string)
	if !ok {
		return
	}

	if err := ch.UserUsecase.JoinGroup(client.UserID, newGroupID); err != nil {
		response := WebSocketMessage{
			Type: "join_group_error",
			Data: map[string]string{
				"user_id":  client.UserID,
				"group_id": newGroupID,
				"error":    err.Error(),
			},
		}
		if responseData, marshalErr := json.Marshal(response); marshalErr == nil {
			client.Send <- responseData
		}
		return
	}

	// すでに別グループへ所属していれば退出通知
	if client.RoomID != "" {
		leaveMessage := WebSocketMessage{
			Type: "user_left",
			Data: map[string]string{
				"user_id":  client.UserID,
				"group_id": client.RoomID,
			},
		}
		if data, err := json.Marshal(leaveMessage); err == nil {
			ch.Hub.BroadcastToGroup(data, client.RoomID)
		}
	}

	// 所属グループを更新
	client.RoomID = newGroupID

	// 新グループへ参加通知
	joinMessage := WebSocketMessage{
		Type: "user_joined",
		Data: map[string]string{
			"user_id":  client.UserID,
			"group_id": client.RoomID,
		},
	}
	if data, err := json.Marshal(joinMessage); err == nil {
		ch.Hub.BroadcastToGroup(data, client.RoomID)
	}

	// 新規グループ作成時に AI エージェントの自己紹介メッセージを送信
	// client.UserID が ai_agent 自身でない場合のみ送る
	if client.UserID != "ai_agent" && newGroupID != "" {
		aiMessage := WebSocketMessage{
			Type: "new_message",
			Data: map[string]interface{}{
				"id":        "ai_intro",
				"content":   "こんにちは!私はこのグループをサポートするAIです。皆さんが仲良くなれるようにお手伝いします。何か質問があれば気軽に話しかけてくださいね!",
				"user_id":   "ai_agent",
				"group_id":  newGroupID,
				"username":  "AIエージェント",
				"timestamp": 0,
			},
		}
		if data, err := json.Marshal(aiMessage); err == nil {
			ch.Hub.BroadcastToGroup(data, newGroupID)
		}
	}
}
