from fastapi import APIRouter, HTTPException
from google.oauth2 import id_token
from google.auth.transport import requests
import jwt
import os
from datetime import datetime, timedelta
from dotenv import load_dotenv
from redis import Redis

load_dotenv()
router = APIRouter(prefix="/auth", tags=["result"])

redis = Redis.from_url(os.getenv("REDIS_URL"))
# redis = Redis(host="host.docker.internal", port=6379)


@router.post("/")
def auth(req: dict):

    google_token = req.get("token")

    if not google_token:
        raise HTTPException(400, "Token missing")

    try:

        info = id_token.verify_oauth2_token(
            google_token, requests.Request(), os.getenv("GOOGLE_CLIENT_ID")
        )

    except Exception:
        raise HTTPException(401, "Invalid Google token")

    google_id = info["sub"]

    redis.set(f"user:{google_id}", "1")

    jwt_token = jwt.encode(
        {"user_id": google_id, "exp": datetime.utcnow() + timedelta(days=7)},
        os.getenv("JWT_SECRET"),
        algorithm="HS256",
    )

    return {"token": jwt_token}
