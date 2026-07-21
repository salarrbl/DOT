> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/bypass/00-index|← Back to Bypass]]

# 403 Bypass

1. **Header manipulation**: `X-Forwarded-For: 127.0.0.1`, `X-Real-IP`, `X-Original-URL`, `X-Custom-IP-Authorization`
2. **Referer spoof**: `Referer: target.com/allowed-page`
3. **Path normalization**: `/Admin`, `/%61dmin`, `/admin/`, `/admin/.`, `/./admin`, `//admin`
4. **Verb tampering**: GET → POST, PUT, OPTIONS, PATCH, DELETE, TRACE
5. **Parameter injection**: `?admin=true`, `?debug=1`
6. **Wayback Machine**: same path was public before
7. **Backup files**: `file.php~`, `file.php.bak`, `.swp`, `.git/HEAD`
8. **Host header**: `Host: internal-target.com`

The 403 → Goldmine pattern: bypass → find `/reset-password` → XSS in reset params.
