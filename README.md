# MuchUp – AI が参加するリアルタイムグループチャット

MuchUp は、**5 人をランダムでマッチング**し、  
**AI が会話内容を理解して場の空気を良くする**ことを目的とした  
リアルタイムチャットアプリケーションです。

Envoy を API Gateway としたマイクロサービス構成により、  
高いスケーラビリティ・リアルタイム性・可観測性を備えています。

---

## ✨ コンセプト

- 👥 **5 人ランダムマッチング**
- 💬 **WebSocket によるリアルタイムチャット**
- 🤖 **AI が会話を解析し、空気を和らげる / 盛り上げる**
- 🔐 **API Gateway 集約型の認証・制御**
- 📈 **Observability による運用前提設計**

---

## 🏗️ 全体アーキテクチャ

```text
[ Browser (SPA) ]
        |
        | HTTPS / WebSocket
        v
[ Envoy API Gateway ]
  - 認可 (ext_authz)
  - レート制限
  - WebSocket Upgrade
        |
        +-------------------+
        |                   |
        v                   v
[ Backend API ]        [ AI Service ]
   (Go)                   (Python)
        |                   |
        +---------+---------+
                  |
               [ Redis ]
         (Streams / PubSub / State)
```
