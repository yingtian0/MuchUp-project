import os


SECRET_KEY = os.getenv("DJANGO_SECRET_KEY", "dev-llm-service-secret-key")
DEBUG = os.getenv("DJANGO_DEBUG", "false").lower() == "true"
ALLOWED_HOSTS = ["*"]
INSTALLED_APPS = [
    "llm_service",
]
DATABASES = {
    "default": {
        "ENGINE": "django.db.backends.sqlite3",
        "NAME": os.getenv("SQLITE_PATH", ":memory:"),
    }
}
USE_TZ = True
TIME_ZONE = os.getenv("TIME_ZONE", "Asia/Tokyo")
LLM_BOT_NAME = os.getenv("LLM_BOT_NAME", "AIエージェント")
GRPC_BIND_ADDR = os.getenv("LLM_GRPC_ADDR", "0.0.0.0:50052")
