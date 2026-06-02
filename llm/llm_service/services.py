from uuid import uuid4

from django.conf import settings
from django.utils import timezone


class ReplyGenerator:
    def generate(self, request):
        last_user_message = ""
        for message in request.messages:
            if getattr(message, "role", "") == "user" and getattr(message, "content", ""):
                last_user_message = message.content

        bot_name = getattr(settings, "LLM_BOT_NAME", "AIエージェント")
        if last_user_message:
            content = (
                f"{bot_name}です。『{last_user_message}』をきっかけに、"
                "まずは最近ハマっていることを一人ずつ話してみませんか。"
            )
        else:
            content = (
                f"{bot_name}です。新しいルームができました。"
                "最初は名前と最近気になっていることを一言ずつ話すと入りやすいです。"
            )

        return {
            "reply_id": str(uuid4()),
            "room_id": request.room_id,
            "model": request.model or "template-facilitator-v1",
            "content": content,
            "finish_reason": "stop",
            "created_at": int(timezone.now().timestamp()),
            "prompt_tokens": max(1, len(request.messages) * 16),
            "completion_tokens": max(1, len(content) // 4),
        }
