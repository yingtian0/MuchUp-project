import logging
import os
from concurrent import futures

import django
import grpc
from django.conf import settings

os.environ.setdefault("DJANGO_SETTINGS_MODULE", "llm_service.settings")
django.setup()

from llm_service.protobuf import GenerateReplyRequest, GenerateReplyResponse, TokenUsage
from llm_service.services import ReplyGenerator


logger = logging.getLogger("llm_service")


class LLMService:
    def __init__(self):
        self.generator = ReplyGenerator()

    def GenerateReply(self, request, context):
        payload = self.generator.generate(request)
        usage = TokenUsage(
            prompt_tokens=payload["prompt_tokens"],
            completion_tokens=payload["completion_tokens"],
            total_tokens=payload["prompt_tokens"] + payload["completion_tokens"],
        )
        return GenerateReplyResponse(
            reply_id=payload["reply_id"],
            room_id=payload["room_id"],
            model=payload["model"],
            content=payload["content"],
            finish_reason=payload["finish_reason"],
            created_at=payload["created_at"],
            usage=usage,
        )


def serve():
    logging.basicConfig(level=logging.INFO)
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    handler = grpc.unary_unary_rpc_method_handler(
        LLMService().GenerateReply,
        request_deserializer=GenerateReplyRequest.FromString,
        response_serializer=lambda message: message.SerializeToString(),
    )
    server.add_generic_rpc_handlers(
        (
            grpc.method_handlers_generic_handler(
                "llm.v1.LLMService",
                {"GenerateReply": handler},
            ),
        )
    )
    bind_addr = getattr(settings, "GRPC_BIND_ADDR", "0.0.0.0:50052")
    server.add_insecure_port(bind_addr)
    server.start()
    logger.info("llm gRPC server started on %s", bind_addr)
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
