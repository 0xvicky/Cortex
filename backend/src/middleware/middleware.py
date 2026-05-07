from fastapi import Depends, HTTPException
from fastapi.security import HTTPBearer
import jwt
import os

security = HTTPBearer()

SECRET = os.getenv("JWT_SECRET")


def verify_jwt(credentials = Depends(security)):

    token = credentials.credentials

    try:

        payload = jwt.decode(
            token,
            SECRET,
            algorithms=["HS256"]
        )

        return payload

    except:
        raise HTTPException(
            status_code=401,
            detail="Invalid token"
        )