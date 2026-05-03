import os
import fnmatch
import mimetypes
from src.constants import constants

SKIP_EXTENSIONS = constants.SKIP_EXTENSIONS
SKIP_FILENAME_PATTERNS = constants.SKIP_FILENAME_PATTERNS
SKIP_FILENAMES = constants.SKIP_FILENAMES


def is_binary_file(path: str) -> bool:
    try:
        with open(path, "rb") as f:
            buf = f.read(512)
    except Exception:
        return True

    # 1. MIME type detection (best-effort in Python)
    mime_type, _ = mimetypes.guess_type(path)
    if mime_type is not None and not mime_type.startswith("text"):
        return True

    # 2. Null byte check (strong fallback)
    if b"\x00" in buf:
        return True

    return False


def should_process(path: str) -> bool:
    name = os.path.basename(path)
    ext = os.path.splitext(name)[1]

    # Skip files larger than 500KB
    try:
        size = os.path.getsize(path)
        if size > 500 * 1024:
            return False
    except Exception:
        return False

    # Skip extensions
    if ext in SKIP_EXTENSIONS:
        return False

    # Skip filenames
    if name in SKIP_FILENAMES:
        return False

    # Skip filename patterns
    for pattern in SKIP_FILENAME_PATTERNS:
        if fnmatch.fnmatch(name, pattern):
            return False

    # Skip binary files
    if is_binary_file(path):
        return False

    return True
